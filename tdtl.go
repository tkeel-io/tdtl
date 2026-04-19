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
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// CompiledExpr is an immutable, goroutine-safe expression produced by Compile.
// It can be evaluated concurrently by multiple goroutines via Eval.
type CompiledExpr = Expr

// Compile parses an expression string once and returns an immutable CompiledExpr.
// Prefer Compile over ParseExpr for production use: parse once, eval many times.
func Compile(exprStr string) (CompiledExpr, error) {
	return ParseExpr(exprStr)
}

// Eval evaluates a compiled expression against the provided context.
// DefaultValue built-ins (abs, base64, sum, avg, min, max, count) are always available.
func Eval(ctx Context, expr CompiledExpr) Node {
	return eval(MutilContext{DefaultValue, ctx}, expr)
}

var _ TDTL = (*tdtl)(nil)

type tdtl struct {
	target   string
	sources  map[string][]string
	listener *TDTLListener
	extFunc  map[string]ContextFunc
	fields   map[string]string
}

type TDTL interface {
	Target() string
	Entities() map[string][]string
	Fields() map[string]string
	Exec(map[string]Node) (map[string]Node, error)
}

func NewTDTL(sql string, extFunc map[string]ContextFunc) (TDTL, error) {
	parse, listener := parse(sql)
	antlr.ParseTreeWalkerDefault.Walk(listener, parse.Root())
	err := listener.error()
	if err != nil {
		return nil, err
	}
	return &tdtl{
		listener: listener,
		target:   listener.target,
		sources:  listener.sources,
		fields:   listener.fields,
		extFunc:  extFunc,
	}, nil
}

func (Q *tdtl) Target() string {
	return Q.target
}

func (Q *tdtl) Entities() map[string][]string {
	return Q.sources
}

func (Q *tdtl) Fields() map[string]string {
	return Q.fields
}

func (Q *tdtl) expr() Expr {
	return Q.listener.Expr()
}

func (Q *tdtl) Exec(input map[string]Node) (map[string]Node, error) {
	ctx := NewMapContext(input, Q.extFunc)
	result := EvalRuleQL(ctx, Q.expr())
	retCtx := NewJSONContext(result.String())
	ret := map[string]Node{}
	for k, _ := range Q.listener.fields {
		ret[k] = retCtx.Value(k)
	}
	return ret, nil
}
