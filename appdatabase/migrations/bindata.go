// Code generated by go-bindata. DO NOT EDIT.
// sources:
// 0001_app.down.sql (356B)
// 0001_app.up.sql (2.967kB)
// 0002_tokens.down.sql (19B)
// 0002_tokens.up.sql (248B)
// 0003_settings.down.sql (118B)
// 0003_settings.up.sql (1.311kB)
// 0004_pending_stickers.down.sql (0)
// 0004_pending_stickers.up.sql (61B)
// 0005_waku_mode.down.sql (0)
// 0005_waku_mode.up.sql (146B)
// doc.go (74B)

package migrations

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var __0001_appDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\xcd\xcb\xaa\xc2\x40\x0c\xc6\xf1\x7d\x9f\xa2\xef\xd1\xd5\x39\xb4\x0b\x41\x54\xc4\x85\xbb\x21\x4e\x63\x1b\x6c\x27\x63\x92\x7a\x79\x7b\x41\xf0\x32\xea\x6c\x7f\xf9\xf8\xa7\x5e\x2f\x57\xe5\xe6\xef\x7f\xde\x94\x8a\x66\x14\x3a\xad\x8a\x37\x04\xef\x79\x0a\x96\xe2\x4e\xf8\xac\x28\xbf\xd1\xf5\xa4\xc6\x72\x4d\x8e\x2d\xc4\x98\xce\x23\xca\x48\xaa\xc4\x21\x75\x13\x08\xba\xff\x8a\x0f\xec\x0f\x29\x8d\x40\x83\xa2\x9c\x3e\xa7\x2f\x77\x82\xc7\x09\xd5\x5c\x07\xcf\xe7\xb3\x45\xdd\x6c\x73\x1b\xe7\x7b\x30\x47\xad\xa3\xf6\x92\x6b\x1a\x47\xf2\xd9\x8f\xf7\xc0\x23\x29\x10\x3a\xd4\xaa\xb8\x05\x00\x00\xff\xff\xf6\xca\x86\xce\x64\x01\x00\x00")

func _0001_appDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__0001_appDownSql,
		"0001_app.down.sql",
	)
}

func _0001_appDownSql() (*asset, error) {
	bytes, err := _0001_appDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "0001_app.down.sql", size: 356, mode: os.FileMode(0644), modTime: time.Unix(1579789333, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xb5, 0x25, 0xa0, 0xf8, 0x7d, 0x2d, 0xd, 0xcf, 0x18, 0xe4, 0x73, 0xc3, 0x95, 0xf5, 0x24, 0x20, 0xa9, 0xe6, 0x9e, 0x1d, 0x93, 0xe5, 0xc5, 0xad, 0x93, 0x8f, 0x5e, 0x40, 0xb5, 0x30, 0xaa, 0x25}}
	return a, nil
}

