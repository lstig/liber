package web

import "embed"

//go:embed assets
var Assets embed.FS

//go:embed dist
var Dist embed.FS
