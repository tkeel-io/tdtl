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
)

//Expr
type Expr interface {
	expr()
}

func (*SelectStatementExpr) expr() {}
func (FieldsExpr) expr()           {}
func (*FieldExpr) expr()           {}
func (TopicExpr) expr()            {}
func (*FilterExpr) expr()          {}
func (*BinaryExpr) expr()          {}
func (*JSONPathExpr) expr()        {}
func (*SwitchExpr) expr()          {}
func (CaseListExpr) expr()         {}
func (*CaseExpr) expr()            {}

func (BoolNode) expr()   {}
func (IntNode) expr()    {}
func (FloatNode) expr()  {}
func (StringNode) expr() {}
func (*CallExpr) expr()  {}
func (JSONNode) expr()   {}

//BinaryExpr
type BinaryExpr struct {
	Op  int
	LHS Expr
	RHS Expr
}

//JSONPathExpr xpath
type JSONPathExpr struct {
	val string
}

//CallExpr
type CallExpr struct {
	raw  string
	key  string
	args []Expr
}

func (e *CallExpr) String() string {
	return fmt.Sprintf("%s", e.raw)
}

func (e *CallExpr) FuncName() string {
	return fmt.Sprintf("%s", e.key)
}

func (e *CallExpr) Args() []Expr {
	return e.args
}

//SwitchExpr
type SwitchExpr struct {
	exp  Expr
	list []*CaseExpr
	last Expr
}

//CaseListExpr
type CaseListExpr []*CaseExpr

//CaseExpr
type CaseExpr struct {
	when Expr
	then Expr
}

//FieldExpr
type FieldExpr struct {
	exp   Expr
	alias string
}

func (r *FieldExpr) String() string {
	panic("todo")
}

//FieldExpr
type FieldsExpr []*FieldExpr

func (r FieldsExpr) String() string {
	panic("todo")
}

//TopicExpr
type TopicExpr []string

func (r TopicExpr) String() string {
	panic("todo")
}

//FilterExpr
type FilterExpr struct {
	exp Expr
}

func (r *FilterExpr) String() string {
	panic("todo")
}

//GroupExpr
type DimensionsExpr struct {
	exprs  []*JSONPathExpr
	window *WindowExpr
}

func (*DimensionsExpr) expr() {}

type DimensionExpr struct {
	exp Expr
}

func (*DimensionExpr) expr() {}

type WindowType int

const (
	NOT_WINDOW WindowType = iota
	TUMBLING_WINDOW
	HOPPING_WINDOW
	SLIDING_WINDOW
	SESSION_WINDOW
)

type WindowLength int

func (WindowLength) expr() {}

type WindowInterval int

func (WindowInterval) expr() {}

type WindowExpr struct {
	WindowType WindowType
	Length     WindowLength
	Interval   WindowInterval
}

func (WindowExpr) expr() {}

//SelectStatementExpr
type SelectStatementExpr struct {
	fields     FieldsExpr
	topic      TopicExpr
	filter     *FilterExpr
	dimensions *DimensionsExpr
}

func (r *SelectStatementExpr) String() string {
	return "Root Expr"
}
