//go:build mage
package main

import (
    "github.com/magefile/mage/sh"
)

// Runs the server for local development
func Server() error {
    return sh.Run("go", "run", "./cmd/liber", "server")
}