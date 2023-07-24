package gocode

import (
	"go/parser"
	"go/token"
	"io/fs"
)

func ParseDir(path string, filter func(fs.FileInfo) bool, mode parser.Mode) (pkgs map[string]*Package, e error) {
	fset := token.NewFileSet()
	keys, e := parser.ParseDir(fset, path, filter, mode)
	if e != nil || len(keys) == 0 {
		return
	}
	pkgs = make(map[string]*Package, len(keys))
	for k, v := range keys {
		pkgs[k] = NewPackage(v)
	}
	return
}
