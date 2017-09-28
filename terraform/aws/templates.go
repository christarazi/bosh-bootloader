// Code generated by go-bindata.
// sources:
// templates/base.tf
// templates/cf_dns.tf
// templates/cf_lb.tf
// templates/concourse_lb.tf
// templates/lb_subnet.tf
// templates/ssl_certificate.tf
// DO NOT EDIT!

package aws

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

var _templatesBaseTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe4\x5b\x5f\x6f\xdb\x38\x12\x7f\x8e\x3f\x05\x21\xf4\xa1\xed\x39\x6a\x92\x6d\xb2\xbd\x60\xf3\x90\xb6\xb9\xbb\x1e\x7a\xdd\x22\x09\xf6\x1e\x8a\x42\xa0\x29\x5a\xe6\x45\x22\x05\x92\x72\xea\x1a\xfa\xee\x07\xfe\xd3\x7f\x59\xb2\x9b\x34\x09\x9a\xc5\xee\x26\xe4\x70\x66\xf8\xe3\x70\x66\x48\x0d\x39\x16\x2c\xe3\x08\x03\x0f\xde\x8a\x00\x93\xd4\x03\xde\xff\xb2\x24\x9d\xb1\x6f\xe6\xaf\xf5\x04\x80\x10\xa7\x98\x86\x22\x60\x14\x9c\x81\x2f\x9a\x92\x50\x89\x39\xc5\x32\x88\xa0\xc4\xb7\x70\xe5\x93\xc8\xfb\x3a\x01\x60\x99\x22\xa0\x7f\xce\x80\xe4\x19\x9e\xe4\x93\x49\x29\x42\xc6\x22\x48\x39\x59\x42\x89\x83\x1b\xbc\xf2\x80\x37\x63\x62\x11\x2c\x13\x61\xe4\xc0\x38\x62\x9c\xc8\x45\x02\xce\x80\x77\x79\x75\xee\x4d\x00\xe0\x02\x06\x33\x22\x05\x38\x03\xaf\x0f\xfe\x7e\x52\x67\xa8\x34\xb9\xc1\xab\x20\x85\x84\xb7\xb8\xa9\x0e\x0a\x13\xac\x98\x3d\x5b\x2f\x21\xf7\x31\x5d\x06\x24\xcc\x83\x82\x6e\x02\x40\x9a\xcd\x62\x82\x14\x17\x43\xd7\xd0\xd1\x77\xb4\x7e\x49\x18\xb0\x14\x53\x21\x16\xb9\xa7\xb4\x61\x99\x4c\x33\x59\x0a\x0f\x9c\x5c\xa3\xc5\x12\xc6\x99\x55\xa1\xaa\x6d\xc9\xd7\x91\xf7\x70\xab\xe1\xd5\x60\xd8\xaf\x6b\xd9\x18\xa4\x38\xc9\xd5\x44\x05\xa6\x82\x48\xb2\xc4\x95\xa5\x71\xd2\xf0\x37\xb5\x9a\x30\x0e\xdc\x8a\x37\xb4\xc6\x24\xf5\x2b\x56\xe1\xb0\x20\x69\x5d\x69\x47\x92\xf1\xd8\xb0\xd9\x82\xd1\xe9\xd1\x51\x8d\x57\x48\x38\x46\x92\xf1\x00\x86\x21\xc7\x42\x34\xf4\x5a\x48\x99\x8a\xd3\x57\xaf\x86\xd9\x1e\x1f\x1f\x1f\x7b\x6d\xb3\x21\x30\x09\x38\x8b\xb1\x35\x1b\xc3\x7e\x83\xb9\x68\x5a\x65\x2f\x50\x2e\x14\xc9\x2b\xf5\x47\x4c\xe6\x18\xad\x50\x8c\xed\x6c\x11\xc7\x0a\xf6\x19\x9e\x33\x8e\x83\x10\x0b\xc9\xd9\xca\xe1\x0d\x40\x3e\x51\x46\x2e\x44\x96\x60\xcd\x2f\x48\x59\x4c\x90\x22\xf8\xe3\x8f\x8b\x3f\xff\x31\x51\x4c\xbc\xbf\x30\x17\x84\x51\xef\x14\x78\x47\x07\x87\x47\xfb\x87\x07\xfb\x87\xbf\x7b\x53\xd5\x75\x25\xa1\xc4\x09\xa6\xd2\x3b\x05\x5f\xb4\x40\x23\x16\x00\xef\x1c\x49\x3b\x48\x48\x71\x7a\xae\x65\x5c\x2a\x95\xa7\x8e\xe2\x33\x27\x14\x91\x14\xc6\xde\x69\x31\x4c\xf1\xc4\x7c\x49\x10\x56\x23\x31\x3a\xf2\x61\x02\xbf\x33\x0a\x6f\x85\x8f\x58\xe2\x59\xb2\xbc\x60\x72\x31\x9f\x63\xa4\xc4\x7b\xe7\x71\xcc\x6e\x4b\xee\x57\x24\x54\xad\x66\x44\x3e\x01\xe0\xeb\x24\x9f\xa8\x39\x75\x02\x6f\xe6\xdd\x86\x1e\xf4\x80\x6f\xe9\x1d\xfc\xa0\x58\x80\x7b\x00\xf0\x4b\x89\x0d\x46\x47\x0a\x4a\x86\x08\x94\xf8\xdc\xda\xe1\xb4\xd1\x2f\x25\x44\x8b\xbf\x58\x9c\x25\xb8\xd9\xf7\x4e\x9b\x43\x77\xdf\x7b\x1c\x63\x89\xaf\x28\x4c\xc5\x82\xc9\xee\xde\xbe\x91\x02\x71\x32\x73\x0a\xe1\x96\x4a\x8e\xe0\x43\x02\xa3\x0d\xbd\x54\x48\x48\x51\x3f\xc1\x25\x8e\x08\xa3\xbd\xdd\x57\x18\x65\x9c\xc8\xd5\x3f\x39\xcb\xd2\x7e\x2a\x3b\xc1\x7e\x82\x6c\x46\x71\x7f\xb7\x81\xa0\xa3\x7b\x08\xf5\x3e\x64\x4d\xef\x35\x8c\x5a\x3c\x2f\x33\xda\x8b\xc9\x35\xe6\x09\xa1\x50\xf6\xa3\xa6\xd0\x12\x12\x73\x0d\x7a\x5b\x5d\x5e\xeb\x9e\xec\x01\xf0\x75\xaa\xfe\xdb\xb1\xa3\x54\xeb\xa5\xdd\x32\xaa\xfd\xa5\xdd\x54\xd3\xc9\xde\x5a\x77\x56\x4c\x75\x4f\x8b\x20\x30\x39\xfd\x0c\x85\xd0\x1b\x7e\x2b\xde\xc6\x7d\x3a\x5f\xa8\x03\x88\x0f\x39\xcd\xbd\xc9\xde\x06\x79\x38\x86\x42\x12\x14\x33\x18\xce\x60\x0c\x29\x22\x34\x3a\x7d\xb9\xf5\xac\xf6\x06\xfd\x44\xc5\x49\x06\x50\x6f\x34\xbd\x79\xab\x7e\x43\x91\x0c\xb9\x6c\xcb\x80\xd3\x32\x10\x95\x5e\xa8\x3a\xe9\x4e\x25\x88\x5d\xf2\x20\xe5\x6c\x4e\x1a\x11\xa3\x14\xdf\x06\xb2\x27\xaa\x77\xf3\xec\x88\xba\x5d\x84\x4d\xce\x4b\xc8\x09\x9c\xc5\x18\x78\x14\xca\x00\x26\x24\x48\xa0\x8d\xe1\x72\x95\x6a\x66\xaa\x61\xa2\xb3\xb8\x39\xcc\x62\x09\xce\x74\xef\x7a\xcd\x21\x8d\x30\x78\x76\x83\x57\x53\xf0\xcc\x88\x3e\x3d\x03\xfe\xf9\x7f\xaf\x3e\x9d\x5f\x9f\xff\xe7\x83\xc8\x73\x45\xa6\x08\xf2\x5c\x31\x5a\xaf\x0d\x59\xae\xf3\x89\xf5\x1a\xd3\x30\xcf\xf3\x36\x68\xc2\x7a\x86\x20\x52\xae\xc1\x33\xaa\x35\x1b\x4d\x5e\xa9\x36\x79\xaa\xac\xcb\xf0\xf7\x3f\x9d\x5f\xbf\x2f\x1b\x8d\xa0\x65\x8a\x02\x12\x9a\xdd\x54\x60\xb3\x4c\x91\xaf\xfe\x25\x61\xae\x27\x47\x68\xa4\x9c\xa1\xf5\xe7\x29\x67\x92\x21\x16\xdb\x21\x12\xa5\x66\x0f\xcd\x39\x53\xcb\xce\xa5\x6e\x3f\xd0\x6d\x92\xb9\x16\xd5\x76\x72\x7c\xfc\xdb\xb1\x6e\xaf\x2b\x2c\x74\xe6\x6b\x64\xd7\x7b\x7c\x93\x0a\xc3\xb8\xd5\x1e\xe6\x3a\x25\xce\x07\xf5\xcb\xc2\xc7\xad\x1f\x41\x49\xa7\x82\xfb\x87\x1d\x1a\xda\xc6\xbb\x55\x0f\x57\xb5\x2b\x95\x68\x62\xe4\xfe\x2e\xf4\x3f\x03\xde\xfe\xa1\x51\x1d\x91\x90\x07\xb3\x98\xa1\x1b\xa3\xcc\x81\xaf\xff\x79\x75\x50\x4a\x91\x30\x72\x32\x3e\x75\x65\x82\xfb\x14\xca\x7d\xa7\xe6\xbe\x31\x64\x3d\xb6\xed\x37\xec\xb6\x35\xc6\x6f\xac\xdd\x25\xe5\x24\x05\xd5\x9f\x33\xe0\x1d\x1a\x5d\x7e\xf7\xf4\x42\xd8\x0d\xaf\xb7\x6f\x95\x4a\x1e\xf9\x09\x0e\x49\xa6\xd3\x32\xa1\xc3\x66\xb1\x2f\x2a\x64\x16\x64\xdd\x6f\x4e\x04\xf6\x77\xbd\x55\x00\x30\x6a\xea\xe4\x34\x40\x0b\x8c\x6e\xdc\xc8\x39\x8c\x85\xca\x52\x61\x42\x40\xc7\x8f\x66\x1d\x33\x76\x93\xa5\xcf\x15\x26\x15\x7f\x33\x05\xaa\x81\xeb\x7c\xe1\x45\xb1\x67\xeb\xeb\x19\x90\x70\x83\x11\xb4\x3d\x84\x5d\xff\x91\xab\x62\x03\xee\x05\x5d\x7e\x78\xdf\x22\xe8\x59\x23\x73\xd8\x55\x92\x77\x39\xe8\xba\x75\xaa\xb8\x6b\xdb\xa2\x26\xe3\xc0\xee\x38\x0e\xbb\x58\x50\x13\xdc\x71\x4c\xb2\xfd\xcd\xb3\x56\xe9\xf0\x21\x42\x58\x88\xf2\x60\xe8\xfc\xbd\x90\x9c\xd0\xa8\x41\x2c\x30\xe2\x58\x8e\x24\x36\x2b\xd9\x4b\x98\x72\xb6\x24\x21\xe6\x1a\x46\x7b\x72\x2f\x74\x29\xd1\x2f\xdb\xec\xf9\xd3\x69\x50\x92\x94\x6d\x9a\xc4\xc8\x2d\xad\xad\xb4\xaa\xae\xd8\x6c\xe3\x59\x3b\xdc\xf4\x75\xac\x27\x7b\x36\x96\x74\x87\x91\xe1\x40\xd6\xe3\xab\xfa\xa2\xd9\x07\x4b\xbe\x5b\x48\x1b\xb4\x7b\xa7\xcd\x38\x97\xd4\xd8\x8e\x3c\xd3\x99\x4c\xcf\x8c\x74\x77\xa0\xe2\xa6\xd6\xa0\xb5\x95\xdb\xfe\x66\xbc\x53\x77\x46\xd5\xe3\x64\x6c\x24\xd2\x89\x5b\x25\x0c\x35\xc9\x6c\x4c\xaf\x06\xa4\x06\x89\x0a\x05\xd5\xd0\xd4\xe8\x76\x91\x54\xe0\x78\xde\xa3\x4b\xfb\x06\x6b\x47\x20\x55\x80\x7f\xac\x40\xda\xe4\xe3\x69\x00\xa9\x33\x91\xc7\x8a\xa4\x4b\x93\x36\x40\xa9\x93\xa3\x0d\x58\xea\xfe\x6a\xa6\xd2\xe8\xaf\xa7\x2d\x77\x81\x28\x54\xe7\xb3\x22\xca\xfd\x7c\x6c\xf1\x28\x68\x4d\x12\xb7\xbb\x8d\x1e\xfc\x6c\x58\x85\x3b\x20\x3e\x42\x3b\xbd\x7e\xf7\x79\x00\xcd\xa3\xa3\xcd\x70\xea\x7e\x9b\x48\xb6\x27\xd8\x37\x33\x7b\x43\x5b\x04\x5b\x97\x09\x6d\x8c\xaa\x3a\x33\x3a\xdb\x01\xaa\x5a\x46\x63\x0e\xde\x74\xc6\x32\x1a\x06\xca\x10\x5c\xc8\x76\x47\xe2\x8a\x01\x8c\xc8\x03\x4c\x5a\x3d\x2a\x07\x78\xfb\xe7\xd5\xbf\xee\x29\xfe\x2b\x2d\xfa\x62\x7f\xed\xc6\x61\x5b\x5c\x3b\x06\x8d\xca\x90\xdc\xce\xe8\x18\x5f\x24\x14\x3f\xb0\x33\x7a\xd5\xfa\x49\x09\xc5\xa8\x5d\xb1\xd1\xcb\x98\xf5\x6b\x19\x63\x3e\xde\xe9\x6c\x84\x56\x77\xc2\x48\x5f\x91\x3d\x49\x84\x4f\xde\x9c\xbc\x19\x48\x36\x0c\xc5\x43\xa1\x9c\x41\xf8\x44\xa1\x7d\xf3\xfa\xf5\x6f\x9b\xa1\xb5\x14\x0f\x69\xc0\xe5\x47\xbf\x94\x3c\x51\x9c\xf5\xf7\xc6\x01\x3f\x61\x49\x1e\x10\xe9\x27\x0a\xee\xd8\x93\xc8\xb6\x99\xc9\x50\x22\xf1\x63\x3e\x63\xe7\xc3\xdf\xfd\xc2\x7d\x77\x07\xbf\x47\x05\xf7\x9d\x1c\x68\x76\x44\xfe\xe9\x1d\x66\xca\x72\x9f\xce\x04\x16\x66\x92\x25\x50\x12\x04\xe3\x78\x65\xcb\x1b\x42\x60\x47\x80\xd9\x0a\xbc\x7d\xfb\xf1\xee\x12\x5a\xcb\x77\x28\xa7\x75\x95\x1e\xdb\xa6\xb5\xcd\xf3\xc7\x18\x33\x2b\x64\xed\x9c\xb5\xd6\xa4\xfe\x42\x99\xaa\x43\xee\x47\xf2\xd1\x87\xc0\xee\xb1\xe4\xa0\x0e\x3f\x97\x13\x3d\x21\x08\x1f\x4f\x0e\x54\x14\x85\x45\x65\x0d\xd9\x7d\x42\xf8\xf4\x02\x40\x2d\x9e\xb6\x03\xf3\x2f\xf4\x1d\x60\xdb\x2c\xe6\x4e\x6e\x49\x7a\x10\xff\x35\x3e\x18\xdc\x25\xe2\x8d\xbb\x3e\xfb\x7d\xbc\xbc\xea\x73\x13\xaf\x7c\xca\xac\xdd\xff\x1d\xba\x3d\x73\xf4\xba\x8b\x1f\x5c\x42\x12\xc3\x19\x89\x95\xe4\xef\x8c\xe2\xde\xaf\xa3\x8d\xa5\xd7\x7a\x78\x35\xad\x6c\xa6\x52\x49\x99\x1a\x4b\x5a\x4b\x9c\xaa\x5b\xbd\x46\x59\xf8\xc6\xca\x5c\xb7\xba\x39\x34\xca\x74\x7f\x2d\xe4\x2c\x93\x38\x90\x0a\x01\xa7\x7c\xad\xa9\x32\x83\x91\x1f\x53\xf5\xf0\x5e\x5e\x21\x16\x92\x50\xa8\xf2\xce\xa0\x32\xdf\xfa\xc5\x2c\x00\xf6\x23\x7c\x4d\x6c\xc7\x17\x7a\x87\x5c\x45\x4c\x6d\x48\xa5\xdd\x6f\xea\xb3\x49\x7d\xcb\x0a\xda\x9a\x50\xfd\x91\xdc\x33\x3d\x95\x95\x70\x71\xa6\x5e\xa2\x31\xa2\x34\xe3\x87\xd4\xad\xdf\xf8\x3a\xd9\x9d\xe5\x05\x7d\x1a\xf4\x70\xe9\xb1\xfd\x61\xa6\xad\x81\xad\x12\x86\x26\x81\xa8\x6f\xab\x98\x08\xb9\x69\x53\x95\xce\xae\x0a\x3c\x62\x19\x95\x4d\xaf\xf5\x6c\x1d\x63\x1a\xc9\x85\xae\x5d\x69\xcb\x7d\xd1\xba\x97\xdf\x69\x4f\xaa\x66\xa3\xcb\xf3\xd2\xa3\x1c\x9e\x78\x53\xf0\x7a\x6a\xf4\xf2\x09\x0d\xf1\xb7\xbf\x1d\x1a\x81\x2d\x45\x0c\x1b\x1c\xeb\x2a\xe5\x1e\x5d\x6b\x9c\x5e\x6c\x5d\x2c\xa0\xd5\x7b\xb6\xae\xf0\xb0\x15\x32\x1d\x05\xed\x24\xa2\x8c\xe3\x00\x2d\x20\x8d\xb0\xa9\xdf\x29\x67\xee\x4d\x3b\x56\xd0\x56\x53\x0d\x38\x94\x62\xe1\xee\xc8\xa9\xf4\xf3\x1b\xe9\x58\x8a\xb2\xab\xba\x67\x69\x57\xf4\x8c\xd9\xa3\x5d\xda\xec\xe8\x56\x46\x59\xf8\x58\xf3\xee\xf2\x48\xce\xd4\x2a\x5b\xb8\x29\xd3\x7f\xe9\x93\xb0\x65\x74\x77\x00\x45\xeb\x03\x20\xfc\x5e\x7a\xae\x20\x81\x69\xaa\x82\xaa\xae\xdb\x29\x7d\xcd\x64\x0f\x80\xef\x24\x4d\x60\xfa\xbc\xee\x79\x3a\xd4\xee\x70\x40\x53\x30\x38\x4a\xa9\xf7\x62\xb2\x37\xa8\xa3\x36\xa7\x07\xd3\xb2\x34\xe6\x8a\xb6\xa5\x6b\x35\xfb\x7e\x4c\x65\xd8\x82\x71\x19\x8c\x26\x57\xfb\xb3\x91\x55\x0d\xa6\x54\x87\x27\x1d\x96\xbf\x4c\x91\xa7\xd9\x59\x93\x6e\x39\xd4\x6a\x9e\xe3\xa4\xe6\xf5\x12\x49\x4c\x21\x45\x2b\x47\x6a\x45\x2b\x12\x4c\xb5\x55\x86\x54\x04\x0b\x26\x24\x85\x89\xf6\x5e\xba\x0e\x65\x8c\xb7\x54\x6a\xf5\x55\x76\xd6\xb3\x0d\xe5\x7c\xa2\x71\xae\xcb\x99\x92\xa1\xeb\x0c\xa2\x9b\xbd\xdd\x3c\x66\xb7\x41\xcc\x22\x95\x45\xcd\xec\x13\xab\x98\x45\x36\x71\x2e\x1f\x2f\x29\x5a\x14\xb3\x2c\xbc\x85\x12\x2d\x82\x82\xc4\x9f\xcd\x62\x57\x3b\x0e\x40\x51\x60\x0f\x39\xad\x46\xba\xa2\x88\xdd\x89\x13\xb6\x3a\xbe\x15\x1f\xfb\x82\xa3\xe4\x70\x3e\x27\xc8\x95\xb0\x9e\x01\xef\xf2\xe2\xdf\x17\xef\xae\x3b\xa6\xd4\xa5\x66\x75\x7a\x4a\xdb\x20\xe5\x78\x4e\xbe\x55\xca\x06\x2b\x26\x9b\xef\xc7\x2c\x72\x17\x85\x9b\x5e\x79\x15\xb3\xd9\xf0\xd4\x6b\x5f\x11\x29\x86\x62\xdf\x3c\x1e\xb8\xb7\xf7\x5a\xee\xbd\xd4\xf0\xcb\xaa\xe1\x77\x5b\xcb\x14\x95\x8a\x0f\xbd\xe0\xea\x7d\x28\x36\xee\xe5\x56\x05\x86\xed\x31\x2d\x9f\x71\xf5\x3c\x9b\x28\x2d\xce\xdd\x19\xdf\xef\x03\x2f\x25\xca\xbe\x08\xfa\xc8\x22\xfd\x92\xa9\xfa\x74\xa7\xde\x7d\x25\x39\x86\x49\xab\xff\x73\x26\x3f\xb2\xe8\x62\x89\x69\xfd\x15\x93\xee\x74\xcf\x98\x1c\xf7\x8d\x14\x46\x80\x70\x6b\xf6\x75\xd8\x36\xba\x9e\x09\x6d\x5a\xc1\x9b\xc4\xd6\x0b\x7b\xc5\x6f\xeb\xd2\x5b\xde\xe0\x55\xc0\x99\x84\xf6\xf2\xbf\x59\xb0\x6c\x87\x28\x77\xd1\xfd\xb4\xd5\xf4\xfb\xee\xff\xee\x49\xcd\xff\x03\x00\x00\xff\xff\x6b\x90\xe2\x47\x63\x3c\x00\x00")

func templatesBaseTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesBaseTf,
		"templates/base.tf",
	)
}

func templatesBaseTf() (*asset, error) {
	bytes, err := templatesBaseTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/base.tf", size: 15459, mode: os.FileMode(420), modTime: time.Unix(1506547663, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesCf_dnsTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x91\xc1\x6a\xf3\x30\x0c\xc7\xef\x79\x0a\x61\xbe\xd3\x07\x35\x85\xb2\x63\x0f\x65\xec\xb8\xbe\xc0\x18\xc6\xb1\xb5\xc6\xc3\x89\x8d\x65\xa7\xeb\x8a\xdf\x7d\x38\xc9\xa0\xdd\xc6\xc8\x20\xbd\x25\x46\xfa\xff\xf4\x93\x7a\x19\x8c\xac\x2d\x02\xa3\x13\x45\x6c\x85\x76\xad\x34\x1d\x83\x73\x05\x10\x4f\x1e\x61\x0b\x8c\x62\x30\xdd\x81\x55\xb9\xaa\x02\x92\x4b\x41\x21\x30\x79\x24\x11\x5c\x8a\x78\xb7\x11\xef\xae\x43\x06\x0c\xbb\x5e\xe8\x8e\xa6\xdf\x92\xd0\xc9\x76\x48\xf8\x77\xee\x65\xe0\x57\x88\xcc\xaa\x82\x90\x07\x1a\x2a\x01\xf6\x57\xb5\x25\xcb\xe8\xbc\x6a\x1c\x45\xd4\xab\x21\xb2\x02\xc8\x65\x08\x97\xa2\x4f\xf1\x9a\x27\x0a\x4a\x10\x86\x1e\x03\x8d\xf0\x5e\xda\x34\x25\x7e\x1d\x96\x5f\xb6\xf2\xcb\xd6\xfc\x8b\x66\x40\xe5\x82\x66\xc0\x8e\xc6\x6a\x25\x83\x2e\x11\x23\x6b\x18\xc1\xe8\x39\x34\xa3\x33\xfb\x5c\x0d\x40\xe9\xf8\xcf\x7f\xde\xcf\x74\x81\xb1\xe8\x7e\xbf\x7b\x7c\x18\xde\xa2\x85\xf1\x6d\xb3\x5e\x97\x1d\x8e\x63\x11\x6c\xe1\x69\x82\xa3\xad\xb9\x7a\x19\x67\x08\xc2\xd6\xbc\xc0\x0b\x30\xb3\xe7\x19\x7a\x44\xcd\x02\x56\x44\xcd\x8d\xbc\x88\x9a\xbf\x4b\xd5\x6e\x11\xab\x12\x33\x47\x6b\x37\x57\xc9\x78\xfe\x9a\x5a\x5f\xbb\xb7\xe1\xdb\xa7\xda\x1a\x25\x8c\x9f\x67\x15\x95\x5f\x40\x2a\x2a\x7f\xa3\x53\x45\xe5\xbf\x9f\xea\x23\x00\x00\xff\xff\x30\x40\x15\x44\x75\x04\x00\x00")

