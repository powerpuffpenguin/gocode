package gocode

import (
	"go/ast"
)

// 表示了一個切片
type ArrayType struct {
	Type *ast.ArrayType
}

func NewArrayType(tp *ast.ArrayType) *ArrayType {
	return &ArrayType{Type: tp}
}
func (t *ArrayType) TypeString() string {
	return `[]` + typeString(`ArrayType.Elt.(type)`, t.Type.Elt)
}