var __0001_appUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x56\x4f\x73\xa2\x30\x14\xbf\xf3\x29\x72\xd4\x19\x2e\x7b\xee\x09\x35\x5a\x66\x29\xec\x22\x6e\xdb\x53\x26\x42\xc4\x8c\x40\xd2\x24\xd4\xfa\xed\x77\x02\x04\x50\x41\xeb\xce\xde\x4c\xde\xcb\xe3\xf7\xe7\xe5\xc5\x79\x08\x9d\x08\x82\xc8\x99\x79\x10\xb8\x4b\xe0\x07\x11\x80\x6f\xee\x3a\x5a\x03\x49\x94\xa2\x45\x2a\xc1\xc4\x52\x27\x4e\xc0\x1f\x27\x9c\x3f\x3b\x21\xf8\x15\xba\x2f\x4e\xf8\x0e\x7e\xc2\x77\xdb\xfa\xc4\x59\x49\xc0\xcc\x0b\x66\xd6\x14\xbc\xba\xd1\x73\xb0\x89\x40\x18\xbc\xba\x8b\x27\xcb\xba\x51\x1c\xc7\x31\x2b\x0b\xa5\x8b\xe3\x24\x11\x44\xca\xe1\xfa\x47\x9c\x65\x44\x81\x59\x10\x78\xd0\xf1\x6d\x2b\xde\xe3\xde\xaa\xc2\x15\xc1\xb7\xc8\xb6\xa4\x62\x02\xa7\x66\xc5\xcb\xed\x81\x9c\x2a\x5c\xb6\xc5\xb1\xda\x37\xfb\x05\xce\x4d\x4a\xcc\x32\x26\xcc\x6f\x41\xb0\x22\x09\xc2\x0a\x2c\x9c\x08\x46\xee\x0b\xac\xc0\xfa\x1b\xcf\xb3\xad\x92\x27\xa3\xd1\x71\xd6\x1b\xdf\xfd\xbd\x81\xc0\xf5\x17\xf0\x0d\x94\x05\xfd\x28\x09\xaa\xd9\x20\xc3\x38\xf0\x7b\x3a\xd4\xb1\x29\x78\x7d\x86\x21\x6c\x97\x4f\xb7\xca\x69\x31\x86\x8b\xe9\x48\x5b\xaa\x5a\xb4\x85\xea\x0a\x1d\x63\xd4\x9c\xba\x28\xd0\xc6\xbb\x32\xdd\xd6\x6d\x6f\xb7\x82\x1d\x25\x11\xda\x5b\x9a\x54\x0a\x9f\x7b\xda\x9a\xd0\xd3\x58\xd1\x9c\x48\x85\x73\x0e\x36\xeb\x95\xbb\xf2\xe1\x02\xcc\xdc\x95\xeb\x47\xb6\x95\x60\xce\x8d\xe5\x60\x01\x97\xce\xc6\x8b\xc0\x0e\x67\x92\xd8\xd6\x9e\x6a\xdf\x4f\x6e\x91\x90\x2f\xb0\xf1\xd7\xf5\x49\xd7\x8f\x1e\xeb\x46\x83\x18\x35\xf5\xc0\xc4\x6a\xb6\x90\x61\xd0\x41\x35\x39\x75\xeb\x2c\x83\x10\xba\x2b\x5f\x33\x9b\x74\x67\xa6\x20\x84\x4b\x18\x42\x7f\x0e\xbb\xea\x13\xbd\x1f\x68\x0e\x1e\x8c\x20\x98\x3b\xeb\xb9\xb3\x80\xd6\x1d\x35\x35\x7d\x2d\x65\xa7\x5a\x4f\xcc\xc7\x68\x72\x22\x72\x2a\x25\x65\x85\x2e\xa8\x0b\xa3\x21\x2f\xba\xb4\xcb\x48\x9f\x6c\x7b\xfc\x8c\x6b\x85\x76\x52\x6f\x0f\x53\xbd\x05\x50\x09\x5c\xc8\x5d\xdd\x3a\x05\x51\x47\x26\x0e\xda\x80\xd6\xd8\xba\x25\xfa\x5e\x60\xb9\x6f\x07\x47\xb7\x7d\x39\x52\xba\xc8\x36\x3b\xa0\x91\x43\xea\xab\x99\x17\x92\x14\x09\x11\x26\xc3\xb6\x04\x89\x09\xe5\xaa\x89\x66\x2c\x6d\x7e\x9d\x4d\xc5\xf3\x4f\x14\x65\xbe\x25\xe2\x1a\x6f\xaf\xcd\x47\x39\x65\x0c\x27\x24\xa9\x3a\xbe\x6d\xf7\x1f\xe7\xda\x77\xda\xd8\x0d\x55\xdb\x10\x3b\xef\xbc\x8c\xc5\x07\x79\x3b\xfd\xca\x25\xdb\x9a\x07\xfe\x3a\x0a\x1d\x0d\xab\x99\x34\xc6\x18\xc4\x89\x30\x13\xa7\xfa\xdd\x94\x36\xe3\x69\xa2\x6b\xb6\x1f\xe9\xbe\x3b\xbd\xd7\xe5\x35\xd2\xef\xda\x7e\xdb\xdf\x31\xf1\x5b\xef\xbf\x25\xf9\xd2\xf1\xd6\x83\x5a\xe4\x98\x73\x5a\xa4\x68\xc7\x84\x99\x9d\x48\x31\x54\x31\x18\xd4\xe4\x52\xf3\xc7\x75\x41\x02\x17\x29\xf9\x4f\xf2\xec\x04\xcb\x87\xc5\x51\xec\x72\xff\x1e\xbc\x1c\xd3\x4c\x12\xf1\x59\x5f\x59\x00\x00\xa0\xc9\xf0\x43\xae\x63\xd5\xb0\xb9\x06\xa5\x43\xe3\x90\x75\x94\x63\x29\x8f\x4c\x24\xdd\x9d\xd4\xbb\xbb\x8c\x10\x75\x75\xe2\xb1\x91\xd8\x11\x40\x82\x7c\x94\x44\x2a\x94\x62\x6e\xc8\xa4\x98\xd7\x72\xf5\x9f\x16\xb8\x82\x97\xf8\x74\x9e\x62\xf7\xb2\x06\x1f\x43\x1d\xa8\xde\xf1\xcb\x87\x66\x9c\x47\xfd\x82\x8f\x20\x47\x4d\x31\x44\x93\x2f\x7d\xb7\x47\x09\x36\x79\xdf\x36\x18\x29\xc6\x69\x6c\x94\xa9\x16\xe3\x4e\x37\xc5\xe5\xb9\x61\x19\x96\xca\xa0\x68\x35\xea\x8d\x38\x9d\x93\x50\x19\xb3\x4f\x22\x4e\x57\x4f\x7e\x73\x21\xab\x46\x22\x29\x53\x54\xff\x1b\x19\xce\xfa\xe7\x1e\xa8\x70\x1b\x9d\xda\x4b\xd7\xf7\x68\x94\x72\xc6\x8e\xa4\xa3\x57\xb7\x4d\xc3\xb1\x4e\xd8\xd3\x74\xdf\xcf\x50\xcc\xc4\xaf\xe1\xfe\x0d\x00\x00\xff\xff\xe8\x42\x77\x9b\x97\x0b\x00\x00")

