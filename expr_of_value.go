package gocode

import (
	"fmt"
	"go/ast"
	"reflect"
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
		return NewCallExpr(t).String()
	case *ast.BasicLit:
		return t.Value
	case *ast.BinaryExpr:
		//  iota
		return NewBinaryExpr(t).String()
	case *ast.UnaryExpr:
		return t.Op.String() + NewValueExpr(t.X).String()
	case *ast.CompositeLit:
		return NewTypeExpr(t.Type).TypeString() + `{}`
	case *ast.ParenExpr:
		return `(` + NewValueExpr(t.X).String() + `)`
	case *ast.IndexExpr: // 模板
		return fmt.Sprintf(`%s[%s]`, NewTypeExpr(t.X).TypeString(), NewTypeExpr(t.Index).TypeString())
	case *ast.FuncLit:
		return NewFuncType(t.Type).TypeString()
	case *ast.ArrayType:
		return NewArrayType(t).TypeString()
	case *ast.MapType:
		return NewMapType(t).TypeString()
	case *ast.ChanType:
		return NewChanType(t).TypeString()
	default:
		panic(`unknow value type` + reflect.TypeOf(t).String())
	}
}
