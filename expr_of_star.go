package gocode

import (
	"go/ast"
)

// 表示一個指針表達式
type StarExpr struct {
	Expr *ast.StarExpr
}

func NewStarExpr(expr *ast.StarExpr) *StarExpr {
	return &StarExpr{
		Expr: expr,
	}
}

// 返回類型字符串
func (e *StarExpr) TypeString() string {
	return `*` + typeString("StarExpr.X.(type)", e.Expr.X)
}
