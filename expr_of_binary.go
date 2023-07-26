package gocode

import (
	"fmt"
	"go/ast"
	"reflect"
)

// 表示一個值
type BinaryExpr struct {
	Expr *ast.BinaryExpr
}

func NewBinaryExpr(expr *ast.BinaryExpr) *BinaryExpr {
	return &BinaryExpr{
		Expr: expr,
	}
}
func (b *BinaryExpr) String() string {
	var (
		exprs = []ast.Expr{b.Expr.X, b.Expr.Y}
		strs  = []string{"", ""}
		s     string
	)
	for i, expr := range exprs {
		switch t := expr.(type) {
		case *ast.Ident:
			s = t.Name
		case *ast.BasicLit:
			s = t.Value
		case *ast.BinaryExpr:
			s = NewBinaryExpr(t).String()
		case *ast.UnaryExpr:
			s = t.Op.String() + NewValueExpr(t.X).String()
		default:
			panic(`unknow x type` + reflect.TypeOf(t).String())
		}
		strs[i] = s
	}
	return fmt.Sprintf(`%s %s %s`, strs[0], b.Expr.Op, strs[1])
}
