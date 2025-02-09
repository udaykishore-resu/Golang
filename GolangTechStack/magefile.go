//go:build mage
// +build mage

package main

import (
	"fmt"
	"os"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

// Default target to run when none is specified
var Default = Build.All

// Build namespace
type Build mg.Namespace

// All builds the project for all platforms
func (Build) All() error {
	fmt.Println("Building for all platforms...")
	if err := os.MkdirAll("build", 0755); err != nil {
		return fmt.Errorf("failed to create build directory: %w", err)
	}

	platforms := []struct {
		os   string
		arch string
		ext  string
	}{
		{"linux", "amd64", ""},
		{"darwin", "arm64", ""},
		{"windows", "amd64", ".exe"},
	}

	for _, p := range platforms {
		binary := fmt.Sprintf("build/golang-techstack-%s_%s%s", p.os, p.arch, p.ext)
		fmt.Printf("Building for %s/%s...\n", p.os, p.arch)

		env := map[string]string{
			"GOOS":   p.os,
			"GOARCH": p.arch,
		}

		if err := sh.RunWith(env, "go", "build", "-o", binary, "./cmd/golang-techstack"); err != nil {
			return fmt.Errorf("error building for %s/%s: %w", p.os, p.arch, err)
		}
	}
	return nil
}

// Clean removes build artifacts
func (Build) Clean() error {
	fmt.Println("Cleaning build directory...")
	return os.RemoveAll("build")
}

// Release namespace
type Release mg.Namespace

// Create creates a new release using goreleaser in snapshot mode
func (Release) Create() error {
	fmt.Println("Creating local release with GoReleaser...")

	// Clean the dist directory
	if err := os.RemoveAll("dist"); err != nil {
		return fmt.Errorf("failed to clean dist directory: %w", err)
	}

	// Set environment variable to indicate this is a snapshot
	env := map[string]string{
		"SNAPSHOT": "true",
	}

	// Create a local release
	return sh.RunWith(env, "goreleaser", "release", "--snapshot", "--clean")
}

// Test runs go tests
func Test() error {
	fmt.Println("Running tests...")
	return sh.Run("go", "test", "./...")
}
