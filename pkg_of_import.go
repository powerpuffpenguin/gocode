package gocode

import (
	"go/ast"
)

type Import struct {
	AST *ast.ImportSpec
}

func NewImport(imt *ast.ImportSpec) *Import {
	return &Import{
		AST: imt,
	}
}

// 返回 import 名稱
func (imt *Import) Name() string {
	if imt.AST.Name == nil {
		return ``
	}
	return imt.AST.Name.Name
}

// 返回 import 路徑
func (imt *Import) Path() string {
	return imt.AST.Path.Value
}
func (imt *Import) String() string {
	name, path := imt.Name(), imt.Path()
	if name == `` {
		return `import ` + path
	} else {
		return `import ` + name + ` ` + path
	}
}
