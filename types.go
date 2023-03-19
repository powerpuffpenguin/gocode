package gocode

import (
	"errors"
	"go/ast"
	"reflect"
	"strconv"
)

type Type interface {
	Expr() ast.Expr
	X() ast.Expr
	// return type import from "net" "http" "" ...
	Import() string
	// return type name: int uint16 ...
	Name() string
	// is it a pointer type
	Star() bool
	// if is [] return item type, else return nil
	Array() Type
	// if is chan return item type, else return nil
	Chan() Type

	// reutrn type decl: int *int bytes.Buffer *bytes.Buffer
	String() string
}

func typeString(t Type) string {
	var s string
	if t.Star() {
		if ch, ok := t.X().(*ast.ChanType); ok {
			ty, e := NewType(ch)
			if e != nil {
				panic(e)
			}
			s = `*` + typeString(ty)
			return s
		} else {
			s = `*`
		}
	}
	if it := t.Array(); it != nil {
		s += `[]` + typeString(it)
	} else if ch := t.Chan(); ch != nil {
		ty := t.Expr().(*ast.ChanType)
		send := (ty.Dir & ast.SEND) != 0
		recv := (ty.Dir & ast.RECV) != 0
		if send && recv {
			s += `chan ` + typeString(ch)
		} else if send {
			s += `chan<- ` + typeString(ch)
		} else if recv {
			s += `<-chan ` + typeString(ch)
		} else {
			panic(`unknow chan dir: ` + strconv.Itoa(int(ty.Dir)))
		}
	} else {
		imt := t.Import()
		if imt == `` {
			s += t.Name()
		} else {
			s += imt + `.` + t.Name()
		}
	}
	return s
}
func NewType(expr ast.Expr) (ty Type, e error) {
	switch t := expr.(type) {
	case *ast.Ident:
		ty = NewIdentType(t)
	case *ast.StarExpr:
		ty = NewStarType(t)
	case *ast.SelectorExpr:
		ty = NewSelectorType(t)
	case *ast.ArrayType:
		ty = NewArrayType(t)
	case *ast.InterfaceType:
		ty = NewInterfaceType(t)
	case *ast.ChanType:
		ty = NewChanType(t)
	default:
		e = errors.New(`NewType unknow expr: ` + reflect.TypeOf(expr).String())
	}
	return
}

type exprType struct {
	expr ast.Expr
	x    ast.Expr
	star bool
	ch   bool
}

func (t *exprType) Expr() ast.Expr {
	return t.expr
}
func (t *exprType) X() ast.Expr {
	return t.x
}
func (t *exprType) Import() string {
	if t.x == nil {
		return ``
	}
	ty, e := NewType(t.x)
	if e != nil {
		panic(e)
	}
	return ty.Import()
}
func (t *exprType) Name() string {
	if t.x == nil {
		return ``
	}
	ty, e := NewType(t.x)
	if e != nil {
		panic(e)
	}
	return ty.Name()
}
func (t *exprType) Star() bool {
	return t.star
}
func (t *exprType) Array() Type {
	if t.x == nil {
		return nil
	}
	ty, e := NewType(t.x)
	if e != nil {
		panic(e)
	}
	return ty.Array()
}
func (t *exprType) Chan() Type {
	if t.ch {
		ty, e := NewType(t.x)
		if e != nil {
			panic(e)
		}
		return ty
	}
	return nil
}
func (t *exprType) String() string {
	return typeString(t)
}

type IdentType struct {
	exprType
}

func NewIdentType(expr *ast.Ident) *IdentType {
	return &IdentType{
		exprType: exprType{
			expr: expr,
		},
	}
}
func (t *IdentType) AST() *ast.Ident {
	return t.expr.(*ast.Ident)
}
func (t *IdentType) Name() string {
	return t.AST().Name
}
func (t *IdentType) String() string {
	return typeString(t)
}

type StarType struct {
	exprType
}

func NewStarType(expr *ast.StarExpr) *StarType {
	return &StarType{
		exprType: exprType{
			expr: expr,
			x:    expr.X,
			star: true,
		},
	}
}
func (t *StarType) AST() *ast.StarExpr {
	return t.expr.(*ast.StarExpr)
}

type SelectorType struct {
	exprType
}

func NewSelectorType(expr *ast.SelectorExpr) *SelectorType {
	return &SelectorType{
		exprType: exprType{
			expr: expr,
			x:    expr.X,
		},
	}
}
func (t *SelectorType) AST() *ast.SelectorExpr {
	return t.expr.(*ast.SelectorExpr)
}
func (t *SelectorType) Import() string {
	return t.x.(*ast.Ident).Name
}
func (t *SelectorType) Name() string {
	return t.expr.(*ast.SelectorExpr).Sel.Name
}
func (t *SelectorType) String() string {
	return typeString(t)
}

type ArrayType struct {
	exprType
}

func NewArrayType(expr *ast.ArrayType) *ArrayType {
	return &ArrayType{
		exprType: exprType{
			expr: expr,
			x:    expr.Elt,
		},
	}
}
func (t *ArrayType) AST() *ast.ArrayType {
	return t.expr.(*ast.ArrayType)
}
func (t *ArrayType) Array() Type {
	ty, e := NewType(t.x)
	if e != nil {
		panic(`ArrayType Array error: ` + e.Error())
	}
	return ty
}
func (t *ArrayType) String() string {
	return typeString(t)
}

type InterfaceType struct {
	exprType
}

func NewInterfaceType(expr *ast.InterfaceType) *InterfaceType {
	return &InterfaceType{
		exprType: exprType{
			expr: expr,
		},
	}
}
func (t *InterfaceType) AST() *ast.InterfaceType {
	return t.expr.(*ast.InterfaceType)
}
func (t *InterfaceType) Name() string {
	return `interface{}`
}
func (t *InterfaceType) String() string {
	return typeString(t)
}

type ChanType struct {
	exprType
}

func NewChanType(expr *ast.ChanType) *ChanType {
	return &ChanType{
		exprType: exprType{
			expr: expr,
			x:    expr.Value,
			ch:   true,
		},
	}
}
func (t *ChanType) AST() *ast.ChanType {
	return t.expr.(*ast.ChanType)
}
