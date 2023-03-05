package main

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"log"

	"github.com/davecgh/go-spew/spew"
)

// sample.go
const src = `package main

import "fmt"

func main() {
    const message = "hello, world"
    fmt.Println(message)
}`

func main() {
	// https://qiita.com/tenntenn/items/ac5940dfbca703183fdf
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "sample.go", src, 0)
	if err != nil {
		log.Fatalln("Error:", err)
	}

	conf := types.Config{Importer: importer.Default()}
	pkg, err := conf.Check("main", fset, []*ast.File{f}, nil)
	if err != nil {
		log.Fatalln("Error:", err)
	}

	scopes := map[*types.Scope]struct{}{}
	ast.Inspect(f, func(n ast.Node) bool {
		if ident, ok := n.(*ast.Ident); ok {
			innerMost := pkg.Scope().Innermost(ident.Pos())
			s, _ := innerMost.LookupParent(ident.Name, ident.Pos())
			if s != nil {
				scopes[s] = struct{}{}
			}
		}
		return true
	})

	fmt.Println("────", len(scopes), "scopes ────")
	for s := range scopes {
		spew.Dump(s)
	}
}
