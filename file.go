package gocode

import (
	"bytes"
	"go/ast"
	"go/token"
	"io"
	"reflect"
)

// 包中的檔案信息
type File struct {
	AST *ast.File
	// 檔案名稱
	Name string
	// 導入包
	Imports []*Import
	// 定義的 struct
	Structs []*Struct
	// 定義的 interface
	Interfaces []*Interface
	// 定義的別名 type xxx xxx
	Alias []*Alias
	// 定義的常量
	Consts []*Const
	// 定義的變量
	Vars []*Var
	// 定義的函數
	Funcs []*Func
}

func NewFile(name string, f *ast.File) *File {
	var (
		size       = len(f.Imports)
		imports    []*Import
		structs    []*Struct
		interfaces []*Interface
		alias      []*Alias
		consts     []*Const
		vars       []*Var
		funcs      []*Func
	)
	if size > 0 {
		imports = make([]*Import, size)
		for i, v := range f.Imports {
			imports[i] = NewImport(v)
		}
	}

	for _, decl := range f.Decls {
		switch node := decl.(type) {
		case *ast.GenDecl:
			switch node.Tok {
			case token.IMPORT:
			case token.CONST:
				consts = append(consts, NewConst(node.Specs))
			case token.VAR:
				vars = append(vars, NewVar(node.Specs))
			case token.TYPE:
				for _, spec := range node.Specs {
					tspec := spec.(*ast.TypeSpec)
					name := tspec.Name.Name
					switch t := tspec.Type.(type) {
					case *ast.StructType:
						structs = append(structs, NewStruct(name, t))
					case *ast.InterfaceType:
						interfaces = append(interfaces, NewInterface(name, t))
					default:
						alias = append(alias, NewAlias(name, t))
						// panic(`unknow type: ` + reflect.TypeOf(t).String())
					}
				}
			default:
				panic(`unknow token: ` + node.Tok.String())
			}
		case *ast.FuncDecl:
			funcs = append(funcs, NewFunc(node))
		default:
			panic(`unknow decl: ` + reflect.TypeOf(decl).String())
		}
	}
	return &File{
		AST:        f,
		Name:       name,
		Imports:    imports,
		Structs:    structs,
		Interfaces: interfaces,
		Alias:      alias,
		Consts:     consts,
		Vars:       vars,
		Funcs:      funcs,
	}
}
func (f *File) String() string {
	w := bytes.NewBuffer(make([]byte, 0, 1024))
	_, e := f.Output(w, ``, ``, false)
	if e != nil {
		panic(e)
	}
	b := w.Bytes()
	return BytesToString(b)
}
func (f *File) Output(writer io.Writer, prefix, indent string, all bool) (n int64, e error) {
	w := writerTo{w: writer}
	_, e = w.WriteString(prefix + `file: ` + f.Name + "\n")
	if e != nil {
		n = w.n
		return
	}
	prefix += indent
	for _, node := range f.Imports {
		_, e = w.WriteString(prefix + node.String() + "\n")
		if e != nil {
			n = w.n
			return
		}
	}
	for _, node := range f.Consts {
		if !all && !node.IsExport() {
			continue
		}
		_, e = node.Output(&w, prefix, indent, all)
		if e != nil {
			n = w.n
			return
		}
	}
	for _, node := range f.Vars {
		if !all && !node.IsExport() {
			continue
		}
		_, e = node.Output(&w, prefix, indent, all)
		if e != nil {
			n = w.n
			return
		}
	}
	for _, node := range f.Interfaces {
		if !all && !node.IsExport() {
			continue
		}
		_, e = node.Output(&w, prefix, indent)
		if e != nil {
			n = w.n
			return
		}
	}
	for _, node := range f.Alias {
		if !all && !node.IsExport() {
			continue
		}
		_, e = w.WriteString(prefix + node.String() + "\n")
		if e != nil {
			n = w.n
			return
		}
	}
	for _, node := range f.Structs {
		if !all && !node.IsExport() {
			continue
		}
		_, e = node.Output(&w, prefix, indent, all)
		if e != nil {
			n = w.n
			return
		}
	}
	for _, node := range f.Funcs {
		if !all && !node.IsExport() {
			continue
		}
		_, e = w.WriteString(prefix + node.String() + "\n")
		if e != nil {
			n = w.n
			return
		}
	}
	n = w.n
	return
}
