// +build production
// Code generated by go-bindata.
// sources:
// templates/error.tmpl
// templates/main.tmpl
// templates/partials/footer.tmpl
// templates/partials/header.tmpl
// redirects/redirects.csv
// DO NOT EDIT!

package assets

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _templatesErrorTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x52\xc1\x8e\xd3\x30\x10\xbd\xe7\x2b\x46\xbe\x3b\xd5\x5e\x51\x1a\x21\x01\x12\x48\x70\x5b\xc4\x71\x35\xb1\x27\xb5\x55\xc7\x63\x79\x26\xe9\x56\x4b\xfe\x1d\xa5\xb0\x2c\xad\x7a\xe2\xf4\xc6\xcf\x6f\x34\xe3\xf7\xdc\xf9\xb8\x80\x4b\x28\xb2\x37\x05\x0f\x64\x03\xa1\xa7\x0a\x92\x50\xc9\xf4\xcd\xbf\xf7\xa7\x8a\xa5\x50\xdd\xd8\xf0\x70\xa7\xe9\xe9\x49\xa3\x26\x82\x31\xcd\x12\xac\x0d\x98\x46\x6b\x07\x56\xe5\xc9\xf4\x2f\x2f\xd0\x7e\xaa\x95\x6b\xfb\x78\x11\xad\x6b\xb7\x0b\x0f\x7d\xd3\xed\x7c\x5c\xfa\xe6\x15\xef\xcc\x83\x82\x99\x92\xb5\x13\xd6\x9b\x85\x0e\x35\x7a\xbb\xa9\xee\xf1\x8e\x13\x78\x92\xa3\x72\xb1\x17\x62\x9c\x53\xb2\xa7\xe8\x35\x6c\x72\xac\x1a\x5d\xa2\xeb\xc6\x81\x9f\x61\xe0\x67\x6b\x0b\x7a\x4f\xfe\x77\x5d\xc9\x5b\x8f\xf5\x78\x7d\xb2\x56\xa8\x60\x45\x25\x6f\x13\x8d\x0a\x42\x4e\x23\x67\xf2\x66\x7b\xcd\x7f\x6f\xa9\x27\xb6\x1a\x62\xf5\x62\xfa\xe6\xcd\xb4\x8f\x24\xae\xc6\xb2\x4d\x80\x9f\x20\x38\xd2\xe7\xc7\x6f\x5f\x61\x5d\x9b\xae\xf4\x5f\x46\x38\xf3\x0c\xa2\x31\x25\xa0\xec\x78\xce\xba\xf9\x56\x79\x48\x34\x09\x94\x44\x28\x04\x1d\x42\xa8\x34\xee\xcd\x84\x31\x29\xbf\x3b\xd1\xd0\x3a\x9e\x26\xca\x2a\xef\x39\x4b\x7b\xe0\xa5\x9d\x8f\x06\x2e\x31\xee\xcd\x07\xce\x8a\x4e\xe1\xbb\x98\xde\xfd\xa9\x67\xe9\x76\xd8\xb7\xf0\x83\x00\x0b\x27\x3e\x44\x21\x18\xb9\x02\xe6\x33\xc4\xec\x38\x2f\x94\x23\x65\x47\xa0\x21\x0a\x4c\x78\x86\x80\x0b\x81\xc3\x59\xc8\xb7\xdd\xae\xfc\xcd\xfc\x16\xde\x22\xb9\xfe\x13\xaf\xf8\x2b\x00\x00\xff\xff\x38\x7a\x05\x98\xac\x02\x00\x00")

func templatesErrorTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesErrorTmpl,
		"templates/error.tmpl",
	)
}

