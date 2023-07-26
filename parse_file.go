package gocode

import (
	"go/parser"
	"go/token"
)

func ParseFile(path string, src any, mode parser.Mode) (file *File, e error) {
	fset := token.NewFileSet()
	f, e := parser.ParseFile(fset, path, src, mode)
	if e != nil {
		return
	}
	file = NewFile(f.Name.Name, f)
	return
}