func _0001_appUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__0001_appUpSql,
		"0001_app.up.sql",
	)
}

func _0001_appUpSql() (*asset, error) {
	bytes, err := _0001_appUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "0001_app.up.sql", size: 2967, mode: os.FileMode(0644), modTime: time.Unix(1579789333, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xf7, 0x3a, 0xa7, 0xf2, 0x8f, 0xfa, 0x82, 0x7c, 0xc5, 0x49, 0xac, 0xac, 0xf, 0xc, 0x77, 0xe2, 0xba, 0xe8, 0x4d, 0xe, 0x6f, 0x5d, 0x2c, 0x2c, 0x18, 0x80, 0xc2, 0x1d, 0xe, 0x25, 0xe, 0x18}}
	return a, nil
}

var __0002_tokensDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\x28\xc9\xcf\x4e\xcd\x2b\xb6\xe6\x02\x04\x00\x00\xff\xff\xf0\xdb\x32\xa7\x13\x00\x00\x00")

func _0002_tokensDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__0002_tokensDownSql,
		"0002_tokens.down.sql",
	)
}

func _0002_tokensDownSql() (*asset, error) {
	bytes, err := _0002_tokensDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "0002_tokens.down.sql", size: 19, mode: os.FileMode(0644), modTime: time.Unix(1576226192, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xd1, 0x31, 0x2, 0xcc, 0x2f, 0x38, 0x90, 0xf7, 0x58, 0x37, 0x47, 0xf4, 0x18, 0xf7, 0x72, 0x74, 0x67, 0x14, 0x7e, 0xf3, 0xb1, 0xd6, 0x5f, 0xb0, 0xd5, 0xe7, 0x91, 0xf4, 0x26, 0x77, 0x8e, 0x68}}
	return a, nil
}

var __0002_tokensUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8e\xcd\x6a\x85\x30\x10\x46\xf7\x79\x8a\x6f\x79\x05\xdf\xa0\xab\xa8\xa9\x0e\xb5\xb1\xc4\xb1\xea\xaa\x58\x93\x85\xf8\x13\x30\x42\xe9\xdb\x17\x4b\x4b\x2b\xdc\xed\x37\x67\x0e\x27\x35\x4a\xb2\x02\xcb\xa4\x54\xa0\x47\xe8\x8a\xa1\x3a\xaa\xb9\xc6\xe1\x67\xb7\x05\xdc\x04\x30\x58\xbb\xbb\x10\xf0\x2a\x4d\x5a\x48\xf3\x4d\xe9\xa6\x2c\x63\x01\x6c\xee\xf8\xf0\xfb\xfc\x36\x59\x34\xba\xa6\x5c\xab\x0c\x09\xe5\xa4\xf9\x8a\x0d\xab\x03\xab\xee\xba\x86\xcf\xf5\xdd\x2f\x77\xbd\xd6\x8d\xd3\x3a\x2c\xe1\xcf\x4a\x9a\xcf\xc3\xe8\x17\xbf\xff\xbe\x9c\xc3\x8b\xa1\x67\x69\x7a\x3c\xa9\x1e\xb7\x9f\xd4\xf8\x5f\x57\x24\x22\xb4\xc4\x45\xd5\x30\x4c\xd5\x52\xf6\x20\xc4\x57\x00\x00\x00\xff\xff\x73\xf3\x87\xe5\xf8\x00\x00\x00")

