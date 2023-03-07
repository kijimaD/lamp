package main

import (
	"os"

	"github.com/kijimaD/lamp/myscope"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}

	c := myscope.NewClient(file, os.Stdout)
	c.Run()
}
