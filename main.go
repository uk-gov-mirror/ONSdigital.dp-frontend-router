package main

import (
	"math/rand"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"time"

	"github.com/ONSdigital/dp-frontend-router/handlers/homepage"
	"github.com/ONSdigital/go-ns/handlers/requestID"
	"github.com/ONSdigital/go-ns/handlers/timeout"
	"github.com/ONSdigital/go-ns/log"
	"github.com/gorilla/pat"
	"github.com/justinas/alice"
)

type config struct {
	BindAddr    string
	BabbageURL  string
	RendererURL string
}

func main() {
	cfg := config{
		BindAddr:    ":20000",
		BabbageURL:  "https://www.ons.gov.uk",
		RendererURL: "http://localhost:20010",
	}

	if v := os.Getenv("BIND_ADDR"); len(v) > 0 {
		cfg.BindAddr = v
	}
	if v := os.Getenv("BABBAGE_URL"); len(v) > 0 {
		cfg.BabbageURL = v
	}
	if v := os.Getenv("RENDERER_URL"); len(v) > 0 {
		cfg.RendererURL = v
	}

	log.Namespace = "dp-frontend-router"

	router := pat.New()
	alice := alice.New(
		timeout.Handler(10*time.Second),
		log.Handler,
		requestID.Handler(16),
	).Then(router)

	babbageURL, err := url.Parse(cfg.BabbageURL)
	if err != nil {
		log.Error(err, nil)
		os.Exit(1)
	}

	reverseProxy := createReverseProxy(babbageURL)

	router.Handle("/", abHandler(http.HandlerFunc(homepage.Handler(cfg.RendererURL)), reverseProxy))
	router.Handle("/{uri:.*}", reverseProxy)

	log.Debug("Starting server", log.Data{"bind_addr": cfg.BindAddr})

	server := &http.Server{
		Addr:         cfg.BindAddr,
		Handler:      alice,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Error(err, nil)
		os.Exit(2)
	}
}

//abHandler ... percentA is the percentage of request that handler 'a' is used
func abHandler(a, b http.Handler, percentA int) http.Handler {
	if percentA < 0 || percentA > 100 {
		panic("Percent 'a' but be between 0 and 100")
	}
	rand.Seed(time.Now().UnixNano())

	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// Detect cookie
		cookie, _ := req.Cookie("homepage-version")

		if cookie == nil {
			var cookieValue string
			if rand.Intn(100) < percentA {
				cookieValue = "A"
			} else {
				cookieValue = "B"
			}

			expiration := time.Now().Add(365 * 24 * time.Hour)
			cookie = &http.Cookie{Name: "homepage-version", Value: cookieValue, Expires: expiration}
			http.SetCookie(w, cookie)
		}

		// Use cookie value to direct to a or b handler
		if cookie.Value == "A" {
			a.ServeHTTP(w, req)
		}
		if cookie.Value == "B" {
			b.ServeHTTP(w, req)
		}
	})
}

func createReverseProxy(babbageURL *url.URL) http.Handler {
	proxy := httputil.NewSingleHostReverseProxy(babbageURL)
	director := proxy.Director
	proxy.Transport = &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   5 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   5 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
	proxy.Director = func(req *http.Request) {
		director(req)
		req.Host = babbageURL.Host
	}
	return proxy
}
