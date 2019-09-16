// Code generated for package docs by go-bindata DO NOT EDIT. (@generated)
// sources:
// swagger.json
package docs

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

var _swaggerJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x58\x4b\x6f\xe3\x36\x10\xbe\xe7\x57\x0c\xd8\x1e\x37\x96\x93\x6d\x0f\xcd\x6d\x9b\x7d\xd4\x40\x17\x30\x8a\xf4\xd4\xee\x81\x91\xc6\x12\x37\x12\xc9\xce\x90\xf6\x1a\x41\xfe\x7b\x41\x4a\xb2\xf5\x72\xed\x64\x83\x6d\xf6\x66\x71\x1e\xfc\x38\x8f\x6f\x06\xbe\x3f\x03\x10\xa9\xd1\xec\x2b\x64\x71\x05\x7f\x9d\x01\x00\x08\x69\x6d\xa9\x52\xe9\x94\xd1\xc9\x67\x36\x5a\x9c\x01\x7c\x7a\x15\x74\x2d\x99\xcc\xa7\xa7\xe9\x72\x5a\x60\xcf\x6d\xe1\x9c\xe5\x8e\x7c\x23\xf3\x1c\x49\x5c\x81\xb8\x9c\xcd\x45\x3c\x53\x7a\x65\xc4\x15\xdc\xd7\x06\x19\x72\x4a\xca\x06\xdf\x41\xeb\xa6\x40\xb0\x9e\xac\x61\x04\xb3\x02\x57\x28\x86\xcc\xa4\xbe\x42\xed\x22\x00\x50\x0c\xce\x80\x25\xb3\x56\x19\x42\x86\x6b\x2c\x8d\x45\x62\x90\x1a\x94\x66\x95\x17\x2e\x58\x16\x66\x03\xce\xfc\xad\x95\x76\x48\x32\x75\xb0\x51\xae\x80\x8f\x5b\x76\x48\xca\x57\xf0\x1e\x31\xbb\x95\xe9\x1d\xbc\x59\x2e\x22\x2e\x00\xe1\x94\x2b\x31\xa0\x98\x12\x96\x2a\x45\xcd\xb8\x43\x0e\x20\xb4\xac\xa2\xfa\x87\xe5\xef\xeb\xd7\x8d\x1e\x80\xf0\x54\x86\xd3\x18\x8a\xab\x24\xd9\x6c\x36\xb3\x5c\xfb\x99\xa1\x3c\x69\x9c\x70\x92\xdb\xf2\xfc\xf5\x6c\x3e\x2b\x5c\x55\x8a\x68\xf8\xd0\xdc\xb3\x46\xe2\x26\x18\xf3\xd9\x7c\x76\x11\xa4\x51\x26\x0a\xc3\x2e\x1c\x97\x26\x95\x65\xfc\x88\xc7\xb7\x92\x71\x29\x5d\x11\x44\x89\xb4\x2a\x59\x5f\xd4\x02\x2b\x5d\xc1\xfb\x48\x27\xb9\x72\x85\xbf\xed\x3e\xc0\xd6\x1e\xdb\xef\x71\x3a\x2e\x80\xf0\x1f\x8f\xec\xc0\x22\x41\xa5\xb4\x77\x18\x32\x20\xcb\xd2\x6c\x30\xdb\xbd\x79\xaa\xc4\x9a\xf3\xca\x97\x4e\x59\x49\x2e\x59\x19\xaa\xce\x33\xe9\xa4\xd8\xc9\x3f\x75\x1c\x8c\xea\xae\x39\x9f\xaa\xbe\xb1\x35\xfb\xaa\x92\xb4\x0d\xa0\xaf\x09\xa5\x43\x06\x09\x1a\x37\xf0\x21\xbe\x1a\x14\xb3\xc7\xba\x08\x3c\x23\x01\xa1\x35\xe4\xba\x0f\x08\x55\x14\x6f\x59\x64\xc1\x4b\x1a\xbd\xd4\xd6\x8b\x60\xdc\xd5\xb5\x92\x64\x85\x0e\x69\x88\xf6\xbe\xf3\x3b\x54\xd4\xd6\xc6\x0a\x61\x47\x4a\xe7\x1d\x0f\x51\xfa\xe5\x3c\x37\xe7\x6d\x11\xfd\xc9\x48\x8b\x6c\xa8\xd2\x4a\xfd\xa4\x54\xc5\x24\x85\xb8\xbe\x0d\x61\x1d\x48\x43\xea\x14\x61\x78\x8d\x23\x8f\x1d\xe1\xc3\xab\x67\x81\xfc\xb6\x53\x2c\x07\x70\x67\x87\x55\xfe\x67\xf0\xef\x2a\xa9\xca\x43\xb0\x71\x4a\x38\x00\xfc\x48\x48\x2b\x55\xe2\x7f\x02\x7a\x3f\xa1\xd0\xca\xa6\x8c\x9f\x1c\xbf\xc9\xfe\x21\x64\x6b\x02\x35\xf5\xf8\x00\x40\x5c\xce\xe7\x83\xa3\x31\x4d\xc4\xfe\x80\xba\x63\x32\x50\xba\x69\xba\x21\xa6\x38\x2d\xe4\xc8\x1b\x80\xf8\x91\x70\x15\x1c\xfd\x90\x64\xb8\x52\x5a\x05\xc7\x9c\x5c\x0f\x5b\xf0\x8f\x06\xa4\xe8\xd9\x3f\x1c\xca\x84\xf8\xe9\x04\xec\xbf\xca\xac\x25\xb9\xaf\xc5\xfb\x66\xb9\x78\x47\x64\xe8\xd1\x30\x2f\x7f\x39\x0a\xf3\xc6\x18\xa8\xa4\xde\xb6\x58\xf9\x60\xfd\x89\x9f\x4f\xc9\x58\x18\x8b\x5a\x96\xc0\x48\x6b\x24\xc0\x80\xfb\x1b\xbe\xff\x6c\xf8\xeb\xa1\x37\x04\x13\x1b\xda\xb7\x33\xa9\x72\x1c\x0c\xaa\x0e\xdf\xd7\xd7\x65\xf5\x5e\xd0\x6f\xfb\x21\xa9\x0f\xc5\x5f\x55\xf5\x4b\xa5\x73\x60\x9f\xa6\xc8\xbc\xf2\x23\xba\x78\x6c\xf4\x82\xbb\x27\x44\x6c\xb7\x1e\x74\x7c\xed\x87\x7e\x9b\x92\x6e\x2c\x07\xcf\x68\x55\xc2\x48\x24\x64\xd4\x2e\x8c\x4e\x56\x3a\x2f\xb1\xae\x8b\xd0\xd1\x52\xc3\x28\xbd\xbb\x8d\xa7\xe5\x38\x73\xfb\x19\xd3\x7d\x1b\x85\x91\x6e\x91\x9c\x1a\xc4\x57\x5c\x4b\xdf\xdb\xa4\x60\x82\xbb\xf7\xef\xed\xe4\xab\x42\x66\x99\x1f\x33\xed\x75\x43\x8f\x63\x3f\x36\xf6\xe3\x60\xee\x30\x47\x7d\x2b\xd3\xbb\xfa\x1e\x51\x6f\x4d\xb3\xd4\x54\x49\xd5\x6e\x8f\x1a\xdd\xc6\xd0\x5d\xb2\x6a\xf6\xc4\x44\xe9\x15\xc9\xb0\x7c\xd5\x7d\xd4\xab\xe5\x51\xd8\x8e\xa7\xa2\x55\xdd\xa7\x04\x7e\xbb\xb9\x59\x36\xd9\xb0\x72\x5b\x1a\x99\x3d\x3d\xfc\xd1\xcd\xa8\xe4\x5b\x2f\x92\x48\x6e\xfb\x31\x54\x0e\xab\xa1\xfe\x51\x16\x38\x4c\x50\xfd\x39\x5c\x83\xf9\x96\x19\x39\x3c\x57\x0e\xa7\xe6\xa0\xcd\xa0\x6d\x76\x74\x00\x79\x77\xf1\x8c\xb3\xb1\xbb\x04\x3d\x3a\x69\xd1\x4f\x64\xb1\xa7\x95\xfe\xa2\xb1\x7f\xee\x40\xb7\x3f\xfa\x21\x3e\x46\x3a\xa7\x31\xce\x33\xd1\xcd\x4b\x65\x8d\xb4\x54\xa8\xdd\x44\xe0\x4e\xa8\xc7\xef\x9f\x27\x5e\x00\x49\x4c\x25\x60\x39\xd8\x3c\xa6\x86\x7e\xa7\x6e\x6d\xfd\xfd\x52\x4a\xb4\xa7\x81\x5f\x64\x65\xeb\x7f\x34\xac\xe9\x8d\xd4\x67\xe4\x58\xb1\x5f\x43\xce\x1e\xfe\x0d\x00\x00\xff\xff\x36\xf9\x5a\x59\x6d\x12\x00\x00")

func swaggerJsonBytes() ([]byte, error) {
	return bindataRead(
		_swaggerJson,
		"swagger.json",
	)
}

func swaggerJson() (*asset, error) {
	bytes, err := swaggerJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "swagger.json", size: 4717, mode: os.FileMode(420), modTime: time.Unix(1568642595, 0)}
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
	"swagger.json": swaggerJson,
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
	"swagger.json": &bintree{swaggerJson, map[string]*bintree{}},
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
