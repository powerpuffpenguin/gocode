package gocode

import (
	"go/ast"
	"io"
)

type Const struct {
	Specs []ast.Spec
}

func NewConst(specs []ast.Spec) *Const {
	return &Const{
		Specs: specs,
	}
}
func (c *Const) IsExport() bool {
	for _, a := range c.Specs {
		spec := a.(*ast.ValueSpec)
		for _, name := range spec.Names {
			if IsExport(name.Name) {
				return true
			}
		}
	}
	return false
}
func (c *Const) Output(writer io.Writer, prefix, indent string, all bool) (n int64, e error) {
	w := writerTo{w: writer}
	specs := c.Specs
	var s string
	if len(specs) == 1 {
		_, e = w.WriteString(prefix + `const `)
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
		_, e = w.WriteString(prefix + "const (\n")
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
