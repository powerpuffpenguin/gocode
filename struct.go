package gocode

import (
	"bytes"
	"go/ast"
	"io"
)

type Struct struct {
	AST    *ast.StructType
	Name   string
	Fields []*Field
}

func NewStruct(name string, st *ast.StructType) *Struct {
	var (
		size   = len(st.Fields.List)
		fields []*Field
	)
	if size > 0 {
		fields = make([]*Field, 0, size)
	}
	for _, f := range st.Fields.List {
		if len(f.Names) == 0 { // 組合的匿名字段
			fields = append(fields, NewField(``, f))
		} else {
			for _, name := range f.Names {
				fields = append(fields, NewField(name.Name, f))
			}
		}
	}
	return &Struct{
		AST:    st,
		Name:   name,
		Fields: fields,
	}
}
func (s *Struct) String() string {
	w := bytes.NewBuffer(make([]byte, 0, 1024))
	_, e := s.Output(w, ``, ``)
	if e != nil {
		panic(e)
	}
	b := w.Bytes()
	return BytesToString(b)
}
func (s *Struct) Output(writer io.Writer, prefix, indent string) (n int64, e error) {
	w := writerTo{w: writer}
	_, e = w.WriterString(prefix + `type ` + s.Name + " struct {\n")
	if e != nil {
		n = w.n
		return
	}
	p := prefix
	prefix += indent
	for _, node := range s.Fields {
		w.WriterString(prefix + node.Output("\t") + "\n")
	}
	_, e = w.WriterString(p + "}\n")
	n = w.n
	return
}
