//go:build mage

package main

import (
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

var (
	// commands
	airCmd   = "air"
	goCmd    = "go"
	templCmd = "templ"
	yarnCmd  = "yarn"

	// build variables
	bin    = "out/liber"
	module = "github.com/lstig/liber/cmd/liber"
)

type Dev mg.Namespace

// Builds the binary for development in out/liber
func (Dev) Build() error {
	mg.Deps(css, js, templ)
	return sh.RunV(goCmd, "build", "-o", bin, module)
}

// Runs a development server with hot-reloading
func (Dev) Run() error {
	return sh.RunV(airCmd, "--build.bin", bin, "server", "--dev")
}

func css() error {
	return sh.RunV(yarnCmd, "postcss", "-o", "web/dist/main.css", "web/src/*.css")
}

func js() error {
	return sh.RunV(yarnCmd, "rollup", "--config", "--silent")
}

func templ() error {
	return sh.RunV(templCmd, "generate")
}
