package gocode

import (
	"bytes"
	"fmt"
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
}

func NewFile(name string, f *ast.File) *File {
	var (
		size       = len(f.Imports)
		imports    []*Import
		structs    []*Struct
		interfaces []*Interface
	)
	if size > 0 {
		imports = make([]*Import, size)
		for i, v := range f.Imports {
			imports[i] = NewImport(v)
		}
	}

	keysStruct := make(map[string]*Struct)
	keysInterface := make(map[string]*Interface)
	for _, decl := range f.Decls {
		switch node := decl.(type) {
		case *ast.GenDecl:
			switch node.Tok {
			case token.IMPORT:
			case token.CONST:
				fmt.Println(`CONST`)
			case token.VAR:
				fmt.Println(`VAR`)
			case token.TYPE:
				for _, spec := range node.Specs {
					tspec := spec.(*ast.TypeSpec)
					name := tspec.Name.Name
					switch t := tspec.Type.(type) {
					case *ast.StructType:
						if _, ok := keysStruct[name]; ok {
							panic(`struct alreay exists: ` + name)
						}
						keysStruct[name] = NewStruct(name, t)
					case *ast.InterfaceType:
						if _, ok := keysInterface[name]; ok {
							panic(`interface alreay exists: ` + name)
						}
						keysInterface[name] = NewInterface(name, t)
					default:
						fmt.Println(reflect.TypeOf(t))
					}
				}
			default:
				panic(`unknow token: ` + node.Tok.String())
			}
		case *ast.FuncDecl:
		default:
			panic(reflect.TypeOf(decl))
		}
	}
	size = len(keysStruct)
	if size > 0 {
		structs = make([]*Struct, 0, size)
		for _, v := range keysStruct {
			structs = append(structs, v)
		}
	}
	size = len(keysInterface)
	if size > 0 {
		interfaces = make([]*Interface, 0, size)
		for _, v := range keysInterface {
			interfaces = append(interfaces, v)
		}
	}
	return &File{
		AST:        f,
		Name:       name,
		Imports:    imports,
		Structs:    structs,
		Interfaces: interfaces,
	}
}
func (f *File) String() string {
	w := bytes.NewBuffer(make([]byte, 0, 1024))
	_, e := f.Output(w, ``, ``)
	if e != nil {
		panic(e)
	}
	b := w.Bytes()
	return BytesToString(b)
}
func (f *File) Output(writer io.Writer, prefix, indent string) (n int64, e error) {
	w := writerTo{w: writer}
	_, e = w.WriterString(prefix + `file: ` + f.Name + "\n")
	if e != nil {
		n = w.n
		return
	}
	prefix += indent
	for _, node := range f.Imports {
		_, e = w.WriterString(prefix + node.String() + "\n")
		if e != nil {
			n = w.n
			return
		}
	}
	for _, node := range f.Structs {
		_, e = node.Output(&w, prefix, indent)
		if e != nil {
			n = w.n
			return
		}
	}
	// for _, node := range f.Interfaces {
	// 	_, e = node.Output(&w, prefix, indent)
	// 	if e != nil {
	// 		n = w.n
	// 		return
	// 	}
	// }

	n = w.n
	return
}
