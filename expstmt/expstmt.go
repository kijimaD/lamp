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
			case *ast.DeclStmt:
				fmt.Fprintf(c.Output, "DeclStmt:\n")
			case *ast.EmptyStmt:
				fmt.Fprintf(c.Output, "EmptyStmt:\n")
			case *ast.LabeledStmt:
				fmt.Fprintf(c.Output, "LabeledStmt:\n")
			case *ast.ExprStmt:
				fmt.Fprintf(c.Output, "ExprStmt:\n")
			case *ast.SendStmt:
				fmt.Fprintf(c.Output, "SendStmt:\n")
			case *ast.IncDecStmt:
				fmt.Fprintf(c.Output, "IncDecStmt:\n")
			case *ast.AssignStmt:
				fmt.Fprintf(c.Output, "AssignStmt:\n")
			case *ast.GoStmt:
				fmt.Fprintf(c.Output, "GoStmt:\n")
			case *ast.DeferStmt:
				fmt.Fprintf(c.Output, "DeferStmt:\n")
			case *ast.ReturnStmt:
				fmt.Fprintf(c.Output, "ReturnStmt:\n")
			case *ast.BranchStmt:
				fmt.Fprintf(c.Output, "BranchStmt:\n")
			case *ast.BlockStmt:
				fmt.Fprintf(c.Output, "BlockStmt:\n")
			case *ast.IfStmt:
				fmt.Fprintf(c.Output, "IfStmt:\n")
			case *ast.CaseClause:
				fmt.Fprintf(c.Output, "CaseClause:\n")
			case *ast.SwitchStmt:
				fmt.Fprintf(c.Output, "SwitchStmt:\n")
			case *ast.TypeSwitchStmt:
				fmt.Fprintf(c.Output, "TypeSwitchStmt:\n")
			case *ast.CommClause:
				fmt.Fprintf(c.Output, "CommClause:\n")
			case *ast.SelectStmt:
				fmt.Fprintf(c.Output, "SelectStmt:\n")
			case *ast.ForStmt:
				fmt.Fprintf(c.Output, "ForStmt:\n")
			case *ast.RangeStmt:
				fmt.Fprintf(c.Output, "RangeStmt:\n")
			default:
				panic(d)
			}
		case ast.Expr:
			fmt.Fprintf(c.Output, "  Expr: ")
			switch d := v.(type) {
			case *ast.Ident:
				fmt.Fprintf(c.Output, "  Ident: %#v\n", d.Name)
			case *ast.Ellipsis:
				fmt.Fprintf(c.Output, "  Ellipsis:\n")
			case *ast.BasicLit:
				fmt.Fprintf(c.Output, "  BasicLit: %#v\n", d.Value)
			case *ast.FuncLit:
				fmt.Fprintf(c.Output, "  FuncLit:\n")
			case *ast.CompositeLit:
				fmt.Fprintf(c.Output, "  CompositeLit\n")
			case *ast.ParenExpr:
				fmt.Fprintf(c.Output, "  ParenExpr\n")
			case *ast.SelectorExpr:
				fmt.Fprintf(c.Output, "  SelectorExpr\n")
			case *ast.IndexExpr:
				fmt.Fprintf(c.Output, "  IndexExpr\n")
			case *ast.IndexListExpr:
				fmt.Fprintf(c.Output, "  IndexListExpr\n")
			case *ast.SliceExpr:
				fmt.Fprintf(c.Output, "  SliceExpr\n")
			case *ast.TypeAssertExpr:
				fmt.Fprintf(c.Output, "  TypeAssertExpr\n")
			case *ast.CallExpr:
				fmt.Fprintf(c.Output, "  CallExpr\n")
			case *ast.StarExpr:
				fmt.Fprintf(c.Output, "  StarExpr\n")
			case *ast.UnaryExpr:
				fmt.Fprintf(c.Output, "  UnaryExpr\n")
			case *ast.BinaryExpr:
				fmt.Fprintf(c.Output, "  BinaryExpr\n")
			case *ast.KeyValueExpr:
				fmt.Fprintf(c.Output, "  KeyValueExpr\n")
			case *ast.StructType:
				fmt.Fprintf(c.Output, "  StructType\n")
			case *ast.FuncType:
				fmt.Fprintf(c.Output, "  FuncType\n")
			case *ast.InterfaceType:
				fmt.Fprintf(c.Output, "  InterfaceType\n")
			case *ast.MapType:
				fmt.Fprintf(c.Output, "  MapType\n")
			case *ast.ChanType:
				fmt.Fprintf(c.Output, "  ChanType\n")
			default:
				panic(d)
			}
		}

		return true
	})
}
