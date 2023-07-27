package gocode

import (
	"fmt"
	"go/ast"
	"reflect"
	"strings"
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
	case *ast.StructType:
		// NewStruct(``, t).Output()
		return "struct{}"
	case *ast.Ellipsis:
		return `...` + NewTypeExpr(t.Elt).TypeString()
	case *ast.InterfaceType:
		return `interface{}`
	case *ast.ChanType:
		return NewChanType(t).TypeString()
	case *ast.ParenExpr:
		return `(` + NewTypeExpr(t.X).TypeString() + `)`
	case *ast.IndexListExpr: // 模板
		strs := make([]string, len(t.Indices))
		for i, v := range t.Indices {
			strs[i] = NewTypeExpr(v).TypeString()
		}
		return fmt.Sprintf(`%s[%s]`, NewTypeExpr(t.X).TypeString(), strings.Join(strs, ", "))
	case *ast.BasicLit:
		return t.Value
	case *ast.IndexExpr: // 模板
		return fmt.Sprintf(`%s[%s]`, NewTypeExpr(t.X).TypeString(), NewTypeExpr(t.Index).TypeString())
	default:
		panic(`unknow ` + tag + `: ` + reflect.TypeOf(t).String())
	}
}

// 返回類型字符串
func (e *TypeExpr) TypeString() string {
	return typeString(`field`, e.Expr)
}
