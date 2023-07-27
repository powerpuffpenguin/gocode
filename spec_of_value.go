package gocode

import (
	"bytes"
	"go/ast"
)

type ValueSpec struct {
	Spec *ast.ValueSpec
}

func NewValueSpec(spec *ast.ValueSpec) *ValueSpec {
	return &ValueSpec{
		Spec: spec,
	}
}
func (v *ValueSpec) IsExport() bool {
	for _, name := range v.Spec.Names {
		if IsExport(name.Name) {
			return true
		}
	}
	return false
}
func (v *ValueSpec) Output(all bool) (s string, e error) {
	if !all && !v.IsExport() {
		return
	}
	var w bytes.Buffer
	spec := v.Spec
	for i, name := range spec.Names {
		if i != 0 {
			_, e = w.WriteString(`, `)
			if e != nil {
				return
			}
		}
		if all || IsExport(name.Name) {
			_, e = w.WriteString(name.Name)
		} else {
			_, e = w.WriteString(`_`)
		}
		if e != nil {
			return
		}
	}
	if spec.Type != nil {

		_, e = w.WriteString(` ` + NewTypeExpr(spec.Type).TypeString())
		// _, e = w.WriteString(` ` + spec.Type.(*ast.Ident).Name)
		if e != nil {
			return
		}
	}
	// == 0 時通常是由 iota 指定
	if len(spec.Values) != 0 {
		_, e = w.WriteString(" = ")
		if e != nil {
			return
		}

		for i, expr := range spec.Values {
			if i != 0 {
				_, e = w.WriteString(`, `)
				if e != nil {
					return
				}
			}
			_, e = w.WriteString(NewValueExpr(expr).String())
			if e != nil {
				return
			}
		}
	}
	_, e = w.WriteString("\n")
	if e != nil {
		return
	}
	s = w.String()
	return
}
