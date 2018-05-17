// Code generated by go-bindata.
// sources:
// .DS_Store
// bindata.go
// generate.go
// migrations.go
// DO NOT EDIT!

package migrations

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

var _Ds_store = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\xd8\x31\x0a\x02\x31\x10\x85\xe1\x37\x31\x45\xc0\x26\xa5\x65\x1a\x0f\xe0\x0d\xc2\xb2\x9e\xc0\x0b\x58\x78\x05\xfb\x1c\x5d\x96\x79\x60\x60\xd5\x4e\x8c\xcb\xfb\x40\xfe\x05\x37\x2a\x16\x31\x23\x00\x9b\xee\xb7\x13\x90\x01\x24\x78\x71\xc4\x4b\x89\x8f\x95\xd0\x5d\x1b\x5f\x43\x44\x44\x44\xc6\x66\x9e\xb4\xff\xf5\x07\x11\x91\xe1\x2c\xfb\x43\x61\x2b\xdb\xbc\xc6\xe7\x03\x1b\xbb\x35\x99\x2d\x6c\x65\x9b\xd7\x78\x5f\x60\x23\x9b\xd8\xcc\x16\xb6\xb2\xcd\xcb\x4d\xcb\x38\x7c\x18\xdf\xd9\x38\xa1\x18\xa7\x10\x2b\x6c\xfd\xce\x77\x23\xf2\xef\x76\x9e\xbc\xfc\xfe\x9f\xdf\xcf\xff\x22\xb2\x61\x16\xe7\xcb\x3c\x3d\x07\x82\xf5\x0d\x00\xae\xdd\xf5\xa7\x43\x40\xf0\x3f\x0b\x0f\xdd\x5a\x1d\x04\x44\x06\xf3\x08\x00\x00\xff\xff\x6a\x00\x88\x6d\x04\x18\x00\x00")

func Ds_storeBytes() ([]byte, error) {
	return bindataRead(
		_Ds_store,
		".DS_Store",
	)
}

