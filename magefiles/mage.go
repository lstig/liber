//go:build mage

package main

import (
	"github.com/magefile/mage/sh"
)

// Runs the server for local development
func Dev() error {
	return sh.RunV("air", "server")
}
