// Code generated for package new by go-bindata DO NOT EDIT. (@generated)
// sources:
// templates/gin_go_mod.tmpl
// templates/gin_main.tmpl
// templates/gin_mirc_main.tmpl
// templates/gin_mirc_routes_site.tmpl
// templates/gin_mirc_routes_site_v1.tmpl
// templates/gin_mirc_routes_site_v2.tmpl
// templates/makefile.tmpl
// templates/readme.tmpl
package new

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

func bindataRead(data, name string) ([]byte, error) {
	gz, err := gzip.NewReader(strings.NewReader(data))
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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _gin_go_modTmpl = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xca\xcd\x4f\x29\xcd\x49\x55\x48\xcf\x2c\xc9\x28\x4d\xd2\x4b\xce\xcf\xd5\x4f\xcc\xc9\xcc\xad\xd4\xcf\xcd\x2c\xd2\x4d\xad\x48\xcc\x2d\xc8\x49\x2d\xe6\xe2\x4a\xcf\x57\x30\xd4\x33\x34\xe4\xe2\x2a\x4a\x2d\x2c\xcd\x2c\x42\x51\x9f\x9e\x99\xa7\x9b\x9e\x9f\x97\x99\x0c\x62\x29\x94\x19\xea\x99\xe8\x19\x00\x02\x00\x00\xff\xff\x66\xe7\x3e\xf4\x56\x00\x00\x00"

func gin_go_modTmplBytes() ([]byte, error) {
	return bindataRead(
		_gin_go_modTmpl,
		"gin_go_mod.tmpl",
	)
}

func gin_go_modTmpl() (*asset, error) {
	bytes, err := gin_go_modTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "gin_go_mod.tmpl", size: 86, mode: os.FileMode(436), modTime: time.Unix(1582267611, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _gin_mainTmpl = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x50\xbd\x0a\xdb\x30\x10\xde\xf5\x14\x87\x86\x60\x43\x62\x91\x8c\x86\x4e\xa5\xa4\x43\x69\x02\x1d\x4b\x07\x55\xb9\xc8\x47\xad\x93\x39\xcb\xa1\xc5\xf8\xdd\x8b\x6c\xa7\x3f\x50\x52\x12\x4d\xf7\xfd\xa1\xbb\xaf\xb3\xee\x9b\xf5\x08\xc1\x12\x2b\x45\xa1\x8b\x92\xa0\x50\x00\x00\xba\x8d\x5e\xab\x65\xf4\x94\x9a\xe1\x6b\xe5\x62\x30\x9e\x78\xe7\x23\x93\xcb\xd3\x3f\x74\xdb\x52\xf8\x61\x02\x89\xb9\x1d\x8c\x8b\x82\xfa\xb1\x05\xd9\x13\x3f\x32\xed\xf0\xbb\x0d\x5d\x8b\x7d\x06\xce\x48\x1c\x12\xf6\xcf\xfa\xcd\x6d\xff\x7c\xe4\xa0\x55\xa9\x94\x31\x3e\xd6\x1e\x19\xc5\x26\x04\x1f\x41\x06\x9e\xfb\xaa\x7c\x54\xd7\x81\xdd\x0c\x8a\x12\xc6\xf9\x83\x36\xfa\xea\x2c\xc4\xa9\xe5\x42\xff\x8a\xb9\x78\x41\xe8\x93\x95\xa4\xcb\xd9\x86\x9c\x84\xb0\x87\xfa\x0d\x04\x92\x77\x0b\x2a\xee\x5a\xae\xa4\x3a\xae\xe1\x62\xf5\x6e\x61\x93\xfb\xcc\xfc\xa9\x4b\xfd\xf8\xd1\x06\xac\xe1\x4e\x65\x6b\x94\x23\xf1\x16\x4e\x43\x3a\xdb\xd4\xd4\xa0\x2b\xe3\x91\xf5\x54\xfe\x6f\xb3\x2b\x31\xf5\x8d\x2e\xd5\xa4\xd6\x93\xfe\xd8\x09\x3e\x7f\x21\x4e\x28\x57\xeb\x70\x9c\xd6\x33\x05\xd3\x20\xfc\xb7\xb4\x28\xf9\x6d\x96\x0a\xab\x4f\x94\x70\x7c\xdb\x58\xe2\x1a\x3c\x71\xf5\xde\xf2\xa5\x45\xe9\x67\x6a\xcc\xcc\x87\xe8\x3d\x4a\x51\x4e\xd3\xf6\x77\xfa\xb6\x7f\x39\x79\x78\x21\x39\xa9\xe9\x67\x00\x00\x00\xff\xff\x7e\x8b\x5b\x35\x08\x03\x00\x00"

func gin_mainTmplBytes() ([]byte, error) {
	return bindataRead(
		_gin_mainTmpl,
		"gin_main.tmpl",
	)
}

func gin_mainTmpl() (*asset, error) {
	bytes, err := gin_mainTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "gin_main.tmpl", size: 776, mode: os.FileMode(436), modTime: time.Unix(1582267612, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _gin_mirc_mainTmpl = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x50\xbd\x0a\xdb\x30\x10\xde\xf5\x14\x87\x86\x60\x43\x62\x91\x8c\x86\x4e\xa5\xa4\x43\x69\x02\x1d\x4b\x07\x55\xb9\xc8\x47\xad\x93\x39\xcb\xa1\xc5\xf8\xdd\x8b\x6c\xa7\x3f\x50\x52\x12\x4d\xf7\xfd\xa1\xbb\xaf\xb3\xee\x9b\xf5\x08\xc1\x12\x2b\x45\xa1\x8b\x92\xa0\x50\x00\x00\xba\x8d\x5e\xab\x65\xf4\x94\x9a\xe1\x6b\xe5\x62\x30\x9e\x78\xe7\x23\x93\xcb\xd3\x3f\x74\xdb\x52\xf8\x61\x02\x89\xb9\x1d\x8c\x8b\x82\xfa\xb1\x05\xd9\x13\x3f\x32\xed\xf0\xbb\x0d\x5d\x8b\x7d\x06\xce\x48\x1c\x12\xf6\xcf\xfa\xcd\x6d\xff\x7c\xe4\xa0\x55\xa9\x94\x31\x3e\xd6\x1e\x19\xc5\x26\x04\x1f\x41\x06\x9e\xfb\xaa\x7c\x54\xd7\x81\xdd\x0c\x8a\x12\xc6\xf9\x83\x36\xfa\xea\x2c\xc4\xa9\xe5\x42\xff\x8a\xb9\x78\x41\xe8\x93\x95\xa4\xcb\xd9\x86\x9c\x84\xb0\x87\xfa\x0d\x04\x92\x77\x0b\x2a\xee\x5a\xae\xa4\x3a\xae\xe1\x62\xf5\x6e\x61\x93\xfb\xcc\xfc\xa9\x4b\xfd\xf8\xd1\x06\xac\xe1\x4e\x65\x6b\x94\x23\xf1\x16\x4e\x43\x3a\xdb\xd4\xd4\xa0\x2b\xe3\x91\xf5\x54\xfe\x6f\xb3\x2b\x31\xf5\x8d\x2e\xd5\xa4\xd6\x93\xfe\xd8\x09\x3e\x7f\x21\x4e\x28\x57\xeb\x70\x9c\xd6\x33\x05\xd3\x20\xfc\xb7\xb4\x28\xf9\x6d\x96\x0a\xab\x4f\x94\x70\x7c\xdb\x58\xe2\x1a\x3c\x71\xf5\xde\xf2\xa5\x45\xe9\x67\x6a\xcc\xcc\x87\xe8\x3d\x4a\x51\x4e\xd3\xf6\x77\xfa\xb6\x7f\x39\x79\x78\x21\x39\xa9\xe9\x67\x00\x00\x00\xff\xff\x7e\x8b\x5b\x35\x08\x03\x00\x00"

func gin_mirc_mainTmplBytes() ([]byte, error) {
	return bindataRead(
		_gin_mirc_mainTmpl,
		"gin_mirc_main.tmpl",
	)
}

func gin_mirc_mainTmpl() (*asset, error) {
	bytes, err := gin_mirc_mainTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "gin_mirc_main.tmpl", size: 776, mode: os.FileMode(436), modTime: time.Unix(1582267612, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _gin_mirc_routes_siteTmpl = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\xcd\xb1\x8e\xc2\x30\x0c\x80\xe1\x3d\x4f\x61\x65\xb9\xbb\xe1\x6a\x89\xb1\x1b\x62\x40\xcc\xbc\x40\x43\x6a\x52\x8b\xa6\xa9\x1c\x17\x51\x21\xde\x1d\xb5\x0d\x13\xa3\xf5\x7f\xb6\x47\xe7\x6f\x2e\x10\x48\x9a\x94\xb2\x31\x1c\xc7\x24\x0a\xbf\x06\x00\x20\xb2\x80\x0d\xac\xdd\x74\xa9\x7c\x8a\xe8\x7a\x8e\x33\x46\x16\xbc\xef\xac\xf9\x33\x06\x11\xce\xac\xb4\xc0\x9f\x0c\x59\x65\xf2\x0a\xea\x02\xb4\x74\xe5\x81\x8c\xce\x23\x6d\xa2\xb4\xe7\x7a\xf7\xd0\x39\x1e\xca\x83\x6a\x1b\x9a\xc8\x52\xdb\x7f\xdb\xac\xe0\x34\xb4\xf4\xf8\x80\x23\x29\x40\x01\xc8\x4b\xc1\xc2\xf6\xa2\xec\x7b\xca\xdf\xcc\x95\x82\xb5\x77\x4a\x21\xc9\xbc\xec\xbc\xde\x01\x00\x00\xff\xff\x70\x65\x95\x92\xee\x00\x00\x00"

func gin_mirc_routes_siteTmplBytes() ([]byte, error) {
	return bindataRead(
		_gin_mirc_routes_siteTmpl,
		"gin_mirc_routes_site.tmpl",
	)
}

func gin_mirc_routes_siteTmpl() (*asset, error) {
	bytes, err := gin_mirc_routes_siteTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "gin_mirc_routes_site.tmpl", size: 238, mode: os.FileMode(436), modTime: time.Unix(1582267612, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _gin_mirc_routes_site_v1Tmpl = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\xcd\xb1\x4e\xc4\x30\x0c\xc6\xf1\x3d\x4f\x61\x65\x01\x06\x6a\x95\xb1\x1b\x62\x40\xcc\xbc\x40\x43\x6a\x52\x8b\xa6\x89\x5c\xb7\xa2\x42\xbc\xfb\xa9\x4d\xee\x96\x1b\x3f\xfd\x7f\x89\xb3\xf3\x3f\x2e\x10\x6c\xad\x31\x1c\x73\x12\x85\x47\x03\x00\x10\x59\xc0\x06\xd6\x71\xfd\x6a\x7c\x8a\xe8\x26\x8e\x3b\x46\x16\xdc\x5e\xac\x79\x32\x06\x11\x3e\x59\xe9\x80\x0f\x0b\x2c\x2a\xab\x57\x50\x17\x60\xa0\x6f\x9e\xc9\xe8\x9e\xa9\x88\xda\xfe\xce\x7f\xdf\x46\xc7\x73\x3d\xd0\x94\xd1\x47\x96\xce\x3e\xdb\xfe\x04\xef\x92\xd6\x7c\x05\x65\x14\xb0\xb5\x55\x7c\xcc\x03\xfd\xde\x04\x29\x40\x15\xc8\x47\xc1\xca\x5e\x45\xd9\x4f\xb4\xdc\x33\x57\x0b\x76\xde\x29\x85\x24\xfb\xf1\xe6\xff\x12\x00\x00\xff\xff\xe6\xe7\xea\xf7\x0c\x01\x00\x00"

func gin_mirc_routes_site_v1TmplBytes() ([]byte, error) {
	return bindataRead(
		_gin_mirc_routes_site_v1Tmpl,
		"gin_mirc_routes_site_v1.tmpl",
	)
}

func gin_mirc_routes_site_v1Tmpl() (*asset, error) {
	bytes, err := gin_mirc_routes_site_v1TmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "gin_mirc_routes_site_v1.tmpl", size: 268, mode: os.FileMode(436), modTime: time.Unix(1582267611, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _gin_mirc_routes_site_v2Tmpl = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\xcd\xb1\x4e\xc4\x30\x0c\x06\xe0\xdd\x4f\x61\x65\x01\x06\xce\x52\xc7\xdb\xd0\x0d\x88\x99\x17\xb8\x90\x33\x39\x8b\x4b\x13\xb9\x6e\x45\x85\x78\x77\xd4\x26\x45\x27\x75\xfc\xf5\x7f\xbf\x5d\x7c\xf8\xf2\x91\x71\xea\x00\x24\x95\xac\x86\x8f\x80\x88\x98\x44\xd1\x45\xb1\xeb\xf8\x71\x08\x39\x91\xbf\x49\x9a\x29\x89\xd2\xd4\x39\x78\x02\x20\xc2\x77\x31\x5e\xe0\xc3\x80\x83\xe9\x18\x0c\xcd\x47\xbc\xf0\xa7\xf4\x0c\x36\x17\xae\xa2\x75\x3f\xeb\xdd\xd3\xd5\x4b\xdf\x1e\x1c\x6a\x38\x27\xd1\xa3\x7b\x76\xe7\x15\xbc\x6a\x1e\xcb\x06\x6a\xa8\x60\xea\x9a\x78\xeb\x2f\xfc\xfd\x2f\xd8\x10\x9b\x20\x59\x1a\x6a\xec\x45\x4d\xc2\x8d\x87\x3d\xf3\xad\xa1\x63\xf0\xc6\x31\xeb\xbc\x6d\x4e\x2d\xef\x37\xf7\xf2\x17\xfe\x02\x00\x00\xff\xff\x60\x81\x5c\x84\x37\x01\x00\x00"

func gin_mirc_routes_site_v2TmplBytes() ([]byte, error) {
	return bindataRead(
		_gin_mirc_routes_site_v2Tmpl,
		"gin_mirc_routes_site_v2.tmpl",
	)
}

func gin_mirc_routes_site_v2Tmpl() (*asset, error) {
	bytes, err := gin_mirc_routes_site_v2TmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "gin_mirc_routes_site_v2.tmpl", size: 311, mode: os.FileMode(436), modTime: time.Unix(1582267611, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _makefileTmpl = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\x8f\xcd\x8a\x83\x30\x14\x46\xd7\xde\xa7\xf8\x10\x17\x3a\x90\xcc\x3e\x20\xb3\x1a\x9d\x81\xb6\x16\xda\x4d\x97\xb6\x26\x69\xc0\x24\x25\xa6\x94\xbe\x7d\x31\x22\x6e\x2e\x87\x73\xff\xdb\xae\xd9\x9f\xf1\x53\x43\x7b\x65\x23\xd8\x04\xf6\xa2\xb6\x6b\xfe\x77\xbf\x27\x88\x1a\x45\x39\xdd\xe5\x38\x42\x19\x37\x80\x83\xb9\xde\x4a\xe4\x5f\x5c\xfb\x1c\x2c\xbe\x1f\x12\xaa\x22\xe2\xc7\xbf\xee\x70\x11\xb8\x3e\xcd\x38\x50\x8a\x02\xca\x46\xca\xb4\x5f\x24\xf8\x56\xa5\xa5\x93\xa1\x8f\x92\x56\x10\x94\xb1\x60\xc1\x82\x82\x35\xe1\xf6\xad\xa5\x4b\x9d\x6b\x7e\xb1\xb6\x37\x8e\x6b\xbf\xcd\x99\x17\x28\x1b\x05\x65\x45\x99\xfe\xa8\x90\x60\xbe\xbd\xa2\x4f\x00\x00\x00\xff\xff\xbf\x2a\x49\x7b\xda\x00\x00\x00"

func makefileTmplBytes() ([]byte, error) {
	return bindataRead(
		_makefileTmpl,
		"makefile.tmpl",
	)
}

func makefileTmpl() (*asset, error) {
	bytes, err := makefileTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "makefile.tmpl", size: 218, mode: os.FileMode(420), modTime: time.Unix(1582263685, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _readmeTmpl = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x34\xca\x41\x8e\xc2\x30\x0c\x46\xe1\x7d\x4e\xf1\x4b\xd6\x68\x56\x70\x1b\xf6\x71\xc1\xb4\x29\x49\x1c\xd9\x89\x28\xb7\x47\x41\xea\xf2\xe9\x7d\x44\x28\xc9\x2e\x72\x70\x69\x59\x3c\xec\xc3\x3b\x18\x9e\x66\xce\xf5\xef\x38\x27\x9e\x6a\x90\xa3\x65\x4e\x15\x9b\xbe\xd1\x15\xc3\x7f\x0a\xa9\xe2\xa3\xc3\xd0\x4c\x77\xb9\xf7\x6b\x08\x44\x44\xb8\x39\xaf\x12\x62\x8c\x0b\xfb\x16\xfe\x50\xf8\x25\x58\xa5\x8a\x71\x97\xb3\x97\x91\xf2\x63\xa2\x6f\x00\x00\x00\xff\xff\xe6\xd9\x54\xfa\x8d\x00\x00\x00"

func readmeTmplBytes() ([]byte, error) {
	return bindataRead(
		_readmeTmpl,
		"readme.tmpl",
	)
}

func readmeTmpl() (*asset, error) {
	bytes, err := readmeTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "readme.tmpl", size: 141, mode: os.FileMode(436), modTime: time.Unix(1582266099, 0)}
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
	"gin_go_mod.tmpl":              gin_go_modTmpl,
	"gin_main.tmpl":                gin_mainTmpl,
	"gin_mirc_main.tmpl":           gin_mirc_mainTmpl,
	"gin_mirc_routes_site.tmpl":    gin_mirc_routes_siteTmpl,
	"gin_mirc_routes_site_v1.tmpl": gin_mirc_routes_site_v1Tmpl,
	"gin_mirc_routes_site_v2.tmpl": gin_mirc_routes_site_v2Tmpl,
	"makefile.tmpl":                makefileTmpl,
	"readme.tmpl":                  readmeTmpl,
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
	"gin_go_mod.tmpl":              &bintree{gin_go_modTmpl, map[string]*bintree{}},
	"gin_main.tmpl":                &bintree{gin_mainTmpl, map[string]*bintree{}},
	"gin_mirc_main.tmpl":           &bintree{gin_mirc_mainTmpl, map[string]*bintree{}},
	"gin_mirc_routes_site.tmpl":    &bintree{gin_mirc_routes_siteTmpl, map[string]*bintree{}},
	"gin_mirc_routes_site_v1.tmpl": &bintree{gin_mirc_routes_site_v1Tmpl, map[string]*bintree{}},
	"gin_mirc_routes_site_v2.tmpl": &bintree{gin_mirc_routes_site_v2Tmpl, map[string]*bintree{}},
	"makefile.tmpl":                &bintree{makefileTmpl, map[string]*bintree{}},
	"readme.tmpl":                  &bintree{readmeTmpl, map[string]*bintree{}},
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
