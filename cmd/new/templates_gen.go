// Code generated for package new by go-bindata DO NOT EDIT. (@generated)
// sources:
// template/go_mod.tmpl
// template/main.tmpl
// template/mirc_main.tmpl
// template/mirc_routers_site.tmpl
// template/mirc_routers_site_v1.tmpl
// template/mirc_routers_site_v2.tmpl
// template/readme_md.tmpl
// template/servants_site.tmpl
// template/servants_site_v1.tmpl
// template/servants_site_v2.tmpl
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

var _go_modTmpl = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00"

func go_modTmplBytes() ([]byte, error) {
	return bindataRead(
		_go_modTmpl,
		"go_mod.tmpl",
	)
}

func go_modTmpl() (*asset, error) {
	bytes, err := go_modTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "go_mod.tmpl", size: 0, mode: os.FileMode(420), modTime: time.Unix(1581296466, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _mainTmpl = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00"

func mainTmplBytes() ([]byte, error) {
	return bindataRead(
		_mainTmpl,
		"main.tmpl",
	)
}

func mainTmpl() (*asset, error) {
	bytes, err := mainTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "main.tmpl", size: 0, mode: os.FileMode(420), modTime: time.Unix(1581296400, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _mirc_mainTmpl = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00"

func mirc_mainTmplBytes() ([]byte, error) {
	return bindataRead(
		_mirc_mainTmpl,
		"mirc_main.tmpl",
	)
}

func mirc_mainTmpl() (*asset, error) {
	bytes, err := mirc_mainTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "mirc_main.tmpl", size: 0, mode: os.FileMode(420), modTime: time.Unix(1581296104, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _mirc_routers_siteTmpl = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00"

func mirc_routers_siteTmplBytes() ([]byte, error) {
	return bindataRead(
		_mirc_routers_siteTmpl,
		"mirc_routers_site.tmpl",
	)
}

func mirc_routers_siteTmpl() (*asset, error) {
	bytes, err := mirc_routers_siteTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "mirc_routers_site.tmpl", size: 0, mode: os.FileMode(420), modTime: time.Unix(1581296070, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _mirc_routers_site_v1Tmpl = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00"

func mirc_routers_site_v1TmplBytes() ([]byte, error) {
	return bindataRead(
		_mirc_routers_site_v1Tmpl,
		"mirc_routers_site_v1.tmpl",
	)
}

func mirc_routers_site_v1Tmpl() (*asset, error) {
	bytes, err := mirc_routers_site_v1TmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "mirc_routers_site_v1.tmpl", size: 0, mode: os.FileMode(420), modTime: time.Unix(1581296038, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _mirc_routers_site_v2Tmpl = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00"

func mirc_routers_site_v2TmplBytes() ([]byte, error) {
	return bindataRead(
		_mirc_routers_site_v2Tmpl,
		"mirc_routers_site_v2.tmpl",
	)
}

func mirc_routers_site_v2Tmpl() (*asset, error) {
	bytes, err := mirc_routers_site_v2TmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "mirc_routers_site_v2.tmpl", size: 0, mode: os.FileMode(420), modTime: time.Unix(1581296055, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _readme_mdTmpl = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00"

func readme_mdTmplBytes() ([]byte, error) {
	return bindataRead(
		_readme_mdTmpl,
		"readme_md.tmpl",
	)
}

func readme_mdTmpl() (*asset, error) {
	bytes, err := readme_mdTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "readme_md.tmpl", size: 0, mode: os.FileMode(420), modTime: time.Unix(1581296441, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _servants_siteTmpl = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00"

func servants_siteTmplBytes() ([]byte, error) {
	return bindataRead(
		_servants_siteTmpl,
		"servants_site.tmpl",
	)
}

func servants_siteTmpl() (*asset, error) {
	bytes, err := servants_siteTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "servants_site.tmpl", size: 0, mode: os.FileMode(420), modTime: time.Unix(1581296142, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _servants_site_v1Tmpl = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00"

func servants_site_v1TmplBytes() ([]byte, error) {
	return bindataRead(
		_servants_site_v1Tmpl,
		"servants_site_v1.tmpl",
	)
}

func servants_site_v1Tmpl() (*asset, error) {
	bytes, err := servants_site_v1TmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "servants_site_v1.tmpl", size: 0, mode: os.FileMode(420), modTime: time.Unix(1581296142, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _servants_site_v2Tmpl = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00"

func servants_site_v2TmplBytes() ([]byte, error) {
	return bindataRead(
		_servants_site_v2Tmpl,
		"servants_site_v2.tmpl",
	)
}

func servants_site_v2Tmpl() (*asset, error) {
	bytes, err := servants_site_v2TmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "servants_site_v2.tmpl", size: 0, mode: os.FileMode(420), modTime: time.Unix(1581296142, 0)}
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
	"go_mod.tmpl":               go_modTmpl,
	"main.tmpl":                 mainTmpl,
	"mirc_main.tmpl":            mirc_mainTmpl,
	"mirc_routers_site.tmpl":    mirc_routers_siteTmpl,
	"mirc_routers_site_v1.tmpl": mirc_routers_site_v1Tmpl,
	"mirc_routers_site_v2.tmpl": mirc_routers_site_v2Tmpl,
	"readme_md.tmpl":            readme_mdTmpl,
	"servants_site.tmpl":        servants_siteTmpl,
	"servants_site_v1.tmpl":     servants_site_v1Tmpl,
	"servants_site_v2.tmpl":     servants_site_v2Tmpl,
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
	"go_mod.tmpl":               &bintree{go_modTmpl, map[string]*bintree{}},
	"main.tmpl":                 &bintree{mainTmpl, map[string]*bintree{}},
	"mirc_main.tmpl":            &bintree{mirc_mainTmpl, map[string]*bintree{}},
	"mirc_routers_site.tmpl":    &bintree{mirc_routers_siteTmpl, map[string]*bintree{}},
	"mirc_routers_site_v1.tmpl": &bintree{mirc_routers_site_v1Tmpl, map[string]*bintree{}},
	"mirc_routers_site_v2.tmpl": &bintree{mirc_routers_site_v2Tmpl, map[string]*bintree{}},
	"readme_md.tmpl":            &bintree{readme_mdTmpl, map[string]*bintree{}},
	"servants_site.tmpl":        &bintree{servants_siteTmpl, map[string]*bintree{}},
	"servants_site_v1.tmpl":     &bintree{servants_site_v1Tmpl, map[string]*bintree{}},
	"servants_site_v2.tmpl":     &bintree{servants_site_v2Tmpl, map[string]*bintree{}},
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