func _0002_tokensUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__0002_tokensUpSql,
		"0002_tokens.up.sql",
	)
}

func _0002_tokensUpSql() (*asset, error) {
	bytes, err := _0002_tokensUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "0002_tokens.up.sql", size: 248, mode: os.FileMode(0644), modTime: time.Unix(1576226192, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xcc, 0xd6, 0xde, 0xd3, 0x7b, 0xee, 0x92, 0x11, 0x38, 0xa4, 0xeb, 0x84, 0xca, 0xcb, 0x37, 0x75, 0x5, 0x77, 0x7f, 0x14, 0x39, 0xee, 0xa1, 0x8b, 0xd4, 0x5c, 0x6e, 0x55, 0x6, 0x50, 0x16, 0xd4}}
	return a, nil
}

var __0003_settingsDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x3c\xca\xb1\x0a\xc2\x40\x0c\x06\xe0\x3d\x4f\xf1\x8f\x0a\xbe\x41\xa7\x5c\x1b\x69\xb0\x9a\x92\x46\x6b\x47\x87\x43\x04\x11\xe1\x4e\xc1\xb7\x77\x11\xd7\x8f\xaf\x73\x1b\x11\x9c\x06\x41\xc9\xb5\xde\x1e\xd7\xd2\x50\xeb\xc2\x21\x3f\xd6\x2d\x0e\x16\x90\xb3\x4e\x31\xfd\x13\x56\x04\xd4\xcf\x33\xe3\xc4\xde\xf6\xec\x18\x5d\xf7\xec\x0b\x76\xb2\x6c\x08\x78\x5f\xee\xaf\x8c\x34\x58\xa2\x35\x66\x8d\xde\x8e\x01\xb7\x59\xbb\x86\xe8\x1b\x00\x00\xff\xff\x49\x2e\x16\x6c\x76\x00\x00\x00")

func _0003_settingsDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__0003_settingsDownSql,
		"0003_settings.down.sql",
	)
}

func _0003_settingsDownSql() (*asset, error) {
	bytes, err := _0003_settingsDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "0003_settings.down.sql", size: 118, mode: os.FileMode(0644), modTime: time.Unix(1578050942, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xe5, 0xa6, 0xf5, 0xc0, 0x60, 0x64, 0x77, 0xe2, 0xe7, 0x3c, 0x9b, 0xb1, 0x52, 0xa9, 0x95, 0x16, 0xf8, 0x60, 0x2f, 0xa5, 0xeb, 0x46, 0xb9, 0xb9, 0x8f, 0x4c, 0xf4, 0xfd, 0xbb, 0xe7, 0xe5, 0xe5}}
	return a, nil
}

