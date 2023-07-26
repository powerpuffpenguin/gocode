package gocode

import (
	"go/ast"
	"reflect"
	"strings"
)

// 表示一個值
type ValueExpr struct {
	Expr ast.Expr
}

func NewValueExpr(expr ast.Expr) *ValueExpr {
	return &ValueExpr{
		Expr: expr,
	}
}
func (v *ValueExpr) String() string {
	switch t := v.Expr.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.SelectorExpr:
		return NewSelectorExpr(t).TypeString()
	case *ast.CallExpr:
		return v.call(t.Fun, t.Args)
	case *ast.BasicLit:
		return t.Value
	case *ast.BinaryExpr:
		//  iota
		return NewBinaryExpr(t).String()
	case *ast.UnaryExpr:
		return t.Op.String() + NewValueExpr(t.X).String()
	case *ast.CompositeLit:
		return NewTypeExpr(t.Type).TypeString() + `{}`
	default:
		panic(`unknow value type` + reflect.TypeOf(t).String())
	}
}
func (v *ValueExpr) call(f ast.Expr, args []ast.Expr) string {
	var fname string
	switch t := f.(type) {
	case *ast.Ident:
		fname = t.Name
	case *ast.SelectorExpr:
		fname = t.X.(*ast.Ident).Name + `.` + t.Sel.Name
	default:
		panic(`unknow f type` + reflect.TypeOf(t).String())
	}
	count := len(args)
	strs := make([]string, count)
	for i, arg := range args {
		strs[i] = NewValueExpr(arg).String()
	}
	return fname + `(` + strings.Join(strs, ",") + `)`
}
