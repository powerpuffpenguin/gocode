package gocode_test

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"reflect"
	"sort"
	"strings"
	"testing"

	"github.com/powerpuffpenguin/gocode"
	"github.com/stretchr/testify/assert"
)

func getStructFields(f *ast.File, name string) *ast.FieldList {
	for _, decl := range f.Decls {
		d, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		} else if d.Tok != token.TYPE {
			continue
		}
		for _, spec := range d.Specs {
			s, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			} else if s.Name.Name != `Cat` {
				continue
			}
			t, ok := s.Type.(*ast.StructType)
			if !ok {
				continue
			}
			return t.Fields
		}
	}
	return nil
}

func TestTypes(t *testing.T) {
	tests := map[string]struct {
		String string
		Import string
		Name   string
		Star   bool
		Array  bool
		Chan   bool
	}{
		`x`: {
			Name:   `int`,
			String: `int`,
		},
		`x_pointer`: {
			Name:   `int`,
			Star:   true,
			String: `*int`,
		},
		`x_array`: {
			Name:   `int`,
			String: `[]int`,
			Array:  true,
		},
		`x_array2`: {
			Name:   `int`,
			String: `[][]int`,
			Array:  true,
		},
		`x_array3`: {
			Name:   `int`,
			String: `[][]*int`,
			Array:  true,
		},
		`x_array4`: {
			Name:   `int`,
			String: `[]*[]*int`,
			Array:  true,
		},
		`x_array5`: {
			Name:   `int`,
			Star:   true,
			String: `*[]*[]*int`,
			Array:  true,
		},

		`y`: {
			Name:   `int`,
			Import: `ea`,
			String: `ea.int`,
		},
		`y_pointer`: {
			Name:   `int`,
			Import: `ea`,
			Star:   true,
			String: `*ea.int`,
		},
		`y_array`: {
			Name:   `int`,
			Import: `ea`,
			String: `[]ea.int`,
			Array:  true,
		},
		`y_array2`: {
			Name:   `int`,
			Import: `ea`,
			String: `[][]ea.int`,
			Array:  true,
		},
		`y_array3`: {
			Name:   `int`,
			Import: `ea`,
			String: `[][]*ea.int`,
			Array:  true,
		},
		`y_array4`: {
			Name:   `int`,
			Import: `ea`,
			String: `[]*[]*ea.int`,
			Array:  true,
		},
		`y_array5`: {
			Name:   `int`,
			Import: `ea`,
			Star:   true,
			String: `*[]*[]*ea.int`,
			Array:  true,
		},
		`a`: {
			Name:   `any`,
			Star:   false,
			String: `any`,
		},
		`a0`: {
			Name:   `interface{}`,
			Star:   true,
			String: `*interface{}`,
		},
		`a1`: {
			Name:   `interface{}`,
			String: `interface{}`,
		},
		`a2`: {
			Name:   `interface{}`,
			String: `[]interface{}`,
			Array:  true,
		},
		`a3`: {
			Name:   `interface{}`,
			String: `[][]*interface{}`,
			Array:  true,
		},
		`a4`: {
			Name:   `interface{}`,
			Star:   true,
			String: `*[][]*interface{}`,
			Array:  true,
		},
		`c`: {
			Name:   `any`,
			String: `chan any`,
			Chan:   true,
		},
		`c1`: {
			Name:   `interface{}`,
			String: `chan interface{}`,
			Chan:   true,
		},
		`c2`: {
			Name:   `abc`,
			Import: `ko`,
			String: `chan ko.abc`,
			Chan:   true,
		},
		`c2p`: {
			Name:   `abc`,
			Import: `ko`,
			String: `chan *ko.abc`,
			Chan:   true,
		},
		`c2pp`: {
			Name:   `abc`,
			Import: `ko`,
			String: `*chan ko.abc`,
			Star:   true,
		},
		`c2pp2`: {
			Name:   `abc`,
			Import: `ko`,
			String: `*chan *ko.abc`,
			Star:   true,
		},
		`c3`: {
			Name:   `abc`,
			Import: `ko`,
			String: `*chan chan *ko.abc`,
			Star:   true,
		},
		`cw`: {
			Name:   `int`,
			String: `chan<- int`,
			Chan:   true,
		},
		`cw3`: {
			Name:   `abc`,
			Import: `ko`,
			String: `<-chan chan<- *ko.abc`,
			Chan:   true,
		},
		`cr3`: {
			Name:   `abc`,
			Import: `ko`,
			String: `*<-chan *chan<- *ko.abc`,
			Star:   true,
		},
		// `map`: {},
		// `func`: {},
	}
	var w bytes.Buffer
	for k, v := range tests {
		w.WriteString("\t" + k + " " + v.String + "\n")
	}
	fset := token.NewFileSet()
	f, e := parser.ParseFile(fset, `main.go`, `package main
type Cat struct{
`+w.String()+`}`, 0)
	if !assert.Nil(t, e) {
		t.FailNow()
	}
	fields := getStructFields(f, `Cat`)

	sort.Slice(fields.List, func(i, j int) bool {
		l := fields.List[i].Names[0]
		r := fields.List[j].Names[0]
		return strings.Compare(l.Name, r.Name) < 0
	})
	for _, f := range fields.List {
		key := f.Names[0].Name
		test, ok := tests[key]
		if !ok {
			fmt.Println(reflect.TypeOf(f.Type))
			continue
		}
		ty, e := gocode.NewType(f.Type)
		if !assert.Nil(t, e) {
			t.FailNow()
		}
		if !assert.Equal(t, test.Name, ty.Name(), `%s Name %s`, key, test.String) {
			t.FailNow()
		}
		if !assert.Equal(t, test.Star, ty.Star(), `%s Star %s`, key, test.String) {
			t.FailNow()
		}
		arr := ty.Array() != nil
		if !assert.Equal(t, test.Array, arr, `%s Array %s`, key, test.String) {
			t.FailNow()
		}
		ch := ty.Chan() != nil
		if !assert.Equal(t, test.Chan, ch, `%s Chan %s`, key, test.String) {
			t.FailNow()
		}
		if !assert.Equal(t, test.Import, ty.Import(), `%s Import %s`, key, test.String) {
			t.FailNow()
		}
		if !assert.Equal(t, test.String, ty.String(), `%s String %s`, key, test.String) {
			t.FailNow()
		}
	}
}