var __0003_settingsUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x93\xbd\x6e\xdb\x30\x10\xc7\x77\x3f\x05\xb7\xb4\x40\x87\x66\x28\x50\x20\x93\x1c\xab\x89\x50\x57\x0a\x54\xb9\x41\xa6\x03\x4d\x9e\xad\x83\x29\x92\xe0\x51\x0e\xf4\xf6\x85\x1c\x45\x56\x53\xd9\x23\x79\xbf\xfb\xbe\xff\xaa\x2c\x9e\x44\x95\x2c\xd7\xa9\x60\x8c\x91\xec\x9e\xef\x16\xf7\x65\x9a\x54\xe9\x87\x6f\xf1\x69\x21\x84\xd4\x3a\x20\xb3\xf8\x93\x94\xf7\x8f\x49\x29\xf2\xa2\x12\xf9\x66\xbd\xfe\xb2\x10\x42\xd5\xd2\x31\x34\x4e\xa3\x58\x16\xc5\x3a\x4d\x72\xb1\x4a\x7f\x24\x9b\x75\x25\x76\xd2\x30\x9e\x98\x36\x04\xb4\xaa\x1b\x03\xbc\x13\x37\x2d\xeb\x9b\x33\x11\xc1\x62\x7c\x75\xe1\x30\x9f\xa9\xe5\xe8\x1a\xd8\x3a\x17\xad\xd3\xc8\x62\xb9\x2e\x96\x73\x06\x40\x2b\xb7\x06\xf5\x08\x68\xe9\x3d\xc3\xb5\x2e\x90\xfc\xed\xb7\xef\xb7\x1f\x99\xde\xb4\x33\x88\x71\xfa\x51\x93\x46\xa8\x5d\x83\x10\x9d\x33\x91\xfc\xe5\xc6\xc9\x72\x94\xc6\xc8\x48\xce\x02\xe9\xd9\xd4\x07\xec\xa0\xbd\x6c\x53\x32\x68\x38\xc5\xb1\x0a\xa7\xe0\xd4\xee\x25\x05\xd4\xe0\xac\xd8\xe4\xbf\xb3\x87\x3c\x5d\x89\x65\xf6\x90\xe5\xd5\x47\x88\xec\x7e\xea\x6f\x24\x47\x68\xbd\x96\x11\xf5\x9c\xab\x91\x11\x39\x82\xc6\x40\x47\xec\x23\xc4\xfa\x8c\x65\x79\x35\x76\xfc\xf5\x44\xbb\x3d\x18\x3c\xa2\x99\xa6\x68\x2c\x36\xce\x92\x9a\xfe\x59\xd9\xe0\x6c\xbf\xc3\xfa\xdf\x56\xfb\xaf\xc5\x69\x04\xe5\xec\x8e\xf6\xe3\x5a\xad\x8b\xb4\x23\x75\x9a\xee\x64\xe9\x97\x96\xe1\x6b\x17\xdd\x5b\x0f\xff\x85\xf7\x64\x2d\x6a\x68\x24\x19\xc6\x70\xc4\x70\xbe\x2e\x1f\x70\x87\xa1\x1f\xef\xb4\xec\xc1\x72\x24\x7c\x05\x1f\xe8\x28\x55\x77\x25\x73\xbb\x35\xa4\xe0\x80\xdd\x6c\xd7\x01\x1b\x6c\xb6\x18\x80\x3b\xab\xc8\xee\x41\xd5\x8e\xd4\x15\x3d\x31\xed\x6d\xcf\xf9\x3a\x48\x9e\x9f\x24\x47\x52\x07\x0c\x0c\x5e\xaa\x03\xc3\x70\x88\x13\x4d\x8c\x40\x40\xd5\x2b\xef\xfd\x7d\x06\x86\x62\x9c\x85\xc6\x6d\xc9\xe0\x28\xce\xcb\x75\x75\x36\xd6\x18\x49\x4d\x6f\x7d\xd4\x3a\xe9\x1b\xf1\x54\x66\xbf\x92\xf2\x45\xfc\x4c\x5f\x7a\x87\x96\x31\xf4\x53\x3d\x67\x7d\xed\xab\x8c\x10\x9c\x8b\x57\x05\x3b\x70\x8c\xfd\xfd\x82\x97\xcc\xd7\x56\x3f\xd0\x47\x62\xda\x9a\x5e\xb7\x07\xb4\x63\xdc\xc5\x67\xf1\x9c\x55\x8f\xc5\xa6\x12\x65\xf1\x9c\xad\xee\x16\x7f\x03\x00\x00\xff\xff\xa5\xa1\x7b\x78\x1f\x05\x00\x00")

func _0003_settingsUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__0003_settingsUpSql,
		"0003_settings.up.sql",
	)
}

func _0003_settingsUpSql() (*asset, error) {
	bytes, err := _0003_settingsUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "0003_settings.up.sql", size: 1311, mode: os.FileMode(0644), modTime: time.Unix(1578050942, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xea, 0x35, 0x0, 0xeb, 0xe2, 0x33, 0x68, 0xb9, 0xf4, 0xf6, 0x8e, 0x9e, 0x10, 0xe9, 0x58, 0x68, 0x28, 0xb, 0xcd, 0xec, 0x74, 0x71, 0xa7, 0x9a, 0x5a, 0x77, 0x59, 0xb1, 0x13, 0x1c, 0xa1, 0x5b}}
	return a, nil
}

