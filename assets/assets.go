package assets

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assets4acb00d51ec2010e84c91383c371dd5238ecd483 = "<!doctype html>\r\n<html lang=\"zh-CN\">\r\n<head>\r\n    <meta charset=\"UTF-8\">\r\n    <meta name=\"viewport\"\r\n          content=\"width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0\">\r\n    <meta http-equiv=\"X-UA-Compatible\" content=\"ie=edge\">\r\n    <title>Hello</title>\r\n</head>\r\n<body>\r\n<p>Hello {{ .Name }}</p>\r\n</body>\r\n</html>"
var _Assets1a5e8ea491317ea9b66346403f3439020b55361c = "   ______      ______         ____                     \r\n  / ____/___  / ____/___     / __ \\___  ____ ___  ____ \r\n / / __/ __ \\/ / __/ __ \\   / / / / _ \\/ __ `__ \\/ __ \\\r\n/ /_/ / /_/ / /_/ / /_/ /  / /_/ /  __/ / / / / / /_/ /\r\n\\____/\\____/\\____/\\____/  /_____/\\___/_/ /_/ /_/\\____/ \r\n\r\n        GOGO v1.0.1\r\n"
var _Assets6fd96f1754d0069b388ae11a51b82b57c4117a23 = "database:\n  dbtype: mysql\n  url: root:123456@(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local"
var _Assets00617c304d6ad2c25c296dd330f4f763678f1e24 = "database:\r\n  dbtype: mysql\r\n  url: root:123456@(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local"
var _Assets40f82363455a857ff0703a4d297829c6e7086ce3 = "server:\r\n  port: 8081\r\nlog:\r\n  path: log/gogo-demo\r\nredis:\r\n  host: localhost\r\n  port: 6379\r\n  password:\r\nexternal: "

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"resources"}, "/resources": []string{"banner.txt", "config-dev.yml", "config-prod.yml", "config.yml"}, "/resources/templates": []string{"test.html"}}, map[string]*assets.File{
	"/resources": &assets.File{
		Path:     "/resources",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1606976968, 1606976968502558900),
		Data:     nil,
	}, "/resources/banner.txt": &assets.File{
		Path:     "/resources/banner.txt",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1606978535, 1606978535816392300),
		Data:     []byte(_Assets1a5e8ea491317ea9b66346403f3439020b55361c),
	}, "/resources/config-dev.yml": &assets.File{
		Path:     "/resources/config-dev.yml",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1606981622, 1606981622797042300),
		Data:     []byte(_Assets6fd96f1754d0069b388ae11a51b82b57c4117a23),
	}, "/resources/config-prod.yml": &assets.File{
		Path:     "/resources/config-prod.yml",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1606981626, 1606981626062764600),
		Data:     []byte(_Assets00617c304d6ad2c25c296dd330f4f763678f1e24),
	}, "/resources/config.yml": &assets.File{
		Path:     "/resources/config.yml",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1606959880, 1606959880216449400),
		Data:     []byte(_Assets40f82363455a857ff0703a4d297829c6e7086ce3),
	}, "/resources/templates": &assets.File{
		Path:     "/resources/templates",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1606981308, 1606981308913017600),
		Data:     nil,
	}, "/resources/templates/test.html": &assets.File{
		Path:     "/resources/templates/test.html",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1606981310, 1606981310470937800),
		Data:     []byte(_Assets4acb00d51ec2010e84c91383c371dd5238ecd483),
	}, "/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1606981058, 1606981058146664000),
		Data:     nil,
	}}, "")