func templatesErrorTmpl() (*asset, error) {
	bytes, err := templatesErrorTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/error.tmpl", size: 684, mode: os.FileMode(420), modTime: time.Unix(1513248350, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesMainTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x57\x7f\x6f\xdb\x38\x12\xfd\x7b\xf3\x29\x66\xb5\xc0\xc9\x6e\x44\xc9\xda\x5c\x7c\x4e\x6c\x05\x48\xb7\xc1\x5e\x80\xb6\x09\xae\x29\xee\x0e\x41\x80\xd2\xd2\x48\x62\x4b\x91\x2a\x39\xb2\x6b\x38\xfa\xee\x07\xca\x3f\xe2\xa4\x69\x2e\x8b\xa0\xff\x88\xa2\x66\xe6\xbd\x37\xc3\x21\x4d\x4f\x7e\x7d\x73\xf1\xc7\xd5\x7f\x2f\xcf\xa0\xa4\x4a\x9e\xec\x4d\x36\x03\xf2\xec\x64\x0f\x00\x60\x42\x82\x24\x9e\x9c\x2b\x42\xa3\xb8\x04\x8b\x66\x86\x06\xd0\x18\x6d\x80\xc1\x45\x9e\x8b\x14\x21\xd7\x06\xde\x73\x12\xda\xb9\x7c\x20\x4e\xc2\x92\x48\xed\x24\x5a\x45\xaf\x90\x2a\x24\x0e\x25\x51\xcd\xf0\x6b\x23\x66\x89\xf7\x1f\xf6\xf1\x94\xfd\xa1\xab\x9a\x93\x98\x4a\xf4\x20\xd5\x8a\x50\x51\xe2\x9d\x9f\x25\x98\x15\x18\xa4\xa5\xd1\x15\x26\xb1\x07\xd1\x2e\x48\x5a\x72\x63\x91\x12\xaf\xa1\x9c\x8d\xbc\xfb\xb6\x0d\xc6\x5c\x64\x54\x26\x19\xce\x44\x8a\xac\x9b\x04\x42\x09\x12\x5c\x32\x9b\x72\x89\x49\x1c\x0e\x82\xc6\xa2\xe9\xa6\x7c\x2a\x3b\x1e\xc5\x2b\x4c\xbc\x99\xc0\x79\xad\x0d\x79\xbb\xc8\x2b\x53\xae\x4d\xc5\x89\x65\x48\x98\xba\x84\x77\x54\x13\x4a\xac\x4b\xad\x30\x51\xfa\x91\x48\x2a\xb1\x42\x96\x6a\xa9\xcd\x4e\xd0\x6f\x87\xa3\xc3\xa3\xc3\xd7\x8f\xf8\xf3\xba\x96\xc8\x2a\x3d\x15\x12\xd9\x1c\xa7\x8c\xd7\x35\xb3\xc4\xa9\xb1\x6c\xca\x0d\xb3\xb4\xb8\x57\xb4\x07\x48\x52\xa8\x2f\x60\x50\x26\x9e\x2d\xb5\xa1\xb4\x21\x10\xa9\xd3\x5b\x1a\xcc\x13\xcf\x2d\x84\x3d\x8e\xa2\x34\x53\xa1\x56\x36\x2c\xf4\x2c\x6c\xbe\x44\xdc\x5a\x24\x1b\x89\x8a\x17\x68\xa3\x9c\xcf\x5c\x4c\x28\x52\xed\x01\x2d\x6a\x4c\xbc\xce\x12\x7d\x63\x1d\xd6\xc9\xde\xde\x2f\xbf\x2c\x97\x9f\x26\xbf\x32\x76\x2d\x72\x90\x84\x70\x7e\x06\xa3\x9b\x93\x4f\x70\x0b\x96\xe7\xf8\xcf\xab\x77\x6f\xdb\xb6\x53\x74\x5f\x93\x13\x6f\x4b\x44\xda\x08\x5a\x2e\xc3\x4b\x4e\xae\xcd\xde\x8a\xa9\xe1\x66\x71\xda\x49\xb9\xe4\x54\xb6\x6d\x94\x5a\x1b\x69\x99\x31\x81\x61\x6a\xad\x77\xb2\xe1\xbd\x46\x95\x89\xfc\x86\xb1\xc7\x18\x77\x94\x9d\x9f\xc1\xd1\xcf\x51\x25\x90\x1d\xad\x35\x6d\x39\x7f\xa8\xea\x5e\xb5\x0a\x5a\xcb\x72\x1f\x7e\x8a\xb6\x8a\x0b\xf5\x50\x1b\x63\x4f\xe8\x9b\x44\xab\xad\x3f\x99\xea\x6c\x01\xa9\xe4\xd6\x3a\x1e\x08\xaf\x16\x35\x42\xdb\xba\x15\x5f\x21\x01\x61\x55\x4b\x4e\x08\x5e\xcd\x8d\xdb\x56\xb6\x8b\x45\xe3\x41\x08\x6d\xbb\xb7\xee\x67\x2e\x14\x88\x2c\xf1\xdc\x8b\x07\x46\x4b\xdc\xbc\x13\x9f\x0a\x95\xe1\xb7\xc4\x63\xf1\x5a\xe0\x0a\x7a\x21\x50\x66\xeb\x12\x4c\xba\x1c\x9e\x64\xcd\xb5\x26\xc7\xba\xe5\xb4\xa9\x11\x35\x6d\x10\x7b\x79\xa3\xba\x9d\xda\x13\x81\x0d\x74\x50\x04\x26\xe0\x41\xd5\x5f\x8a\x6b\xff\x4f\xad\x0b\x89\xa7\x8a\xcb\x85\x3b\xaf\x2e\xa6\x9f\x31\x25\xff\x26\x31\x63\x71\x6d\x6e\x12\xf7\xb8\xbd\xdd\xc6\xf7\x97\x1b\x48\x67\x08\xbf\x26\xab\xe1\xf6\xf6\xfa\xa6\x1f\xd6\x8d\x2d\x7b\xdc\x14\x4d\x85\x8a\x6c\xbf\x0d\x3a\xa3\x4c\xe2\x57\x0a\xe7\xf0\x86\x13\xf6\xfa\x63\x9e\xd8\x30\x35\xc8\x09\xcf\x24\x3a\xc7\x9e\xee\x07\x6b\xd0\x2a\xb1\x61\x81\xb4\x36\xd8\xd7\x8b\x2b\x5e\xbc\xe7\x15\xf6\x74\xff\x7a\x70\x33\xe6\x21\xb7\x0b\x95\x26\xf1\x98\x87\xd6\xa4\x49\x31\xae\xc2\x9a\x1b\x54\xf4\x5e\x67\x18\x0a\x65\xd1\xd0\x6b\xcc\xb5\xc1\x9e\x4b\x6f\x8d\xda\xf6\x7b\x73\xa1\x32\x3d\x0f\x32\x9d\x76\xda\x02\x7f\x55\x1f\x3f\xf0\xa3\x68\x3e\x9f\x87\x45\x57\x04\xc6\x37\x55\x08\x53\x5d\x45\x77\xb3\xcf\xd6\x0f\xfc\x82\xfb\xfd\xf1\x1a\xb2\xe0\x3d\x7f\x95\x84\x1f\x80\xff\xf1\x94\x1d\x0e\x47\x47\xbf\x0f\x0e\xfe\xc1\x0e\xfc\x00\x96\x3e\x97\x52\xcf\x4f\x55\x5a\x6a\xe3\x1f\x03\x99\x06\xdb\x7b\xb1\x16\x55\xe6\x22\x6b\x5e\xa0\x3b\x68\xbb\x20\x37\xf1\x8f\x41\xea\xb4\xfb\x19\x09\x6b\x4e\xa5\x3b\x06\x61\x1f\x0a\xa4\x0f\xc8\x4d\x5a\xf6\xfa\xb0\x7f\xe7\x51\x72\x5b\xde\x01\x6f\x16\x69\xd7\x7b\xb9\xed\x29\xd7\x3a\xd1\x2b\xb8\xba\x78\x73\x01\x0c\xfe\x5d\xa2\x02\xdb\x39\x81\xb0\x50\xe9\x19\x66\x40\x1a\x0c\xaa\x0c\x0d\x1a\x98\xa3\x2f\x25\x28\x5c\x7d\xe6\x59\xb6\xf1\x26\x34\x15\x08\x45\x1a\x9c\x5e\xf8\xf3\x14\x0c\xda\x5a\x2b\x8b\x3b\x54\x51\x04\x22\xef\x7d\x9f\x49\x92\x24\xe0\x47\x2b\x24\xff\x9e\xb8\x28\xea\x86\x19\x37\xa0\x9a\x6a\x8a\xe6\x22\xff\x17\xda\x46\x92\x85\x04\x96\xcb\xdf\x44\xee\x78\x1a\x49\xe1\x03\x73\xdb\xc2\x72\xf9\x84\x09\xa5\xc5\xb6\x85\x81\x4b\x5f\xe4\x6d\x3b\xfe\x9e\xd4\x20\x35\x46\xdd\x95\x75\x9d\xe9\x3e\xf8\x7f\x7b\x80\x98\xf8\xb0\xff\x50\xdf\x03\xc0\x16\x1c\x23\x3c\x92\xdb\xe3\x34\xdf\x85\xbf\x8a\xb6\x47\xdf\x5d\xd4\xa6\x95\xf7\xf6\xb6\x9e\xef\x90\xdb\xc6\x20\x90\xa8\x10\xb4\x5a\x2d\x07\x83\xd4\x60\x26\xc8\x2d\x9a\xfb\x79\x3b\x8e\xa2\x12\x65\x1d\x6e\x7b\xd9\xdd\x29\xba\xee\x5e\xb5\xfc\xf6\x7b\x54\x75\x70\x42\x15\xcc\x01\x32\xad\xd8\x54\x37\x2a\x45\xe6\x70\xa3\x87\x2d\xe6\x7c\x4c\x1c\xf7\xfa\xcb\xdd\x76\xc6\x19\x2a\x72\x2f\x57\xa2\xc2\x0b\x75\xe9\x1a\x3a\x00\x3f\xee\x1e\x31\x3b\x18\x80\xc5\x54\xab\xcc\xba\x7e\x07\x5f\x69\xd5\x5d\xaa\x78\x87\xe9\x1f\x43\x0c\x6d\x7f\xdc\x3e\xca\x75\xf0\x4c\xae\xdf\xdd\xe3\x20\x66\xc3\x17\x70\x0d\x9f\xc9\xe5\xf6\xba\x3f\x8c\x59\x3c\x7a\x01\x59\x3c\x7a\x26\xdb\xdf\xbb\x2a\x8e\x5c\x6a\x2f\xc9\x6d\xf0\x4c\xba\xc3\x2e\xb9\x41\x97\xdd\x8b\xd2\x7b\x2e\xe1\x70\x95\xdf\x20\xde\xff\x2b\x6c\xcf\x00\x1e\x74\x0f\x16\x3f\x2b\x8b\x35\xac\x45\x72\x20\xba\xa1\xde\xba\xd5\x83\x38\x1e\x0c\x06\x3f\x74\x38\x88\x83\x83\x27\x1d\x86\x71\x30\x7c\xd2\x21\x1e\xc5\x41\x3c\x7a\x1a\x63\x10\x07\xc3\xc1\xff\x41\x19\x38\x98\x5d\xa7\x49\xb4\xb9\x0d\x4c\x22\x77\xa7\x71\xe3\xea\x4f\xce\xff\x02\x00\x00\xff\xff\x2b\xf0\xc6\xc8\xfc\x0c\x00\x00")

func templatesMainTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesMainTmpl,
		"templates/main.tmpl",
	)
}

func templatesMainTmpl() (*asset, error) {
	bytes, err := templatesMainTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/main.tmpl", size: 3324, mode: os.FileMode(420), modTime: time.Unix(1513248350, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesPartialsFooterTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x56\xc1\x6e\xe3\x36\x10\xbd\xfb\x2b\x08\x9e\x2b\x73\xb1\x01\x7a\x28\x14\x01\x41\xd0\x6c\x17\x0d\xe2\x43\xb6\xe8\xd1\x18\x93\x23\x69\x6a\x6a\x28\x90\x23\x29\xee\xd7\x17\xb4\xe2\x38\xad\xbd\xe8\xc6\xd9\xf6\x62\xcb\x33\x7c\x6f\x9e\x39\x8f\x43\x95\x75\x08\x82\x51\x59\x0f\x29\x5d\xeb\x3e\x12\x4b\x51\xb4\xe4\x50\x57\x8b\xb2\xfd\x78\x48\x8c\x94\x06\xf0\x7e\xd7\x92\x73\xc8\xba\xba\x9b\x61\x9e\x78\x9b\x4a\xd3\x7e\xac\x16\xa5\xa3\xf1\xb0\x7a\x26\xd5\x7f\x0f\x4e\x11\xfa\x7e\x8e\x32\x8c\xe7\x00\x05\xc3\xa8\x6c\xf0\x45\x5e\xfa\x0f\xb4\x0d\x7e\x9f\x2a\x7c\x53\x04\xc6\x42\x5a\x8a\x6e\x8e\x74\xee\x18\xd9\xab\xbe\x3a\xa5\x5d\xaf\x5b\x04\x47\xdc\xe8\xea\x17\xf4\x7d\x69\xda\xab\x6a\x51\x0e\xfe\xdc\x4a\x4f\x49\x32\x8f\xa7\x73\x59\x12\xec\x72\x16\x54\x1b\xb1\xbe\xd6\xa6\x45\xdf\x1b\xb0\x16\x53\xa2\x0d\x79\x92\x9d\xae\x6e\x5e\xff\x2c\x0d\x54\x8b\xd2\x78\x7a\x23\xa7\x0d\x61\x4b\x98\x80\x5d\x1f\x69\x04\xbb\xd3\xd5\xed\x1c\x52\xc0\x4e\x3d\x07\x2f\x65\x17\x8c\x5d\xe6\xb6\x81\x1d\x09\x05\x4e\xba\xfa\x92\x63\x7b\xf2\x63\xf4\x35\xbf\x19\x7c\xfe\x74\x34\xfe\x57\xcd\xb9\xd9\x84\x41\xd4\xea\xe1\xf1\x7b\x76\x08\x32\xe9\x90\xcc\xd4\x82\x4c\xe8\x82\xae\x7e\x6f\x41\xd4\x84\xca\x85\x0b\xb6\xef\x40\x67\x21\x22\xc6\xa4\xab\xdb\xf9\xe1\x3d\x54\x81\x05\xac\x0c\x99\x6c\x7e\x54\xc3\x25\x7c\x8c\x53\xd2\xd5\x03\x4e\xef\x11\x23\x11\x38\xf5\x10\x91\xed\x0e\xd8\x35\x61\xc4\xc8\xc0\x16\x4d\x1d\x11\x5d\xe8\x42\x4d\x5c\x87\xd8\x41\x36\x48\x1d\x48\x57\x77\x73\x42\x85\x5a\x7d\x3e\xa6\xfe\x5f\xeb\xdc\x06\x66\xb4\xa2\x26\x92\x76\xbf\x7b\xdf\xcd\x40\xad\x48\x9f\x7e\x32\x46\x26\x12\xc1\xb8\xb4\xa1\x33\xab\x87\x47\x7d\xc0\x92\x0d\xfc\x3c\x34\x95\x40\x6c\x50\xae\xf5\x7a\xe3\x81\xb7\xba\xfa\x32\x63\xde\xde\x8e\x43\xd1\x69\x9a\x96\x35\x58\xdc\x84\xb0\x7d\x53\xe5\xbb\x67\xd0\xfb\x4a\xe7\x09\x8f\x8e\x78\x5f\xda\x86\xae\x07\xde\x99\x50\xd7\x64\xb1\xa8\x43\x26\xc8\x9d\x06\x5f\x24\x01\xa1\x24\x64\xd3\x37\xa9\xbb\xdf\xf3\x7e\xe6\xcb\xd5\xf5\xc3\xc6\x93\x5d\x36\x61\x74\xe8\x69\xc4\xb8\xdb\x6b\x04\x6b\xc3\xc0\x92\xcc\x6f\xbf\xae\x1e\x1e\x4d\x1a\x36\xc9\x46\xda\x60\x4c\xf9\x74\x7c\x93\xb6\x9f\x3b\x20\xaf\xc0\x63\x94\xaf\x8f\xbf\xc3\xd7\x7c\x97\x9d\x1a\xfb\xd5\x75\x77\x7a\xd1\x79\xb2\xc8\x69\x7f\xc7\x52\xd7\x9c\x4f\xae\xd7\xd4\x35\x5a\x81\x97\x6b\xbd\xfa\x74\xaf\xd5\x44\x4e\xda\x6b\xfd\xe3\x07\xad\x52\xb4\xc7\x8d\xb0\x8e\x97\x81\x53\xde\x89\xe5\xb0\x35\x90\x12\x4a\x32\xd4\x41\x83\xc9\x84\xc6\x2f\xfb\x7c\x3c\x16\x65\xff\xb5\x3a\x82\x4f\xa2\x3a\x88\x0d\x71\xe1\xb1\x96\x22\x75\x45\xf1\x41\x57\x8b\x1b\x9f\x0f\x27\x0b\xb2\x28\x4a\x0a\x46\x20\x0f\x1b\x8f\x6a\x60\x87\x51\x49\x8b\xaa\x84\x73\x3b\x7a\xec\xd4\xb3\x8d\x0e\x36\x81\x68\x5b\x1a\xf1\x45\xac\x0b\xd6\x84\x1e\xb9\x98\xa7\x4c\x87\x2c\xb3\x2e\x8b\x66\xc4\x98\x28\xb0\xb9\x32\xa7\x2d\x5a\xf5\xc8\xea\xd3\x0b\x46\xdd\xcf\x18\x35\x5e\x2d\x3f\xe4\x96\xa9\x32\xf5\xc0\xaf\xa5\xa9\xbd\x3e\x7c\x92\x3c\xcc\xf2\xbc\xa1\xa6\xcd\xff\x14\xbc\xd7\x55\x69\xf2\xf2\xea\x07\x85\x4f\x16\x7b\x51\x53\x8b\x11\x55\x90\x16\xe3\x44\x09\x55\x36\x37\xba\x45\x69\xfa\x53\x03\x1c\x3b\x4f\x2e\xbf\x28\xe1\xd4\x87\x98\x99\x5f\xbc\xf6\x47\x2a\x8e\x61\xfa\x13\x73\xbd\x8c\x3a\x05\x75\xee\x02\x90\x6f\xfe\x1d\x64\xe6\x9e\x57\x8b\xbf\x02\x00\x00\xff\xff\x00\x2c\x01\x11\xf1\x09\x00\x00")

func templatesPartialsFooterTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesPartialsFooterTmpl,
		"templates/partials/footer.tmpl",
	)
}

func templatesPartialsFooterTmpl() (*asset, error) {
	bytes, err := templatesPartialsFooterTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/partials/footer.tmpl", size: 2545, mode: os.FileMode(420), modTime: time.Unix(1513248350, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesPartialsHeaderTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x57\x4d\x6f\xdb\x46\x13\xbe\xe7\x57\xcc\xbb\xef\x25\x3d\xac\x58\xf4\xd2\x26\xa0\x08\xb4\xae\x81\x18\x48\xd2\xc2\x6e\x0e\x45\x60\x38\x63\xee\x90\xdc\x6a\xb9\xcb\xee\xae\x64\x0b\x2a\xff\x7b\xb1\xe2\xa7\x25\x52\xb6\xd1\xa0\x4d\x74\x22\x87\xf3\xf1\xcc\xcc\xb3\xb3\xa3\xb8\x20\x14\x64\x93\x17\x00\x00\x31\x42\xaa\xd0\xb9\x25\x73\x2b\x59\x29\xa9\x57\x0c\x0a\x4b\xd9\x92\xfd\xbf\x44\xa9\x19\x78\xbc\x95\x5a\xd0\xfd\x92\x7d\xcb\x92\xab\x95\xac\xc0\x1b\x08\x9f\x20\x35\xda\x93\xf6\x71\x84\xad\x2b\x21\x37\x20\xc5\x92\x55\x98\xd3\xaf\xe8\x0b\xd6\xb9\x2e\xa4\x20\x96\xec\x76\x8b\x0f\x97\x17\x75\x1d\x47\x42\x6e\x92\x17\x83\x4d\xab\x75\x67\xb1\xaa\xc8\xb2\xc6\xdb\xe1\xd7\x06\x34\xa4\x46\xf1\xa0\x38\xd2\x3a\xd4\x4c\x8d\xda\xab\x71\x95\x73\xa3\x89\xfb\x42\x5a\xd1\x48\x4a\x31\x48\x0e\x3c\xb4\xc5\x08\xf8\x95\xc9\x0d\x1f\x97\x22\x9a\xd0\x0d\xbf\xdd\xee\x53\xfc\x3f\xce\x3f\xca\x0c\x94\x27\xb8\x38\x87\x1f\xae\x93\x4f\xf0\x17\x38\xcc\xe8\xcd\x6f\xef\xde\xd6\xf5\xa4\xdd\x3e\x96\x2c\xf3\x0e\x71\x08\xc8\xc0\xd9\x74\xc9\x0a\xef\x2b\xf7\x3a\x8a\x52\xa1\x17\x46\xbb\x45\x6e\x36\x8b\xf5\x2a\x42\xe7\xc8\xbb\x48\x96\x98\x93\x8b\x8c\x76\x3c\xd8\x44\x9b\xef\xfa\xe7\x45\xa5\x73\x06\xa8\xfc\x92\xfd\x92\x65\x32\x25\xc8\x8c\x85\xf7\xe8\xa5\xd1\xa8\xe0\xca\xa3\x97\xce\xcb\xd4\x9d\xca\xe5\x23\x69\x21\xb3\x6b\xce\x9f\x92\xc5\x28\xfb\xbc\xc9\xfe\xd5\x75\x12\x24\xff\x59\x09\xdc\xe6\x5f\x2c\x41\xcf\xfb\x41\xb0\xe7\xf5\x03\xd1\x0c\x2d\xfd\x9d\x69\x48\xe8\x7a\x5e\x8e\x44\xe1\xb8\x70\xee\x4a\xa8\xac\xd4\x9e\xf3\xe6\xf8\x1c\x03\x10\xaa\xa1\x2b\xea\x7c\x8d\x39\xf5\xc7\xad\x17\x4c\xa7\x19\x0b\x7f\xa8\x79\x73\xe3\xa5\x57\xc4\x92\xb3\x02\x75\x4e\xd0\xc9\x5f\xc7\x91\xf0\x73\x5e\xc4\xb1\x17\xe9\xa9\x9c\x89\xda\x14\x58\x66\x40\x7f\xc2\xe2\x6d\x6b\x01\x8c\x34\xab\xeb\x61\x08\x0d\xae\xc2\xf9\x03\x99\x1a\xdd\xe6\xdf\x1d\xc5\x28\xdd\x2e\x76\xbb\xc5\x95\xf4\xf4\xb3\x09\x73\xa8\xae\xbb\xc1\xc2\x92\xb3\x6d\x69\x91\x72\x78\x79\xf6\xfb\x37\xa1\x3f\xbb\x1d\x69\x71\x82\x81\x13\x80\xd2\xed\x73\x01\xcd\xa2\x39\xd7\xb9\x92\xae\x80\x97\xe7\xef\x1f\x45\x13\x47\x42\x4c\xb4\x38\x12\xea\x59\x24\x73\x94\x1a\x2d\xd0\x6e\xb9\xc6\x0d\x3c\x9d\x72\x8f\x30\x6d\xdd\x30\x4d\xe3\x86\xf7\x11\xd8\x64\xcc\x50\x27\xe7\xe1\x0f\x17\x5e\x78\xaa\x8c\xa6\x46\x34\x47\x46\x25\x67\xfc\x3c\xc2\xa5\xd1\xc5\x75\x18\x5f\xaf\xe0\x08\xc0\x68\x98\x5b\x52\x84\x8e\x52\x54\x14\xcc\x58\x72\xd9\x08\xa0\x93\x1c\x9d\xec\xa1\xf8\x4a\x7e\x31\x69\x94\xe4\x0b\x23\x8c\x32\xf9\x96\x25\xef\x86\x97\xaf\x03\xbd\xa6\x3b\x17\x60\x0b\x89\x5f\x07\x60\xbc\x35\x6b\xbf\x76\x2c\xf9\x31\x3c\x7c\x39\x98\x87\x89\x34\x89\xde\xa3\xcd\xc9\x2f\xd9\xcd\xad\xc2\xf0\x6e\x49\x2d\x99\x36\xa6\x22\x4d\x16\xb4\xb1\x94\x91\xb5\x64\xbb\x3c\xbb\xcb\xf7\x56\x99\x7c\x7c\xfb\xb2\xe4\x27\x65\xf2\x67\xa6\x1d\x47\xeb\xd3\xe3\x6b\xf4\x3a\x7e\x1c\x0d\xb4\xca\xca\xb2\x1b\x67\xd3\x43\x2a\xd6\x78\x38\x10\xd7\xaa\x33\x0f\xf5\xe0\x61\x51\xb5\x46\x4d\x5d\xfb\xa3\xa6\x3c\x50\x3d\xd9\x94\x18\xbb\x15\x39\xd8\xb4\x08\xd9\x7e\x42\x96\xa4\xd7\xdc\x9b\x3c\x57\xc4\x00\xad\xc4\xde\x63\x13\xa0\x57\x9e\x0e\x1a\xcc\x4f\x31\xc1\x55\xa8\x67\x4c\x3d\xdd\xfb\x70\xa0\xf4\x3a\x8e\x82\xda\x5c\x9f\x26\x1a\x38\xd3\xbc\x93\xa5\x81\x27\xd5\xc6\x11\xda\xb4\x68\x4a\xd3\x3c\x9f\x2a\x4e\xa7\x3d\x1d\xb5\xfd\xfa\x4f\xaa\x73\xb5\x77\xf1\x19\xea\x33\x41\xec\xe3\x3f\x32\x10\x30\x84\x9d\x60\x4f\x59\x41\x9a\xf5\x97\x68\x4f\x83\x7d\x15\xe8\xbe\x42\x2d\x48\x2c\x59\x86\xca\xcd\x5d\xc0\xc7\x07\xe2\xa9\x17\xeb\x03\x93\x7d\xf3\xfa\xbd\xe7\xc3\xe5\x05\xb0\x88\xd5\xf5\x91\x0e\xe7\x98\x7a\xb9\xa1\x76\x69\x69\x87\xcb\xb0\x9b\x1a\x25\xb8\x24\xce\x85\x74\x95\xc2\x2d\xbf\x55\x26\x5d\xcd\x76\x67\x98\x61\x07\xf8\xf5\x6a\xd8\x52\x4a\xc1\xbf\xef\xf7\x95\x57\xa3\x3f\x5f\x6f\x4c\x49\x9f\x7f\xf4\xf4\x73\xa3\xfd\x12\x47\xdd\x1f\xe3\xbf\x03\x00\x00\xff\xff\x92\xea\xe8\xf4\x22\x0f\x00\x00")

func templatesPartialsHeaderTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesPartialsHeaderTmpl,
		"templates/partials/header.tmpl",
	)
}

func templatesPartialsHeaderTmpl() (*asset, error) {
	bytes, err := templatesPartialsHeaderTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/partials/header.tmpl", size: 3874, mode: os.FileMode(420), modTime: time.Unix(1513248350, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _redirectsRedirectsCsv = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x5b\xcd\x92\xdb\x38\x0e\xbe\xfb\x59\x5a\x51\x92\x9d\xd9\xad\xda\x6b\x6a\x0f\x73\x99\x39\xe4\x01\xba\x68\x12\x92\x10\x53\x84\x1a\x24\xd5\xf1\x3e\xfd\x16\x49\xfd\xda\xb2\x2d\xdb\xea\x99\xd9\x4b\xd2\xe6\x0f\xbe\x0f\x24\x01\x82\x20\x95\x83\x24\x43\xf5\x31\x2f\x99\xac\x55\x54\x83\x75\x28\x1b\x26\xe5\xa5\x2b\x55\x93\xbf\xd5\x68\x2f\x55\xbe\xd5\xf8\x72\x53\x40\x0d\xae\x22\x45\x9a\x4a\x84\x6b\x92\x76\x79\x03\xd4\x68\x68\xa8\xf1\x5a\x38\x24\x23\x8c\x92\x54\xd7\xde\xa0\x3b\xe6\x15\x79\x8b\xa6\x4c\x7c\xc2\x0f\x68\x18\x25\x58\x27\x1c\x06\x59\xb6\x20\xb6\xb5\xd0\x5a\x30\x08\x1b\x99\xad\x92\x37\xa7\xb7\x46\xf0\x6e\x50\xd9\x24\xb1\x5a\x48\x49\xde\x38\x9b\xfb\x83\x05\xe9\x88\x87\x82\xc8\x16\x8d\x82\x9f\x54\x58\xe0\x36\x08\x9e\x8d\xda\x6d\x11\x73\x82\xe7\xb2\x76\x39\xd4\x8d\xa6\x63\x0d\xc6\x09\xa3\xb4\xd8\x93\xe7\x5a\xf0\x01\x5c\x37\x02\x68\xde\x89\x0f\x79\xaa\xe9\xc6\x1c\xdb\x30\x08\x91\xde\x79\x79\x62\xf8\xb0\xd4\x39\xe3\x45\xf1\xe3\x18\x96\xd4\x02\x9b\x00\xd3\xf8\xbd\x46\xd9\x69\x6f\x94\x13\x3f\xc1\xe6\xd3\xc2\x02\x8d\x30\x12\x12\xeb\x85\x0a\x2b\x8c\x1a\xc5\x29\x28\x50\x62\x20\xaf\x60\xef\xbc\x51\xc0\xae\x82\x5a\x08\xeb\x18\x65\xe5\x1c\x83\xe8\x55\x7d\x9c\xcb\x5c\xd7\xed\x48\x8d\x03\x84\xa6\x18\x96\x6f\x5c\x99\x68\x54\x98\xfa\x6e\x18\xe2\xb8\x02\x4f\x6b\x66\x3a\x5d\xec\x7e\xc2\x7c\x59\xce\x2e\xdf\x07\x33\x01\x6b\xd1\x28\x6f\x1d\x1f\xc3\x68\xb0\x50\x90\x33\x38\x81\xba\x2f\x4e\x6c\x52\x99\x15\x1a\x6c\x5c\xa8\x91\xc9\x5a\x09\x73\x42\xa7\xa2\xd8\xe2\x1d\xc3\x32\xda\x71\xec\x5c\x35\xf8\xc8\xa0\xf8\xc3\xa2\x9c\x47\x1d\x40\x6f\x03\x64\x9e\x72\x01\xfe\xb0\x24\xef\x0a\xa9\xbd\xd0\x61\x25\x52\xd1\x88\x68\xcf\x1d\x2b\x7f\x88\xd3\x70\x9d\xcb\x79\xdf\x53\x32\xbd\x90\x1b\xfe\x7b\x56\x5c\x63\xc9\xf1\xef\x49\x71\xd8\x11\x6a\xe1\xfa\x19\xfc\xfa\xf9\xcb\x17\x09\xc6\x7a\x3b\xf3\xc4\x60\x4a\x2d\x8c\x12\x46\xbd\x87\xa5\x51\x0b\x96\x55\x68\x7a\xdb\xe1\xdf\x41\x60\xae\xe2\xfd\x4c\x6e\x6d\x65\x20\xb4\xab\x84\x51\x96\x24\x0a\x2d\x05\x43\x2e\x85\xb7\x60\xa9\x50\x20\x5c\x95\x46\x40\x68\x49\x15\x69\x06\x2d\x1c\xa8\x58\x61\xd1\xb8\x0a\xfc\x61\xc5\xf6\x76\x13\x63\xae\xe4\x75\xb0\x2b\x6e\xa0\xaf\x18\xfe\xe8\xb7\xa6\x4e\x09\x63\xbc\xd0\x43\x9d\xe7\x16\x8e\xd7\xbd\xc2\x65\x81\x27\x8c\x97\x25\x6f\xbc\x0c\x13\x4a\x8d\xea\x08\x82\x17\x5a\x7d\xe0\xba\x5b\x01\xbd\x7a\xef\x9f\x35\x4b\x3f\x00\xdc\xb1\x99\xab\x39\x82\xa4\xe1\x14\x8d\xbd\x27\x14\xb8\x02\xb2\xa4\xd9\x32\xda\x46\xb6\xd3\x12\x2a\xb1\xd7\x50\x13\x3b\xa1\xd1\x1d\xd1\x9c\x18\xec\x07\xd8\xd0\x2a\xd0\x1b\x0a\xee\x91\x5d\x65\x93\x01\x86\x05\x23\x98\x51\x94\x60\xf3\x54\x94\xd4\x8b\x8d\x82\xc1\xb2\x73\x01\x30\xf8\x23\x34\x85\x08\x01\x46\x68\x75\x5b\xb7\x5b\x30\x73\xcd\x56\xe0\xad\x71\x11\xa2\x0b\x00\x2d\xfe\x17\xc2\x4a\x22\x99\x6c\x21\xe9\xd4\xb5\x52\x50\x53\xc9\xa2\xa9\xd6\xfa\x89\x4b\x52\x4f\x54\x58\x12\x3f\x0d\x27\x5a\xb0\x2e\xee\x70\x0d\x18\x1b\x16\x65\x44\xf2\xb6\xdf\x2e\x47\x22\x7d\xcb\x93\x90\xe2\x8a\x80\x65\x2a\x33\x49\x9b\x99\x72\x2f\x9d\xa1\x44\xeb\x80\xc7\xf6\xc9\xcc\xf6\x0c\x1f\x63\xd5\x2b\x81\x1f\x31\x70\x32\x0a\x5d\x37\xa4\x0a\x2d\x08\xdb\x6b\x2b\x43\x74\xc2\x09\x32\x79\xd6\x71\x93\x8e\x5a\x3e\x3f\x2f\xdb\xd3\x0d\xe3\x81\xad\xd0\x03\xd5\xb0\x8f\x69\x0d\x0e\xcd\xa3\x6e\x69\x11\x72\xae\xde\x2a\xec\xbb\x63\x5b\x29\x1a\x74\x41\x1c\xc9\x43\x00\xef\x7e\xcb\xe0\xd5\xeb\xe6\xe9\x60\x77\x95\xf8\x2d\x5c\xaa\xac\x50\xab\xc1\x71\x9f\x2c\xa2\x0d\x5d\xe9\x15\x9c\xd5\x2e\x60\x7a\xcc\x6c\x80\x2d\x19\x03\xba\x53\x03\x5b\xd4\x7d\xc4\x34\x57\xe2\x29\xe1\x27\x4a\x5c\x42\xd9\xcc\x8b\x49\x2d\xb0\x16\xc6\xc5\x45\xf1\x21\xee\xea\x14\xe1\xd1\x45\x24\xc9\x48\x68\xba\xf6\x05\xb0\xc3\x30\xb1\x3c\x86\x90\x63\x83\xd3\x09\xb9\x64\x13\x56\x38\xd0\x1a\x1d\x5c\x32\x8a\xb8\xfa\x81\x1d\x83\x51\x37\xac\xf6\x5c\x56\x4f\x6a\x2e\x62\x5b\x3a\xcf\xdb\x63\x7f\x04\x69\x49\xb7\x68\x4a\xa9\xc9\x3a\x46\x85\xbe\x56\x58\x14\x28\x51\xc3\xca\x30\xee\x3e\xfb\xbc\x03\xf7\x03\xd4\xac\xd9\x8a\x3f\x43\xad\x29\xce\x76\x6a\x74\xa7\x47\x47\x8a\x7d\xd9\x10\x5a\x32\x68\xca\x07\x22\xef\x47\x74\x5b\x0f\xfe\xb0\xc2\xd8\x12\xf7\xf9\xc7\xee\x87\xdd\x54\xbb\x0e\xe0\x44\xbd\x2b\x48\x57\xa2\xee\x98\x65\x44\xe3\x80\x4d\xd8\xdc\xa7\x69\x3a\x08\x5c\x80\x25\x5c\x8f\xb0\x97\x25\xcc\xc9\x4d\x45\x8d\x5e\x08\x4c\x8b\x4c\x31\xe3\x39\xf1\x23\x09\x7b\xa9\x8a\x8c\x40\x86\x1a\x6d\x8c\xcf\x66\xde\x68\x59\xd4\x09\x89\x55\x32\x9f\xa1\x07\x06\xb8\x3c\x6e\x45\x6c\x90\xf6\x1c\xa5\x49\x71\x4c\x50\x6f\x47\x6f\x41\xf2\x16\x5e\x02\x7e\x4a\xb0\xf6\x3d\x2e\xa9\x67\x0e\xe6\xf7\xb9\x87\xd5\xa8\xf7\x64\xb9\xd0\x18\x6a\x27\xa7\xd7\x82\x18\xb0\x34\x0a\x19\xa4\x1b\x0f\x10\x85\xc2\x3b\xb3\x5d\x13\xc1\x73\x35\xae\x22\x6c\x93\x35\x29\x53\x30\x47\xc6\x36\x20\x31\x6c\x76\xf1\x8c\x3f\x0c\xda\xf6\x29\x93\x5b\x88\xcb\xd3\xb4\x36\x14\x14\x1c\xb6\x80\xb0\x44\xc2\x6f\x34\x65\x45\x9e\xe7\x39\xf9\xd4\x5b\x92\x0d\xc7\x43\x0e\xf5\x77\x05\x9b\x17\x10\x16\xef\xec\x16\xa1\xd6\x5e\x68\x0c\x29\x7f\x6c\x85\x83\xee\xf2\x92\xa3\x85\xa6\xeb\xca\x47\xee\x37\xd6\x08\xbd\xe7\x5a\x36\x98\x58\x19\x22\x5f\x15\xef\x4d\x2a\xd2\x2a\x1d\xfd\x71\x4d\x4a\x6a\xf9\x46\xf6\xa2\xcc\x87\x73\xbd\xc3\x5f\xef\xe8\xaa\x2e\xc7\x3d\xf2\x67\x23\xf4\xd0\xe2\xf9\x34\xef\x22\xd8\xb9\x86\x17\x51\x27\x3b\x44\xfc\x1f\x25\x79\xd7\x78\x17\x27\x76\x72\xe7\x3a\xfd\x51\x83\xb0\x9e\xc7\x95\x13\xc5\xa7\x98\x5e\x52\xdd\x08\x0e\xb1\x91\x1d\xef\x8f\xc6\x2b\xdf\xc7\xa1\x16\x55\xba\x8d\x79\x63\x06\x35\x60\x10\x1f\xbc\x25\x79\x46\x5b\x2f\xa8\xd4\x08\x6b\xc1\x94\x29\xb3\x01\x47\x6c\x56\x4c\xd6\xb9\xdc\x2b\xfc\x4f\x00\x6e\x44\x5d\xd3\x9e\xa9\xe8\x9c\x72\x2c\x47\x33\x7b\x10\x70\x8f\xc4\x2b\x64\xcf\x45\xaf\x30\xe0\x68\x55\xb2\x12\x2c\xa4\x03\x4e\x27\xd4\xbc\xa2\x1a\xfa\xe0\x6f\x70\xe8\x35\x28\x14\xde\x8a\x72\xa6\x16\x38\x21\xc3\x1e\x3b\x5a\x68\x0c\x1c\x15\xb6\xa8\xbc\xd0\x2b\x6d\xff\x21\x16\x4b\x43\xb1\x82\xce\xea\x3d\x24\xfc\xd3\x68\x21\x41\xa1\x6d\xbc\x83\x89\xab\x1f\x73\x6e\xd3\x77\x14\x7d\xbb\xc9\x55\xd7\xa6\x50\x4b\x8f\x2b\xce\x30\x37\x4b\xc3\xa4\x9e\x45\x3c\x7f\x44\xe1\xba\xf8\x98\xe4\xf1\x05\xa0\x8d\x15\x49\x3d\x3f\x50\x81\x01\xe0\x81\x80\x4c\xa1\x15\xfb\x98\x3c\xea\x28\x63\x01\xd4\x34\xc4\x2e\xf4\xc1\xd9\x92\x7a\x46\xf8\x09\xef\x8b\x28\xb7\x76\xd7\x98\x1b\x14\x21\x38\x1b\x0c\xad\x7f\xf3\x92\xa3\x91\x54\x07\xff\xfa\x1e\xb9\xf4\x0a\xc5\x8c\x4a\x08\x80\x84\x51\x05\x91\x5a\xab\xd1\x5d\x58\xa7\xfa\x5d\x04\x7d\xf6\xaa\xb8\x61\xfa\x01\x72\xe2\x01\x6a\xc1\x29\x59\x2d\x9c\xb7\x8b\xed\x9e\xbd\x2b\x9e\x42\xce\xf5\xbc\x8d\x7d\x7d\xe3\x92\x95\x30\x21\x1a\xa7\xe1\x44\x52\x03\x87\x4d\x2f\x04\xdf\xf2\xcd\xa3\x9d\xfa\xba\xe5\xba\x5a\x5c\xdf\xca\x56\x63\x9c\xa8\x76\x19\xec\xf1\x87\x09\x35\x19\x57\xe9\xe3\x47\xbc\x4c\xb8\x24\x7a\x33\x67\x16\x1a\x45\x57\xf9\x83\xf6\x1f\xe3\x8e\x4f\x11\x1e\x33\x95\x59\x64\x32\x16\x47\x1d\x6c\xbc\x4c\x06\xae\x97\x1b\x0d\x31\x70\x41\xac\x49\x0a\x2d\xbc\xab\x88\xa3\x87\x7a\x38\x18\xbf\xc0\x67\xae\xfa\x1d\xc4\xee\x4e\x2d\x7e\xec\x38\xdd\x9d\x47\xf9\x0b\x87\xe9\xf9\x9b\xae\x69\xd5\x28\x6d\xeb\x9b\xae\x8b\x28\xff\x67\x9e\xe7\xee\x80\x61\x28\x0b\x01\x02\xfc\x6c\x40\x3a\x61\x24\xf6\x4a\x50\x83\xa6\xbb\x6e\x0e\x0d\xac\x3b\x6a\x78\x26\x3e\xb9\x0c\x37\xd7\xf0\x2a\xee\xc3\x6a\x86\x3f\xed\xd1\x3a\xe8\x0e\x98\xfd\x72\xde\x03\x83\x68\x61\x7c\xc3\xd0\xd2\x70\x60\x7b\x1e\x6a\xae\xd9\x6d\xcc\xcd\xf6\x8f\x56\x48\x61\xe4\xf1\xfe\x23\xca\xea\xfd\xe3\x14\x61\x75\x1a\xa3\x33\xb7\xb4\xba\xcf\xdf\xaa\xcf\xea\xa7\xd5\x83\xbf\x01\xe5\xd3\xfb\x9f\x7b\x32\x1a\x97\x51\x97\xbd\xc1\x3a\xf8\xb5\xb9\xbd\x39\xf0\xf3\x6f\xb9\x2f\xc8\xeb\xfa\xfe\xa9\xb3\xe1\xc8\x09\x3d\x17\xf3\xd7\xcc\xcb\x32\x91\xcd\x8f\x38\xae\x02\x28\x0a\x90\xce\x52\x11\x6f\x4e\x84\x51\x7b\x30\x50\xa0\xb3\x64\x06\x21\xa9\xef\x87\x9e\x7b\xae\x33\xe9\x5e\xb6\x8f\x5c\x6e\x8c\x04\xe8\xfe\xd4\x11\xff\x22\x16\x7a\xfa\xe6\xaa\x7f\x50\x3e\x54\xde\xf3\x7e\xe6\x96\xec\xd3\x07\xe7\xcb\x20\xf7\x3c\xe5\xab\xdf\x7e\x1d\x1b\xed\x8f\x68\xac\xe7\x30\xa6\x31\x53\x69\x10\x86\x7e\x85\x37\x6a\xec\xfc\xe8\x7b\xbf\xc7\xe1\x46\xa5\xc2\x88\xcc\x1e\x65\xa4\x0f\xa9\xd0\x36\x64\xc5\x5e\xc3\xc9\x74\xf6\x9f\x63\xa4\x4e\xd7\xdb\xce\xd4\xba\x17\xe7\xf4\xa3\x8d\x95\x80\xf7\xcc\x16\xbd\x1b\x60\x5b\x61\x43\x85\x3f\xbc\x79\x72\xa0\x6c\x25\xf8\xcc\x37\xae\x9e\x8f\x6b\x02\x9f\x7e\xc6\xea\x0f\x7d\xbb\x4d\x9f\xaf\x8e\x62\x2f\xb4\x7f\xf2\xfe\xd2\x1f\x66\x5f\xe6\x85\x08\xcc\x28\x74\x9e\x81\x0c\x83\x05\xc1\xe1\xf0\xae\x14\xb4\xa0\xa9\x19\x22\xed\x0d\xee\x35\x1f\x42\xde\xcd\x3f\x3b\x6c\x85\xf6\x20\x94\x02\x55\xb6\x62\x61\xed\x8f\xf5\xe3\x37\x33\xa2\x69\x98\x84\xac\xce\x3f\x62\x9c\x4b\xbb\xb2\xc2\x6f\x88\xdd\xad\xfb\xb8\xf2\xcd\x0b\x76\xc0\xfa\x38\xcd\x6e\xa4\x67\x9c\xc1\x7e\xb4\xa5\xa2\x7b\xf5\x28\xac\x05\x37\xcb\x2c\xdf\xf5\xed\xe5\xc3\x40\xbb\x3c\xfd\x69\xf3\x12\x5c\x66\x5d\x10\xa3\x5e\x2a\xe7\x1a\xfb\xef\xbc\xaf\xfb\x44\xc6\x7e\x2a\xa9\xfd\xe4\x0f\xbb\xbc\x45\xeb\x85\x46\x9b\xee\x76\xf3\xb0\x00\x24\x19\x07\xc6\xe5\x0a\xda\x7f\xe4\x7f\xfc\xfe\xfd\xf5\xfb\x1f\xdf\x5e\x49\x4a\xdf\xc4\x46\xaf\x92\x14\x9a\xf2\xd5\x11\xe9\x4f\x95\xab\xf5\x20\x9f\x8c\x55\x58\x06\x5a\x9f\x4a\x74\x95\xdf\x7f\x42\xca\x55\x93\xc9\x40\x13\x0b\x4c\xeb\x3f\x0b\x1d\x6d\x6e\x43\x10\x21\x58\x65\xa3\x64\xa1\x4f\x9a\xae\x41\x5f\xa7\xc2\xef\xdf\xbf\xff\xe7\xdb\xab\x42\x2b\xa9\x05\x3e\xfe\x69\xe4\x2f\xe2\x3e\xf9\x4a\xe3\xc4\x32\x2d\x70\xe9\x51\x81\xa3\x85\x07\xab\x9b\xbe\xd1\xb8\x0e\xf5\xc3\xeb\xe3\xd7\xcf\x5f\xfe\x75\x75\x52\xbe\xff\xf6\x2d\x4d\xec\x6f\xdf\x5e\x2b\x04\x0e\x7e\xe3\xf8\xda\x22\xbc\x6f\x31\x23\x9d\x83\xc3\x0b\x8b\x69\x19\x73\x97\x33\xe8\xee\x29\x76\x3a\x99\x17\x58\x7a\x06\x4b\xa6\x7f\xb3\xd6\x3d\x95\x57\xfb\xa3\x60\x10\x54\xf8\xa0\x1f\x83\x45\x05\x46\xc2\x49\xe6\xe4\xeb\xe7\x2f\xbf\xa6\x6f\x67\x5e\x36\x96\xfc\xc3\x1b\xf8\xfa\xf9\xcb\x3f\x1b\xa6\x16\x6d\x5c\x77\xbb\x3c\x7d\xf4\x96\x57\xa0\x9b\xdc\x01\xd7\xf1\xfd\xf5\x70\x13\xf5\xd2\xd7\xa7\xff\x1c\x0b\x63\x0b\xe2\xba\x4f\x3b\x97\x2c\xea\x1a\x16\x3a\xce\x05\x4b\xa2\x03\xc6\xe8\x34\x3e\x43\x90\xc7\x95\x72\xbb\xd6\x61\x91\xc4\xb3\xf2\x4c\xea\x4a\x19\x61\x4d\x75\x2d\xc0\xba\xdd\x70\x1a\x91\x64\x8a\x97\x3c\x9c\x82\x9d\xb7\xf9\x7b\x25\xdc\x3b\x28\xca\xa1\x8d\x1f\x40\x4e\x5b\x01\x87\xc1\xdc\xe5\xc6\x56\xef\x2f\x83\x97\x44\x33\x80\x15\xc4\x97\xae\x21\xc7\x4b\xcf\x79\x79\x27\xa4\xcf\x07\x58\xe7\xd5\x91\x8a\x21\xa3\xf0\x0e\x5a\xef\x01\x4d\x19\x5f\x92\x33\x04\xb3\x3b\x92\x37\x65\x32\xc7\x47\xfb\x25\x1d\x2a\xfe\x3b\x6a\x11\xe0\xcf\x3a\xce\xe9\x4d\xfa\xa6\x1b\x1f\x34\x95\xd0\x61\x4b\x1b\x96\xfc\x2e\x98\x7e\xb0\x9b\xf1\x51\xc6\x97\xf1\xcf\x6c\x74\x37\xd9\xb0\x51\x66\x0c\x0d\xb1\xcb\x0b\xd8\xb3\x17\x7c\xcc\x82\x8d\xe4\xd6\xed\xb3\xda\xbe\x71\x36\x2b\x4e\x5e\x66\xcb\x84\xeb\xf0\x5d\xc6\xc8\x72\x72\xe4\xe9\x39\x9e\x50\x0c\x54\x36\xd3\x34\x3e\x10\xfa\x5b\x2a\xa6\x6b\x9b\xc7\x39\xcf\x12\xaf\xcc\x90\xcb\xd0\x64\x43\xfe\x25\xcb\xc6\x64\x55\x46\x9c\x39\x16\x68\xd0\x94\x59\x66\x00\x9c\xcd\xce\xe7\xb4\xd3\xf3\x56\x2a\xcc\x90\xeb\xb2\x61\xde\x8c\x6d\x27\x3a\x4d\x96\x62\x6c\x3b\x50\x1a\x5b\x13\xf7\x74\x02\x99\x0b\x2a\x1a\x21\xbe\xe6\x36\x78\x1b\x95\xf5\x79\x8c\x8c\x8a\x2c\x06\x8e\xbf\x04\xda\xbf\xc6\xa5\xf8\x35\xd5\x87\xf2\xac\x2b\xef\x55\xb9\x11\x1e\x8e\x9c\x13\x4c\x8f\x42\x45\xc4\x48\x13\xf1\x0b\x49\xe7\x48\x81\x0c\x82\x47\x76\x7b\x6f\xb3\x74\xe6\x1a\x42\xfc\x6c\x3c\x83\xf5\x04\xb3\xc9\x86\x92\x31\x58\xaf\x9d\x9d\xad\xaa\xf5\x14\xcf\xbf\xce\x5a\x64\x38\x01\xec\xf0\xe6\x8b\x26\xcd\x68\x96\xa6\x74\x62\x0a\x57\x56\xfe\xf3\xc9\xd1\x51\x0b\x7f\x98\x09\x38\x9f\x79\xd1\x60\x97\xa0\xca\xdf\x61\x3f\xfd\x59\x51\x0d\x2f\x69\xd3\x8c\x37\x62\x23\xf7\xd4\x73\x4a\xf9\x7f\x01\x00\x00\xff\xff\xa0\xe2\x62\xc1\xd7\x45\x00\x00")

