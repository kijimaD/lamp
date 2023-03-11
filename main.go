package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"

	"github.com/kijimaD/lamp/expstmt"
	"github.com/kijimaD/lamp/myscope"
)

const cmdCount = 2

func main() {
	myscopeCmd := flag.NewFlagSet("myscope", flag.ExitOnError)

	if len(os.Args) < cmdCount {
		panic("expected 'myscope' | 'astdump' | 'expstmt' subcommands")
	}

	switch os.Args[1] {
	case "myscope":
		if err := myscopeCmd.Parse(os.Args[2:]); err != nil {
			panic(err)
		}
		file, err := os.Open(os.Args[2])
		if err != nil {
			panic(err)
		}
		c := myscope.NewClient(file, os.Stdout)
		c.Run()
	case "astdump":
		if err := myscopeCmd.Parse(os.Args[2:]); err != nil {
			panic(err)
		}
		fset := token.NewFileSet()
		astf, err := parser.ParseFile(fset, os.Args[2], nil, parser.Mode(0))
		if err != nil {
			panic(err)
		}

		for _, d := range astf.Decls {
			ast.Print(fset, d)
			fmt.Println() //nolint:forbidigo
		}
	case "expstmt":
		if err := myscopeCmd.Parse(os.Args[2:]); err != nil {
			panic(err)
		}
		file, err := os.Open(os.Args[2])
		if err != nil {
			panic(err)
		}
		c := expstmt.NewClient(file, os.Stdout)
		c.Run()
	default:
		panic("not found subcommand")
	}
}
