package gocode

import (
	"go/ast"
)

// 表示一個chan
type ChanType struct {
	Type *ast.ChanType
}

func NewChanType(tp *ast.ChanType) *ChanType {
	return &ChanType{
		Type: tp,
	}
}

// 返回類型字符串
func (e *ChanType) TypeString() string {
	var s string
	if e.Type.Dir&ast.SEND != 0 {
		if e.Type.Dir&ast.RECV != 0 {
			s = `chan `
		} else {
			s = `chan<- `
		}
	} else {
		s = `<-chan `
	}

	return s + NewTypeExpr(e.Type.Value).TypeString()
}
