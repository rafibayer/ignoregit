package source

import (
	"embed"
	"errors"
	"fmt"
	"io/fs"
	"strings"
)

const (
	repoRoot string = "repo/gitignore-main"
	suffix   string = ".gitignore"
)

//go:embed repo/gitignore-main/*
var Repo embed.FS

func Find(language string) ([]byte, error) {
	var result []byte
	target := language + suffix

	err := fs.WalkDir(Repo, repoRoot, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("error reading %s %w", repoRoot, err)
		}
		name := strings.ToLower(d.Name())
		if name != target {
			return nil
		}

		contents, err := Repo.ReadFile(path)
		if err != nil {
			return fmt.Errorf("error reading %s %w", path, err)
		}

		result = contents

		// stop walking if we found our target
		return fs.SkipAll
	})
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, errors.New("not found")
	}

	return result, nil
}

func List(all bool) ([]string, error) {
	var result []string

	err := fs.WalkDir(Repo, repoRoot, func(filename string, d fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("error reading %s %w", repoRoot, err)
		}

		// skip non-root dirs if all is false
		if !all && d.IsDir() && filename != repoRoot {
			return fs.SkipDir
		}

		if !d.IsDir() && strings.HasSuffix(d.Name(), suffix) {
			clean := strings.ToLower(strings.TrimSuffix(d.Name(), suffix))
			result = append(result, clean)
		}
		return nil
	})
	return result, err
}
