package main

import (
	"fmt"
	"io/fs"
	"log"
	"strings"

	"github.com/powerpuffpenguin/gocode"
)

func main() {
	pkgs, e := gocode.ParseDir(`demo`, func(fi fs.FileInfo) bool {
		return !strings.HasSuffix(fi.Name(), `_test.go`)
	}, 0)
	if e != nil {
		log.Fatalln(e)
	}

	for _, p := range pkgs {
		fmt.Println(p.String())
	}
}
