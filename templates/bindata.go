// Code generated by go-bindata.
// sources:
// templates/common/base.go.html
// templates/common/head.go.html
// templates/common/menu.go.html
// templates/monitor/index.go.html
// templates/monitor/menu.go.html
// templates/monitor/processor.go.html
// templates/query/index.go.html
// DO NOT EDIT!

package templates

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

var _templatesCommonBaseGoHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xaa\xae\x4e\x49\x4d\xcb\xcc\x4b\x55\x50\x4a\x4a\x2c\x4e\x55\xaa\xad\xe5\xb2\x51\x74\xf1\x77\x0e\x89\x0c\x70\x55\xc8\x28\xc9\xcd\xb1\xe3\xaa\xae\x2e\x49\xcd\x2d\xc8\x49\x2c\x49\x55\x50\xca\x48\x4d\x4c\x51\x52\xd0\x03\xa9\x82\x48\xda\x24\xe5\xa7\x54\xa2\xaa\xc9\x4d\xcd\x2b\x85\xa8\x41\x16\x4d\xce\xcf\x2b\x49\xcd\x2b\x81\x6a\xd6\x87\x68\xb3\xd1\x87\x59\x91\x9a\x97\x52\x5b\xcb\x05\x08\x00\x00\xff\xff\xfd\x8f\xc0\x67\x8d\x00\x00\x00")

func templatesCommonBaseGoHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesCommonBaseGoHtml,
		"templates/common/base.go.html",
	)
}

func templatesCommonBaseGoHtml() (*asset, error) {
	bytes, err := templatesCommonBaseGoHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/common/base.go.html", size: 141, mode: os.FileMode(420), modTime: time.Unix(1490714385, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesCommonHeadGoHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x54\x4b\x73\xe2\x38\x18\xbc\xe7\x57\xb8\x74\xe1\x90\xb5\x04\x98\x0d\x24\x85\xb3\x95\x25\xbb\x3c\xf2\x22\x90\x40\xd8\xcb\x96\x90\x3e\xdb\x72\x64\xc9\xb1\x64\xc0\xc3\xf0\xdf\xa7\x1c\x86\x84\xca\x64\x1e\x39\xcc\xcd\x5f\xbb\xd5\xee\x6e\x7f\xa5\xf5\x9a\x43\x20\x14\x38\x28\x02\xca\xd1\x66\x73\xd0\x2e\x1f\x4e\x0f\x1c\xc7\x71\xda\x09\x58\xea\xb0\x88\x66\x06\xac\x8f\x72\x1b\xb8\x2d\xb4\xff\x2a\xb2\x36\x75\xe1\x29\x17\x0b\x1f\xad\xdc\x9c\xba\x4c\x27\x29\xb5\x62\x2e\x01\x39\x4c\x2b\x0b\xca\xfa\x48\x80\x0f\x3c\x84\xdd\x49\x2b\xac\x84\xd3\xf5\x1a\xa7\x34\x84\xff\x9f\xa7\xcd\xa6\x4d\xb6\xf0\x9e\xb8\xa2\x09\xf8\x88\x83\x61\x99\x48\xad\xd0\x6a\x4f\x12\x7d\x4b\x5c\x08\x58\xa6\x3a\xb3\x7b\xac\xa5\xe0\x36\xf2\x39\x2c\x04\x03\xf7\x79\xf8\xc3\x11\x4a\x58\x41\xa5\x6b\x18\x95\xe0\xd7\x76\x42\x52\xa8\x47\x27\xca\x20\xf0\x2b\x65\x28\x73\x42\x48\xa0\x95\x35\x38\xd4\x3a\x94\x40\x53\x61\x30\xd3\x09\x61\xc6\xfc\x15\xd0\x44\xc8\xc2\xbf\x49\x41\x1d\x8e\xa9\x32\x87\x1d\xad\x38\x28\x03\xfc\xc4\xab\x56\x3f\xbf\xe0\x15\x27\x03\xe9\x57\x8c\x2d\x24\x98\x08\xc0\x56\x1c\x5b\xa4\xe0\x57\x2c\xac\x6c\xa9\x54\xd9\xff\x78\xc9\x45\xaf\x5c\xb4\x75\x83\x76\x6e\x12\xba\x62\x5c\xe1\xb9\xd6\xd6\xd8\x8c\xa6\xe5\x50\x1a\x7a\x01\x88\x87\x3d\xdc\x2c\x65\x5f\x31\x9c\x08\x85\x99\x31\xc8\x11\xca\x42\x98\x09\x5b\xf8\xc8\x44\xd4\x6b\x35\xdc\xbf\x27\x33\x21\xc6\xfd\x7f\xe1\xa2\xc6\xbb\xc9\x60\x74\xf6\x58\xb0\xbc\x77\xd6\x1b\x85\x5e\xfd\x26\xb9\x67\xcb\x65\x53\x2b\x6f\x34\xe3\x61\x63\x42\x0f\x87\xc9\xf8\xce\x7c\x22\x17\x47\xad\xc5\x9c\xff\x13\x47\x8d\x1c\x39\x2c\xd3\xc6\xe8\x4c\x84\x42\xf9\x88\x2a\xad\x8a\x44\xe7\x06\xfd\xee\x50\xae\x8d\x20\x81\x1f\x45\xcb\x7a\x85\xbe\xae\x89\x91\x99\x3c\x4c\x1a\xea\xbc\x3a\xc8\xad\x54\x5d\x6a\x64\x67\x90\x77\x9a\xf9\x32\xe6\xf9\xf4\x78\x3c\xc9\x2e\x17\xa3\x99\xd6\xc3\xb4\x3e\x9f\xce\xc2\x24\x1c\xdc\xf6\x1f\x96\x92\x8c\xd3\x9f\x45\xdb\x6e\xa4\x63\x32\xe6\x23\x42\x68\x4c\x57\x6f\xd7\xa4\xc4\x88\x14\x73\x43\xe2\xa7\x1c\xb2\x82\xd4\x70\xad\x86\xab\x5f\xa7\x67\xef\xb1\x41\xa7\x6d\xb2\x95\x7a\x47\xf7\xa3\x15\xc5\x6f\x7f\x7b\xfc\x6e\x35\x77\xec\xcf\xfe\xad\x98\x57\xeb\xcd\xa7\x45\x11\x8f\xaf\x82\x5e\x7c\x73\x45\x2f\x1f\x83\x7c\x3a\x59\xfd\xb7\xba\x1f\xaa\xce\xe0\xac\x29\xeb\x49\x67\x7a\xdd\x4f\xbb\xc7\x49\xb7\x73\xde\x5a\x76\xaf\xfb\x6c\x78\xde\xbc\x5b\xd1\xef\x57\xf3\x0b\x59\x18\x57\xb1\xc1\x4c\xea\x9c\x07\x92\x66\xf0\xa6\x2a\xa9\x39\x35\x11\x8e\x0d\x69\xe0\xda\x11\x6e\xec\x80\x0f\xb4\xc5\xbd\xd8\x60\x9d\x85\x84\x7b\x78\xd1\x78\xe7\x64\x9b\x6c\xaf\xb7\xf5\x1a\x14\xdf\x6c\x0e\xbe\x04\x00\x00\xff\xff\xee\x02\xd0\x16\x00\x05\x00\x00")

func templatesCommonHeadGoHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesCommonHeadGoHtml,
		"templates/common/head.go.html",
	)
}

func templatesCommonHeadGoHtml() (*asset, error) {
	bytes, err := templatesCommonHeadGoHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/common/head.go.html", size: 1280, mode: os.FileMode(420), modTime: time.Unix(1490714385, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesCommonMenuGoHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x64\x90\x41\x6e\x84\x30\x0c\x45\xf7\x9c\xc2\xca\x7e\xc8\x05\x32\x5c\x05\x19\x62\x44\x54\xd7\x8c\x42\x92\x8d\xe5\xbb\x57\xd3\x01\xda\x4e\x57\x89\xf5\x9f\xf2\x7e\xac\x1a\x69\x49\x42\xe0\x3e\x49\xaa\x33\xeb\x82\x60\x83\x99\x71\xdf\xef\x4e\xb0\x4d\x98\xe1\x75\xdc\x22\x2d\x58\xb9\xb8\xa1\x03\x08\x31\x5d\xd4\xbc\x49\xc1\x24\x94\x6f\x0b\xd7\x14\xbf\xf3\xbf\xc4\xf1\xc0\x4a\x18\x29\x1f\x39\x40\xc0\xb7\x7c\xca\x28\xd1\xc1\x9a\x69\xb9\x3b\xd5\x7e\xc2\x9d\xc6\x07\x96\xd5\xcc\xbb\x41\xb5\x7f\x76\x1c\x4b\x2a\x4c\x66\xc1\xe3\x21\xf2\x31\xb5\xff\xce\x79\x63\xc6\xc7\x4e\x67\xfb\x73\xfe\xd1\x57\xfe\xe5\x3f\x31\xc1\x76\x11\xaa\x13\x6f\xf3\xc7\x6b\x37\xe3\xf3\x9b\x24\xc5\x41\x6f\x76\x01\x24\xf1\x9a\x82\xaf\xfc\x56\xe9\xb8\x04\x2f\xd8\x86\x4e\x15\x48\x22\x98\x75\x5f\x01\x00\x00\xff\xff\x50\x2f\x43\xbe\x77\x01\x00\x00")

func templatesCommonMenuGoHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesCommonMenuGoHtml,
		"templates/common/menu.go.html",
	)
}

