package cli

import (
	"github.com/rafibayer/ignoregit/source"
	"github.com/spf13/cobra"
)

func Root() *cobra.Command {
	root := &cobra.Command{
		Use: "ignoregit",
	}

	source := source.New()

	root.AddCommand(newCmd(source))
	root.AddCommand(listCmd(source))

	return root
}
