package gocode

import (
	"bytes"
	"go/ast"
	"io"
	"reflect"
	"strings"
)

type Interface struct {
	AST  *ast.InterfaceType
	Name string
}

func NewInterface(name string, it *ast.InterfaceType) *Interface {
	return &Interface{
		AST:  it,
		Name: name,
	}
}
func (it *Interface) String() string {
	w := bytes.NewBuffer(make([]byte, 0, 1024))
	_, e := it.Output(w, ``, ``)
	if e != nil {
		panic(e)
	}
	b := w.Bytes()
	return BytesToString(b)
}
func (it *Interface) Output(writer io.Writer, prefix, indent string) (n int64, e error) {
	w := writerTo{w: writer}
	_, e = w.WriterString(prefix + `type ` + it.Name + " interface {\n")
	if e != nil {
		n = w.n
		return
	}
	p := prefix
	prefix += indent
	for _, node := range it.AST.Methods.List {
		var name string
		if len(node.Names) != 0 {
			name = node.Names[0].Name
		}
		switch t := node.Type.(type) {
		case *ast.SelectorExpr:
			_, e = w.WriterString(prefix +
				t.X.(*ast.Ident).Name + `.` + t.Sel.Name +
				"\n",
			)
			if e != nil {
				n = w.n
				return
			}
		case *ast.FuncType:
			var s string
			var params []string
			if t.Params != nil {
				for _, f := range t.Params.List {
					if len(f.Names) == 0 { // 匿名參數
						params = append(params, NewField(``, f).String())
					} else {
						for _, name := range f.Names {
							params = append(params, NewField(name.Name, f).String())
						}
					}
				}
			}

			// var results []string
			// if t.Results != nil {
			// 	for _, f := range t.Results.List {
			// 		if len(f.Names) == 0 { // 匿名返回值
			// 			results = append(results, NewField(``, f).String())
			// 		} else {
			// 			for _, name := range f.Names {
			// 				results = append(results, NewField(name.Name, f).String())
			// 			}
			// 		}
			// 	}
			// }
			w.WriterString(prefix + name +
				s + `(` + strings.Join(params, `, `) + `)` +
				"\n",
			)
		default:
			panic(`unknow Interface type: ` + reflect.TypeOf(t).String())
		}
	}
	_, e = w.WriterString(p + "}\n")
	n = w.n
	return
}
