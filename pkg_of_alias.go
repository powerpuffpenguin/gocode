package gocode

import (
	"go/ast"
)

type Alias struct {
	Expr ast.Expr
	Name string
}

func NewAlias(name string, expr ast.Expr) *Alias {
	return &Alias{
		Expr: expr,
		Name: name,
	}
}
func (a *Alias) String() string {
	return `type ` + a.Name + ` ` + NewTypeExpr(a.Expr).TypeString()
}
