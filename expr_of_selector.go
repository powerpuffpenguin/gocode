package gocode

import (
	"go/ast"
)

// 表示一個選擇表達式
type SelectorExpr struct {
	Expr *ast.SelectorExpr
}

func NewSelectorExpr(expr *ast.SelectorExpr) *SelectorExpr {
	return &SelectorExpr{
		Expr: expr,
	}
}

// 返回類型字符串
func (e *SelectorExpr) TypeString() string {
	t := e.Expr
	return t.X.(*ast.Ident).Name + `.` + t.Sel.Name
}
