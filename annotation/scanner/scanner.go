package scanner

import (
	"go/ast"
	"go/parser"
	"go/token"
)

type Store struct {
}

func Scan(path string) {
	fSet := token.NewFileSet()
	f, err := parser.ParseFile(fSet, path, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	ast.Print(fSet, f)
}
