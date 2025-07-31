package main

import (
	"fmt"
	"os"

	"github.com/rafibayer/ignoregit/source"
	"github.com/spf13/cobra"
)

const (
	defaultSrc string = "https://raw.githubusercontent.com/github/gitignore/refs/heads/main/"
	defaultOut string = ".gitignore"
)

type flags struct {
	src string
	out string
}

func main() {
	flags := &flags{}
	root := &cobra.Command{
		Use:  "ignoregit [language]",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			language := args[0]

			content, err := source.Find(language)
			if err != nil {
				return err
			}

			return os.WriteFile(flags.out, content, 0644)
		},
	}

	root.Flags().StringVarP(&flags.src, "source", "s", defaultSrc, ".gitignore source url")
	root.Flags().StringVarP(&flags.out, "out", "o", defaultOut, ".gitignore output filename")

	if err := root.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "ignoregit error: %q\n", err)
		os.Exit(1)
	}
}
