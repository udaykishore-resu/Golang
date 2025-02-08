//go:build mage
// +build mage

package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

type Build mg.Namespace

// Build all platforms
func (Build) All() error {
	targets := []struct{ OS, Arch string }{
		{"linux", "amd64"},
		{"darwin", "arm64"},
		{"windows", "amd64"},
	}

	for _, t := range targets {
		if err := buildBinary(t.OS, t.Arch); err != nil {
			return err
		}
	}
	return nil
}

func buildBinary(goos, goarch string) error {
	outputDir := filepath.Join("build", fmt.Sprintf("%s_%s", goos, goarch))
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return err
	}

	outputPath := filepath.Join(outputDir, "myapp")
	if goos == "windows" {
		outputPath += ".exe"
	}

	return sh.RunWith(
		map[string]string{
			"GOOS":        goos,
			"GOARCH":      goarch,
			"CGO_ENABLED": "0", // Disable CGO for static binaries
		},
		"go", "build", "-o", outputPath, "./cmd/app",
	)
}
