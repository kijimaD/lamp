package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kijimaD/lamp/myscope"
)

func main() {
	myscopeCmd := flag.NewFlagSet("myscope", flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Println("expected 'myscope' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {

	case "myscope":
		myscopeCmd.Parse(os.Args[2:])
		file, err := os.Open(os.Args[2])
		if err != nil {
			panic(err)
		}
		c := myscope.NewClient(file, os.Stdout)
		c.Run()
	default:
		os.Exit(1)
	}
}
