package main

import (
	"fmt"
	"log"
	"os"

	"github.com/powerpuffpenguin/gocode"
	"github.com/spf13/cobra"
)

func fileCommand() *cobra.Command {
	var (
		all bool
	)
	cmd := &cobra.Command{
		Use:   "file",
		Short: "Analyze go code file",
		Long: `Analyze go code file, and print code statement

gocode file a.go b.go`,
		Run: func(cmd *cobra.Command, args []string) {
			for i, path := range args {
				if i != 0 {
					fmt.Println()
				}

				f, e := gocode.ParseFile(path, nil, 0)
				if e != nil {
					log.Fatalln(e)
				}
				_, e = f.Output(os.Stdout, "- ", "  ", all)
				if e != nil {
					log.Fatalln(e)
				}
			}
		},
	}
	flags := cmd.Flags()
	flags.BoolVarP(&all, `all`, `a`, false, `Print all declarations, by default only the exported content`)

	return cmd
}
