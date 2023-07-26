package gocode

import (
	"bytes"
	"go/ast"
	"io"
	"sort"
	"strings"
)

// 包信息
type Package struct {
	AST *ast.Package
	// 檔案
	Files []*File
}

func NewPackage(p *ast.Package) *Package {
	var (
		files []*File
		size  = len(p.Files)
	)
	if size > 0 {
		files = make([]*File, 0, size)
		for name, file := range p.Files {
			files = append(files, NewFile(name, file))
		}
		if size > 1 {
			sort.Slice(files, func(i, j int) bool {
				return strings.Compare(files[i].Name, files[j].Name) < 0
			})
		}
	}
	return &Package{
		AST:   p,
		Files: files,
	}
}

// 返回包名
func (p *Package) Name() string {
	return p.AST.Name
}

// 返回人類友好的字符串
func (p *Package) String() string {
	w := bytes.NewBuffer(make([]byte, 0, 1024))
	_, e := p.Output(w, ` - `, `  `)
	if e != nil {
		panic(e)
	}
	b := w.Bytes()
	return BytesToString(b)
}

// 將人類友好字符串寫入到 writer
func (p *Package) Output(writer io.Writer, prefix, indent string) (n int64, e error) {
	w := writerTo{w: writer}
	_, e = w.WriteString(prefix + `package ` + p.Name() + "\n")
	if e != nil {
		n = w.n
		return
	}
	prefix += indent
	for _, f := range p.Files {
		_, e = f.Output(&w, prefix, indent)
		if e != nil {
			n = w.n
			return
		}
	}
	n = w.n
	return
}

type writerTo struct {
	w io.Writer
	n int64
}

func (w *writerTo) WriteString(s string) (n int, e error) {
	n, e = w.w.Write(StringToBytes(s))
	w.n += int64(n)
	return
}
func (w *writerTo) Write(b []byte) (n int, e error) {
	n, e = w.w.Write(b)
	w.n += int64(n)
	return
}
