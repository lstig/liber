package views

import (
	"path"

	"github.com/lstig/liber/web"
)

//go:generate go tool -modfile=../../tools/go.mod templ generate

var (
	js  = path.Join("/dist", web.JsBundle)
	css = path.Join("/dist", web.CssBundle)
)
