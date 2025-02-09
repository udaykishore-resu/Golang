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

// Build namespace
type Build mg.Namespace

// Build the application for multiple platforms with correct cross-compilation
func (Build) All() error {
	fmt.Println("Building all targets...")

	// Create build directory if not exists
	if err := os.MkdirAll("build", 0755); err != nil {
		return err
	}

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
		output := fmt.Sprintf("build/golang-techstack-%s_%s%s", cmd.os, cmd.arch, cmd.ext)
		err := sh.RunWith(
			map[string]string{"GOOS": cmd.os, "GOARCH": cmd.arch},
			"go", "build", "-o", output, "./cmd/golang-techstack",
		)
		if err != nil {
			return err
		}
	}
	return nil
}

// Release namespace
type Release mg.Namespace

// checkGitState checks if the git repository is clean
func checkGitState() error {
	// Check if there are uncommitted changes
	output, err := sh.Output("git", "status", "--porcelain")
	if err != nil {
		return fmt.Errorf("failed to check git status: %w", err)
	}

	if strings.TrimSpace(output) != "" {
		return fmt.Errorf("git repository is not clean. Please commit all changes before releasing")
	}
	return nil
}

// Full release process using Goreleaser (let it handle builds)
func (Release) Full() error {
	fmt.Println("Creating full release with Goreleaser...")

	// Check git state before proceeding
	if err := checkGitState(); err != nil {
		return err
	}

	return sh.Run("goreleaser", "release", "--clean")
}

// Docker namespace
type Docker mg.Namespace

// Build Docker image
func (Docker) Build() error {
	fmt.Println("Building Docker image...")
	return sh.Run("docker", "build", "-t", "golang-techstack:latest", ".")
}
