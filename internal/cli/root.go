package cli

import (
	"github.com/spf13/cobra"
)

func Root() *cobra.Command {
	root := &cobra.Command{
		Use: "ignoregit",
	}

	root.AddCommand(newCmd())
	root.AddCommand(listCmd())

	return root
}
