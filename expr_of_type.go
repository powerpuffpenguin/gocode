package gocode

import (
	"go/ast"
	"reflect"
)

// 表示一個型別
type TypeExpr struct {
	Expr ast.Expr
}

func NewTypeExpr(expr ast.Expr) *TypeExpr {
	return &TypeExpr{
		Expr: expr,
	}
}

func typeString(tag string, expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return NewIdent(t).TypeString()
	case *ast.SelectorExpr:
		return NewSelectorExpr(t).TypeString()
	case *ast.StarExpr:
		return NewStarExpr(t).TypeString()
	case *ast.ArrayType:
		return NewArrayType(t).TypeString()
	case *ast.MapType:
		return NewMapType(t).TypeString()
	case *ast.FuncType:
		return NewFuncType(t).TypeString()
	default:
		panic(`unknow ` + tag + `: ` + reflect.TypeOf(t).String())
	}
}

// 返回類型字符串
func (e *TypeExpr) TypeString() string {
	switch t := e.Expr.(type) {
	case *ast.Ident: // 標識，基礎類型
		return NewIdent(t).TypeString()
	case *ast.SelectorExpr:
		return NewSelectorExpr(t).TypeString()
	case *ast.StarExpr: // 指針
		return NewStarExpr(t).TypeString()
	case *ast.ArrayType: // 切片
		return NewArrayType(t).TypeString()
	case *ast.MapType: // hash 表
		return NewMapType(t).TypeString()
	case *ast.FuncType: // 函數
		return NewFuncType(t).TypeString()
	default:
		panic(`unknow field t type: ` + reflect.TypeOf(t).String())
	}
}
