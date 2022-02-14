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

package ruleql

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/tkeel-io/core/pkg/tql"
)

var _ TQL = (*tkql)(nil)

type tkql struct {
	target    string
	sources   map[string][]string
	tentacles []tql.TentacleConfig
	listener  *QingQLListener
}

type TQL interface {
	Target() string
	Entities() map[string][]string
	Tentacles() []tql.TentacleConfig
	Exec(map[string]Node) (map[string]Node, error)
}

func NewTKQL(sql string) (TQL, error) {
	parse, listener := parse(sql)
	antlr.ParseTreeWalkerDefault.Walk(listener, parse.Root())
	err := listener.error()
	if err != nil {
		return nil, err
	}
	return &tkql{
		listener: listener,
		target:   listener.target,
		sources:  listener.sources,
		//fields:    listener.fields,
		tentacles: nil,
	}, nil
}

func (Q *tkql) Target() string {
	return Q.target
}

func (Q *tkql) Entities() map[string][]string {
	return Q.sources
}

func (Q *tkql) Tentacles() []tql.TentacleConfig {
	return Q.tentacles
}

func (Q *tkql) expr() Expr {
	return Q.listener.Expr()
}

func (Q *tkql) Exec(input map[string]Node) (map[string]Node, error) {
	ctx := NewMapContext(input, map[string]ContextFunc{})
	result := EvalRuleQL(ctx, Q.expr())
	retCtx := NewJSONContext(result.String())
	ret := map[string]Node{}
	for k, _ := range Q.listener.fields {
		ret[k] = retCtx.Value(k)
	}
	return ret, nil
}