func templatesCf_dnsTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesCf_dnsTf,
		"templates/cf_dns.tf",
	)
}

func templatesCf_dnsTf() (*asset, error) {
	bytes, err := templatesCf_dnsTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/cf_dns.tf", size: 1141, mode: os.FileMode(420), modTime: time.Unix(1506547663, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesCf_lbTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\xdc\x5b\x6b\xe3\x46\x14\x07\xf0\x77\x7f\x0a\x61\xfa\x54\x88\xeb\xb1\xee\x05\xbf\x74\xfb\xd0\x42\x29\x4b\x77\xdf\x4a\x11\xb2\x3c\x89\xc5\x2a\x92\x99\x19\xbb\x6c\x83\xbf\x7b\xd1\xd5\xce\xc6\xd6\xe5\xbf\xff\x25\x69\x53\xfa\x10\x49\x67\xe6\x8c\x74\xf4\xd3\xb1\x89\x56\x49\x5d\x1c\x54\x22\xad\x79\xfc\xb7\x8e\xb4\x4c\x0e\x2a\x35\x9f\xa3\x07\x55\x1c\xf6\x73\x6b\x9e\xdc\x47\x5a\xef\xa2\x6c\xf3\x62\xd7\xd3\xcc\xb2\xb6\x52\x27\x2a\xdd\x9b\xb4\xc8\xad\xb5\x35\x7f\x7a\x5a\x7c\xf8\xf0\xcb\x6f\x3f\xfd\x7c\xde\x7c\x3a\xcd\x67\x96\x75\xdc\x27\x51\xba\xb5\xaa\x9f\xb5\x35\xff\xee\xa9\x9c\xeb\xb8\x4f\x16\xe5\xff\xe9\xf6\x34\x9f\xcd\x2c\x2b\xcd\x1f\x94\xd4\xba\x1a\xd8\xb2\x92\x74\xab\xa2\x4d\x56\x24\x9f\xb4\xb5\xb6\xfe\x9c\x2f\x17\xd5\x7f\x3f\x2c\xe7\x7f\x55\xfb\xf7\xaa\x30\x45\x52\x64\xcd\x90\x26\xd9\xcf\xab\xed\xf7\xaa\x78\x8c\xf6\x85\x32\xd5\xf6\xd5\x6a\xb5\xaa\x36\x9b\xa2\xdd\x78\xb1\xf9\x54\x4e\x2b\x2f\x67\x3d\x47\xaf\xad\xe5\xb3\xc0\xf6\xf7\x6e\xde\xb5\x35\xbf\x13\xf3\x11\xb9\x56\xb3\x98\xf8\xa1\x9d\xe3\xf7\xf8\x51\xd6\x67\xe1\x18\xab\x85\xcc\x8f\x51\xba\x3d\xdd\x25\xf7\x77\x5a\xef\xee\xb2\xcd\x5d\x7b\xa2\xef\xea\x13\x5d\x8d\x70\x9a\xcd\x8a\x83\xd9\x1f\xcc\xd0\x15\x39\xc6\xd9\x41\xae\x9b\x33\xfc\xfc\x80\xc5\xad\xc8\xfa\x0a\x9c\x66\xb3\xd1\xb5\x90\xe6\x46\xaa\x3c\xce\xa6\x14\xc5\xaf\x4d\x0c\xa3\x38\x9e\x4f\x5b\x9f\xf4\xe9\x4b\xfe\xbf\x17\x52\x7b\x95\xc6\x57\x54\xef\x75\x1d\x57\x5a\x37\x86\xb8\x51\x63\x32\xdb\x5c\x16\x56\x3d\x51\x5e\xae\xeb\xea\x4f\xb7\x58\xbd\x2b\x94\x89\x5e\x2c\xb9\x5c\x5a\xa2\x0a\xad\xa3\x7f\x8a\x5c\x46\x59\x11\x6f\xa3\x4d\x9c\xc5\x79\x92\xe6\x0f\xd6\xda\x32\xea\x20\xcb\x93\xb8\x93\x71\x66\x76\x51\xb2\x93\xc9\xa7\xe6\x64\xd6\x9b\x3e\x47\x66\xa7\xa4\xde\x15\xd9\xb6\x9a\xce\xad\xf6\x1d\xf2\x97\x7b\xd7\x56\x5d\x0f\xd5\x7a\x8f\x71\xf6\x3c\x4d\xaf\xbe\xe4\xb1\x7a\x90\xe6\xc5\x12\x3e\xbe\x7b\xff\x63\x59\x38\xf5\x35\x37\xe9\xa3\x2c\x0e\xe6\x8b\x83\xba\xaa\xca\x52\x6d\x64\x2e\x55\x93\x66\x9a\x6b\x13\xe7\x89\xbc\x52\x84\x97\x3b\x2f\x6a\xab\x2b\xe8\x6c\x73\x0e\xb2\xbe\x0c\x2d\x77\x5e\xdc\x08\xcf\xee\x85\x2a\x0f\xde\x2d\xa7\x0f\x9b\x5c\x1a\x7d\x91\x45\x37\x52\xb5\x67\x51\x86\xd6\xc7\x2c\xbe\x6f\xa2\xae\x56\x6b\x59\x27\x17\xa5\x79\xc6\x43\x66\x9b\x73\x1a\x8b\xf2\xb0\xba\xf6\x5e\x0e\x71\x50\xd9\x88\x11\xb6\xb9\x8e\xce\xa3\x0c\x2b\xa9\x8a\x83\x91\x6a\xfc\x43\xf3\x8f\xea\xf8\xb7\xf3\xd4\x0c\x96\x57\xa8\xab\x36\x9e\xbe\xd5\x94\x8e\x63\x5f\x99\xb3\xde\xfa\x0d\x27\xbd\x31\xeb\x79\xda\xb7\x83\x7a\x5d\x54\xe3\x1a\x84\xfe\x02\x1c\x80\xfc\x56\xf0\x84\x36\xe1\x3c\xc4\xc4\x4e\xa1\xbe\x13\x5e\xab\x55\xe8\x5d\x39\xf1\x06\x7a\x8b\x45\x35\xa1\x59\x18\x79\x75\x47\x97\x19\xd8\x32\x74\x03\xe0\x5d\x43\xb7\xfc\x37\xd3\x38\x88\xd5\x50\xe7\x10\x2c\x59\x7d\x43\x53\xa5\x57\xbb\x86\x9d\x31\x3d\x6d\x43\x13\x79\xb5\x69\x68\x23\xc7\x65\xd1\x97\xc6\x50\x1e\x17\x8f\x8d\x97\x99\xb4\xc1\xba\x8e\xd6\x3a\x8b\x12\xa9\x4c\x7a\x9f\x26\xb1\x91\xa5\x22\x1d\x20\x69\xfc\x18\x69\xa9\x8e\x52\x5d\x1e\x52\xb6\x21\xe5\xaf\x8b\x58\xe5\x27\xde\x82\x7a\xda\xb1\xcb\x27\xd2\xf5\x05\x69\x9d\x71\x97\x43\xd5\xf1\xeb\x1b\xbb\xf3\x14\x43\xbd\x5d\x77\xe4\xf5\xf6\xee\x3c\xd0\x40\x87\x77\x1e\x67\x6a\x93\x67\x92\xfd\xf8\x0e\xef\xe3\xbb\xf7\x6f\xe9\x6b\x11\xb1\x5c\x39\x57\x9e\x50\x42\xac\xde\x60\xe3\x63\x92\xfd\xb8\xae\xa7\xe7\x8a\x0c\x3c\x8b\xae\x46\x4e\xe8\x77\x9a\xf8\x89\xcd\x4e\x55\x14\xaf\xd5\xeb\xdc\x5e\x32\xb9\x90\x5e\x3b\xc5\xff\x4c\x2f\xd6\xd4\xf9\x84\x46\x6c\x4c\xd9\x8d\xab\x7c\xb0\x05\xab\xa3\xf1\xfe\xab\x5e\x32\xbd\xf9\xf2\x7a\x9a\x2f\xbb\xa7\xf9\x72\xbf\xae\xf7\xb2\x27\xf4\x5e\xdd\x8d\x33\xfd\x3b\x9b\x2e\x74\xf0\x3b\x9b\x71\x79\xb8\x78\x1e\x2e\x33\x0f\x0f\xcf\xc3\x63\xe6\xe1\xe3\x79\xf8\xcc\x3c\x02\x3c\x8f\x80\x99\x47\x88\xe7\x11\x12\xf3\xb0\x7b\x3e\xad\x0c\xe4\x61\xf7\x7c\x5c\x99\x9e\x87\xc0\xf3\x10\xcc\x3c\xd0\xef\x7c\xbb\x50\x52\x1e\x36\x9e\xc7\xad\xcf\x3a\x50\x1e\xb8\xa7\x36\xd3\x53\x1b\xf7\xd4\x66\x7a\x6a\xe3\x9e\xda\x4c\x4f\x6d\xdc\x53\x9b\xe9\xa9\x8d\x7b\x6a\x33\x3d\xb5\x71\x4f\x6d\xa6\xa7\x0e\xee\xa9\xc3\xf4\xd4\xc1\x3d\x75\x98\x9e\x3a\xb8\xa7\x0e\xd3\x53\x07\xf7\xf4\xe6\x77\x47\x50\x1e\xb8\xa7\x0e\xd3\x53\x07\xf7\xd4\x61\x7a\xea\xe0\x9e\x3a\x4c\x4f\x1d\xdc\x53\x87\xe9\xa9\x83\x7b\xea\x30\x3d\x75\x70\x4f\x1d\xa6\xa7\x2e\xee\xa9\xcb\xf4\xd4\xc5\x3d\x75\x99\x9e\xba\xb8\xa7\x2e\xd3\x53\x17\xf7\xd4\x65\x7a\xea\xe2\x9e\xba\x4c\x4f\x5d\xdc\x53\x97\xe9\xa9\x8b\x7b\xea\x32\x3d\x75\x71\x4f\x5d\xa6\xa7\x2e\xee\xa9\xcb\xf4\xd4\xc5\x3d\x75\x99\x9e\x7a\xb8\xa7\x1e\xd3\x53\x0f\xf7\xd4\x63\x7a\xea\xe1\x9e\x7a\x4c\x4f\x3d\xdc\x53\x8f\xe9\xa9\x87\x7b\xea\x31\x3d\xf5\x70\x4f\x3d\xa6\xa7\x1e\xee\xa9\xc7\xf4\xd4\xc3\x3d\xf5\x98\x9e\x7a\xb8\xa7\x1e\xd3\x53\x0f\xf7\xd4\x63\x7a\xea\xe3\x9e\xfa\x4c\x4f\x7d\xdc\x53\x9f\xe9\xa9\x8f\x7b\xea\x33\x3d\xf5\x71\x4f\x7d\xa6\xa7\x3e\xee\xa9\xcf\xf4\xd4\xc7\x3d\xf5\x99\x9e\xfa\xb8\xa7\x3e\xd3\x53\x1f\xf7\xd4\x67\x7a\xea\xe3\x9e\xfa\x4c\x4f\x7d\xdc\x53\x9f\xe9\x69\xdf\x5f\xcf\x0d\xe4\xd1\xf7\xe7\x73\xd3\xf3\xc0\x3d\x0d\x98\x9e\x06\xb8\xa7\x01\xd3\xd3\x00\xf7\x34\x60\x7a\x1a\xe0\x9e\x06\x4c\x4f\x03\xdc\xd3\x80\xe9\x69\x80\x7b\x1a\x30\x3d\x0d\x70\x4f\x03\xa6\xa7\x01\xee\x69\xc0\xf4\x34\xc0\x3d\x0d\x98\x9e\x86\xb8\xa7\x21\xd3\xd3\x10\xf7\x34\x64\x7a\x1a\xe2\x9e\x86\x4c\x4f\x43\xdc\xd3\x90\xe9\x69\x88\x7b\x1a\x32\x3d\x0d\x71\x4f\x43\xa6\xa7\x21\xee\x69\xc8\xf4\x34\xc4\x3d\x0d\x99\x9e\x86\xb8\xa7\x21\xd3\xd3\x10\xf7\x34\x24\x7a\x2a\x96\xb0\xa7\x6d\x28\x29\x0f\xd8\xd3\x36\x94\x94\x07\xec\x69\x1b\x4a\xca\x03\xf6\xb4\x0d\x25\xe5\x01\x7b\xda\x86\x92\xf2\x80\x3d\x6d\x43\x49\x79\xc0\x9e\xb6\xa1\xa4\x3c\x60\x4f\xdb\x50\x52\x1e\xb0\xa7\x6d\x28\x29\x0f\xd8\xd3\x36\x94\x93\x87\xc0\x3d\x15\x4c\x4f\x05\xee\xa9\x60\x7a\x2a\x70\x4f\x05\xd3\x53\x81\x7b\x2a\x98\x9e\x0a\xdc\x53\xc1\xf4\x54\xe0\x9e\x0a\xa6\xa7\x02\xf7\x54\x30\x3d\x15\xb8\xa7\x82\xe9\xa9\xc0\x3d\x15\x4c\x4f\x05\xee\xa9\x60\x7a\xba\xc2\x3d\x5d\x31\x3d\x5d\xe1\x9e\xae\x98\x9e\xc2\xff\x86\x4b\x17\x4a\xca\x03\xf7\x74\x35\xd2\x53\xde\x4b\x80\x5f\xff\xca\x71\x33\xfe\xd0\xfb\xc6\xf5\x61\xd7\x5f\x36\x6e\x86\x18\x78\xd3\xb8\x19\xe1\xd9\x6b\xc6\xff\x06\x00\x00\xff\xff\x86\x35\x6c\xe5\x7d\x4d\x00\x00")

