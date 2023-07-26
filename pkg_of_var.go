package gocode

import (
	"go/ast"
	"io"
)

type Var struct {
	Specs []ast.Spec
}

func NewVar(specs []ast.Spec) *Var {
	return &Var{
		Specs: specs,
	}
}
func (v *Var) IsExport() bool {
	for _, a := range v.Specs {
		spec := a.(*ast.ValueSpec)
		for _, name := range spec.Names {
			if IsExport(name.Name) {
				return true
			}
		}
	}
	return false
}
func (v *Var) Output(writer io.Writer, prefix, indent string, all bool) (n int64, e error) {
	if !all && !v.IsExport() {
		return
	}

	w := writerTo{w: writer}
	specs := v.Specs
	var s string
	if len(specs) == 1 {
		_, e = w.WriteString(prefix + `var `)
		if e != nil {
			n = w.n
			return
		}
		spec := specs[0].(*ast.ValueSpec)
		s, e = NewValueSpec(spec).Output(all)
		if e != nil {
			n = w.n
			return
		}
		_, e = w.WriteString(s)
		if e != nil {
			n = w.n
			return
		}
	} else {
		_, e = w.WriteString(prefix + "var (\n")
		if e != nil {
			n = w.n
			return
		}
		for _, spec := range specs {
			s, e = NewValueSpec(spec.(*ast.ValueSpec)).Output(all)
			if e != nil {
				n = w.n
				return
			}
			_, e = w.WriteString(prefix + indent + s)
			if e != nil {
				n = w.n
				return
			}
		}

		_, e = w.WriteString(prefix + ")\n")
		if e != nil {
			n = w.n
			return
		}
	}
	n = w.n
	return
}
