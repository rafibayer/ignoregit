package cli

import (
	"os"

	"github.com/rafibayer/ignoregit/source"
	"github.com/spf13/cobra"
)

const (
	defaultOut string = ".gitignore"
)

func newCmd() *cobra.Command {
	var out string

	cmd := &cobra.Command{
		Use:  "new [language]",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			language := args[0]
			content, err := source.Find(language)
			if err != nil {
				return err
			}

			return os.WriteFile(out, content, 0644)
		},
	}

	cmd.Flags().StringVarP(&out, "out", "o", defaultOut, ".gitignore output filename")

	return cmd
}
