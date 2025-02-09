//go:build mage
// +build mage

package main

import (
	"fmt"
	"os"
	"strings"

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

// Full creates a full production release and pushes Docker image
func (Release) Full() error {
	mg.Deps(Release.createProduction, Docker.Push)
	return nil
}

// Create production release (without snapshot)
func (Release) createProduction() error {
	fmt.Println("Creating production release with GoReleaser...")
	return sh.Run("goreleaser", "release", "--clean")
}

// Test runs go tests
func Test() error {
	fmt.Println("Running tests...")
	return sh.Run("go", "test", "./...")
}

// Docker namespace
type Docker mg.Namespace

// Build Docker image with version tag
func (Docker) Build() error {
	fmt.Println("Building Docker image...")
	version, err := getVersion()
	if err != nil {
		return err
	}

	return sh.Run("docker", "build",
		"-t", "golang-techstack:latest",
		"-t", fmt.Sprintf("golang-techstack:%s", version),
		".",
	)
}

// Push Docker image to registry
func (Docker) Push() error {
	fmt.Println("Pushing Docker image...")
	version, err := getVersion()
	if err != nil {
		return err
	}

	mg.Deps(Docker.Build)

	return sh.Run("docker", "push",
		fmt.Sprintf("golang-techstack:%s", version),
	)
}

// Helper to get current version from git tag
func getVersion() (string, error) {
	version, err := sh.Output("git", "describe", "--tags", "--abbrev=0")
	if err != nil {
		return "", fmt.Errorf("failed to get version: %w", err)
	}
	return strings.TrimSpace(version), nil
}
