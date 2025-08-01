package source

import (
	"embed"
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

type Source struct {
	// name -> path
	// prevents us from having to walk Repo if we've
	// already seen a language on a previous call to find
	cache map[string]string
}

func New() *Source {
	return &Source{
		cache: make(map[string]string),
	}
}

func (s *Source) Find(language string) ([]byte, error) {
	var result []byte
	targetName := language + suffix

	// check the cache before walking Repo
	if content, err := s.findCached(targetName); err != nil || content != nil {
		return content, err
	}

	err := fs.WalkDir(Repo, repoRoot, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("error reading %s %w", repoRoot, err)
		}
		currentName := strings.ToLower(d.Name())
		s.putCache(currentName, path)

		if currentName != targetName {
			return nil
		}

		contents, err := Repo.ReadFile(path)
		if err != nil {
			return fmt.Errorf("error reading %s %w", path, err)
		}

		result = contents
		return fs.SkipAll
	})
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, fmt.Errorf("not found: %s", language)
	}

	return result, nil
}

func (s *Source) List(all bool) ([]string, error) {
	var result []string

	err := fs.WalkDir(Repo, repoRoot, func(filename string, d fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("error reading %s %w", repoRoot, err)
		}

		// skip non-root dirs if "all" is false
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

// returns the content at a path if we've already seen it.
// nil content if not found, returns an error if reading Repo fails.
func (s *Source) findCached(name string) (content []byte, err error) {
	path, ok := s.cache[name]
	if !ok {
		return nil, nil
	}

	return Repo.ReadFile(path)
}

func (s *Source) putCache(name string, path string) {
	s.cache[name] = path
}
