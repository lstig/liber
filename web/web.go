package web

import "embed"

//go:embed assets
var Assets embed.FS

//go:embed views
var Views embed.FS