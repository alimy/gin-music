// Code generated by go-bindata.
// sources:
// assets/README.md
// assets/albums.json
// assets/config/README.md
// assets/config/app.toml
// DO NOT EDIT!

// +build release

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

var _assetsReadmeMd = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x52\x56\x56\x48\xce\xcf\x2b\x49\xcc\xcc\x2b\x56\x48\x2c\x2e\x4e\x2d\x29\x56\xc8\x4f\x53\x48\x2c\x28\xc8\xc9\x4c\x4e\x2c\xc9\xcc\xcf\x03\x04\x00\x00\xff\xff\x01\xe6\x8a\x6c\x21\x00\x00\x00"

func assetsReadmeMdBytes() ([]byte, error) {
	return bindataRead(
		_assetsReadmeMd,
		"assets/README.md",
	)
}

func assetsReadmeMd() (*asset, error) {
	bytes, err := assetsReadmeMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/README.md", size: 33, mode: os.FileMode(436), modTime: time.Unix(1545785454, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsAlbumsJson = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x98\x4f\x6f\xdb\x46\x10\xc5\xef\xf9\x14\x03\x5d\x7c\x09\x84\x26\x6d\xe4\xb8\x97\x82\x74\xec\xa4\xae\x95\xb8\xa4\x13\x23\x2d\x8a\x62\xc8\x1d\x93\x03\x2d\x77\x84\xfd\xa3\x8a\x28\xfa\xdd\x0b\xb2\x05\x22\x04\x4b\x32\x17\x8b\xd2\x41\x07\x2d\x49\xf0\x87\x37\x6f\xde\x8c\x7e\x7f\x06\x00\xf0\x77\xff\xdd\x7d\x16\x7f\x96\x1a\x9d\x5b\xfc\x08\x0b\xb1\xd5\xb2\xd4\x12\xd4\xa3\x04\xa3\x6c\xbb\x74\xd8\x6c\x35\xb9\x65\x13\x1c\x97\x4b\x25\x0d\xb2\x59\x26\xba\x08\xcd\xe2\xf9\x97\x07\xa0\xf5\xec\x7c\xf7\x80\xf7\x6c\x77\x68\xf0\xf0\xd0\xb3\xd7\xd4\x9f\xd1\x8e\x6c\xc3\x46\x1d\x9e\x5a\xd2\x84\x8e\x3e\x13\xda\xee\x9a\x17\x17\x17\x2f\x0e\x8f\x2b\x32\xb6\xbf\x39\x93\x72\xb3\xe8\x7f\xff\xe7\xf9\x13\xbe\xff\x7d\x4d\x90\x12\x96\x35\xa4\xd2\xba\x28\xc6\x1d\x79\xc8\xbb\xc7\xbb\x51\x8e\xd5\x6a\x4e\x8e\x35\xda\x1d\x1b\x78\x8b\x2d\x45\x21\x1e\x6a\xf4\x67\x0e\xde\x0a\x9b\x0a\x3e\x98\x51\x92\xf3\x59\x15\xb9\xe1\x86\xe1\x1d\x19\x65\x79\x0f\x57\xfb\x2d\x59\x26\x53\xc6\xa9\x12\x4b\xf0\x59\xc2\xc1\x65\xea\xa7\x71\x91\xce\xe7\x44\xfb\xf8\x32\x4a\xd1\xd5\xe0\x8d\xb8\x3a\x20\xdc\x5b\xa2\x51\x80\xd7\xb3\x02\xfc\xef\x16\xaf\x29\x6e\x95\xa4\x28\xa8\x85\x4c\x70\xdc\xf2\xab\x8b\x39\x21\xae\x35\x91\xff\x4b\x44\xc1\x1a\xcb\x28\x46\x16\x1a\x09\x76\xdc\xee\xe7\xb3\x0a\x71\xa5\x77\xec\xe0\xce\x92\xd3\xd4\x46\x19\xf2\x60\x20\x27\xe7\x58\xcc\x04\xc8\xbc\x7d\x8b\xcb\x1a\x49\xc3\x0d\x96\x1b\x27\x66\xc0\x1f\x96\xb5\x26\x3b\x6e\x8c\x97\x51\x8c\x3b\xd9\x1e\xc9\x17\x99\x68\xdd\x35\xd7\xdc\x8b\x19\xb0\xc7\xd5\x9e\x35\x81\x18\x58\x23\x1b\xc8\xbd\x25\xf2\xe3\xda\xc4\xa1\x8e\xa4\x4d\x6a\x43\x49\x90\x6f\x2d\x9b\xca\x79\xa2\xb8\x3a\xa9\x58\x03\x5e\x20\x0b\x13\xa9\xf2\x6a\xee\xce\x75\xa9\xd1\xd5\x51\x86\x5b\x31\x4a\x0c\x5c\x62\x2f\xe1\x78\x9d\x7d\x37\x37\xc6\x15\x56\x43\xfd\xf7\x9d\x78\xd2\x1d\x06\x3f\x8a\x35\x8c\x27\x6c\xfc\x5b\x52\xf0\x1b\x6d\xb7\xa4\x39\x5e\x57\x43\x17\x9c\x58\x9a\x4c\x72\xfc\xfc\xe9\x84\x87\xad\x3b\xd1\x3c\x30\x5b\xe5\xad\x29\x6b\x2b\x86\x4b\xf6\xed\xb8\x25\xbe\x3f\xc1\xa1\x2a\x29\x6b\x1f\x4c\x05\x29\x16\xe3\x6f\x3f\xff\xfe\xf1\x0d\xc9\x71\x4b\x1e\xd8\x43\xaa\x89\x4e\x79\xb4\x9a\x9a\x0f\xb3\x50\x14\x64\xbb\x6d\x4a\x8f\x53\xcc\x9e\x15\x19\x36\x83\x62\x0c\x9c\x9f\x58\x83\xfd\x35\x0c\x25\x76\x02\xef\xb9\xaa\x3d\x24\x1e\x3a\x94\x0f\x5b\xb2\x13\x49\x31\xab\x1c\xa9\x38\x3f\x30\x19\xbe\x11\x73\xe6\xe1\x56\x64\x03\x29\x96\x9b\x71\x88\xd7\xb3\x42\xa4\xf0\xcb\x57\xa3\xc5\x97\x4e\xcb\xa6\x62\x73\xd6\x8b\x91\xea\x30\x51\x55\xaf\xe2\x55\xf5\xdf\x8d\x4f\x0f\x92\xe8\x82\xac\x1f\x86\xe9\xc7\xc1\x8f\x46\x91\x85\x04\x52\x54\x90\x73\x35\x91\xdf\xf1\x4d\xea\x58\x40\xeb\xa0\x54\x0b\x0f\xe8\xc9\xc6\xed\x7e\x2d\x7a\x03\x9d\x46\x13\x1b\xc8\xea\x87\x59\x39\xba\xea\xb9\xc6\x22\x68\x09\x0e\xee\xeb\xd0\x49\x50\xb0\x55\x03\x9d\x58\xca\x0d\x3c\xb0\xaf\x61\x3d\xfe\x87\xc3\x79\x3c\x50\x8e\x45\x95\x49\x5f\x6e\x37\x52\x9b\xa1\xf5\xb0\x2b\x45\x90\x47\xf0\x35\xc1\x1b\xd2\x1e\xbf\xc1\x44\xab\x78\xe8\x1f\x8b\x2a\xf7\xb4\xe3\x2e\x44\x5a\xf8\x84\xa1\xaa\x71\x60\xf1\xa5\x3d\x3a\xb8\xd6\x22\xe3\xa1\x3f\x30\x80\x9d\x18\xcd\xa5\x04\xad\xba\x7e\x9d\x7b\x34\xaa\xef\x76\x0f\x84\xbe\x9e\x5a\xec\xa7\x6d\xf5\xec\x8f\x7f\x03\x00\x00\xff\xff\x52\x7a\x95\x07\xe3\x16\x00\x00"

func assetsAlbumsJsonBytes() ([]byte, error) {
	return bindataRead(
		_assetsAlbumsJson,
		"assets/albums.json",
	)
}

func assetsAlbumsJson() (*asset, error) {
	bytes, err := assetsAlbumsJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/albums.json", size: 5859, mode: os.FileMode(436), modTime: time.Unix(1523784950, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsConfigReadmeMd = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x52\x56\x56\x48\xce\xcf\x2b\x49\xcc\xcc\x2b\x06\x31\xd2\x32\xd3\x4b\x8b\x52\x15\xd2\x32\x73\x52\x8b\x01\x01\x00\x00\xff\xff\x2b\x6f\x3a\x53\x1b\x00\x00\x00"

func assetsConfigReadmeMdBytes() ([]byte, error) {
	return bindataRead(
		_assetsConfigReadmeMd,
		"assets/config/README.md",
	)
}

func assetsConfigReadmeMd() (*asset, error) {
	bytes, err := assetsConfigReadmeMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/config/README.md", size: 27, mode: os.FileMode(436), modTime: time.Unix(1545785370, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsConfigAppToml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00"

func assetsConfigAppTomlBytes() ([]byte, error) {
	return bindataRead(
		_assetsConfigAppToml,
		"assets/config/app.toml",
	)
}

func assetsConfigAppToml() (*asset, error) {
	bytes, err := assetsConfigAppTomlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/config/app.toml", size: 0, mode: os.FileMode(436), modTime: time.Unix(1545785044, 0)}
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
	"assets/README.md":        assetsReadmeMd,
	"assets/albums.json":      assetsAlbumsJson,
	"assets/config/README.md": assetsConfigReadmeMd,
	"assets/config/app.toml":  assetsConfigAppToml,
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
	"assets": {nil, map[string]*bintree{
		"README.md":   {assetsReadmeMd, map[string]*bintree{}},
		"albums.json": {assetsAlbumsJson, map[string]*bintree{}},
		"config": {nil, map[string]*bintree{
			"README.md": {assetsConfigReadmeMd, map[string]*bintree{}},
			"app.toml":  {assetsConfigAppToml, map[string]*bintree{}},
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
