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

// Build the application for multiple platforms
func (Build) All() error {
	// List of targets to build for
	targets := []struct{ OS, Arch string }{
		{"linux", "amd64"},
		{"darwin", "arm64"},
		{"windows", "amd64"},
	}

	// Loop through each target OS/Arch combination and build
	for _, target := range targets {
		if err := buildBinary(target.OS, target.Arch); err != nil {
			return err
		}
	}
	return nil
}

// Helper function to build binary for a specific OS and architecture
func buildBinary(goos, goarch string) error {

	// Define the output binary file path in the build folder
	outputPath := filepath.Join("build", fmt.Sprintf("golang-techstack-%s_%s", goos, goarch))
	if goos == "windows" {
		outputPath += ".exe" // Add .exe for Windows
	}

	// Ensure the build directory exists
	if err := os.MkdirAll("build", 0755); err != nil {
		return err
	}

	// Run the go build command for the specified target OS and architecture
	return sh.RunWith(
		map[string]string{
			"GOOS":        goos,
			"GOARCH":      goarch,
			"CGO_ENABLED": "0", // Disable CGO for static binaries
		},
		"go", "build", "-o", outputPath, "./cmd/golang-techstack",
	)
}