func templatesCommonMenuGoHtml() (*asset, error) {
	bytes, err := templatesCommonMenuGoHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/common/menu.go.html", size: 375, mode: os.FileMode(420), modTime: time.Unix(1490714385, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesMonitorIndexGoHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\x91\x4d\x6e\xc2\x30\x10\x85\xf7\x9c\x62\x64\xb1\x6c\xb0\x8a\x58\x55\xc1\x9b\xb2\x61\xd7\x1b\x54\x86\x19\xb0\x25\xd7\x46\xb6\x4b\x7f\x2c\xdf\xbd\x0a\x24\x0e\x49\x50\xd9\x24\x99\x79\xdf\xc4\xef\x79\x52\x42\x3a\x68\x4b\xc0\xf6\xce\x46\xb2\x91\xe5\x3c\xab\x51\x9f\x61\x6f\x64\x08\x6b\xe6\xdd\x17\x13\x33\x80\xdb\x5e\x83\x4a\x6d\xc9\x5f\x94\xb1\x66\xaa\x0f\xac\x9e\x97\xad\x06\x50\xab\xa5\x78\xf3\x6e\x4f\x21\x38\x1f\x6a\xae\x96\x9d\x92\xd2\x7c\x27\x03\xbd\x9f\x64\x54\xf0\xb2\x86\x45\xa9\x72\x2e\x88\x97\xf6\x48\x30\xd7\x16\xe9\xfb\x09\xe6\xa7\xee\x4f\x97\x81\x52\x85\x32\xf1\xc8\xce\x98\x38\x49\x4b\x06\x2e\xcf\x0a\xe9\x20\x3f\x4d\x1c\xb0\x77\xe8\x4a\x91\x44\x6d\x8f\x23\xae\x89\xba\x1a\x82\x51\x47\x43\x4c\xd4\x12\x94\xa7\xc3\x9a\xdd\x26\xce\x99\x17\xfb\x3c\xa5\x6b\xc2\x9c\x59\x7f\x59\xd0\x77\x6b\x2e\x27\x87\x71\xb5\x12\xb3\xa1\x53\x8e\xfa\xfc\xc8\xfc\xce\xe1\xcf\xd4\x39\x9a\x8e\x42\x53\x29\xe7\xf5\x6f\xb3\x64\x33\x01\x1b\x34\x8a\x57\xa3\xc9\x46\xd8\x6e\x6a\x8e\xf1\x1e\x82\x22\xa5\x7e\x57\x8b\x2b\xbf\xdd\x34\x41\x10\xa7\x49\xd0\x88\xff\x83\x8c\x1a\x83\x32\x25\xb2\xd8\xee\xbf\x08\xed\x47\xfb\xea\x90\xbf\x00\x00\x00\xff\xff\x68\xfb\xe0\xd6\xef\x02\x00\x00")

func templatesMonitorIndexGoHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesMonitorIndexGoHtml,
		"templates/monitor/index.go.html",
	)
}

func templatesMonitorIndexGoHtml() (*asset, error) {
	bytes, err := templatesMonitorIndexGoHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/monitor/index.go.html", size: 751, mode: os.FileMode(420), modTime: time.Unix(1490714385, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesMonitorMenuGoHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x44\xce\x41\xaa\xc3\x30\x0c\x04\xd0\x7d\x4e\x21\x44\x96\x3f\xdf\xfb\xe2\x78\xd3\x83\x04\xc5\x56\x5a\x83\xa3\x80\x9d\x94\x82\xd0\xdd\x0b\x4d\x4b\x57\x03\x0f\x66\x18\xd5\xc4\x4b\x16\x06\x5c\x59\x8e\x29\x6e\xb2\xb3\xec\x68\xd6\xa9\x56\x92\x1b\x43\x9f\x25\xf1\xf3\x0f\xfa\x09\x2e\x23\xfc\xc7\x4d\xda\xb1\x72\x6d\x66\x1d\x80\x2f\x39\x78\x82\x58\xa8\xb5\x11\x85\x1e\x33\xd5\x61\xae\x24\x09\xe1\x5e\x79\x19\xd1\x7d\x0b\x4e\xf5\x9c\x32\xc3\x70\xfd\xe0\xf0\x43\xef\x28\x78\x57\x72\xe8\x54\x59\xd2\xfb\xc1\x99\xaf\x00\x00\x00\xff\xff\xf2\xc3\x81\x2e\xa4\x00\x00\x00")

func templatesMonitorMenuGoHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesMonitorMenuGoHtml,
		"templates/monitor/menu.go.html",
	)
}

