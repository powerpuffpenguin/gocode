package main

import (
	"fmt"
	"io/fs"
	"log"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/powerpuffpenguin/gocode"
	"github.com/spf13/cobra"
)

func parseDir(keys map[string]bool, keysRule map[string]*regexp.Regexp, path string, test, all bool) {
	pkgs, e := gocode.ParseDir(path, func(fi fs.FileInfo) bool {
		name := fi.Name()
		path := filepath.Join(path, name)
		if keys[path] {
			return false
		}
		for _, rule := range keysRule {
			if rule.MatchString(path) {
				return false
			}
		}
		if !test && strings.HasSuffix(name, "test.go") {
			return false
		}
		return true
	}, 0)
	if e != nil {
		log.Fatalln(e)
	}

	for _, p := range pkgs {
		fmt.Println(p.String())
	}
}
func dirCommand() *cobra.Command {
	var (
		all                bool
		test               bool
		recursion          bool
		exclude0, exclude1 []string
	)
	cmd := &cobra.Command{
		Use:   "dir",
		Short: "Analyze go code dir",
		Long: `Analyze go code dir, and print code statement

gocode dir /opt/google/go/src`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				return
			}
			keys := make(map[string]bool)
			for _, s := range exclude0 {
				keys[s] = true
			}
			keysRule := make(map[string]*regexp.Regexp)
			for _, s := range exclude1 {
				keysRule[s] = regexp.MustCompile(s)
			}
			for _, path := range args {
				if recursion {
					filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
						if !d.IsDir() {
							return nil
						}
						parseDir(keys, keysRule, path, test, all)
						return nil
					})
				} else {
					parseDir(keys, keysRule, path, test, all)
				}
			}
		},
	}
	flags := cmd.Flags()
	flags.BoolVarP(&all, `all`, `a`, false, `Print all declarations, by default only the exported content`)
	flags.BoolVarP(&test, `test`, `t`, false, `Analyze *test.go`)
	flags.BoolVarP(&recursion, `recursion`, `r`, false, `Recursively analyze subfolders`)
	flags.StringSliceVarP(&exclude0, "exclude", "E", nil, `Exclude these files`)
	flags.StringSliceVarP(&exclude1, "regexp", "R", nil, `Exclude these files using the regular expression match`)
	return cmd
}
