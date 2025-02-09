//go:build mage
// +build mage

package main

import (
	"fmt"
	"os"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

type Build mg.Namespace

// Build the application for multiple platforms
func (Build) All() error {
	fmt.Println("Building all targets...")

	// Create build directory if not exists
	if err := os.MkdirAll("tmp/build", 0755); err != nil {
		return err
	}

	// Build commands using tmp/build directory
	commands := []struct {
		os   string
		arch string
		ext  string
	}{
		{"linux", "amd64", ""},
		{"darwin", "arm64", ""},
		{"windows", "amd64", ".exe"},
	}

	for _, cmd := range commands {
		output := fmt.Sprintf("build/golang-techstack-%s_%s%s",
			cmd.os, cmd.arch, cmd.ext)
		err := sh.Run("go", "build",
			"-o", output,
			"./cmd/golang-techstack")
		if err != nil {
			return err
		}
	}
	return nil
}

/*
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
*/
// Release namespace
type Release mg.Namespace

// Full release process
func (Release) Full() error {
	fmt.Println("Creating full release...")
	defer os.RemoveAll("build") // Clean up after release
	mg.Deps(Build{}.All)
	return sh.Run("goreleaser", "release", "--clean")
}

// Docker namespace
type Docker mg.Namespace

// Build Docker image
func (Docker) Build() error {
	fmt.Println("Building Docker image...")
	return sh.Run("docker", "build", "-t", "golang-techstack:latest", ".")
}
