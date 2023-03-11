package expstmt

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"log"
)

type Client struct {
	Input  io.Reader
	Output io.Writer
}

func NewClient(input io.Reader, output io.Writer) Client {
	return Client{input, output}
}

func (c *Client) Run() {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "", c.Input, 0)
	if err != nil {
		log.Fatalln("Error:", err)
	}

	// scopes := map[*types.Scope]struct{}{}
	ast.Inspect(file, func(n ast.Node) bool {
		// if ident, ok := n.(*ast.Ident); ok {
		// 	innerMost := pkg.Scope().Innermost(ident.Pos())
		// 	s, _ := innerMost.LookupParent(ident.Name, ident.Pos())
		// 	if s != nil {
		// 		scopes[s] = struct{}{}
		// 	}
		// }
		// fmt.Fprintf(c.Output, "%#v", n)
		switch v := n.(type) {
		case ast.Decl:
			switch d := v.(type) {
			case *ast.GenDecl:
				fmt.Fprintf(c.Output, "GenDecl:\n")
			case *ast.FuncDecl:
				fmt.Fprintf(c.Output, "FuncDecl:\n")
			default:
				panic(d)
			}
		case ast.Stmt:
			switch d := v.(type) {
			case *ast.BlockStmt:
				fmt.Fprintf(c.Output, "Blockstmt:\n")
			case *ast.AssignStmt:
				fmt.Fprintf(c.Output, "Assignstmt:\n")
			case *ast.ExprStmt:
				fmt.Fprintf(c.Output, "Exprstmt:\n")
			case *ast.ReturnStmt:
				fmt.Fprintf(c.Output, "Returnstmt:\n")
			default:
				panic(d)
			}
		case ast.Expr:
			fmt.Fprintf(c.Output, "  Expr: ")
			switch d := v.(type) {
			case *ast.Ident:
				fmt.Fprintf(c.Output, "  Ident: %#v\n", d.Name)
			case *ast.BasicLit:
				fmt.Fprintf(c.Output, "  BasicLit: %#v\n", d.Value)
			case *ast.StructType:
				fmt.Fprintf(c.Output, "  StructType\n")
			case *ast.FuncType:
				fmt.Fprintf(c.Output, "  FuncType\n")
			case *ast.CompositeLit:
				fmt.Fprintf(c.Output, "  CompositeLit\n")
			case *ast.CallExpr:
				fmt.Fprintf(c.Output, "  CallExpr\n")
			case *ast.SelectorExpr:
				fmt.Fprintf(c.Output, "  SelectorExpr\n")
			default:
				panic(d)
			}
		}

		return true
	})
}
