package main

import (
	"fmt"
	"os"

	"github.com/rafibayer/ignoregit/internal/cli"
)

func main() {
	if err := cli.Root().Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "ignoregit error: %q\n", err)
		os.Exit(1)
	}
}
