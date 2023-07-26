package gocode

import (
	"go/ast"
	"strings"
)

// 表示了一個函數
type FuncType struct {
	Type *ast.FuncType
}

func NewFuncType(tp *ast.FuncType) *FuncType {
	return &FuncType{Type: tp}
}
func (t *FuncType) TypeString() string {
	return `func` + t.ParamsAndResults()
}

// 返回參數和返回值定義
func (t *FuncType) ParamsAndResults() string {
	var params []string
	if t.Type.Params != nil {
		for _, f := range t.Type.Params.List {
			if len(f.Names) == 0 { // 匿名參數
				params = append(params, NewField(``, f).String())
			} else {
				for _, name := range f.Names {
					params = append(params, NewField(name.Name, f).String())
				}
			}
		}
	}
	var results []string
	var named = false
	if t.Type.Results != nil {
		for _, f := range t.Type.Results.List {
			if len(f.Names) == 0 { // 匿名返回值
				results = append(results, NewField(``, f).String())
			} else {
				named = true
				for _, name := range f.Names {
					results = append(results, NewField(name.Name, f).String())
				}
			}
		}
	}
	s := `(` + strings.Join(params, `, `) + `)`
	switch len(results) {
	case 0:
	case 1:
		if named {
			s += ` (` + results[0] + `)`
		} else {
			s += ` ` + results[0]
		}
	default:
		s += ` (` + strings.Join(results, `, `) + `)`
	}
	return s
}
