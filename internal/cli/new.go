package cli

import (
	"fmt"
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
		Use:   "new [language...]",
		Short: "create a new .gitignore",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			all := make([]byte, 0)
			for i, lang := range args {
				content, err := source.Find(lang)
				if err != nil {
					return fmt.Errorf("failed to find .gitignore for arg %d %q: %w", i, lang, err)
				}

				all = append(all, content...)
			}

			return os.WriteFile(out, all, 0644)
		},
	}

	cmd.Flags().StringVarP(&out, "out", "o", defaultOut, ".gitignore output filename")

	return cmd
}
