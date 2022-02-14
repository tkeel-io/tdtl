/*
Copyright 2021 The tKeel Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package tdtl

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/tkeel-io/tdtl/json/jsonparser"
	"github.com/tkeel-io/tdtl/parser"
	"strings"
)

//func Parse(expr string) Expr {
//	ex, err := ParseWithError(expr)
//	if err != nil {
//		return nil
//	}
//	return ex
//}

func Parse(expr string) (Expr, error) {
	parse, listener := parse(expr)
	antlr.ParseTreeWalkerDefault.Walk(listener, parse.Root())

	return listener.Expr(), listener.error()
}

func ParseField(expr string) (Expr, error) {
	parse, listener := parse(expr)
	antlr.ParseTreeWalkerDefault.Walk(listener, parse.Fields())

	return listener.Expr(), listener.error()
}

func ParseFilter(expr string) (Expr, error) {
	parse, listener := parse(expr)
	antlr.ParseTreeWalkerDefault.Walk(listener, parse.Filter())

	return listener.Expr(), listener.error()
}

func ParseExpr(expr string) (Expr, error) {
	parse, listener := parse(expr)
	antlr.ParseTreeWalkerDefault.Walk(listener, parse.Expr())
	return listener.Expr(), listener.error()
}

func parse(expr string) (*parser.TDTLParser, *TDTLListener) {
	//expr = strings.ReplaceAll(expr, "\n", " ")
	// Setup the input
	is := antlr.NewInputStream(expr)

	// Create the Lexer
	lexer := parser.NewTDTLLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	parse := parser.NewTDTLParser(stream)
	parse.RemoveErrorListeners()

	// Finally parseField the expression (by walking the tree)
	var listener TDTLListener
	parse.AddErrorListener(&listener)
	return parse, &listener
}

func updateJSON(node JSONNode, xpath string, setValue StringNode) (string, error) {
	keys := strings.Split(xpath, ".")
	value, err := jsonparser.Set([]byte(node), []byte(setValue), keys...)
	return string(value), err
}

func ParseFunc(x Expr) []*CallExpr {
	ret := &callList{make([]*CallExpr, 0)}
	ret.walkFunc(x)
	return ret.list
}

type callList struct {
	list []*CallExpr
}

func (c *callList) walkFunc(x Expr) {
	switch x := x.(type) {
	case *SelectStatementExpr:
		c.walkFunc(x.fields)
		c.walkFunc(x.filter)
	case FieldsExpr:
		for i, n := 0, len(x); i < n; i++ {
			if elem := x[i]; true {
				c.walkFunc(elem)
			}
		}
	case *FieldExpr:
		c.walkFunc(x.exp)
	case *FilterExpr:
		c.walkFunc(x.exp)
	case *BinaryExpr:
		c.walkFunc(x.LHS)
		c.walkFunc(x.RHS)
	case *SwitchExpr:
		c.walkFunc(x.exp)
		for i, n := 0, len(x.list); i < n; i++ {
			elem := x.list[i]
			c.walkFunc(x.exp)
			c.walkFunc(elem.when)
			c.walkFunc(elem.then)
		}
	case *CaseExpr:
		c.walkFunc(x.then)
	case *CallExpr:
		c.list = append(c.list, x)
	default:
		// default
		fmt.Printf("%v", x)
	}
}