func templatesCf_lbTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesCf_lbTf,
		"templates/cf_lb.tf",
	)
}

func templatesCf_lbTf() (*asset, error) {
	bytes, err := templatesCf_lbTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/cf_lb.tf", size: 19837, mode: os.FileMode(420), modTime: time.Unix(1506547663, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesConcourse_lbTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x55\xc1\x6e\x9c\x30\x10\xbd\xf3\x15\x96\xd5\x53\xa5\xa5\x9b\x4d\x2a\x55\x95\x38\xa5\x97\x5e\xaa\x1e\x7a\xab\x2a\xcb\x98\xc9\x62\xc5\x6b\xa3\xb1\x4d\x95\xae\xf8\xf7\x0a\x1b\x58\x08\x6c\xc2\x56\x89\x1a\x56\x7b\x60\xc6\x33\xef\xcd\xf8\xcd\x80\x60\x8d\x47\x01\x84\xf2\xdf\x96\x59\x10\x1e\xa5\x7b\x60\x7b\x34\xbe\xa2\x84\x0a\xa3\x85\xf1\x68\x81\xa9\x7c\xe6\x3d\x26\x84\x14\x60\x05\xca\xca\x49\xa3\x49\x46\xe8\xf1\x98\xde\xf6\x21\x5f\x4e\xae\xa6\xa1\x09\x21\x75\x25\x98\x2c\x48\x78\x32\x42\xdf\x1d\x5b\xc8\xba\x12\x69\xfb\x97\x45\x43\x93\x84\x10\xa9\xf7\x08\xd6\x86\xe4\x84\x08\x59\x20\xcb\x95\x11\xf7\x96\x64\xe4\x27\xdd\xa6\xe1\xf7\x61\x4b\x7f\x05\x7f\x85\xc6\x19\x61\x54\x97\xd2\x89\x8a\x06\xfb\x1d\x9a\x03\xab\x0c\xba\x60\xff\xb4\x0d\x46\x67\x7a\xd3\x60\x6c\x5e\x0b\x72\xb7\xdb\xed\x16\x40\x3b\xf3\xab\xc1\xde\xdc\x5c\x2f\xa0\x46\x6b\x00\x85\x31\xe6\x29\x36\x23\xd3\x16\xf5\xef\x03\x6a\x46\xe8\xe6\x8a\xae\x60\x1a\x50\x1c\xdf\xf7\x18\xdf\xf8\x01\xe2\x6d\xd7\x1c\x53\xd0\x35\x93\x45\xb3\x19\x74\xb5\x51\xf9\xa6\xd7\xd5\x26\xea\x2a\x24\x69\x92\xe4\x12\x69\x4a\xed\x00\x35\x57\x97\x6a\xf4\x6b\x17\xf7\x12\x5a\x9d\x42\xc7\xde\xc4\xb8\xa9\x27\x7d\x62\xa8\x42\xea\xcb\xa4\x7d\x46\xdc\xe7\xe4\xfd\x5f\x58\xae\x98\x86\xb7\x28\xcc\x5e\x55\xe7\x14\x6a\xbc\xab\xbc\xbb\x44\x8a\x35\x57\x1e\xb2\x15\x0d\x3f\x93\x25\x4a\x6f\x36\x1c\xa0\xf2\x47\x13\x11\xe1\x74\x5b\xe3\xe2\x33\x14\x6e\x4b\x83\x8e\x2d\x95\xdf\x96\x29\xd0\x58\xcb\xfe\x18\x0d\x4c\x19\x5e\xb0\x9c\x2b\xae\x85\xd4\x7b\x92\x11\x87\x1e\xda\x9e\x96\xc0\x95\x2b\x99\x28\x41\xdc\x77\xbd\x8d\xa6\x07\xe6\x4a\x04\x5b\x1a\x55\xc4\xeb\x0e\x3e\xaf\xe7\xde\x8c\x5c\xc5\x6b\x0d\x65\xd7\x5c\x4d\xa9\x5e\x77\x1a\xe0\xb8\x07\x37\xab\xe3\xc7\xed\xf7\xcf\xad\xde\xa3\x08\x9c\x3c\x80\xf1\xee\xd1\xa1\x8f\xbd\x00\x94\xb4\x0e\x34\x60\x47\x54\x6a\xeb\xb8\x16\xb0\x30\x3b\x63\xe7\x48\x6c\x83\xc2\x55\x7e\x0a\x22\xe3\x6f\x4a\x74\x8d\xe6\x62\x32\x1a\xeb\x58\x8c\x67\x66\x4e\xe3\x19\x1e\xe3\xe0\x39\x95\x7f\xe1\xf2\x44\x4b\x9e\xe7\xd2\x7f\x93\x96\xa9\x58\xab\x62\xac\xb5\x8a\x09\x40\x27\xef\xa4\xe0\x0e\xda\xed\x3b\x2c\x5e\xc9\x0f\xcc\x02\xd6\x80\xe3\x23\xa9\xca\xc3\x6b\xca\x51\x37\x43\x3d\x2f\xba\xe0\xac\xcf\x35\x38\x3b\xaa\x66\x48\x16\x3c\x2d\x85\xee\x4c\xfa\xbe\x8b\x3a\xb7\x18\xda\x61\x1c\x6d\x81\x53\x75\xa0\xf2\x09\x99\xb4\x3d\x19\xc7\x7c\x31\x91\x47\xb5\x2e\x4f\xa1\x2d\x1b\x72\xfd\x0d\x00\x00\xff\xff\xfb\x9d\x40\x3e\xea\x09\x00\x00")

