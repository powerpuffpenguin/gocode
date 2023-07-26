package gocode

import (
	"go/ast"
)

type Func struct {
	Expr *ast.FuncDecl
}

func NewFunc(expr *ast.FuncDecl) *Func {
	return &Func{
		Expr: expr,
	}
}
func (f *Func) String() string {
	var (
		expr = f.Expr
		decl string
	)
	if expr.Recv == nil {
		decl = `func `
	} else {
		field := expr.Recv.List[0]
		decl = `func (` + NewField(``, field).String() + `) `
	}
	return decl + expr.Name.Name + NewFuncType(expr.Type).ParamsAndResults()

}
