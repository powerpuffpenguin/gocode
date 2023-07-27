package gocode

import (
	"go/ast"
	"reflect"
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
	switch tx := t.X.(type) {
	case *ast.Ident:
		return tx.Name + `.` + t.Sel.Name
	case *ast.SelectorExpr:
		return NewSelectorExpr(tx).TypeString() + `.` + t.Sel.Name
	default:
		panic(`unknow SelectorExpr.X.type ` + reflect.TypeOf(tx).String())
	}
}
