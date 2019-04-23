package templates

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assets4548c4de5a42a759758868d96443b68ee4bacf04 = "<!DOCTYPE html>\r\n<html lang=\"en\">\r\n<head>\r\n    <meta charset=\"UTF-8\">\r\n    {{ .title }}\r\n</head>\r\n<body>\r\n</body>\r\n</html>"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"index.tmpl"}}, map[string]*assets.File{
	"/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1556031249, 1556031249062296800),
		Data:     nil,
	}, "/index.tmpl": &assets.File{
		Path:     "/index.tmpl",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1556024802, 1556024802532217900),
		Data:     []byte(_Assets4548c4de5a42a759758868d96443b68ee4bacf04),
	}}, "")
