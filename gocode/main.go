package main

import (
	"fmt"
	"log"
	"os"

	"github.com/powerpuffpenguin/gocode"
	"github.com/spf13/cobra"
)

func main() {
	var (
		all bool
	)
	var rootCmd = &cobra.Command{
		Use:   "gocode",
		Short: "analyze go code",
		Long: `Analyze go code, and print code statement

gocode a.go b.go`,
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
	flags := rootCmd.Flags()
	flags.BoolVarP(&all, `all`, `a`, false, `Print all declarations, by default only the exported content`)

	rootCmd.Execute()
}
