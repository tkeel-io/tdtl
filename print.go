// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file contains printing support for ASTs.

package tdtl

import (
	"fmt"
	"github.com/tkeel-io/tdtl/parser"
	"go/token"
	"io"
	"os"
	"reflect"
	"strings"
)

var SymbolicNames []string

func init() {
	ps := parser.NewTDTLParser(nil)
	SymbolicNames = ps.SymbolicNames
}

func symbolicNames(op int) string {
	if op < len(SymbolicNames) && op >= 0 {
		name := SymbolicNames[op]
		if name != "" {
			return name
		}
	}
	return fmt.Sprintf("op[%d]", op)
}

// Fprint prints the (sub-)tree starting at AST node x to w.
// If fset != nil, position information is interpreted relative
// to that file set. Otherwise positions are printed as integer
// values (file set specific offsets).
//
// A non-nil FieldFilter f may be provided to control the output:
// struct fields for which f(fieldname, fieldvalue) is true are
// printed; all others are filtered from the output. Unexported
// struct fields are never printed.
func Fprint(w io.Writer, x Expr) error {
	return fprint(w, x)
}

func fprint(w io.Writer, x Expr) (err error) {
	// setup printer
	p := printer{
		output: w,
		ptrmap: make(map[interface{}]int),
		last:   '\n', // force printing of line number on first line
	}

	// install error handler
	defer func() {
		if e := recover(); e != nil {
			err = e.(localError).err // re-panics if it's not a localError
		}
	}()

	// print x
	if x == nil {
		p.printf("nil\n")
		return
	}
	p.print(x)
	p.printf("\n")

	return
}

// Print prints x to standard output, skipping nil fields.
// Print(fset, x) is the same as Fprint(os.Stdout, fset, x, NotNilFilter).
func Dump(x Expr) error {
	return Fprint(os.Stdout, x)
}

func DumpMore(x ...Expr) error {
	for _, e := range x {
		err := Fprint(os.Stdout, e)
		if err != nil {
			return err
		}
	}
	return nil
}

type printer struct {
	output io.Writer
	fset   *token.FileSet
	ptrmap map[interface{}]int // *T -> line number
	indent int                 // current indentation level
	last   byte                // the last byte processed by Write
	line   int                 // current line number
}

var indent = []byte(".  ")

func (p *printer) Write(data []byte) (n int, err error) {
	var m int
	for i, b := range data {
		// invariant: data[0:n] has been written
		if b == '\n' {
			m, err = p.output.Write(data[n : i+1])
			n += m
			if err != nil {
				return
			}
			p.line++
		} else if p.last == '\n' {
			_, err = fmt.Fprintf(p.output, "%6d  ", p.line)
			if err != nil {
				return
			}
			for j := p.indent; j > 0; j-- {
				_, err = p.output.Write(indent)
				if err != nil {
					return
				}
			}
		}
		p.last = b
	}
	if len(data) > n {
		m, err = p.output.Write(data[n:])
		n += m
	}
	return
}

// localError wraps locally caught errors so we can distinguish
// them from genuine panics which we don't want to return as errors.
type localError struct {
	err error
}

// printf is a convenience wrapper that takes care of print errors.
func (p *printer) printf(format string, args ...interface{}) {
	if _, err := fmt.Fprintf(p, format, args...); err != nil {
		panic(localError{err})
	}
}

// Implementation note: Print is written for AST nodes but could be
// used to print arbitrary data structures; such a version should
// probably be in a different package.
//
// Note: This code detects (some) cycles created via pointers but
// not cycles that are created via slices or maps containing the
// same slice or map. Code for general data structures probably
// should catch those as well.

