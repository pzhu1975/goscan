// 19 august 2018

// +build amd64 386

package main

// TODO probably a bug in libui: changing the font away from skia leads to a crash

import (
	// "errors"
	"fmt"
	// https://github.com/raff/walkngo
	//	"github.com/tdewolff/parse/html"
	"github.com/goscan/ast"
	"github.com/goscan/parser"

	// "github.com/goscan/scanner"
	"github.com/goscan/token"
)

func main() {

	const src = `package main
type x int // comment test
`
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", src, parser.ParseComments)
	if err != nil {
		fmt.Println("parser failed")
	}
	comment := f.Decls[0].(*ast.GenDecl).Specs[0].(*ast.TypeSpec).Comment.List[0].Text
	if comment != "// comment" {
		fmt.Printf("got %q, want %q", comment, "// comment")
	}
}