func templatesMonitorMenuGoHtml() (*asset, error) {
	bytes, err := templatesMonitorMenuGoHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/monitor/menu.go.html", size: 164, mode: os.FileMode(420), modTime: time.Unix(1490714385, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesMonitorProcessorGoHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xbc\x59\x5f\x6f\xdb\x38\x12\x7f\xf7\xa7\x98\xd5\xa1\x6b\x69\x6d\x4b\x76\xda\xdb\x87\xc4\xce\xe2\xfe\xe2\x0a\xdc\x76\x8b\x6e\x6f\xef\x21\x08\x0c\x46\x1a\x5b\xec\x52\xa4\x40\xd2\x76\x02\x57\xdf\xfd\x40\xea\x8f\x29\x59\x4a\xdc\xbb\x6b\xfb\xe0\x4a\x9c\x1f\x67\x86\x33\xa3\x1f\x39\xcc\xf1\x98\xe0\x86\x72\x04\x2f\x16\x5c\x23\xd7\x5e\x51\x8c\x96\x09\xdd\x43\xcc\x88\x52\x2b\x4f\x8a\x83\x77\x3b\x02\x70\xc7\x0c\x94\x50\x8e\x72\xb6\x61\x3b\x9a\x58\x79\x17\xc1\x66\x59\x32\x5b\x5c\xf5\xc8\x72\xc2\x91\x81\xfd\x9d\x25\xb8\x21\x3b\xa6\x2b\xd4\x30\x2e\x45\x92\x50\xbe\x6d\x70\x00\xcb\x74\x71\x7b\x3c\x86\xb9\x14\x31\x2a\x25\xe4\x3b\x92\x61\x51\x2c\xa3\x74\xd1\xe8\x8a\x12\xba\x1f\x54\x3c\x7b\x10\xc9\x93\xab\x4f\x93\x07\x86\x35\xa4\x7c\xb1\xbf\x33\xa5\x25\xcd\x31\x71\xb0\x06\x6d\x3c\x72\x47\xcc\x98\x6c\x0f\x58\xd8\xed\x47\xa3\x64\x19\xe9\xb4\x4f\xf8\x9e\x48\x4d\x35\x15\x7c\x08\xf0\xab\x26\x7a\xa7\x86\xa4\x1f\x88\x46\xf0\x33\xb5\x8d\x54\xf0\x8c\x06\xa9\x87\x84\x7f\xd9\x49\x89\x7c\x50\xfc\x8f\x7f\xff\x3c\x24\xfa\xdb\xc7\x0f\xe0\x3f\x63\x95\x93\xfc\x03\xd1\x54\x80\xbf\xcb\x13\xa2\x51\x45\xc8\xb5\x7c\x1a\x9c\xf0\x5e\x8a\xad\x44\xd5\xb3\xd4\x65\xd4\x0e\xac\x41\x74\x82\xbf\xd4\x26\x9b\x40\x13\x93\xdc\x2a\xa2\xbf\x51\x3c\x78\x9d\x79\x06\xe5\xa4\x3c\xb2\x09\xee\xd4\x4b\x59\xaf\x2a\x96\x34\xd7\xa0\x9f\x72\x5c\x79\x1a\x1f\x75\xf4\x89\xec\x49\x39\xda\x68\xdd\x13\x09\x12\x79\x82\xb2\xc9\xa3\x82\x15\x6c\x76\x3c\x36\xcf\x7e\xe3\x8b\x0a\x8e\xa3\xc6\xac\x99\x65\x24\xff\xa4\x4a\xc3\x0a\xd6\xe1\x86\x11\xad\x91\xfb\xeb\x30\x23\xb9\x33\x69\x7a\xd2\xb4\x27\x6c\x87\x53\xa0\x3c\xc1\xc7\xe0\xe8\x2c\x2a\x16\x59\x2e\x38\x72\x6d\x0c\x5b\x71\xa8\x72\x46\xb5\xef\x85\x5e\x10\x2a\x46\x63\xf4\xaf\x82\x91\x33\x23\x8a\x80\x6e\x9c\x79\x21\x43\xbe\xd5\x29\x7c\xb7\x82\xd7\x53\xd0\x29\x42\x86\x5a\xd2\x58\x01\x55\x40\x40\x8b\x9c\xc6\xa0\x72\x8c\xe9\x86\xc6\x95\xac\xad\x4e\x12\x9d\xa2\x04\x9d\x12\x0e\x04\x1a\xff\x41\x0b\x4d\x98\x03\xa5\x1b\xff\xbb\xd2\x45\xca\x63\xb6\x4b\x50\xf9\xde\x56\xfc\x4e\xbc\x00\x3e\x7f\x1e\xf0\x28\x80\x63\xab\x18\x24\xea\x9d\xe4\x70\x77\x7f\xe3\x0c\x17\xce\xb3\x8d\xd4\xdd\xb8\xf4\x93\x93\x0c\xc7\xf7\xb0\x72\x95\x7f\x12\x94\xdb\xe8\xdc\xb8\x51\xa9\xf5\xda\xe9\x8e\xf2\x22\x98\x82\x96\x3b\x74\xd1\x51\x04\xb1\xe0\x4a\x30\x0c\x99\xd8\xfa\xb6\xe4\x80\x51\xa5\x8d\xce\x13\xc8\xc5\xd4\xf9\x76\xd5\xd4\x75\xf0\x51\x22\xda\x3a\x30\xd9\x5f\x87\x5b\x29\x76\xf9\x9f\x9f\x9a\x29\xd3\x76\x3d\xb5\xb2\x5f\x79\x6d\xc6\xc3\xd3\x8a\x43\xb5\x7b\x30\xa4\xc5\xb7\xfe\x7c\x7a\x26\x64\x44\xe9\xb7\x26\x0b\xbf\x6c\x6c\x18\x82\xf6\x62\x5f\x2e\x39\xe3\x77\x5d\x22\x95\xdb\xbf\x19\xac\x72\x9c\xaf\x26\x9f\x94\x05\xc7\xce\x37\x5f\xb9\xbe\xef\xf7\x7b\x3f\xe4\xf1\x38\x1c\x07\x93\x45\x70\xd3\xd2\xd6\xf2\xbb\x9c\x38\x64\xaf\x94\xde\xcd\xef\xbb\x1a\xdc\xf7\x6a\x75\x77\xe3\xa6\x7e\x6c\x14\x7a\x21\x56\xe2\x62\xce\xd3\x53\x61\xdd\x38\xbf\x54\x4d\x5a\x22\xb6\xab\x09\xba\xe5\x64\xca\xa6\x5b\x4e\x25\xd5\x36\x64\xf4\xde\xee\xa0\x7d\x84\xe4\x84\xc7\xcc\x2b\xb7\xbb\xd5\xe9\xdb\x0d\xcb\x7c\x74\x99\x64\x3e\x5d\x04\xd5\x07\xe4\x7a\x57\x57\xf2\x05\x1a\x16\xd3\x2b\x57\x43\x4b\x85\x61\x75\x58\x81\xb7\xd4\xc9\xad\x37\xb1\x3e\x4d\xbc\x65\xa4\x93\xdb\x72\xc4\xe8\xae\x06\xbc\x51\xa7\x20\x2d\xff\x3b\x2b\xed\xe4\x5f\x1d\xa8\x8e\xd3\xd3\xfa\x43\x65\xf7\xd5\xd0\xd6\x69\x07\x1b\x13\x85\x30\xbf\x36\x31\x17\x39\xf2\xfe\x3a\x1a\x1b\x97\x96\x2a\x27\xbc\x3e\x32\x30\xf2\x80\x0c\xec\xef\xe9\x58\xf3\x4b\x8e\x7c\x19\x19\xd8\xad\xf5\x7b\x7c\x73\x6e\x6a\x61\x4d\x49\x8c\xc5\x1e\x4d\xf1\x8f\x3a\x16\xcd\xf2\x36\x52\x64\xad\xe8\x56\xf0\xb5\x32\xbb\xfb\x5a\x6c\x36\x0a\x35\x7c\xff\xfd\x0b\x88\x72\xbd\x70\xbb\x9a\xc3\x4f\x97\x41\xaf\x61\x7e\xd3\xe3\x8f\x34\x07\x8f\x15\x38\x01\x35\xe5\xb9\xcb\x30\x59\x67\xa8\x14\xd9\xa2\x5a\x5b\x8c\x6b\xa6\x1f\x72\x37\x5e\x64\xa1\x79\x18\xdf\x1b\x6b\x41\x9f\x39\x2d\x7a\x17\x9f\x1e\xb2\xde\x75\xa4\x87\xec\x79\xf7\xe3\xf2\xd8\xd3\xab\xb4\x92\xd5\x31\xed\xd3\xdf\x86\xb8\xa6\x3a\xb6\xcc\x67\x4b\x58\xbc\x63\x26\x14\x79\x75\xc2\xe9\xf1\xa7\x16\x99\x90\xfa\xb5\x73\x33\x9b\xf5\xe0\x87\xc5\x7c\x0e\x11\xf8\x5a\xcc\xec\x7b\x10\x6a\xf1\x77\xfa\x88\x89\x7f\x75\x16\x2b\x47\x0f\x55\xef\xc8\x3b\xbf\x1e\x08\x7e\x9a\x5f\xd7\xcf\x7d\x01\x41\x2d\xad\x6d\x2d\x60\x56\x47\x27\x00\xbb\xb5\xe3\xc9\xe0\x79\x72\xdc\xaf\xe1\x43\x53\xc1\x65\xa9\xc3\xa4\x03\x36\xff\x2c\x72\x3c\x31\x7a\x9d\x75\x4c\xc6\x2f\x4f\x31\x8b\xbf\x04\x57\x79\x7f\x09\x54\x8b\x4b\x50\xa8\x65\x0d\x1b\x44\xcd\x5e\x90\xb7\xda\x8f\x2a\x0f\x1e\x28\xfd\xc4\x70\xe5\x65\x44\x6e\x29\x9f\x3d\x08\xad\x45\x76\x3d\xcf\x1f\x6f\xbc\x7e\x55\x46\x59\x8f\xa2\xd9\x03\x91\x1e\x48\x61\x74\xd5\x63\x76\xa8\xd6\x4f\xf9\xec\x40\x13\x9d\x5e\xc3\x1b\xcc\x6e\xa0\x7c\x1e\xc3\xe4\x54\x2f\x13\x18\xbf\x32\x56\x7b\xe3\x70\xaa\xab\xc9\xf8\x55\x79\x4a\xae\x7f\x4b\x4a\x3b\xe7\xb4\x2b\x97\xd3\x40\x69\xc2\x18\x26\xff\x13\x93\x56\xd5\x05\xbf\x96\xba\x5e\x22\xd5\xd7\xae\x03\x1d\xd3\xdf\x94\xc0\xbe\x2a\xdd\x74\x0d\x19\x42\x74\x8d\x34\x4f\x6b\x26\x48\xd2\xc3\x98\x7d\x80\x41\x3e\x2b\xdb\x9d\x0d\x4a\xe4\xb1\x09\xde\xcf\x44\xa7\x61\x46\x79\xcd\x58\xd3\x6a\x84\x3c\xfa\xe9\x21\x9b\x99\x90\xfc\xf0\xe3\x7c\x0a\xf3\xa0\x13\x13\xba\xf1\xad\xa7\x2b\x98\x9f\x9d\xd3\x8c\x11\x46\xb6\xb0\x82\xc5\xbc\xb3\xc2\x02\x99\xc2\x61\xbc\xdf\x58\x77\x28\xb4\xf1\x37\x30\x7e\xd4\x5c\x6a\xdd\x6b\x24\xc3\x8c\x5a\x8c\xfa\x9a\x8f\x16\xdd\x99\x52\x1c\x60\x90\x2f\xe6\x3a\x97\x4a\x60\x02\x51\x64\x37\xe4\x21\xbd\x2f\x13\x5d\x05\x4c\x0f\xcf\x11\xe7\x99\x51\xb3\x19\x5c\x80\x52\x75\x87\xdf\x8f\xfd\x7f\x30\xde\x20\xdb\x81\xfb\x32\x53\xbb\x38\xb6\xda\xff\x1b\x0a\x34\xb5\xf3\x0c\xfb\x31\xb2\x9d\x8c\x5f\x81\x44\x46\x34\xdd\xa3\x79\xdf\xda\x1d\xee\x12\x16\x7c\x53\x92\xd0\x8e\xf3\xb3\x63\xdd\x37\xa3\x85\xd6\x0e\x5d\x7a\xf2\xf2\x96\xf7\x35\x68\xf1\x8b\xb6\xfb\x53\xa5\xd9\x52\xeb\xfd\x0c\xbe\x7c\xc7\x6f\x6b\x4d\x0f\xd9\x65\xc0\xbe\xef\xa1\x2f\x54\xe6\x8b\x50\xa9\xd0\xeb\xb2\x0d\x93\xf6\xfa\xcb\x0d\x54\x1f\xe0\x94\xb7\xbe\x08\xf5\xec\x6c\x7f\xb4\x45\x15\x33\xa1\x4c\x51\xb5\x7d\xaa\x92\xed\x55\x52\xaf\x67\xfa\x8f\xcd\xf4\xa1\x0d\xd9\x2b\x85\x9d\xc9\xd5\x46\x7c\x3d\x30\xe7\x5f\xfc\x77\x2e\x0e\xdc\xa4\x4a\xa3\xd7\x65\x51\xe7\xd9\x6f\xdf\xbb\x94\xdd\x27\x35\x5d\xaa\xe7\xf5\x35\xcf\xb6\x29\x9c\x94\xfd\xdd\xc4\x42\x1d\x54\xd1\x6a\xa3\xcb\xa8\x4e\x01\xb9\x46\x09\x84\x27\x20\x31\x13\x7b\x84\x84\x68\xd2\xea\x34\x4d\x9b\x99\xbc\x0e\x15\x32\x8c\xb5\xef\xfd\xa1\x7d\x69\x18\x54\x82\x3f\x31\xe6\x7b\xa7\x2d\xf2\x41\x3c\x7a\x41\x68\x74\x35\xfd\xb7\x73\xed\x90\x04\xc7\xda\xe5\x24\x2c\xaf\x02\xdc\x4b\x85\x24\x4c\x75\xc6\xfc\xbe\x06\xdd\x45\x7d\x03\x9f\x42\x1b\x1e\x3f\x08\x49\x9e\x23\x4f\x7c\x4f\x4b\x2f\x08\x2d\xc9\x62\xe2\x7b\x2d\xd5\xd5\xdd\xd7\x65\xbe\x87\xf8\x48\xb5\x1f\x84\x65\xd0\xfd\xe6\xd2\xb1\x68\x12\x76\xa0\x3c\x11\x87\x50\xa1\x7e\x6b\x7c\xd8\x13\xe6\xf7\x35\xec\xc9\xeb\xf0\x93\x12\xdc\xf7\x8e\xc7\xf0\x81\x28\x5c\xe7\x44\xa7\x45\x11\x35\x77\xfd\x66\xbd\x91\x7b\xf9\xff\xf6\xaf\x45\xe1\x4d\xcf\xee\x62\x1b\xef\x8a\x29\xbc\x99\xcf\xe7\xa7\xd2\x2b\x1b\x33\x06\x54\x03\xe5\x54\x53\xc2\xd8\xd3\xe8\xeb\x58\x5f\x46\xe5\xad\x71\xf5\x87\x90\xfa\x8f\x12\xd5\x43\xf5\xdf\xf1\x88\x3c\x29\x8a\xd1\x7f\x02\x00\x00\xff\xff\x83\xd5\x09\xb2\x8f\x19\x00\x00")

func templatesMonitorProcessorGoHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesMonitorProcessorGoHtml,
		"templates/monitor/processor.go.html",
	)
}

func templatesMonitorProcessorGoHtml() (*asset, error) {
	bytes, err := templatesMonitorProcessorGoHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/monitor/processor.go.html", size: 6543, mode: os.FileMode(420), modTime: time.Unix(1499244414, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesQueryIndexGoHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xcc\x56\x4d\x6f\xe3\x36\x10\xbd\xe7\x57\x4c\x08\xa3\x39\xb4\xb2\xd0\xe4\xd6\x50\xea\xa5\x39\x15\x28\x50\xf4\xd0\x63\x40\x89\x23\x89\x28\x4d\xaa\xe4\x30\x8e\x21\xe8\xbf\x2f\xf4\x19\x5b\x96\xed\x2c\x16\x0b\xec\x45\xc9\x68\x3e\x38\xef\xcd\xe3\x58\x4d\x23\xb1\x50\x06\x81\xe5\xd6\x10\x1a\x62\x6d\x7b\xc7\xa5\x7a\x83\x5c\x0b\xef\x13\xe6\xec\x9e\xa5\x77\x00\x00\xc7\x6f\xbb\x60\xa1\x0c\xba\xd1\xb7\xf4\x7f\x64\xad\x79\x73\xab\xa3\x9d\x8c\x7e\x7d\x5c\xc4\x00\xf0\xea\x31\xfd\x3b\xa0\x3b\xc0\x9f\x78\xf0\x3c\xae\x1e\x97\x11\x4d\xa3\x0a\xd8\xa2\x73\xd6\xb5\xed\x32\xfb\xe8\x0c\xa1\xd1\x11\xf4\xcf\x48\x0a\x53\xa2\x9b\x0c\xe5\x77\xca\x7b\x95\x69\x64\xe0\xac\xc6\x31\xf6\xac\x17\x00\x9e\x05\x22\x6b\x80\x0e\x35\x26\x6c\x30\xd8\x0c\x42\x5b\x8f\x0c\xa4\x20\x31\xd5\x9c\x2b\x71\x5f\x0b\x93\xfe\x44\x6a\x87\xfe\x99\xc7\xbd\xc5\xe3\xa1\xc0\xca\x31\x9e\x9c\x35\x65\xfa\xd2\x81\xba\xe7\xf1\x68\x42\xd3\x0c\x38\xb7\x2f\xeb\x68\x63\xa9\xde\xce\xe9\x41\x23\xcf\x42\x07\xd2\xf6\xc2\x19\x65\xca\x4f\xd3\x36\xc6\xff\xf0\xbc\xfd\x3b\xf4\x79\xca\xdc\xd8\xfc\xb7\x73\x37\xb1\xe7\x6d\x70\x39\xfa\x15\x3f\x2f\xac\xdb\x81\x35\x3e\x64\x3b\x45\x09\xdb\x2b\x23\xed\x7e\xab\x6d\x2e\x48\x59\x03\x09\x3c\x34\xcd\x36\x13\x1e\x5f\x6b\x41\x55\xdb\xc6\x4d\xb3\xf5\xa8\x31\x27\x94\xaf\x43\xdd\xb6\x8d\x1f\xe0\x67\xf0\x28\x5c\x5e\x6d\xdf\x84\x0e\xf8\x0c\x0e\x29\x38\x03\x85\xd0\x1e\x9f\x57\x78\x5e\x0e\x4f\x99\x3a\x50\x54\x3a\x1b\x6a\x38\xfa\x3f\xd2\xe5\x85\xe4\x2b\x05\xa2\x8c\xcc\x95\xac\x9b\x63\xce\xc8\x40\x46\x26\x92\x58\x88\xa0\x09\xa4\xb3\xb5\xb4\x7b\x13\x91\x2d\x4b\x3d\x09\x60\x30\x12\x36\x79\x59\xba\x46\x4d\x2f\x8b\x59\x3f\xc2\x61\x27\x95\x5b\xea\x38\xe9\x34\xe8\x29\x7d\xee\x63\x87\x26\xdc\x00\x08\xfd\xf0\x37\xf3\xe4\xe0\xb7\x04\x8e\xe7\xf8\x89\x64\xd7\xad\x1e\xd8\x28\x23\xf1\xfd\x17\xd8\x0c\x88\xfa\x3a\x97\xf5\x74\xd6\xbd\x56\x29\x17\x50\x39\x2c\x12\x76\xdc\x4f\xaf\xa4\xcd\xc4\x52\xc7\xdd\x6c\xf0\x58\xa4\x3c\xd6\xea\x33\x00\xd7\x55\x7f\xd2\x41\x1c\xf4\x15\x05\xad\x5d\xa6\x23\x77\xaf\x2a\x50\x32\x61\x83\xbc\xd9\xa8\x18\xc2\x77\x9a\xf5\xd2\xdd\xa1\xa8\xfb\x59\x71\x56\x33\x70\xf8\x7f\x50\x0e\xe5\xf7\x97\xed\x70\x69\x2f\xc9\x96\xa5\xff\xf4\x2d\xdf\x96\xd9\x15\x12\x2e\xb8\x78\xdc\x61\x3e\x7f\xdf\x34\xa8\x3d\xae\xed\x99\xeb\x7b\x7a\xb1\x95\xff\xb2\x30\x4a\x0c\x0a\x1b\x8c\xbc\x87\x3f\x94\x84\x83\x0d\x50\x58\x57\x22\x01\x59\x10\x44\x22\xaf\x80\x2a\xdc\xfd\x7e\xa1\xcb\x35\x79\x2c\x42\x17\xe6\xb0\x2e\xfb\x15\x76\x94\xc7\x2b\x17\xaf\x7f\x2b\xd4\xc2\xa0\x86\xfe\x39\x6f\x8b\xc1\xf2\x21\xcf\xd1\xfb\x2b\x5f\x12\x43\x5c\x85\x42\x76\x04\xac\x70\x5c\x3d\x9d\x86\x92\x22\x8d\xfd\x96\xf9\x0f\x0f\xdd\x35\xa9\x9e\xd2\x6b\xd8\xd6\x0f\xcc\xac\x3c\xac\x9d\x56\x3b\xec\x4a\x8f\xd8\x79\xdc\xd9\x5f\xc5\xdc\x07\xd7\xa3\x6b\xfc\x33\x79\xbe\x04\x00\x00\xff\xff\x05\xda\xbe\x9a\xaf\x09\x00\x00")

func templatesQueryIndexGoHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesQueryIndexGoHtml,
		"templates/query/index.go.html",
	)
}

func templatesQueryIndexGoHtml() (*asset, error) {
	bytes, err := templatesQueryIndexGoHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/query/index.go.html", size: 2479, mode: os.FileMode(420), modTime: time.Unix(1490714385, 0)}
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
	"templates/common/base.go.html": templatesCommonBaseGoHtml,
	"templates/common/head.go.html": templatesCommonHeadGoHtml,
	"templates/common/menu.go.html": templatesCommonMenuGoHtml,
	"templates/monitor/index.go.html": templatesMonitorIndexGoHtml,
	"templates/monitor/menu.go.html": templatesMonitorMenuGoHtml,
	"templates/monitor/processor.go.html": templatesMonitorProcessorGoHtml,
	"templates/query/index.go.html": templatesQueryIndexGoHtml,
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
	"templates": &bintree{nil, map[string]*bintree{
		"common": &bintree{nil, map[string]*bintree{
			"base.go.html": &bintree{templatesCommonBaseGoHtml, map[string]*bintree{}},
			"head.go.html": &bintree{templatesCommonHeadGoHtml, map[string]*bintree{}},
			"menu.go.html": &bintree{templatesCommonMenuGoHtml, map[string]*bintree{}},
		}},
		"monitor": &bintree{nil, map[string]*bintree{
			"index.go.html": &bintree{templatesMonitorIndexGoHtml, map[string]*bintree{}},
			"menu.go.html": &bintree{templatesMonitorMenuGoHtml, map[string]*bintree{}},
			"processor.go.html": &bintree{templatesMonitorProcessorGoHtml, map[string]*bintree{}},
		}},
		"query": &bintree{nil, map[string]*bintree{
			"index.go.html": &bintree{templatesQueryIndexGoHtml, map[string]*bintree{}},
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

