package myscope_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/kijimaD/lamp/myscope"
	"github.com/stretchr/testify/assert"
)

const src = `package main

import "fmt"

func main() {
    const message = "hello, world"
    fmt.Println(message)
}`

func TestRun(t *testing.T) {
	t.Parallel()

	in := strings.NewReader(src)
	out := &bytes.Buffer{}

	c := myscope.NewClient(in, out)
	c.Run()

	part := "──── 3 scopes ────"
	got := out.String()

	assert.Contains(t, got, part)
}
