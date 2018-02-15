package analytics

import (
	"errors"
	"fmt"
	"net/http"

	"gopkg.in/mgo.v2"

	"github.com/ONSdigital/dp-frontend-router/config"
	"github.com/ONSdigital/go-ns/log"
	"github.com/dgrijalva/jwt-go"
	"time"
	"encoding/json"
	"bytes"
)

const redirUrlParam = "redirUrl"
const pageIndexParam = "pageIndex"
const pageSizeParam = "pageSize"
const linkIndexParam = "linkIndex"
const urlParam = "url"
const termParam = "term"
const searchTypeParam = "type"
const timestampKey = "timestamp"
const sortByParam = "sortBy"
const modelParam = "model"

// Service - defines a Stats Service Interface
type Service interface {
	CaptureAnalyticsData(r *http.Request) (string, error)
}

// ServiceImpl - Implementation of the Analytics Service interface.
type ServiceImpl struct{}

type SearchData struct {
	RedirUrl string `json:"redirUrl"`
	Url string `json:"url"`
	Term string `json:"term"`
	ListType string `json:"listType"`
	PageIndex float64 `json:"pageIndex"`
	LinkIndex float64 `json:"linkIndex"`
	PageSize float64 `json:"pageSize"`
	TimeStamp time.Time `json:"timeStamp"`
	SortBy string `json:"sortBy"`
	Model string `json:"model"`
}

// NewServiceImpl - Creates a new Analytics ServiceImpl.
func NewServiceImpl() *ServiceImpl {
	return &ServiceImpl{}
}

// CaptureAnalyticsData - captures the analytics values
func (s *ServiceImpl) CaptureAnalyticsData(r *http.Request) (string, error) {
	redirUrl := r.URL.Path
	data := r.URL.Query().Get(":data")

	token, err := jwt.Parse(data, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(config.RedirectSecret), nil
	})

	if err != nil {
		log.ErrorR(r, err, nil)
		return "", errors.New("Invalid redirect data")
	}

	log.DebugR(r, "token", log.Data{"token": token})

	var url, term, listType, sortBy, model string
	var pageIndex, linkIndex, pageSize float64

	var claims jwt.MapClaims
	var ok bool
	if claims, ok = token.Claims.(jwt.MapClaims); !ok || !token.Valid {
		log.ErrorR(r, errors.New("error validating token"), nil)
		return "", errors.New("error validating token")
	}

	if s, ok := claims["uri"].(string); ok {
		url = s
	}
	if s, ok := claims["term"].(string); ok {
		term = s
	}
	if s, ok := claims["listType"].(string); ok {
		listType = s
	}
	if s, ok := claims["page"].(float64); ok {
		pageIndex = s
	}
	if s, ok := claims["index"].(float64); ok {
		linkIndex = s
	}
	if s, ok := claims["pageSize"].(float64); ok {
		pageSize = s
	}
	if s, ok := claims["sortBy"].(string); ok {
		sortBy = s
	}

	if sortBy == "ltr" {
		if s, ok := claims["model"].(string); ok {
			model = s
		}
	}

	if len(url) == 0 {
		log.ErrorR(r, errors.New("Failed to redirect to search results as parameter URL was missing."), nil)
		return "", errors.New("400: URL is a mandatory parameter.")
	}

	searchData := SearchData{RedirUrl:redirUrl, Url: url, Term: term, ListType: listType,
		PageIndex: pageIndex, LinkIndex: linkIndex, PageSize: pageSize,
		TimeStamp: time.Now(), SortBy: sortBy, Model: model}

	if config.SearchStatsDatabase == "mongo" {
		// Store the results in mongoDB
		storeAnalyticsData(r, searchData)
	} else if config.SearchStatsDatabase == "elastic" {
		// Store the results in Elasticsearch
		storeAnalyticsDataElasticsearch(r, searchData)
	} else {
		message := fmt.Sprintf("Unknown SearchStatsDatabase: %s", config.SearchStatsDatabase)
		log.ErrorR(r, errors.New(message), nil)
	}

	// Log the data
	log.DebugR(r, "CaptureAnalyticsData", log.Data{
		redirUrlParam:   searchData.RedirUrl,
		urlParam:        searchData.Url,
		termParam:       searchData.Term,
		searchTypeParam: searchData.ListType,
		pageIndexParam:  searchData.PageIndex,
		linkIndexParam:  searchData.LinkIndex,
		pageSizeParam:   searchData.PageSize,
		timestampKey:    searchData.TimeStamp,
		sortByParam:	 searchData.SortBy,
		modelParam:		 searchData.Model,
	})
	return url, nil
}

func storeAnalyticsData(r *http.Request, searchData SearchData) {
	session, err := mgo.Dial(config.MongoURL)
	if err != nil {
		log.ErrorR(r, err, nil)
	}
	defer session.Close()

	c := session.DB(config.MongoDb).C(config.SearchStatsCollection)
	err = c.Insert(&searchData)

	if err != nil {
		log.ErrorR(r, err, nil)
	}

	c.Database.Session.Close()
}

func storeAnalyticsDataElasticsearch(r *http.Request, searchData SearchData) {
	uri := fmt.Sprintf("%s/%s/document/", config.ElasticsearchURL, config.SearchStatsCollection)

	b, err := json.Marshal(searchData)
	if err != nil {
		log.ErrorR(r, err, nil)
	}

	req, err := http.NewRequest(http.MethodPost, uri, bytes.NewBuffer(b))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		log.ErrorR(r, err, nil)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.ErrorR(r, err, nil)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		err := errors.New(fmt.Sprintf("Returned non 201 response: %d", resp.StatusCode))
		log.ErrorR(r, err, nil)
	}
}