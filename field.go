package gocode

import (
	"go/ast"
	"reflect"
	"strings"
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
func (f *Field) IsExport() bool {
	if f.Name == `` {
		s := NewTypeExpr(f.AST.Type).TypeString()
		s = strings.TrimLeft(s, `*`)
		return IsExport(s)
	}
	return IsExport(f.Name)
}

// 輸出字段定義
func (f *Field) Output(indent string) string {
	s := f.Name
	if s != `` {
		s += indent
	}
	node := f.AST
	s += NewTypeExpr(node.Type).TypeString()
	return s
}

// 輸出接口定義
func (f *Field) OutputInterface() string {
	node := f.AST
	switch t := node.Type.(type) {
	case *ast.Ident: // 標識，基礎類型
		return NewIdent(t).TypeString()
	case *ast.SelectorExpr:
		return NewSelectorExpr(t).TypeString()
	case *ast.FuncType: // 函數
		return f.Name + NewFuncType(t).ParamsAndResults()
	default:
		panic(`unknow field t type: ` + reflect.TypeOf(t).String())
	}
}
