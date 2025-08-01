package cli

import (
	"fmt"
	"os"

	"github.com/rafibayer/ignoregit/internal/util"
	"github.com/rafibayer/ignoregit/source"
	"github.com/spf13/cobra"
)

const (
	defaultOut string = ".gitignore"
)

func newCmd(source *source.Source) *cobra.Command {
	var out string
	var stdout bool

	cmd := &cobra.Command{
		Use:   "new [language...]",
		Short: "create a new .gitignore",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			args = util.Dedup(args)
			all := make([]byte, 0)
			for i, lang := range args {
				content, err := source.Find(lang)
				if err != nil {
					return fmt.Errorf("failed to find .gitignore for arg %d %q: %w", i, lang, err)
				}

				all = append(all, content...)
			}

			if stdout {
				fmt.Println(string(all))
				return nil
			}
			return os.WriteFile(out, all, 0644)
		},
	}

	cmd.Flags().StringVarP(&out, "out", "o", defaultOut, ".gitignore output filename")
	cmd.Flags().BoolVarP(&stdout, "stdout", "s", false, "if true, output to stdout instead of a file")
	cmd.MarkFlagsMutuallyExclusive("out", "stdout")

	return cmd
}