func (p *printer) print(x Expr) {
	switch x := x.(type) {
	case *SelectStatementExpr:
		p.printf("Root {")
		p.indent++
		p.printf("\n")
		p.print(x.fields)
		p.printf("\n")
		p.print(x.topic)
		p.printf("\n")
		p.print(x.filter)
		p.printf("\n")
		p.indent--
		p.printf("}")
	case FieldsExpr:
		p.printf("Select {")
		p.indent++
		first := true
		for i, n := 0, len(x); i < n; i++ {
			if elem := x[i]; true {
				if first {
					p.printf("\n")
					first = false
				}
				p.print(elem)
				p.printf("\n")
			}
		}
		p.indent--
		p.printf("}")
	case *FieldExpr:
		p.printf("Field (%s) {", x.alias)
		p.indent++
		p.printf("\n")
		p.print(x.exp)
		p.printf("\n")
		p.indent--
		p.printf("}")
	case TopicExpr:
		p.printf("Topic [%s] {", strings.Join(x, "/"))
		p.indent++
		first := true
		for i, n := 0, len(x); i < n; i++ {
			if elem := x[i]; true {
				if first {
					p.printf("\n")
					first = false
				}
				p.printf("%s", elem)
				p.printf("\n")
			}
		}
		p.indent--
		p.printf("}")
	case *FilterExpr:
		p.printf("Where {")
		p.indent++
		p.printf("\n")
		p.print(x.exp)
		p.printf("\n")
		p.indent--
		p.printf("}")
	case *DimensionsExpr:
		p.printf("Dimensions {")
		p.indent++
		first := true
		for i, n := 0, len(x.exprs); i < n; i++ {
			if elem := x.exprs[i]; true {
				if first {
					p.printf("\n")
					first = false
				}
				p.print(elem)
				p.printf("\n")
			}
		}
		p.indent--
		p.printf("}\n")
		p.print(x.window)
	case *WindowExpr:
		p.printf("Window {")
		p.indent++
		p.printf("type: %v , Length: %d , Interval: %d ", x.WindowType, x.Length, x.Interval)
		p.indent--
		p.printf("}")
	case *BinaryExpr:
		p.printf("Op [%s] {", symbolicNames(x.Op))
		p.indent++
		p.printf("\n")
		p.print(x.LHS)
		p.printf("\n")
		p.print(x.RHS)
		p.printf("\n")
		p.indent--
		p.printf("}")
	case *SwitchExpr:
		p.printf("Switch {")
		p.indent++
		p.printf("\n")
		{
			p.printf("Case {")
			p.indent++
			p.printf("\n")
			p.print(x.exp)
			p.printf("\n")
			p.indent--
			p.printf("}")
			p.printf("\n")
		}
		for i, n := 0, len(x.list); i < n; i++ {
			elem := x.list[i]
			p.printf("When {")
			p.indent++
			p.printf("\n")
			p.print(elem.when)
			p.printf("\n")
			{
				p.printf("Then {")
				p.indent++
				p.printf("\n")
				p.print(elem.then)
				p.printf("\n")
				p.indent--
				p.printf("}")
				p.printf("\n")
			}
			p.indent--
			p.printf("}")
			p.printf("\n")
		}
		{
			p.printf("ELSE {")
			p.indent++
			p.printf("\n")
			p.print(x.last)
			p.printf("\n")
			p.indent--
			p.printf("}")
			p.printf("\n")
		}
		p.indent--
		p.printf("}")
	case CaseListExpr:
		p.printf("CaseListExpr {")
		p.printf("}")
	case *CaseExpr:
		p.printf("CaseExpr {")
		p.indent++
		p.printf("\n")
		p.print(x.then)
		p.printf("\n")
		p.indent--
		p.printf("}")
		p.printf("\n")
	case *CallExpr:
		p.printf("func: %v {", x.key)
		p.indent++
		for _, e := range x.args {
			p.printf("\"%v\"", e)
			p.printf(" ")
		}
		p.indent--
		p.printf("}")
	case StringNode:
		p.printf("\"%v\"", x)
	case *JSONPathExpr:
		p.printf("\"ref:%v\"", x)
	default:
		// default
		p.printf("%v", x)
	}
}

func typeOf(x interface{}) string {
	str := strings.Replace(
		fmt.Sprintf("%v", reflect.TypeOf(x)),
		"ruleql.",
		"",
		-1,
	)
	return str
}
