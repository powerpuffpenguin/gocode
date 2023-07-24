package gocode

import (
	"go/ast"
)

// 表示了一個hash表
type MapType struct {
	Type *ast.MapType
}

func NewMapType(tp *ast.MapType) *MapType {
	return &MapType{Type: tp}
}
func (t *MapType) TypeString() string {
	key := NewTypeExpr(t.Type.Key).TypeString()
	value := NewTypeExpr(t.Type.Value).TypeString()
	return `map[` + key + `]` + value
}
