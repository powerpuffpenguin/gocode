package gocode

import (
	"go/ast"
)

type Field struct {
	AST  *ast.Field
	Name string
}

func NewField(name string, filed *ast.Field) *Field {
	return &Field{
		AST:  filed,
		Name: name,
	}
}
func (f *Field) String() string {
	return f.Output(" ")
}
func (f *Field) Output(indent string) string {
	s := f.Name
	if s != `` {
		s += indent
	}
	node := f.AST
	s += NewTypeExpr(node.Type).TypeString()
	return s
}
