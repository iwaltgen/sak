//go:build mage

package main

import (
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/iwaltgen/magex/git"
	"github.com/iwaltgen/magex/script"
	"github.com/iwaltgen/magex/semver"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

const (
	packageName = "github.com/iwaltgen/sak"
	targetDir   = "build"
)

type BUILD mg.Namespace

var (
	started   int64
	goCmd     string
	workspace string
)

func init() {
	started = time.Now().Unix()
	goCmd = mg.GoCmd()
	workspace, _ = os.Getwd()
}

// Run test cases
func Test() error {
	mg.Deps(Lint)

	return script.ExecStdout(
		"go test ./... -timeout 10s -cover -json",
		"tparse -all",
	)
}

// Run lint
func Lint() error {
	return sh.RunV("golangci-lint", "run")
}

// Remove all artifacts
func Clean() error {
	if err := sh.Rm(targetDir); err != nil {
		return fmt.Errorf("failed to remove dir `%s`: %w", targetDir, err)
	}
	return nil
}

// Build platform artifacts
func Build() {
	mg.Deps(Lint, BUILD.Platform)
}

// Build platform artifacts
func (ns BUILD) Platform() error {
	if err := os.MkdirAll(targetDir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to make dir `%s`: %w", targetDir, err)
	}

	if err := sh.RunWith(ns.buildEnv(), goCmd, ns.buildParameters()...); err != nil {
		return fmt.Errorf("failed to build: %w", err)
	}
	return nil
}

func (BUILD) buildParameters() []string {
	ldflags := "-ldflags=" +
		"-X $PACKAGE/internal.version=$VERSION " +
		"-X $PACKAGE/internal.commitHash=$COMMIT_HASH " +
		"-X $PACKAGE/internal.buildDate=$BUILD_DATE"

	var tags string
	if gotags := os.Getenv("GITHUB_GO_BUILD_TAGS"); gotags != "" {
		tags = gotags
	}

	return []string{"build", "-trimpath", "-tags", tags, ldflags, "-o", targetDir, "./cmd/..."}
}

func (ns BUILD) buildEnv() map[string]string {
	return map[string]string{
		"CGO_ENABLED": "0",
		"PACKAGE":     packageName,
		"WORKSPACE":   workspace,
		"VERSION":     ns.mustCurrentVersion(),
		"COMMIT_HASH": ns.commitHash(),
		"BUILD_DATE":  fmt.Sprintf("%d", started),
	}
}

func (BUILD) commitHash() string {
	if gitsha := os.Getenv("GITHUB_SHA"); gitsha != "" {
		return gitsha
	}

	repo, err := git.NewRepository(".")
	if err != nil {
		panic(err)
	}

	ref, err := repo.Head()
	if err != nil {
		panic(err)
	}

	return ref.Hash().String()
}

func (BUILD) mustCurrentVersion() string {
	if tag := os.Getenv("GITHUB_TAG"); tag != "" {
		return tag
	}

	version, err := semver.LatestTag(".")
	if err != nil {
		panic(err)
	}
	return version
}

// Show current version
func Version() error {
	version, err := semver.LatestTag(".")
	if err != nil {
		return err
	}

	color.Green(version)
	return nil
}

// Release tag version [major, minor, patch]
func Release(div string) error {
	cv, err := semver.LatestTag(".")
	if err != nil {
		return err
	}

	nv, err := semver.Bump(cv, semver.ParseBumpType(div))
	if err != nil {
		return err
	}

	err = git.CreateTag(nv,
		git.WithCreateTagMessage("release "+nv),
		git.WithCreateTagPushProgress(os.Stdout),
	)
	if err == nil {
		color.Green(nv)
	}
	return err
}
