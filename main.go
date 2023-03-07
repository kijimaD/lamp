package main

import (
	"flag"
	"os"

	"github.com/kijimaD/lamp/myscope"
)

const cmdCount = 2

func main() {
	myscopeCmd := flag.NewFlagSet("myscope", flag.ExitOnError)

	if len(os.Args) < cmdCount {
		panic("expected 'myscope' subcommands")
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
	default:
		panic("not found subcommand")
	}
}
