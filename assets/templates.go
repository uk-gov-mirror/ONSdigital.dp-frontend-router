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

	info := bindataFileInfo{name: "templates/error.tmpl", size: 684, mode: os.FileMode(420), modTime: time.Unix(1519374285, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesMainTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x57\x6f\x4f\xdb\x4c\x12\x7f\xfd\xf0\x29\xe6\x71\xa5\x73\x52\xbc\x76\x5c\x8e\x5c\x20\x31\x12\x6d\x51\x0f\xa9\x2d\xe8\x4a\x75\x77\x42\x48\xdd\xd8\x63\x7b\xdb\xf5\xae\xbb\x3b\x4e\x1a\x05\x7f\xf7\xd3\x3a\x09\x04\x4a\x39\x1e\xa1\xbe\xb1\xbd\xde\x99\xdf\x9f\xd9\xf1\x66\x33\xf9\xf3\xed\xd9\x9b\x8b\xff\x9e\x9f\x40\x49\x95\x3c\xda\x99\x6c\x6e\xc8\xb3\xa3\x1d\x00\x80\x09\x09\x92\x78\x74\xaa\x08\x8d\xe2\x12\x2c\x9a\x19\x1a\x40\x63\xb4\x01\x06\x67\x79\x2e\x52\x84\x5c\x1b\xf8\xc8\x49\x68\x17\xf2\x89\x38\x09\x4b\x22\xb5\x93\x68\x95\xbd\x42\xaa\x90\x38\x94\x44\x35\xc3\xef\x8d\x98\x25\xde\x7f\xd8\xe7\x63\xf6\x46\x57\x35\x27\x31\x95\xe8\x41\xaa\x15\xa1\xa2\xc4\x3b\x3d\x49\x30\x2b\x30\x48\x4b\xa3\x2b\x4c\x62\x0f\xa2\x6d\x90\xb4\xe4\xc6\x22\x25\x5e\x43\x39\x1b\x79\x77\xe7\x36\x18\x73\x91\x51\x99\x64\x38\x13\x29\xb2\x6e\x10\x08\x25\x48\x70\xc9\x6c\xca\x25\x26\x71\x38\x08\x1a\x8b\xa6\x1b\xf2\xa9\xec\x78\x14\xaf\x30\xf1\x66\x02\xe7\xb5\x36\xe4\x6d\x23\xaf\xa6\x72\x6d\x2a\x4e\x2c\x43\xc2\xd4\x19\xde\x52\x4d\x28\xb1\x2e\xb5\xc2\x44\xe9\x07\x32\xa9\xc4\x0a\x59\xaa\xa5\x36\x5b\x49\x2f\xf6\x47\xfb\x07\xfb\xaf\x1f\x88\xe7\x75\x2d\x91\x55\x7a\x2a\x24\xb2\x39\x4e\x19\xaf\x6b\x66\x89\x53\x63\xd9\x94\x1b\x66\x69\x71\xa7\x68\xbf\x46\xd2\xca\x1e\xba\x85\x13\xe9\x76\x7c\x56\xb3\xdc\x74\x83\x8c\x19\xdd\x10\x9a\x4d\xaa\x14\xea\x1b\x18\x94\x89\x67\x4b\x6d\x28\x6d\x08\x44\xea\xac\x96\x06\xf3\xc4\x73\x6b\x68\x0f\xa3\x28\xcd\x54\xa8\x95\x0d\x0b\x3d\x0b\x9b\x6f\x11\xb7\x16\xc9\x46\xa2\xe2\x05\xda\x28\xe7\x33\x97\x13\x8a\x54\x7b\x40\x8b\x1a\x13\xaf\x9b\x89\x7e\xb0\x0e\xeb\x68\x67\xe7\x8f\x3f\x96\xcb\x2f\x93\x3f\x19\xbb\x14\x39\x48\x42\x38\x3d\x81\xd1\xd5\xd1\x17\xb8\x06\xcb\x73\xfc\xe7\xc5\x87\xf7\x6d\xdb\x29\xba\xab\xc9\xf9\xb6\x25\x22\x6d\x04\x2d\x97\xe1\x39\x27\xd7\xa1\xef\xc5\xd4\x70\xb3\x38\xee\xa4\x9c\x73\x2a\xdb\x36\x4a\xad\x8d\xb4\xcc\x98\xc0\x30\xb5\xd6\x3b\xda\xf0\x5e\xa2\xca\x44\x7e\xc5\xd8\x43\x8c\x5b\xca\x4e\x4f\xe0\xe0\xf7\xa8\x12\xc8\x0e\xd6\x9a\x6e\x38\x7f\xa9\xea\x4e\xb5\x0a\x5a\xcb\x72\x2f\x7e\x8b\xb6\x8a\x0b\x75\x5f\x1b\x63\x8f\xe8\x9b\x44\xab\x5d\x63\x32\xd5\xd9\x02\x52\xc9\xad\x75\x3c\x10\x5e\x2c\x6a\x84\xb6\x75\x2b\xbe\x42\x02\xc2\xaa\x96\x9c\x10\xbc\x9a\x1b\xf7\x45\xda\x2e\x17\x8d\x07\x21\xb4\xed\xce\xba\x81\xb9\x50\x20\xb2\xc4\x73\x0f\x1e\x18\x2d\x71\xf3\x4c\x7c\x2a\x54\x86\x3f\x12\x8f\xc5\x6b\x81\x2b\xe8\x85\x40\x99\xad\x4b\x30\xe9\x3c\x3c\xca\x9a\x6b\x4d\xf7\x58\x6d\x6a\x44\x4d\x1b\xcc\x5e\xde\xa8\xee\x33\xef\x89\xc0\x06\x3a\x28\x02\x13\xf0\xa0\xea\x2f\xc5\xa5\xff\x4e\xeb\x42\xe2\xb1\xe2\x72\xe1\x36\xbb\xb3\xe9\x57\x4c\xc9\xbf\x4a\xcc\x58\x5c\x9a\xab\xc4\x5d\xae\xaf\x6f\xf2\xfb\xcb\x0d\xa4\x9b\x08\xbf\x27\xab\xdb\xf5\xf5\xe5\x55\x3f\xac\x1b\x5b\xf6\xb8\x29\x9a\x0a\x15\xd9\x7e\x1b\x74\x93\x32\x89\x5f\x2a\x9c\xc3\x5b\x4e\xd8\xeb\x8f\x79\x62\xc3\xd4\x20\x27\x3c\x91\xe8\x02\x7b\xba\x1f\xac\x41\xab\xc4\x86\x05\xd2\x7a\xc2\xbe\x5e\x5c\xf0\xe2\x23\xaf\xb0\xa7\xfb\x97\x83\xab\x31\x0f\xb9\x5d\xa8\x34\x89\xc7\x3c\xb4\x26\x4d\x8a\x71\x15\xd6\xdc\xa0\xa2\x8f\x3a\xc3\x50\x28\x8b\x86\x5e\x63\xae\x0d\xf6\x9c\xbd\x35\x6a\xdb\xef\xcd\x85\xca\xf4\x3c\xc8\x74\xda\x69\x0b\xfc\x55\x7d\xfc\xc0\x8f\xa2\xf9\x7c\x1e\x16\x5d\x11\x18\xdf\x54\x21\x4c\x75\x15\xdd\x8e\xbe\x5a\x3f\xf0\x0b\xee\xf7\xc7\x6b\xc8\x82\xf7\xfc\x95\x09\x3f\x00\xff\xf3\x31\xdb\x1f\x8e\x0e\x5e\x0d\xf6\xfe\xc1\xf6\xfc\x00\x96\x3e\x97\x52\xcf\x8f\x55\x5a\x6a\xe3\x1f\x02\x99\x06\xdb\x3b\xb9\x16\x55\xe6\x32\x6b\x5e\xa0\xdb\xa5\xbb\x24\x37\xf0\x0f\x41\xea\xb4\xfb\x0d\x0a\x6b\x4e\xa5\xdb\xf9\x60\x17\x0a\xa4\x4f\xc8\x4d\x5a\xf6\xfa\xb0\x7b\x1b\x51\x72\x5b\xde\x02\x6f\x16\x69\x3b\x7a\x79\xd3\x55\xae\x79\xa2\x97\x70\x71\xf6\xf6\x0c\x18\xfc\xbb\x44\x05\xb6\x0b\x02\x61\xa1\xd2\x33\xcc\x80\x34\x18\x54\x19\x1a\x34\x30\x47\x5f\x4a\x50\xb8\x7a\xcd\xb3\x6c\x13\x4d\x68\x2a\x10\x8a\x34\x38\xbd\xf0\xee\x18\x0c\xda\x5a\x2b\x8b\x5b\x54\x51\x04\x22\xef\xfd\xec\x24\x49\x12\xf0\xa3\x15\x92\x7f\x47\x5c\x14\x75\xb7\x19\x37\xa0\x9a\x6a\x8a\xe6\x2c\xff\x17\xda\x46\x92\x85\x04\x96\xcb\x17\x22\x77\x3c\x8d\xa4\xf0\xde\x74\xdb\xc2\x72\xf9\xc8\x14\x4a\x8b\x6d\x0b\x03\x67\x5f\xe4\x6d\x3b\xfe\x99\xd4\x20\x35\x46\xdd\x96\x75\xed\x74\x17\xfc\xbf\xdd\x43\x4c\x7c\xd8\xbd\xaf\xef\x1e\x60\x0b\x8e\x11\x1e\xf0\xf6\x30\xcd\x4f\xe9\x2f\xa3\x9b\xcd\xef\x36\x6b\xd3\xca\x3b\x3b\x37\x91\x1f\x90\xdb\xc6\x20\x90\xa8\x10\xb4\x5a\x2d\x07\x83\xd4\x60\x26\xc8\x2d\x9a\xfb\x81\x3b\x8c\xa2\x12\x65\x1d\xde\xf4\xb2\x3b\x90\x74\xdd\xbd\x6a\xf9\x9b\xf7\x51\xd5\xc1\x09\x55\x30\x07\xc8\xb4\x62\x53\xdd\xa8\x14\x99\xc3\x8d\xee\xb7\x98\x8b\x31\x71\xdc\xeb\x2f\xb7\xdb\x19\x67\xa8\xc8\x3d\x5c\x88\x0a\xcf\xd4\xb9\x6b\xe8\x00\xfc\xb8\xbb\xc4\x6c\x6f\x00\x16\x53\xad\x32\xeb\xfa\x1d\x7c\xa5\x55\x77\x22\xe3\x1d\xa6\x7f\x08\x31\xb4\xfd\x71\xfb\x20\xd7\xde\x13\xb9\x5e\xb9\xcb\x5e\xcc\x86\xcf\xe0\x1a\x3e\x91\xcb\x7d\xeb\xfe\x30\x66\xf1\xe8\x19\x64\xf1\xe8\x89\x6c\x7f\xef\xaa\x38\x72\xd6\x9e\xe3\x6d\xf0\x44\xba\xfd\xce\xdc\xa0\x73\xf7\x2c\x7b\x4f\x25\x1c\xae\xfc\x0d\xe2\xdd\xbf\xc2\xf6\x04\xe0\x41\x77\x61\xf1\x93\x5c\xac\x61\x2d\x92\x03\xd1\x0d\xf5\xd6\xad\x1e\xc4\xf1\x60\x30\xf8\x65\xc0\x5e\x1c\xec\x3d\x1a\x30\x8c\x83\xe1\xa3\x01\xf1\x28\x0e\xe2\xd1\xe3\x18\x83\x38\x18\x0e\xfe\x0f\xca\xc0\xc1\x6c\x07\x4d\xa2\xcd\x69\x60\x12\xb9\x53\x8d\xbb\xaf\xfe\x21\xfd\x2f\x00\x00\xff\xff\xee\x50\x5e\x8c\x39\x0d\x00\x00")

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

	info := bindataFileInfo{name: "templates/main.tmpl", size: 3385, mode: os.FileMode(420), modTime: time.Unix(1554396369, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesPartialsFooterTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x57\x4d\x6f\xe3\x36\x10\xbd\xef\xaf\x20\x78\xae\xcc\x60\x03\xf4\x50\xc8\x02\x82\xb4\xd9\x2e\x1a\x24\x87\xa4\xe8\x31\x18\x93\x23\x71\x6a\x8a\x14\xc8\x91\x14\x77\xb1\xff\xbd\xa0\x15\xc7\x8e\xed\x60\x3f\xd0\x36\x46\x6f\xe4\xcc\xbc\xe1\x1b\xf3\x71\x34\x2e\xeb\x10\x18\xa3\xd0\x0e\x52\x9a\xcb\x2e\x92\xe7\xa2\xb0\x64\x50\x56\xef\x84\x10\xa5\x7d\xbf\xf1\x0d\x94\x7a\x70\x6e\x65\xc9\x18\xf4\xb2\xba\x9a\x90\x8e\xfc\x32\x95\xca\xbe\x9f\xe2\x0d\x0d\x1b\xc0\x94\x7a\xca\xb3\xe7\x1a\x23\x74\xdd\x8e\x2f\xbb\x3d\x0c\x3b\xdb\xa3\xb9\x0a\x0f\x83\xd0\xc1\x15\x19\x2f\x5f\x46\xef\x01\x74\x70\xeb\xc8\xc2\x35\x45\xf0\x58\xb0\xa5\x68\x26\x4b\x6b\xb6\x96\xc3\x24\xeb\xa2\xcf\x0f\xcf\x7d\x78\xb0\x08\x86\x7c\x23\xab\x5f\xd1\x75\xa5\xb2\xe7\x47\xb1\xbd\x3b\x86\x75\x94\xf8\xe8\x59\x19\xe2\xe8\x18\x84\x18\xdb\xd7\x20\x19\x05\xc2\x46\xac\xe7\xf2\xd3\xa7\xd9\x3d\x3c\x06\x1f\xda\xd5\xcf\xa1\x05\xf2\x9f\x3f\x2b\x8b\xae\x53\xa0\x35\xa6\x44\x0b\x72\xc4\x2b\x59\x5d\xec\x6e\x4b\x05\xaf\xb1\x51\x8e\xfe\x63\xa2\x3a\x84\x25\x61\x02\x6f\xba\x48\x03\xe8\x95\xac\x2e\x27\x93\x00\x6f\xc4\x93\xf1\xa4\x28\x33\xc6\x36\x13\xd6\xc1\x1b\x62\x0a\x3e\xc9\xea\x3e\xdb\xd6\x8c\xb7\xd6\x6f\x27\x5d\xaa\xde\x1d\x2a\x5b\x19\x1a\xde\x4c\xef\x17\x8b\xd0\xb3\xb8\xbd\xb9\x3b\x6d\xd1\x43\xa6\xd9\x27\x35\x5a\xe0\x11\x4d\x90\xd5\x1f\x16\x58\x8c\x28\x4c\x38\x15\xf1\x6c\x38\x6a\x88\x88\x31\xc9\xea\x72\x5a\x9c\x1c\xbf\xe0\x19\x34\xf7\x99\xe1\xb4\x14\xfd\xc9\x90\xf4\x38\x26\x59\xdd\xe0\x78\x32\x8c\x36\x3f\x1b\x47\xf0\xa9\x83\x88\x5e\xaf\xc0\x9b\x26\x0c\x18\x3d\x78\x8d\xaa\x8e\x88\x26\xb4\xa1\x26\x5f\x87\xd8\x42\xee\x0e\x75\x20\x59\x5d\x4d\x0e\x11\x6a\xf1\x71\xeb\xfa\x5f\xf4\x8d\xcb\xe0\x3d\x6a\x16\x23\xb1\x5d\xcb\xe7\x54\xba\x87\x65\xee\xd2\x4f\x4a\xf1\x48\xcc\x18\x67\x3a\xb4\xea\xf6\xe6\x4e\x6e\x72\x92\x0e\xfe\x69\x0a\x12\x0c\xb1\x41\x9e\xcb\x87\x85\x03\xbf\x94\xd5\xfd\x84\x79\x0b\xe5\x6d\x68\x8f\xe3\x38\xab\x41\xe3\x22\x84\xe5\x37\x71\xbf\x7a\x02\xbd\x35\xf9\x3c\x33\xa2\x21\xbf\x26\xaf\x43\xdb\x81\x5f\xa9\x50\xd7\xa4\xb1\xa8\x43\x4e\x9c\xdf\x00\xb8\x22\x31\x30\x25\x26\x9d\xbe\xaa\xbe\xeb\x75\xde\x8f\xdf\xf1\x7a\xfe\xb9\xfa\xba\x7e\xe1\x48\xcf\x9a\x30\x18\x74\x34\x60\x5c\xad\xab\x04\xad\x43\xef\x39\xa9\xdf\x7f\xbb\xbd\xb9\x53\xa9\x5f\x24\x1d\x69\x81\x31\xe5\x76\xf6\x55\xd5\xfd\xd2\x02\x39\x01\x0e\x23\xff\x8b\x73\xc5\xbe\xa9\x54\x3b\x83\xf9\x0b\xe7\x97\x06\xfa\xc3\xf1\xdd\x91\x46\x9f\x70\xef\xe7\x2c\xa9\x6d\x8e\x07\x3e\x3c\x50\xdb\x48\x01\x8e\xe7\xf2\xf6\xc3\xb5\x14\x23\x19\xb6\x73\xf9\xe3\x99\x14\x29\xea\xb9\x54\xd4\x36\x2a\x34\x6e\xd6\xe5\x7e\xf3\x32\x6b\xf7\x5a\x4e\xc6\x47\x16\x2d\xc4\x86\x7c\xe1\xb0\xe6\x22\xb5\x45\x71\x76\x78\xc7\x17\x2e\xb7\x46\xcf\xe8\x59\x50\x12\x30\x00\x39\x58\x38\x14\xbd\x37\x18\x05\x5b\xcc\xb7\x7f\xe4\xe2\xb6\x82\x78\xd2\xfb\x46\xcf\x10\xb5\xa5\x01\x53\x56\xc7\xac\x5f\x2a\x13\xb4\x0a\x1d\xfa\x62\xfa\x50\xb4\xe8\x79\xa2\xa9\x51\x0d\x18\x13\x05\xaf\xce\xd5\xa1\x12\x6e\x3b\xf4\xe2\xc3\x33\x46\x5c\x4f\x18\x31\x9c\xcf\xce\xb2\x32\x44\x99\x3a\xf0\xbb\xd4\xc4\x9a\x1f\x3e\x72\xfe\x1e\xe5\x6e\x4f\x8d\xcd\x85\x83\x73\xb2\x2a\x55\x0e\xaf\x7e\x10\xf8\xa8\xb1\x63\x31\x5a\x8c\x28\x02\x5b\x8c\x23\x25\x14\xf9\x15\xa2\xd9\x53\x49\xf7\x42\x23\xbb\xb2\x78\xde\xec\xac\xb2\x16\xc8\xe4\x3f\x91\x38\x76\x21\xe6\xa3\x9f\x35\xff\x67\x2a\xb6\x66\xfa\x0b\x33\xa1\xd7\x81\xad\xf9\x4e\xa0\x6b\xbe\x0c\x2c\xd5\x24\x96\xea\xdd\xdf\x01\x00\x00\xff\xff\x67\x15\x8a\xeb\x19\x0f\x00\x00")

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

	info := bindataFileInfo{name: "templates/partials/footer.tmpl", size: 3865, mode: os.FileMode(420), modTime: time.Unix(1554396369, 0)}
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

	info := bindataFileInfo{name: "templates/partials/header.tmpl", size: 3874, mode: os.FileMode(420), modTime: time.Unix(1572260934, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _redirectsRedirectsCsv = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x5c\xcd\x92\xdb\x38\x0e\xbe\xfb\x59\xda\x51\x92\x99\xd9\xad\xda\xdb\x6e\x6a\x0f\x73\x99\x39\xe4\x01\x52\x30\x09\x49\x88\x29\x42\xcd\x1f\x3b\x9e\xa7\xdf\x22\xa9\x5f\x5b\x96\x65\x5b\x9d\xe9\xbd\x24\x96\x48\xe2\xfb\xc0\x1f\x10\x04\xa1\xce\x50\xb0\xe6\xea\x94\x15\x86\xad\x95\x5c\xa1\x75\x24\x6a\xc3\xd2\x0b\x57\xc8\x3a\x7b\xad\xc8\x5e\x2b\x7c\xad\xe8\xe5\xa6\x80\x0a\x5d\xc9\x92\x15\x17\x84\x73\x92\x36\x59\x8d\x5c\x2b\xac\xb9\xf6\x0a\x1c\xb1\x06\x2d\x05\x57\x95\xd7\xe4\x4e\x59\xc9\xde\x92\x2e\x12\x9f\xf0\x80\xb5\x21\x81\xd6\x81\xa3\x20\xcb\xe6\x6c\x6c\x05\x4a\x81\x41\xb0\x91\xd9\x22\x79\x63\x7a\x4b\x04\x6f\x3a\x95\x75\x12\xab\x40\x08\xf6\xda\xd9\xcc\xef\x2d\x0a\xc7\xa6\x7b\x11\xd9\x92\x96\xf8\x83\x73\x8b\xe6\x10\x04\x8f\x7a\xed\xb6\x88\x31\xc1\x4b\x59\x9b\x0c\xab\x5a\xf1\xa9\x42\xed\x40\x4b\x05\x3b\xf6\xa6\x02\xb3\x47\xd7\xf4\x00\xe9\x23\x9b\x7d\x96\x4a\x9a\x3e\xa7\x43\xe8\x84\x48\xef\xf2\x7d\x62\xf8\xb0\xd4\x31\xe3\x49\xf1\x7d\x1f\x16\x7c\x40\xa3\x03\x4c\xed\x77\x8a\x44\xa3\xbd\x96\x0e\x7e\xa0\xcd\x86\x2f\x73\xd2\xa0\x05\x26\xd6\x13\x05\x16\xb4\xec\xc5\x49\xcc\x49\x50\x20\x2f\x71\xe7\xbc\x96\x68\x5c\x89\x15\x80\x75\x86\x44\xe9\x9c\x41\x68\x55\x7d\x9c\xcb\x58\xd7\xf5\x48\xf5\x1d\x44\x3a\xef\xa6\x6f\x9c\x99\xa4\x65\x18\xfa\xa6\x1b\x62\xbf\xa2\x19\x96\x8c\x74\xba\xda\xfc\x8c\xf9\xb4\x9c\x4d\xb6\x0b\xcb\x04\xad\x25\x2d\xbd\x75\xe6\x14\x7a\xc3\x80\xc4\xcc\xa0\x03\x52\xed\xeb\xc4\x26\xbd\xb3\xa0\xd0\xc6\x89\x1a\x99\x2c\x95\x30\x26\x74\x2e\xca\x58\xba\xa3\x5b\xfa\x75\x1c\x1b\x97\x35\x3d\xd2\x29\x7e\x3f\x29\xe7\x51\x03\xd0\xae\x01\xd6\x4f\x99\x00\xbf\x9f\x92\x37\x43\x6a\x07\x2a\xcc\x44\xce\x6b\x88\xeb\xb9\x61\xe5\xf7\x71\x18\xe6\xb9\x5c\xb6\x3d\x27\xd3\x0a\xb9\x61\xbf\x47\xaf\x2b\x2a\x4c\xfc\x3d\x78\x1d\x76\x84\x0a\x5c\x3b\x82\x9f\x3f\x7e\xfa\x24\x50\x5b\x6f\x47\x96\x18\x75\xa1\x40\x4b\xd0\xf2\x18\xa6\x46\x05\x46\x94\xa1\xea\x6d\x83\x7f\x07\x81\xb1\x8a\xf7\x33\xb9\xb5\x95\x21\x28\x57\x82\x96\x96\x05\x81\x12\x60\x30\x13\xe0\x2d\x5a\xce\x25\x82\x2b\x53\x0f\x80\x12\x5c\xb2\x32\xa8\xc0\xa1\x8c\x05\x96\xb4\x2b\xd1\xef\x17\x6c\x6f\x37\x31\xc6\x4a\xce\x83\xcd\x98\x81\xb6\xa0\xfb\xd1\x6e\x4d\x8d\x12\x5a\x7b\x50\x5d\x99\x37\x07\x3c\xcd\x5b\x85\xeb\x02\xcf\x18\x4f\x4b\x5e\x79\x1a\x26\x94\x8a\xe4\x09\xc1\x4c\xd4\x7a\xc3\x79\xb7\x00\x7a\xf1\xde\x3f\xaa\x96\x1e\x10\xdd\xa9\x1e\xab\xd9\x83\xa4\xee\x84\xda\xde\xe3\x0a\xcc\x80\x4c\x69\x36\x8d\xb6\xd2\xda\x39\x30\x49\xd8\x29\xac\xd8\x38\x50\xe4\x4e\xa4\xcf\x16\xec\x1b\xac\xa1\x45\xa0\x37\x14\xdc\x91\x71\xa5\x4d\x0b\x30\x4c\x18\x30\x86\xa0\x40\x9b\xa5\x57\x49\xbd\x58\x29\x2c\x58\xe3\x5c\x00\x0c\xf6\x88\x74\x0e\xc1\xc1\x08\xb5\x6e\xeb\x76\x0b\x66\xac\xd9\x02\xbc\x25\x26\x02\x1a\x07\xd0\xd2\x5f\x18\x66\x12\x8b\xb4\x16\x92\x4e\x4d\x2d\x89\x15\x17\x06\xea\x72\xa9\x9d\xb8\x26\xf5\x4c\x85\x29\xf1\x43\x77\xe2\x80\xd6\xc5\x1d\xae\x46\x6d\xc3\xa4\x8c\x48\xde\xb6\xdb\x65\x4f\xa4\xad\x79\xe6\x52\xcc\x08\x98\xa6\x32\x92\xb4\xda\x52\x6e\xa5\x1b\x2c\xc8\x3a\x34\x7d\xfd\xb4\xcc\x76\x06\xdf\x66\x55\x2f\x04\x7e\x64\x81\xb3\x96\xe4\x9a\x2e\x95\x64\x11\x6c\xab\xad\x08\xde\x89\x49\x90\xc9\xb2\xf6\x9b\xf4\x83\x2b\x7c\x12\x6c\xac\xea\x0d\xd4\xb7\x50\x31\xf4\x21\x1d\x40\x75\x40\x61\xef\x53\x0a\x1d\xe9\x37\x57\xf4\x06\xf6\xdd\xfe\xb0\x80\x9a\x5c\x10\xc7\x62\x1f\xc0\x9b\x67\x11\x76\x82\xaa\x7e\xda\x41\x5e\x24\x7e\x0d\x33\x2c\x4a\x52\xb2\x33\xf6\xf7\x4c\xbc\xfb\xcc\xef\x0c\xce\x62\xb3\x31\x3c\x9a\xd6\x68\x2c\x6b\x8d\xaa\x51\x83\x0e\xa4\x5a\x2f\x6b\xac\xc4\x53\xc2\xcf\x94\xb8\x86\xb2\x9a\xe5\x13\x0a\xa8\x02\xed\xe2\xa4\x78\x13\x13\x77\x8e\xf0\xe8\x24\x12\xac\x05\xd6\x4d\xfd\x1c\x8d\xa3\x30\xb0\xa6\x77\x3b\xfb\x0a\xab\xcc\xaa\x39\xbc\x33\x0d\xa7\x81\x67\x96\xb8\x05\x87\x4a\x91\xc3\xb3\x35\x1e\x57\x1b\x1a\x67\x50\xcb\x1b\x31\xaf\x4b\x11\x17\xa4\xc6\xb2\xd6\x58\xbc\xed\x19\xe7\xc0\xea\x40\xba\x10\x8a\xad\x33\x24\xc9\x57\x92\xf2\x9c\x04\x29\x5c\xe8\x27\xde\xb7\x98\xef\xc0\x7d\x03\x35\x2b\x63\xe1\x67\xa8\x35\xc4\x59\x4f\x8d\xe6\x78\xea\x58\x1a\x5f\xd4\x4c\x96\x35\xe9\xe2\x01\xd7\xfe\x11\xdd\x96\x83\x3f\xac\x30\x1d\xd8\xb4\x01\xce\xe6\xc1\xae\xaa\x5d\x03\x70\xa6\xde\x0c\xd2\x8c\x5b\x1f\xc3\x98\xa4\x1d\x1a\x1d\x3c\x81\x61\x1c\x10\x03\x17\x34\x02\xe7\x5d\xf8\x69\x09\x63\x72\x43\x51\xbd\x15\x42\x7d\x20\xc3\x31\xa4\x3a\xb0\x23\x09\x7b\xaa\x88\x35\x90\xc1\x8a\x6c\x74\xcc\x47\xd6\x68\x5a\xd4\x19\x89\x45\x32\x9f\xa1\x87\x1a\x4d\x71\x5a\x8b\x58\x27\xed\x39\x4a\x83\xd7\x31\x02\xbe\x1e\xbd\x09\xc9\x6b\x58\x09\xfc\x21\xd0\xda\x63\x9c\x52\xcf\x9c\xfc\xef\x33\x0f\x8b\x51\xef\x09\xa3\x91\xd6\x7c\x18\x1c\x8f\x73\x36\x48\x85\x96\x64\x50\xb8\xfe\xe4\x98\x4b\xba\x33\x9c\x36\x10\x3c\x56\x63\x16\x61\x9d\xb0\x4c\x91\x3c\x3f\xd6\xb6\x46\x41\x61\xb3\x8b\x41\x84\xae\xd3\xd6\x8f\xc9\xdc\x42\x9c\x1e\xa6\xa5\x7e\x23\x98\xb0\x05\x84\x29\x12\x9e\x49\x17\x25\x7b\x33\x0e\xfa\xa7\xd6\x82\xad\xb3\x35\x9a\x50\x7e\x97\x67\x7a\x05\x61\xf2\x52\x70\x12\x6a\xe9\x8d\x49\x77\xa7\x40\x07\x70\xd8\xdc\x8e\x9a\xb8\x42\xd3\x7d\xe8\x23\x17\x28\x4b\x84\xde\x73\xef\x1b\x96\x58\x11\xdc\x56\x19\x2f\x66\x4a\x56\x32\xc5\x16\x68\x49\xcc\x6b\xfa\xca\xf7\xaa\xcc\x87\x83\xc9\xdd\xaf\x23\xb9\xb2\x09\xa2\xf7\xfc\x8d\x06\xd5\xd5\x78\x3e\x8e\x3c\x09\x76\xa9\xe1\x55\xd4\xc1\x0e\x11\xff\x27\xc1\xde\xd5\xde\xc5\x81\x1d\x5c\xea\x0e\x1f\x2a\x04\xeb\x4d\x3f\x73\xa2\xf8\xe4\xd3\x0b\xae\x6a\x30\xc1\x37\xb2\xfd\x05\x55\x7f\xa7\xfc\x38\xd4\xa4\x4a\xb7\x31\x6f\x8c\xa0\x42\x0a\xe2\x83\xb5\x64\x6f\xc8\x56\x13\x2a\xd5\x60\x2d\xea\x22\x85\x41\xf0\x44\xf5\x82\xc1\xba\x94\x3b\xc3\xff\x0c\xe0\x86\xd7\x35\x6c\x99\x5e\x5d\x52\x8e\xef\x49\x8f\x32\x0e\xee\x91\x38\x43\xf6\x52\xf4\x82\x05\x1c\x57\x95\x28\xc1\x80\x70\x68\xd2\xa9\x32\x2b\xb9\xc2\xd6\xf9\xeb\x0c\x7a\x85\x92\xc0\x5b\x28\x46\x6a\xa1\x03\x11\xf6\xd8\x7e\x85\x46\xc7\x51\xd2\x81\xa4\x07\xb5\x70\xed\x3f\xc4\x62\xaa\x2b\x16\xd0\x59\xbc\x87\x84\x7f\x6a\x05\x02\x25\xd9\xda\x3b\x1c\x98\xfa\x3e\x40\x37\x4c\xd4\x68\xeb\x0d\xee\xd2\x56\x85\x9a\xca\xde\xb8\xc0\x5c\x2d\x66\x93\x5a\xe6\xf1\xfc\x11\x85\xab\xfc\x6d\xa2\xd3\x57\x80\x56\x56\x24\xb5\x7c\x43\x05\x3a\x80\x07\x1c\x32\x49\x16\x76\x31\xf2\xd3\x50\xa6\x1c\xb9\xae\xd9\xb8\xd0\x86\x46\x53\xea\x19\xe1\x67\xbc\xaf\xa2\xdc\xda\x5d\x63\x20\x11\x82\x73\xd6\x2d\xb4\x36\xa9\x26\x23\x2d\xb8\x0a\xf6\xf5\x18\xb9\xb4\x0a\xc5\x88\x4a\x70\x80\x40\xcb\x9c\x59\x2e\xd5\xe8\x2e\xac\x73\xfd\xae\x82\x3e\x7b\x17\x5d\x1b\xfe\x8e\x62\x60\x01\x2a\x30\x29\xb2\x0d\xce\xdb\xc9\x7a\xcf\x5e\x46\x0f\x21\xc7\x7a\xde\xc6\x9e\xdf\xb8\x44\x09\x3a\x78\xe3\xdc\x9d\x48\x2a\x34\x61\xd3\x0b\xce\xb7\x78\xf5\x64\x87\xb6\x6e\xba\xac\x82\xf9\xad\x6c\x31\xc6\x99\x6a\xd7\xc1\x1e\xcf\x7c\xa8\x58\xbb\x52\x9d\xde\x22\xf5\xe1\x9a\xe8\xd5\x8c\x59\xa8\x14\x4d\xe5\x77\xde\xbd\x8d\x39\x3e\x47\x78\x6c\xa9\x8c\x3c\x93\xfe\x75\xd4\xc1\xc6\xdb\x6a\x34\xd5\x74\xa5\xce\x07\xce\xd9\x28\x16\xa0\xc0\xbb\x92\x4d\xb4\x50\x0f\x3b\xe3\x57\xf8\x8c\x55\xbf\x83\xd8\xdd\xa1\xc5\xb7\xed\xa7\xbb\xe3\x28\x7f\x63\x37\x3d\x7f\x2d\x36\x2c\xea\xa5\xad\x7d\x2d\x76\x15\xe5\xff\xcc\xf2\xdc\xed\x30\x74\xef\x82\x83\x80\x3f\x6a\x14\x0e\xb4\xa0\x56\x09\xae\x49\x37\x77\xd3\xa1\x82\x75\x27\x85\xcf\xf8\x27\xd7\xe1\xc6\x1a\xce\xe2\x3e\xac\x66\xf8\x69\x4f\xd6\x61\x73\xc0\x6c\xa7\xf3\x0e\x0d\xc2\x01\xfb\x24\x89\x03\x77\x07\xb6\xe7\xa1\xc6\x9a\xdd\xc6\x5c\x6d\xff\x38\x80\x00\x2d\x4e\xf7\x1f\x51\x16\xef\x1f\xe7\x08\x8b\xc3\x18\xcd\x72\x4b\xb3\xfb\x32\x19\x7e\x54\x3e\x2c\xee\xec\x0d\x4a\x9f\x12\x8c\xee\x89\x68\x5c\x47\x9d\xb6\x06\xcb\xe0\x97\xc6\xf6\xc6\xc0\xcf\x27\x8b\x5f\x91\xd7\xb4\xfd\xa9\xa3\xe1\xd8\x81\x1a\x8b\xf9\x7b\xc6\x65\x9a\xc8\xea\x47\x1c\x57\x22\xe6\x39\x0a\x67\x39\x8f\x37\x27\xa0\xe5\x0e\x35\xe6\xe4\x2c\xeb\x4e\x48\x6a\xfb\xa6\xe7\x9e\x79\x26\x4d\xea\x7c\xcf\xe5\x46\x4f\xa0\x6a\x4f\x1d\xf1\x17\x1b\x50\xc3\xf4\xaa\x36\x63\xbd\x2b\xbc\x27\x2d\xe2\x96\xec\xf3\x8c\xf6\x69\x90\x7b\x72\x05\xab\xd7\xdf\xfa\x4a\xbb\x13\x69\xeb\x4d\xe8\xd3\x18\xa9\xd4\x84\x5d\xbb\xdc\x6b\xd9\x37\x7e\x34\xa1\xf0\x71\xb8\x5e\xa9\xd0\x23\xa3\xa4\x8c\xf4\xa5\x16\xd9\x9a\x2d\xec\x14\x9e\x0d\x67\xfb\xbd\x47\x6a\x34\x5f\x77\xa4\xd6\xbd\x38\xe7\x5f\x85\x2c\x04\xbc\x67\xb4\xf8\xa8\xd1\xd8\x92\x6a\xce\xfd\xfe\xd5\xb3\x43\x69\x4b\x30\x17\xb6\x71\xf1\x78\xcc\x09\x7c\x3a\x4f\xd6\xef\xdb\x7a\xab\xe6\xc7\xf6\x62\xaf\xd4\x7f\xf2\xfe\xd2\xef\x47\x9f\xfe\x05\x0f\x4c\x4b\x72\xde\x20\x6b\x83\x16\xc1\x84\xc3\xbb\x94\x78\x40\xc5\x75\xe7\x69\xaf\x70\xaf\xf9\x10\xf2\x66\xfc\x5d\xe3\x01\x94\x47\x90\x12\x65\x71\x80\x89\xb9\xdf\x97\xf7\x1f\xe5\x40\x5d\x1b\x06\x51\x5e\x7e\x25\x39\x96\x36\x33\xc3\x6f\x88\xdd\x2c\xfb\x7a\xf3\xd5\x83\x71\x68\xd4\x69\x18\xdd\x48\x39\x9f\x61\xfd\x28\xcb\x79\x93\x22\x09\xd6\xa2\x1b\x45\x96\xef\xfa\xb8\xf3\x61\xa0\x4d\x96\x7e\xda\xac\x40\xb7\xb5\x2e\x88\x91\x2f\xa5\x73\xb5\xfd\x57\xd6\x96\x7d\x60\x6d\x3f\x14\x7c\xf8\xe0\xf7\x9b\xec\x40\xd6\x83\x22\x9b\xee\x76\xb3\x30\x01\x04\x6b\x87\xda\x65\x12\x0f\xbf\x64\x7f\xfe\xf1\xf5\xdb\xd7\x3f\xbf\x7c\x63\x21\x7c\x1d\x2b\x7d\x13\x2c\x49\x17\xdf\x1c\xb3\xfa\x50\xba\x4a\x75\xf2\x59\x5b\x49\x45\xa0\xf5\xa1\x20\x57\xfa\xdd\x07\xe2\x4c\xd6\x5b\x11\x68\x52\x4e\x69\xfe\x6f\x43\x43\x9b\xd9\xe0\x44\x80\x91\xdb\x5e\x32\xa8\xb3\xaa\x4b\xd0\x97\xa9\xf0\xc7\xd7\xaf\xff\xfd\xf2\x4d\x92\x15\x7c\x40\x73\xfa\x69\xe4\xaf\xe2\x3e\x99\xa5\x71\xb6\x32\x2d\x9a\xc2\x93\x44\xc7\x13\xd9\xad\xab\xe6\x68\xcc\x43\x7d\xf7\xea\xf4\xf9\xe3\xa7\x7f\xce\x0e\xca\xd7\xdf\xbf\xa4\x81\xfd\xfd\xcb\xb7\x92\xd0\x04\xbb\x71\xfa\x76\x20\x3c\xae\x31\x22\x8d\x81\xa3\x2b\x93\x69\x1a\x73\x93\x19\x54\x4d\xde\x76\x3a\x99\xe7\x54\x78\x83\x96\x75\x9b\xb3\xd6\xe4\xe2\xcb\xdd\x09\x0c\x02\xe7\x3e\xe8\x67\xd0\x92\x44\x2d\xf0\x2c\x72\xf2\xf9\xe3\xa7\xdf\xd2\xc7\x39\x2f\x2b\x4b\xfe\xee\x35\x7e\xfe\xf8\xe9\x1f\xb5\xe1\x03\xd9\x38\xef\x36\x59\xfa\xaa\x2e\x2b\x51\xd5\x99\x43\x53\xc5\x64\xed\xee\x26\xea\xa5\x2d\x4f\xff\x39\x03\xda\xe6\x6c\xaa\x36\xec\x5c\x18\xa8\x2a\x9c\x68\x38\x16\x2c\x98\xf7\x14\xbd\xd3\x98\x86\x20\x4e\x0b\xe5\x36\xb5\xc3\x24\x89\x67\xe5\x91\xd4\x85\x32\xc2\x9c\x6a\x6a\xa0\x75\x9b\xee\x34\x22\x58\xe7\x2f\x59\x38\x05\x3b\x6f\xb3\x63\x09\xee\x88\x92\x33\x3c\xc4\x2f\x2c\x87\xb5\xd0\x84\xce\xdc\x64\xda\x96\xc7\x97\xce\x4a\x92\xee\xc0\x72\x36\xd7\xae\x21\xfb\x4b\xcf\xf1\xfb\x46\x48\x1b\x0f\xb0\xce\xcb\x13\xe7\x5d\x44\xe1\x88\x4a\xed\x90\x74\x11\xd3\xce\x0d\x86\x65\x77\x62\xaf\x8b\xb4\x1c\x1f\x6d\x97\x74\x28\xcd\x7b\xd4\x22\xc0\x5f\x34\x1c\xd3\x1b\xb4\x4d\x37\x3e\xa4\x4b\x50\x61\x4b\xeb\xa6\xfc\x26\x2c\xfd\xb0\x6e\xfa\xa4\x8c\x4f\xfd\xcf\x6d\x6f\x6e\xb6\xdd\x46\xb9\x35\x58\xb3\x71\x59\x8e\x3b\xe3\xc1\x9c\xb6\x61\x8d\x64\xd6\xed\xb6\x95\x7d\x35\xdb\xd1\xeb\x64\x65\xd6\x0c\xb8\x76\x1f\x71\xf4\x2c\x07\x47\x9e\x96\xe3\x19\xc5\x40\x65\x35\x4d\x63\x82\xd0\xbb\x54\x4c\x55\x36\x8b\x63\xbe\x4d\xbc\xb6\x9a\xdd\x96\xf4\xb6\x8b\xbf\x6c\xb7\x7d\xb0\x6a\xcb\x66\xeb\x0c\x90\x26\x5d\x6c\xb7\x1a\xd1\xd9\xed\xe5\x98\x36\x7a\xde\x0a\x85\x69\x76\x4d\x34\xcc\xeb\xbe\xee\x40\xa7\xc1\x54\x8c\x75\x3b\x4a\x7d\x6d\x36\x2d\x9d\x40\xe6\x8a\x8a\x1a\xe0\x73\x66\x83\xb5\x91\xdb\x36\x8e\xb1\xe5\x7c\x1b\x1d\xc7\x5f\x03\xed\xdf\xe2\x54\xfc\x9c\xca\xc3\xfb\x6d\xf3\xbe\x55\xe5\x86\x7b\xd8\x73\x4e\x30\x2d\x0a\xe7\x11\x23\x0d\xc4\xaf\x2c\x9c\x63\x89\x22\x08\xee\xd9\xed\xbc\xdd\xa6\x33\x57\xe7\xe2\x6f\xfb\x33\x58\x4b\x70\x3b\xd8\x50\xb6\x06\xad\x57\xce\x8e\x66\xd5\x72\x8a\x97\x9f\xe5\x4d\x32\x1c\x00\x36\x78\xe3\x49\x93\x46\x74\x9b\x86\x74\xb0\x14\x66\x66\xfe\xf3\xc1\xd1\x5e\x0b\xbf\x1f\x09\xb8\x1c\x79\xa8\xa9\x09\x50\x65\x47\xdc\x0d\x1f\x4b\xae\xf0\x25\x6d\x9a\xf1\x46\xac\xe7\x9e\x5a\x0e\x29\x2f\x0d\x3d\xb6\x9f\x6e\x5c\x84\x08\xef\x0e\x3a\x4e\x4b\x22\x2d\x94\x97\x68\x41\xa9\x5f\x9a\x76\xa2\xa6\x52\xd4\x04\x5a\x9a\xd1\x9f\x6b\x20\xd1\xae\x89\x89\xad\xb7\xdb\xb3\x93\xa3\x10\xef\xb4\xfb\x9d\xb8\xef\x89\xf1\x4e\x9f\xa1\x2e\xa0\x20\x5d\x1c\xc9\x95\x97\x95\xcf\x37\xf4\x91\x98\xc8\xa4\xe7\x16\x25\x45\x47\xe3\xe7\xb1\xbb\x78\x95\xf8\x0e\x48\xc5\xe7\x77\x40\xa8\xbb\x8a\x59\xcc\x85\xb5\x22\x8d\x93\xb7\x39\x61\xc7\xdf\xac\xeb\x85\x4c\x65\x17\xad\xec\xe8\x5c\x40\xac\xac\xc2\x8d\x9c\xc7\x95\xb5\x99\x43\x5b\x7b\x6c\xae\x24\xe8\x28\x91\xaf\x3d\x44\x57\x90\x56\x56\xa8\x2f\x1d\x44\x51\x4a\x58\x7b\x88\x26\x61\x56\x56\x25\x87\x8a\x82\x23\x66\xd9\x87\x79\x9d\x5e\xe7\x66\x6d\x55\x26\x61\x56\x56\xe5\xfa\xc5\x31\xd7\x7a\x65\x7d\xae\x63\xad\xac\x94\x30\x54\x35\x92\xfb\x3c\x8b\x78\xa0\x9f\x2e\x6a\x4b\x73\x78\x5d\x7b\x0c\x67\xf1\x36\x59\x75\x8a\x76\x7d\xf1\xfe\xd0\xd4\xdf\x64\x0e\xf6\x58\x83\x59\xbe\xed\xa6\x9d\xa5\xd7\x22\x89\xc1\x57\x0f\x2a\x66\x0f\x2d\x16\x24\x50\x3b\x83\x41\x99\xae\x6d\xd4\x5c\x28\x1f\xfc\xcb\x4d\x06\xb6\xc4\x6b\xbd\xd8\x7a\xab\x83\x3f\xbc\xd0\x56\x4c\x11\x9c\x66\xee\xe5\xf1\x83\x91\xe0\x30\xb6\x9f\x92\xd8\x12\x37\x1d\xc5\xb8\x41\xd7\x60\x50\x8b\x53\xf7\x97\xca\xe2\x5f\x36\x93\xe0\xa0\x36\xec\xd2\xdd\x59\x1b\x0e\x39\x1b\xca\x51\x52\x45\xdc\xab\xc3\xb0\x05\xf6\x2f\x2b\x41\x04\xf6\x61\x88\x42\x57\xa3\xb1\x4b\x98\xb3\xb6\x51\x32\x2b\x12\xa3\xf1\xb8\xde\x44\x31\xef\x49\x17\x90\x3b\x34\xa0\x65\xfc\x4e\x24\xc8\xc8\xd9\xa4\x3b\xdd\xe6\x66\x33\x6b\x85\x3e\xc0\x23\x2b\xf9\x78\x44\xc1\x4a\xa1\x70\xa1\xe0\x7d\x11\xdb\x23\xd6\xe1\xad\x45\xe1\x0d\xbe\x2f\x6e\xde\xe2\xfb\xeb\xb0\x28\xce\x96\xc3\x80\xf5\xbb\xa2\x87\xa6\xa2\xe6\x43\x05\xc7\x0d\x46\x7b\xdf\xf5\x6e\xa8\x2a\x2c\xc8\xa6\xa3\x1c\xeb\x37\x1c\xe3\xff\x90\x2e\xbe\x92\xc3\x7f\x7b\x57\x7e\xf8\x31\x88\xd6\x0b\xa9\x07\x17\x4b\x59\xf2\x75\xb2\x03\x28\x92\xe9\xcc\x72\xde\x72\x93\xe5\xf6\xba\x3b\x3c\x63\x96\x53\x5e\x05\xf5\x96\xb9\xbd\x46\x16\x25\x56\x81\x23\xc9\xe5\x9b\xd0\xc8\x7f\x1e\x5c\x5d\x3a\x84\x6a\x93\xed\x30\x5e\xae\xa3\x5c\x2c\xaf\x6b\xb1\xc9\x0c\x6a\x20\x6b\x43\x4f\xd7\x86\xfe\x5a\x7e\xe4\xba\x79\xfc\x73\x25\xb2\xb6\xe7\xf2\xff\x17\x00\x00\xff\xff\x75\x95\x43\xd3\x83\x56\x00\x00")

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

	info := bindataFileInfo{name: "redirects/redirects.csv", size: 22147, mode: os.FileMode(420), modTime: time.Unix(1572260958, 0)}
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

