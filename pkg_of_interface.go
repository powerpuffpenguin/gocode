package gocode

import (
	"bytes"
	"go/ast"
	"io"
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
	_, e = w.WriteString(prefix + `type ` + it.Name + " interface {\n")
	if e != nil {
		n = w.n
		return
	}
	p := prefix
	prefix += indent
	var field *Field
	for _, node := range it.AST.Methods.List {
		if len(node.Names) == 0 { // 組合的匿名字段
			field = NewField(``, node)
		} else {
			field = NewField(node.Names[0].Name, node)
		}
		_, e = w.WriteString(prefix + field.OutputInterface() + "\n")
		if e != nil {
			return
		}
	}
	_, e = w.WriteString(p + "}\n")
	n = w.n
	return
}