func templatesConcourse_lbTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesConcourse_lbTf,
		"templates/concourse_lb.tf",
	)
}

func templatesConcourse_lbTf() (*asset, error) {
	bytes, err := templatesConcourse_lbTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/concourse_lb.tf", size: 2538, mode: os.FileMode(420), modTime: time.Unix(1506547663, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesLb_subnetTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x93\x51\x6e\x9c\x30\x10\x86\xdf\x39\xc5\x68\x94\x87\xb4\xdd\xb8\xdb\xa8\xaf\x7b\x85\x5e\xa0\x8a\x2c\x63\x4f\xc9\xa8\x8e\xbd\xc2\x86\x74\x8b\xb8\x7b\x85\x4d\x0a\x16\x6c\xbb\x52\x40\x48\xc8\xf6\x7c\xf3\x8f\xe7\x9f\x96\x82\xef\x5a\x4d\x80\xea\x35\xc8\xd0\xd5\x8e\x22\x02\xda\x7a\xfe\x0f\x08\x43\x05\xa0\x7d\xe7\x22\xac\x9f\x13\xe0\xdd\x60\xc9\x35\xf1\xf9\xbe\x57\xad\x50\xbd\x62\xab\x6a\xb6\x1c\x2f\xf2\xb7\x77\x14\x3e\x8c\x58\x01\xf4\x67\x2d\xd9\x6c\x22\xa7\x6c\xfd\x59\x8b\xe9\x63\x93\x4e\x6a\x36\xad\xac\xad\xd7\x3f\x8b\x93\xd3\x72\xd6\x72\x8f\x5f\x8e\x22\xbd\x9f\x1f\x8f\x78\x80\xaf\x87\xac\x4b\xb0\x33\xf4\xeb\xd3\x63\x4e\xb8\x11\x92\x31\x64\xe9\x85\x5c\xbc\xa2\xb5\x20\x4d\x9c\x0a\x20\xaa\x26\xa4\xe2\x01\xbe\xa9\x97\x19\x33\x85\x93\xeb\x25\x9b\xf1\xc1\xd6\x0f\x59\xd8\xdd\xb0\x8a\x4e\x22\xc6\x09\x60\xf9\x07\xe9\x8b\xb6\x34\x53\xb8\x71\xbe\x25\xa9\x9f\x95\x6b\x28\xc0\x09\xbe\xe3\x52\x33\x1e\x00\x37\xba\xf0\x29\xb1\xc6\xaa\x2a\xfb\xd4\xfa\x2e\x92\x8c\xaa\xb6\x94\x9b\x55\x2c\x0c\xcb\xb5\xef\xdf\xf5\x3e\xef\x0a\xc9\x50\x88\xec\x54\x64\xef\xe4\xaa\x45\x27\xc0\xb7\x66\x1c\xa7\x8a\x1b\x15\xe9\x55\x5d\x8a\xa4\xec\x22\xb5\x8e\xa2\x9c\x37\x05\x37\x6f\xcd\x5e\xa5\x29\x42\x56\xeb\xa2\x54\xf3\x2f\xe9\x33\x48\x85\xe0\x35\x27\xa9\x08\x98\x77\xfe\xe3\xe4\x5b\x6d\x9c\x19\x7f\x9d\x5c\x58\x6a\x99\x1c\xb1\x64\x13\x1f\x05\x9b\x8d\xad\xde\x55\xb8\xef\xe2\xb9\x8b\xab\xe1\x94\x6c\xe6\xaa\x7a\x65\x3b\x4a\x8e\xca\xb4\x7d\x39\x23\x3e\xed\x73\xb6\x55\xdf\x8e\xdd\xc4\x5e\xcd\x92\x26\xf9\x76\xf0\x62\xb6\x4c\xfc\x13\x00\x00\xff\xff\x29\xc2\x4a\x36\xab\x04\x00\x00")

