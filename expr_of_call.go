package gocode

import (
	"go/ast"
	"reflect"
	"strings"
)

// 表示一個函數調用表達式
type CallExpr struct {
	Expr *ast.CallExpr
}

func NewCallExpr(expr *ast.CallExpr) *CallExpr {
	return &CallExpr{
		Expr: expr,
	}
}

// 返回類型字符串
func (c *CallExpr) String() string {
	var fname string
	expr := c.Expr
	switch t := expr.Fun.(type) {
	case *ast.Ident:
		fname = t.Name
	case *ast.SelectorExpr:
		fname = t.X.(*ast.Ident).Name + `.` + t.Sel.Name
	case *ast.ArrayType:
		fname = NewArrayType(t).TypeString()
	case *ast.FuncLit:
		fname = NewFuncType(t.Type).TypeString() + `{...}`
	default:
		panic(`unknow f type` + reflect.TypeOf(t).String())
	}
	count := len(expr.Args)
	strs := make([]string, count)
	for i, arg := range expr.Args {
		strs[i] = NewValueExpr(arg).String()
	}
	return fname + `(` + strings.Join(strs, ",") + `)`
}
