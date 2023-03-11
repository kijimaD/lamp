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

	ast.Inspect(file, func(n ast.Node) bool {
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
				fmt.Fprintf(c.Output, "BlockStmt:\n")
			case *ast.AssignStmt:
				fmt.Fprintf(c.Output, "AssignStmt:\n")
			case *ast.ExprStmt:
				fmt.Fprintf(c.Output, "ExprStmt:\n")
			case *ast.ReturnStmt:
				fmt.Fprintf(c.Output, "ReturnStmt:\n")
			case *ast.IfStmt:
				fmt.Fprintf(c.Output, "IfStmt:\n")
			case *ast.SwitchStmt:
				fmt.Fprintf(c.Output, "SwitchStmt:\n")
			case *ast.RangeStmt:
				fmt.Fprintf(c.Output, "RangeStmt:\n")
			case *ast.CaseClause:
				fmt.Fprintf(c.Output, "CaseClause:\n")
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
			case *ast.BinaryExpr:
				fmt.Fprintf(c.Output, "  BinaryExpr\n")
			case *ast.IndexExpr:
				fmt.Fprintf(c.Output, "  IndexExpr\n")
			case *ast.SliceExpr:
				fmt.Fprintf(c.Output, "  SliceExpr\n")
			default:
				panic(d)
			}
		}

		return true
	})
}