func templatesLb_subnetTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesLb_subnetTf,
		"templates/lb_subnet.tf",
	)
}

func templatesLb_subnetTf() (*asset, error) {
	bytes, err := templatesLb_subnetTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/lb_subnet.tf", size: 1195, mode: os.FileMode(420), modTime: time.Unix(1506547663, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesSsl_certificateTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x8f\x41\x6e\xc3\x20\x14\x44\xf7\x9c\x62\x84\xba\xee\x0d\x72\x16\x84\xf1\xb8\xf9\x2a\x31\xd1\x87\xd0\xa2\x88\xbb\x57\xc6\x1b\xa7\x92\x37\x61\x83\x04\xf3\x46\x6f\xaa\x57\xf1\x53\x24\x6c\xce\xd1\x05\x6a\x91\x45\x82\x2f\xb4\x78\x1a\xa0\xb4\x3b\x71\x81\xcd\x45\x65\xfd\xb2\xa6\x1b\x73\x4a\xb8\x70\xf5\xb2\xbe\xc1\xdd\x55\xea\x76\x7f\xb3\x9d\xd2\xca\x9c\x1e\x1a\x08\xeb\x7f\xb2\x13\x7f\x73\x99\x5a\xa9\xaf\xca\x36\x4e\xe3\x61\xaf\x59\xfd\x6d\x2b\xe7\x22\xbf\x5b\xdb\xc7\xb3\x7a\xfd\xcc\xd7\xa4\xc5\x71\xad\x4e\xe6\x6e\x8d\x01\x8e\x2a\x53\x9a\x1b\x0e\xe1\x57\xd3\x6e\xff\xc5\xc7\xe2\xd3\xf8\xfe\x3d\xa0\xc3\x44\xec\xe7\x14\x3a\x44\x77\xbf\x28\x0b\x43\x0b\x91\x63\x14\x10\x94\x43\x95\x4b\x52\xba\x99\xb9\x68\x6a\xb8\xa0\xe8\x83\x06\xe8\xa6\x9b\xbf\x00\x00\x00\xff\xff\x4f\x95\x65\x5c\xd6\x01\x00\x00")

