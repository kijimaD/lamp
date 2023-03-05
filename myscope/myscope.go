package myscope

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"io"
	"log"

	"github.com/davecgh/go-spew/spew"
)

type Client struct {
	Input  io.Reader
	Output io.Writer
}

func NewClient(input io.Reader, output io.Writer) Client {
	return Client{input, output}
}

func (c *Client) Run() {
	// https://qiita.com/tenntenn/items/ac5940dfbca703183fdf
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "", c.Input, 0)
	if err != nil {
		log.Fatalln("Error:", err)
	}

	conf := types.Config{Importer: importer.Default()}
	pkg, err := conf.Check("main", fset, []*ast.File{file}, nil)
	if err != nil {
		log.Fatalln("Error:", err)
	}

	scopes := map[*types.Scope]struct{}{}
	ast.Inspect(file, func(n ast.Node) bool {
		if ident, ok := n.(*ast.Ident); ok {
			innerMost := pkg.Scope().Innermost(ident.Pos())
			s, _ := innerMost.LookupParent(ident.Name, ident.Pos())
			if s != nil {
				scopes[s] = struct{}{}
			}
		}

		return true
	})

	fmt.Fprintln(c.Output, "────", len(scopes), "scopes ────")

	for s := range scopes {
		spew.Fdump(c.Output, s)
	}
}
