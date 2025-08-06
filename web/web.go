package web

import (
	"embed"
)

//go:generate pnpm rollup --config --silent

var (
	CssBundle string = "not-built-correctly"
	JsBundle  string = "not-built-correctly"

	//go:embed assets
	Assets embed.FS

	//go:embed dist
	Dist embed.FS
)
