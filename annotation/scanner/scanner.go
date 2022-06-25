package scanner

import (
	"fmt"
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

	ast.Inspect(f, func(n ast.Node) bool {
		switch t := n.(type) {
		case *ast.FuncDecl:
			if t.Doc != nil {
				fmt.Println(t.Doc.Text())
				fmt.Print(t.Name.String())
				typ := t.Type
				if typ.Params != nil {
					for _, p := range typ.Params.List {
						for _, o := range p.Names {
							fmt.Print(o.Name, " ")
						}
						typ := p.Type
						if t, ok := typ.(*ast.StarExpr); ok {
							fmt.Print(t.X, " ")
						}
					}
				}

			}
		case *ast.TypeSpec:
			fmt.Println(t.Doc.Text())
		case *ast.StructType:
			for _, field := range t.Fields.List {
				fmt.Println(field.Doc.Text())
				fmt.Println(field.Comment.Text())
			}
		}
		return true
	})
}
