package main

import (
	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use:   "gocode",
		Short: "analyze go code",
	}

	cmd.AddCommand(
		fileCommand(),
		dirCommand(),
	)

	cmd.Execute()
}