func Ds_store() (*asset, error) {
	bytes, err := Ds_storeBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: ".DS_Store", size: 6148, mode: os.FileMode(420), modTime: time.Unix(1526506262, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _bindataGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func bindataGoBytes() ([]byte, error) {
	return bindataRead(
		_bindataGo,
		"bindata.go",
	)
}

func bindataGo() (*asset, error) {
	bytes, err := bindataGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bindata.go", size: 0, mode: os.FileMode(420), modTime: time.Unix(1526585397, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _generateGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd2\xd7\x4f\xcf\xb7\x4a\x4f\xcd\x4b\x2d\x4a\x2c\x49\x55\x48\xcf\xd7\x4d\xca\xcc\x4b\x49\x2c\x49\x54\xd0\x2d\xc8\x4e\x57\xc8\xcd\x4c\x2f\x4a\x2c\xc9\xcc\xcf\x2b\x56\xd0\xe3\xe2\x2a\x48\x4c\xce\x4e\x4c\x4f\x45\x12\xe5\x02\x04\x00\x00\xff\xff\xd3\xb0\x8e\xc1\x3f\x00\x00\x00")

func generateGoBytes() ([]byte, error) {
	return bindataRead(
		_generateGo,
		"generate.go",
	)
}

func generateGo() (*asset, error) {
	bytes, err := generateGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "generate.go", size: 63, mode: os.FileMode(420), modTime: time.Unix(1526506084, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrationsGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x94\x5d\x6f\xdb\x36\x14\x86\xaf\xc9\x5f\x71\xaa\x0b\x8f\x2a\x1c\xa9\x03\x76\x95\xc2\x03\x82\x66\x01\x7a\xd1\x14\x68\x57\xec\x22\xcb\x5c\x9a\x3a\xa6\x08\x4b\x87\x1a\x49\xd9\x09\xda\xfc\xf7\x81\xd4\x47\x9d\xcc\x45\x83\x61\x17\x86\x40\xf1\x7c\x3c\xef\xf1\x7b\xd4\x49\xb5\x93\x1a\xa1\x35\xda\xc9\x60\x2c\x79\xce\x4d\xdb\x59\x17\x40\x70\x96\x39\xd4\x78\xd7\x65\x9c\x65\xde\xba\x90\x9e\xc1\x29\x4b\xfb\x8c\x73\x96\x69\x13\xea\x7e\x53\x28\xdb\x96\xda\x36\x92\xf4\xd9\x50\x06\xcb\xf1\x99\x71\xb6\x86\x1f\x87\x95\x95\x0c\x72\x23\x3d\x96\xca\xaa\x9d\xb3\x52\xd5\xd5\x26\xe3\x6c\x63\x28\xde\x3c\xa7\x82\xb7\xbd\x53\x58\x6a\x7b\x36\x26\x65\xa7\xf8\x4a\xdd\x58\xfd\xe4\xa6\x95\x77\xad\xaa\x4a\xf5\x4b\xb9\x91\x6a\x87\x54\x95\xdd\x4e\x97\xca\xd2\xd6\xe8\x8c\xe7\x9c\x6f\x7b\x52\xa0\x31\xbc\x9b\x47\x24\x72\x10\x2f\xc7\xce\xc5\xf0\x1a\x97\x70\x73\x6b\x28\xe4\xf0\x85\xb3\xb2\x84\xdf\xdf\x5f\xbe\x3f\x87\x4b\xdc\xf4\x1a\x0e\xf5\x3d\x84\xfa\x78\xc6\x20\x1d\xd2\x4f\x01\x5c\x4f\x64\x48\x03\xee\x91\x20\xd4\xb6\xd7\x75\x8c\xbc\x07\x5f\xdb\x03\xf4\x1d\x18\x82\x51\x0f\x67\x1e\xce\x57\xd3\xa9\xf8\x80\x83\x62\x71\xe1\x3d\x86\x6b\xd9\xa2\x17\xf9\x92\x33\x16\x69\x05\xc9\x16\xc1\x07\x67\x48\xe7\x20\x6e\x6e\x37\xf7\x11\x10\x9d\xb3\x2e\x01\x32\xe6\x30\xf4\x8e\x20\x65\xa7\xf0\x9c\x33\xf6\x90\x73\xb6\x97\xee\x1b\xe8\x7a\x8f\xce\x1b\x4b\xb0\x82\xc1\x09\xc5\xbb\xde\x87\x37\xb6\xed\x4c\x83\xe2\xf3\x5f\x7f\x56\x2f\x3f\x8f\x39\x63\xa4\x1f\xc6\xc0\x59\xf4\x4b\xf1\x31\x21\x78\xe1\x8b\x44\x98\x73\xb6\xb5\x0e\xd6\x4b\x48\x80\xe7\x2b\x70\x92\x34\xc2\x78\x9d\xc8\xcc\x76\x04\x5f\xb7\xfb\x18\xf1\x2f\x96\xe2\xca\x50\x35\x14\x1e\xc0\x5f\x1f\x25\xbc\x58\x41\x96\x0d\x0a\xdb\x7d\x52\x1c\x6b\x8c\x9e\x2d\x2e\x82\x35\x62\x0e\x8e\x8a\x63\xbb\x18\xb4\x5a\x01\x99\x06\x16\x8b\xf8\x8e\x89\x06\x49\x4c\x82\xf2\x78\xf9\x2a\x87\xaf\x5f\x4f\xdd\xfd\x0a\xaf\x60\xb1\x98\xd5\xdf\x3c\xba\x3d\xfb\xf9\x36\x12\xb5\xfb\x71\xe8\x6c\x1e\xd2\x0a\x64\xd7\x21\x55\x73\xec\x12\x46\xa0\x07\x9e\x7e\x0f\xe3\x04\xdf\x52\xf0\xdf\x0a\x72\x56\xcd\xa2\x26\x27\xfc\x61\x42\xfd\x96\x7c\x90\xa4\x50\xc4\x90\x51\xd2\x8b\x41\x52\x6c\x1c\x5d\x5f\x5c\xc9\x20\x1b\x81\xce\xe5\xa9\x7a\x3b\x17\x9a\x8c\x7c\x8d\x87\x58\xeb\x63\xf2\xd5\x5c\x91\x33\x96\x1d\x6d\xd5\x12\xaa\x25\x0c\xcb\x51\x5c\x8e\x5b\xfb\xc9\x35\x22\x9b\x37\x37\xcb\x9f\xcd\x30\x9a\xb0\x5d\xce\xf3\xe3\x0f\xe3\xbe\x7d\xe8\xe9\xd1\xbe\x7d\x49\xc4\xeb\xc8\xfb\x64\x15\xe7\x66\x51\x4a\xf1\xa9\x13\xf9\xeb\xe3\xde\x8b\xc5\x74\x9a\x74\xfe\xe6\xdc\xb5\x7d\x53\x27\xe7\x7d\x87\x6c\xa2\xb8\x74\xb6\xbb\x68\x9a\xff\x44\x12\x73\x9f\xb0\xfc\xa8\x9b\x3d\xd0\xe9\x6e\xb3\x6f\x4e\x35\x2d\xcb\xb9\x5b\x3a\x6c\xad\x53\x98\x3e\x38\x84\x07\xf4\x61\xca\x06\xe9\xe1\x80\xe0\xf0\xef\xde\x38\x84\xca\x1e\xe8\xf8\x93\x14\x6c\xca\x56\x35\xaa\x5d\xac\x01\x78\x67\x7c\x40\x52\xf8\x58\xd7\x55\x2c\x2f\xbe\x6b\xf8\x67\x29\x7e\x3c\x28\x7b\xa0\xff\xe7\x4f\xfb\x27\x00\x00\xff\xff\xec\xb5\x29\x8b\xcd\x06\x00\x00")

func migrationsGoBytes() ([]byte, error) {
	return bindataRead(
		_migrationsGo,
		"migrations.go",
	)
}

func migrationsGo() (*asset, error) {
	bytes, err := migrationsGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations.go", size: 1741, mode: os.FileMode(420), modTime: time.Unix(1526509267, 0)}
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
	".DS_Store": Ds_store,
	"bindata.go": bindataGo,
	"generate.go": generateGo,
	"migrations.go": migrationsGo,
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
	".DS_Store": &bintree{Ds_store, map[string]*bintree{}},
	"bindata.go": &bintree{bindataGo, map[string]*bintree{}},
	"generate.go": &bintree{generateGo, map[string]*bintree{}},
	"migrations.go": &bintree{migrationsGo, map[string]*bintree{}},
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

