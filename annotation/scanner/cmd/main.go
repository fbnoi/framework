package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	fSet := token.NewFileSet()
	f, err := parser.ParseFile(fSet, "../../../example/controller/controller1.go", nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	ast.Inspect(f, func(n ast.Node) bool {
		switch t := n.(type) {
		case *ast.FuncDecl:
			if t.Doc != nil {
				for _, c := range t.Doc.List {
					fmt.Println(c.Text, "PP")
				}
				fmt.Println(t.Name.String())
			}
		}
		return true
	})
}