var __0004_pending_stickersDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func _0004_pending_stickersDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__0004_pending_stickersDownSql,
		"0004_pending_stickers.down.sql",
	)
}

func _0004_pending_stickersDownSql() (*asset, error) {
	bytes, err := _0004_pending_stickersDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "0004_pending_stickers.down.sql", size: 0, mode: os.FileMode(0644), modTime: time.Unix(1581952751, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xe3, 0xb0, 0xc4, 0x42, 0x98, 0xfc, 0x1c, 0x14, 0x9a, 0xfb, 0xf4, 0xc8, 0x99, 0x6f, 0xb9, 0x24, 0x27, 0xae, 0x41, 0xe4, 0x64, 0x9b, 0x93, 0x4c, 0xa4, 0x95, 0x99, 0x1b, 0x78, 0x52, 0xb8, 0x55}}
	return a, nil
}

var __0004_pending_stickersUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x28\x4e\x2d\x29\xc9\xcc\x4b\x2f\x56\x70\x74\x71\x51\x70\xf6\xf7\x09\xf5\xf5\x53\x28\x2e\xc9\x4c\xce\x4e\x2d\x2a\x8e\x2f\x48\x4c\xce\x2e\x8e\x2f\x48\xcd\x4b\xc9\xcc\x4b\x57\x70\xf2\xf1\x77\xb2\xe6\x02\x04\x00\x00\xff\xff\xc9\xc1\xc2\xc6\x3d\x00\x00\x00")

func _0004_pending_stickersUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__0004_pending_stickersUpSql,
		"0004_pending_stickers.up.sql",
	)
}

func _0004_pending_stickersUpSql() (*asset, error) {
	bytes, err := _0004_pending_stickersUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "0004_pending_stickers.up.sql", size: 61, mode: os.FileMode(0644), modTime: time.Unix(1581952751, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x3c, 0xed, 0x25, 0xdf, 0x75, 0x2, 0x6c, 0xf0, 0xa2, 0xa8, 0x37, 0x62, 0x65, 0xad, 0xfd, 0x98, 0xa0, 0x9d, 0x63, 0x94, 0xdf, 0x6b, 0x46, 0xe0, 0x68, 0xec, 0x9c, 0x7f, 0x77, 0xdd, 0xb3, 0x6}}
	return a, nil
}

var __0005_waku_modeDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func _0005_waku_modeDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__0005_waku_modeDownSql,
		"0005_waku_mode.down.sql",
	)
}

func _0005_waku_modeDownSql() (*asset, error) {
	bytes, err := _0005_waku_modeDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "0005_waku_mode.down.sql", size: 0, mode: os.FileMode(0644), modTime: time.Unix(1582132969, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xe3, 0xb0, 0xc4, 0x42, 0x98, 0xfc, 0x1c, 0x14, 0x9a, 0xfb, 0xf4, 0xc8, 0x99, 0x6f, 0xb9, 0x24, 0x27, 0xae, 0x41, 0xe4, 0x64, 0x9b, 0x93, 0x4c, 0xa4, 0x95, 0x99, 0x1b, 0x78, 0x52, 0xb8, 0x55}}
	return a, nil
}

var __0005_waku_modeUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x28\x4e\x2d\x29\xc9\xcc\x4b\x2f\x56\x70\x74\x71\x51\x70\xf6\xf7\x09\xf5\xf5\x53\x28\x4f\xcc\x2e\x8d\x4f\xcd\x4b\x4c\xca\x49\x4d\x51\x70\xf2\xf7\xf7\x71\x75\xf4\x53\x70\x71\x75\x73\x0c\xf5\x09\x51\x48\x4b\xcc\x29\x4e\xb5\xe6\x22\xca\x8c\xa4\x9c\xfc\xfc\xdc\xf8\xb4\xcc\x9c\x92\xd4\xa2\xf8\xdc\xfc\x94\x54\x5c\xa6\x01\x02\x00\x00\xff\xff\x00\x97\x79\x75\x92\x00\x00\x00")

