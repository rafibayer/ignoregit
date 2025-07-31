package source

import (
	"embed"
	"errors"
	"fmt"
	"io/fs"
	"strings"
)

const repoRoot string = "repo/gitignore-main"

//go:embed repo/gitignore-main/*
var Repo embed.FS

func Find(language string) ([]byte, error) {
	var result []byte
	target := language + ".gitignore"

	err := fs.WalkDir(Repo, ".", func(path string, d fs.DirEntry, err error) error {
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
		return fs.SkipAll // stop walking if we found our target
	})
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, errors.New("not found")
	}

	return result, nil
}
