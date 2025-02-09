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

func (Build) All() error {
	fmt.Println("Building all targets...")

	// Use temporary directory
	if err := os.MkdirAll("dist", 0755); err != nil {
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
		outputPath := fmt.Sprintf("dist/golang-techstack-%s_%s%s", cmd.os, cmd.arch, cmd.ext)
		err := sh.Run("go", "build", "-o", outputPath, "./cmd/golang-techstack")
		if err != nil {
			return err
		}
	}
	return nil
}

type Release mg.Namespace

func (Release) Full() error {
	mg.Deps(Build.All)
	fmt.Println("Creating full release...")

	// Clean up after release
	defer os.RemoveAll("dist")

	return sh.Run("goreleaser", "release", "--clean")
}