func templatesSsl_certificateTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesSsl_certificateTf,
		"templates/ssl_certificate.tf",
	)
}

func templatesSsl_certificateTf() (*asset, error) {
	bytes, err := templatesSsl_certificateTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/ssl_certificate.tf", size: 470, mode: os.FileMode(420), modTime: time.Unix(1506547663, 0)}
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
	"templates/base.tf": templatesBaseTf,
	"templates/cf_dns.tf": templatesCf_dnsTf,
	"templates/cf_lb.tf": templatesCf_lbTf,
	"templates/concourse_lb.tf": templatesConcourse_lbTf,
	"templates/lb_subnet.tf": templatesLb_subnetTf,
	"templates/ssl_certificate.tf": templatesSsl_certificateTf,
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
		"base.tf": &bintree{templatesBaseTf, map[string]*bintree{}},
		"cf_dns.tf": &bintree{templatesCf_dnsTf, map[string]*bintree{}},
		"cf_lb.tf": &bintree{templatesCf_lbTf, map[string]*bintree{}},
		"concourse_lb.tf": &bintree{templatesConcourse_lbTf, map[string]*bintree{}},
		"lb_subnet.tf": &bintree{templatesLb_subnetTf, map[string]*bintree{}},
		"ssl_certificate.tf": &bintree{templatesSsl_certificateTf, map[string]*bintree{}},
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