func redirectsRedirectsCsvBytes() ([]byte, error) {
	return bindataRead(
		_redirectsRedirectsCsv,
		"redirects/redirects.csv",
	)
}

func redirectsRedirectsCsv() (*asset, error) {
	bytes, err := redirectsRedirectsCsvBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "redirects/redirects.csv", size: 17879, mode: os.FileMode(420), modTime: time.Unix(1513248811, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"templates/error.tmpl": templatesErrorTmpl,
	"templates/main.tmpl": templatesMainTmpl,
	"templates/partials/footer.tmpl": templatesPartialsFooterTmpl,
	"templates/partials/header.tmpl": templatesPartialsHeaderTmpl,
	"redirects/redirects.csv": redirectsRedirectsCsv,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"redirects": &bintree{nil, map[string]*bintree{
		"redirects.csv": &bintree{redirectsRedirectsCsv, map[string]*bintree{}},
	}},
	"templates": &bintree{nil, map[string]*bintree{
		"error.tmpl": &bintree{templatesErrorTmpl, map[string]*bintree{}},
		"main.tmpl": &bintree{templatesMainTmpl, map[string]*bintree{}},
		"partials": &bintree{nil, map[string]*bintree{
			"footer.tmpl": &bintree{templatesPartialsFooterTmpl, map[string]*bintree{}},
			"header.tmpl": &bintree{templatesPartialsHeaderTmpl, map[string]*bintree{}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

