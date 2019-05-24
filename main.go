// 19 august 2018

// +build amd64 386

package main

// TODO probably a bug in libui: changing the font away from skia leads to a crash

import (
	// "errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	// "github.com/goscan/scanner"
	"github.com/goscan/printer"
	"github.com/goscan/walker"
)

type Walker struct {
	*walker.GoWalker

	outdir string
	prefix string
	ext    string
}

func (w Walker) Walk(path string, info os.FileInfo, err error) error {
	if err != nil {
		fmt.Println(err)
		return nil
	}

	var outpath string

	if len(w.outdir) > 0 {
		outpath = filepath.Join(w.outdir, path[len(w.prefix):])
	}

	if info.IsDir() {
		if strings.HasPrefix(info.Name(), ".") && info.Name() != "." { // assume we want to skip hidden folders
			return filepath.SkipDir
		}

		if len(outpath) > 0 {
			if err := os.MkdirAll(outpath, 0755); err != nil {
				fmt.Println(err)
			}
		}
	} else if strings.HasSuffix(path, ".go") {
		if len(outpath) > 0 {
			outpath = outpath[:len(outpath)-2] + w.ext
			f, err := os.Create(outpath)
			if err != nil {
				fmt.Println(err)
			} else {
				w.SetWriter(f)
				defer f.Close()
			}
		}

		if err := w.WalkFile(path); err != nil {
			fmt.Println(err)
		}
	}

	return nil
}

func main() {
	var p printer.Printer
	p = &printer.GoPrinter{}
	walker := Walker{walker.NewWalker(p, os.Stdout, false), ".", "", "go"} //*debug

	// for _, f := range flag.Args() {
	walker.prefix = "."

	filepath.Walk(".", walker.Walk)
	// }
	// 	const src = `package main
	// type x int // comment test
	// `
	// 	fset := token.NewFileSet()
	// 	f, err := parser.ParseFile(fset, "", src, parser.ParseComments)
	// 	if err != nil {
	// 		fmt.Println("parser failed")
	// 	}
	// 	comment := f.Decls[0].(*ast.GenDecl).Specs[0].(*ast.TypeSpec).Comment.List[0].Text
	// 	if comment != "// comment" {
	// 		fmt.Printf("got %q, want %q", comment, "// comment")
	// 	}
}
