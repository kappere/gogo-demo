package assets

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assetsba30dd761a21783ebe2602f773a63520560ded5b = "database:\n  dbtype: mysql\n  url: root:123456@(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local"
var _Assets9bafd3519abaf1bf004bb6a6e2e8e26901daaccf = "database:\r\n  dbtype: mysql\r\n  url: root:123456@(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local"
var _Assets4b477b42944b87fd4e11e277a534d68f945d83fe = "server:\r\n  port: 8081\r\nlog:\r\n  path: log/test2\r\nredis:\r\n  host: localhost\r\n  port: 6379\r\n  password:\r\nexternal: "

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"config"}, "/config": []string{"config-dev.yml", "config-prod.yml", "config.yml"}}, map[string]*assets.File{
	"/config/config-dev.yml": &assets.File{
		Path:     "/config/config-dev.yml",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1606889845, 1606889845329631500),
		Data:     []byte(_Assetsba30dd761a21783ebe2602f773a63520560ded5b),
	}, "/config/config-prod.yml": &assets.File{
		Path:     "/config/config-prod.yml",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1606889853, 1606889853489407800),
		Data:     []byte(_Assets9bafd3519abaf1bf004bb6a6e2e8e26901daaccf),
	}, "/config/config.yml": &assets.File{
		Path:     "/config/config.yml",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1595410381, 1595410381257674300),
		Data:     []byte(_Assets4b477b42944b87fd4e11e277a534d68f945d83fe),
	}, "/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1606889810, 1606889810552662100),
		Data:     nil,
	}, "/config": &assets.File{
		Path:     "/config",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1594567723, 1594567723653593900),
		Data:     nil,
	}}, "")
