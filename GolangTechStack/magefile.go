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
type Release mg.Namespace

func (Build) All() error {
	fmt.Println("Building all the targets....")

	if err := os.MkdirAll("build", 0755); err != nil {
		return err
	}

	// commands to build binaries
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
		outputPath := fmt.Sprintf("build/golang-techstack-%s_%s%s", cmd.os, cmd.arch, cmd.ext)
		err := sh.Run("go", "build", "-o", outputPath, "./cmd/golang-techstack")
		if err != nil {
			return err
		}
	}
	return nil
}

func (Release) Full() error {
	fmt.Println("Creating full release...")
	defer os.RemoveAll("build") // Clean up after release
	mg.Deps(Build{}.All)
	return sh.Run("goreleaser", "release", "--clean")
}