func _0005_waku_modeUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__0005_waku_modeUpSql,
		"0005_waku_mode.up.sql",
	)
}

func _0005_waku_modeUpSql() (*asset, error) {
	bytes, err := _0005_waku_modeUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "0005_waku_mode.up.sql", size: 146, mode: os.FileMode(0644), modTime: time.Unix(1582129313, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xa6, 0x91, 0xc, 0xd7, 0x89, 0x61, 0x2e, 0x4c, 0x5a, 0xb6, 0x67, 0xd1, 0xc1, 0x42, 0x24, 0x38, 0xd6, 0x1b, 0x75, 0x41, 0x9c, 0x23, 0xb0, 0xca, 0x5c, 0xf1, 0x5c, 0xd0, 0x13, 0x92, 0x3e, 0xe1}}
	return a, nil
}

var _docGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2c\xc9\xb1\x0d\xc4\x20\x0c\x05\xd0\x9e\x29\xfe\x02\xd8\xfd\x6d\xe3\x4b\xac\x2f\x44\x82\x09\x78\x7f\xa5\x49\xfd\xa6\x1d\xdd\xe8\xd8\xcf\x55\x8a\x2a\xe3\x47\x1f\xbe\x2c\x1d\x8c\xfa\x6f\xe3\xb4\x34\xd4\xd9\x89\xbb\x71\x59\xb6\x18\x1b\x35\x20\xa2\x9f\x0a\x03\xa2\xe5\x0d\x00\x00\xff\xff\x60\xcd\x06\xbe\x4a\x00\x00\x00")

func docGoBytes() ([]byte, error) {
	return bindataRead(
		_docGo,
		"doc.go",
	)
}

func docGo() (*asset, error) {
	bytes, err := docGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "doc.go", size: 74, mode: os.FileMode(0644), modTime: time.Unix(1573806410, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xde, 0x7c, 0x28, 0xcd, 0x47, 0xf2, 0xfa, 0x7c, 0x51, 0x2d, 0xd8, 0x38, 0xb, 0xb0, 0x34, 0x9d, 0x4c, 0x62, 0xa, 0x9e, 0x28, 0xc3, 0x31, 0x23, 0xd9, 0xbb, 0x89, 0x9f, 0xa0, 0x89, 0x1f, 0xe8}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"0001_app.down.sql": _0001_appDownSql,

	"0001_app.up.sql": _0001_appUpSql,

	"0002_tokens.down.sql": _0002_tokensDownSql,

	"0002_tokens.up.sql": _0002_tokensUpSql,

	"0003_settings.down.sql": _0003_settingsDownSql,

	"0003_settings.up.sql": _0003_settingsUpSql,

	"0004_pending_stickers.down.sql": _0004_pending_stickersDownSql,

	"0004_pending_stickers.up.sql": _0004_pending_stickersUpSql,

	"0005_waku_mode.down.sql": _0005_waku_modeDownSql,

	"0005_waku_mode.up.sql": _0005_waku_modeUpSql,

	"doc.go": docGo,
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
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"0001_app.down.sql":              &bintree{_0001_appDownSql, map[string]*bintree{}},
	"0001_app.up.sql":                &bintree{_0001_appUpSql, map[string]*bintree{}},
	"0002_tokens.down.sql":           &bintree{_0002_tokensDownSql, map[string]*bintree{}},
	"0002_tokens.up.sql":             &bintree{_0002_tokensUpSql, map[string]*bintree{}},
	"0003_settings.down.sql":         &bintree{_0003_settingsDownSql, map[string]*bintree{}},
	"0003_settings.up.sql":           &bintree{_0003_settingsUpSql, map[string]*bintree{}},
	"0004_pending_stickers.down.sql": &bintree{_0004_pending_stickersDownSql, map[string]*bintree{}},
	"0004_pending_stickers.up.sql":   &bintree{_0004_pending_stickersUpSql, map[string]*bintree{}},
	"0005_waku_mode.down.sql":        &bintree{_0005_waku_modeDownSql, map[string]*bintree{}},
	"0005_waku_mode.up.sql":          &bintree{_0005_waku_modeUpSql, map[string]*bintree{}},
	"doc.go":                         &bintree{docGo, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
