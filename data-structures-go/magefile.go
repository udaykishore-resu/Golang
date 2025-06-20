package main

import (
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

var Aliases = map[string]any{
	"build": Build.All,
	"test":  Test.All,
}

type Build mg.Namespace

func (Build) All() error {
	mg.Deps(Build.Binary, Build.Docker)
	return nil
}

func (Build) Binary() error {
	return sh.RunV("go", "build", "-o", "bin/ds-tool", "./cmd/ds-tool")
}

func (Build) Docker() error {
	return sh.RunV("docker", "build", "-t", "ds-tool:latest", ".")
}

type Test mg.Namespace

func (Test) All() error {
	mg.Deps(Test.Unit, Test.Integration)
	return nil
}

func (Test) Unit() error {
	return sh.RunV("go", "test", "-cover", "./tests/unit/...")
}

func (Test) Integration() error {
	return sh.RunV("go", "test", "./tests/integration/...")
}

func (Test) Coverage() error {
	return sh.RunV("go", "test", "-coverprofile=coverage.out", "./...")
}

type Helm mg.Namespace

func (Helm) Package() error {
	return sh.RunV("helm", "package", "./charts/data-structures")
}

func (Helm) Lint() error {
	return sh.RunV("helm", "lint", "./charts/data-structures")
}

func Clean() error {
	sh.RunV("docker", "system", "prune", "-f")
	return sh.RunV("go", "clean", "-cache")
}
