package cli

import (
	"fmt"

	"github.com/rafibayer/ignoregit/source"
	"github.com/spf13/cobra"
)

func listCmd() *cobra.Command {
	var all bool

	cmd := &cobra.Command{
		Use:   "list",
		Short: "list available .gitignore templates",
		RunE: func(cmd *cobra.Command, args []string) error {
			langs, err := source.List(all)
			if err != nil {
				return err
			}

			for _, lang := range langs {
				fmt.Println(lang)
			}

			return nil
		},
	}

	cmd.Flags().BoolVarP(&all, "all", "a", false, "include community and global .gitignores")

	return cmd
}
