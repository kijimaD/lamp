package myscope_test

import (
	"strings"
	"testing"

	"github.com/kijimaD/lamp/myscope"
)

const src = `package main

import "fmt"

func main() {
    const message = "hello, world"
    fmt.Println(message)
}`

func TestRun(t *testing.T) {
	t.Parallel()

	buf := strings.NewReader(src)
	myscope.Run(buf)
}
