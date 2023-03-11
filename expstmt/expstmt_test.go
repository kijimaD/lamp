package expstmt_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/kijimaD/lamp/expstmt"
	"github.com/stretchr/testify/assert"
)

const src = `
package expstmt

import "fmt"

type S1 struct {
	test string
}

func main() {
	s1 := S1{"test struct"}
	message := test("hello world")
	fmt.Println(message, s1)
}

func test(s string) string {
	return s
}

func (s S1) String() string {
	return s.test
}
`

const expect = `  Expr:   Ident: "expstmt"
GenDecl:
  Expr:   BasicLit: "\"fmt\""
GenDecl:
  Expr:   Ident: "S1"
  Expr:   StructType
  Expr:   Ident: "test"
  Expr:   Ident: "string"
FuncDecl:
  Expr:   Ident: "main"
  Expr:   FuncType
BlockStmt:
AssignStmt:
  Expr:   Ident: "s1"
  Expr:   CompositeLit
  Expr:   Ident: "S1"
  Expr:   BasicLit: "\"test struct\""
AssignStmt:
  Expr:   Ident: "message"
  Expr:   CallExpr
  Expr:   Ident: "test"
  Expr:   BasicLit: "\"hello world\""
ExprStmt:
  Expr:   CallExpr
  Expr:   SelectorExpr
  Expr:   Ident: "fmt"
  Expr:   Ident: "Println"
  Expr:   Ident: "message"
  Expr:   Ident: "s1"
FuncDecl:
  Expr:   Ident: "test"
  Expr:   FuncType
  Expr:   Ident: "s"
  Expr:   Ident: "string"
  Expr:   Ident: "string"
BlockStmt:
ReturnStmt:
  Expr:   Ident: "s"
FuncDecl:
  Expr:   Ident: "s"
  Expr:   Ident: "S1"
  Expr:   Ident: "String"
  Expr:   FuncType
  Expr:   Ident: "string"
BlockStmt:
ReturnStmt:
  Expr:   SelectorExpr
  Expr:   Ident: "s"
  Expr:   Ident: "test"
`

func TestRun(t *testing.T) {
	t.Parallel()

	in := strings.NewReader(src)
	out := &bytes.Buffer{}

	c := expstmt.NewClient(in, out)
	c.Run()

	assert.Equal(t, expect, out.String())
}
