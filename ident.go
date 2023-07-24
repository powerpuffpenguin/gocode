package gocode

import (
	"go/ast"
)

// 表示了一個元素的數據類型
type Ident struct {
	Ident *ast.Ident
}

func NewIdent(ident *ast.Ident) *Ident {
	return &Ident{Ident: ident}
}
func (t *Ident) TypeString() string {
	return t.Ident.Name
}
