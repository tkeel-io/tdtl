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
	"encoding/base64"
	"fmt"
)

//ContextValuable  eval context
type ContextValuable interface {
	Value(key string) Node
}

//ContextCallable   eval context function
//type ContextCallable interface {
//	Call(key string, args []Node) Node
//}
type ContextCallable interface {
	Call(expr *CallExpr, args []Node) Node
}

type ContextCallableFunc func(args ...Node) Node

//Context eval context
type Context interface {
	ContextValuable
	ContextCallable
}

//ContextFunc eval context function
type ContextFunc func(args ...Node) Node

//DefaultValue default eval context
var DefaultValue = &value{
	functions: map[string]ContextFunc{
		"abs":    absFunc,
		"base64": base64Func,
	},
}

var absFunc = func(args ...Node) Node {
	if len(args) == 1 {
		arg := args[0]
		switch arg := arg.(type) {
		case IntNode:
			return absFuncHandle(arg)
		case FloatNode:
			return absFuncHandle(arg)
		case StringNode:
			return absFuncHandle(arg.To(Number))
		}
	}
	return UNDEFINED_RESULT
}

var absFuncHandle = func(args ...Node) Node {
	if len(args) == 1 {
		arg := args[0]
		switch arg := arg.(type) {
		case IntNode:
			if int64(arg) < 0 {
				return -1 * arg
			}
			return arg
		case FloatNode:
			if int64(arg) < 0 {
				return -1 * arg
			}
			return arg
		}
	}
	return UNDEFINED_RESULT
}

var base64Func = func(args ...Node) Node {
	if len(args) < 1 {
		return UNDEFINED_RESULT
	}
	arg := args[0]
	node := ""
	switch arg := arg.(type) {
	case JSONNode:
		node = arg.String()
	case StringNode:
		node = arg.String()
	default:
		fmt.Println(
			"[-]ruleql type error",
		)
		return UNDEFINED_RESULT
	}
	if node == "" {
		fmt.Println(
			"[-]ruleql empty error",
		)
		return UNDEFINED_RESULT
	}
	decodeBytes := base64.StdEncoding.EncodeToString([]byte(node))
	return StringNode(decodeBytes)
}

type value struct {
	functions map[string]ContextFunc
}

func (v *value) Value(key string) Node {
	return UNDEFINED_RESULT
}

func (v *value) Call(expr *CallExpr, args []Node) Node {
	if v.functions == nil {
		return UNDEFINED_RESULT
	}
	if fn, ok := v.functions[expr.key]; ok {
		return fn(args...)
	}

	return UNDEFINED_RESULT
}

//MutilContext mutil-context
type MutilContext []Context

//Value get value from context
func (mc MutilContext) Value(key string) Node {
	for _, v := range mc {
		r := v.Value(key)
		t:= r.Type()
		if t != Undefined {
			return r
		}
	}
	return UNDEFINED_RESULT
}

//Call call function from context
func (mc MutilContext) Call(expr *CallExpr, args []Node) Node {
	for _, v := range mc {
		r := v.Call(expr, args)
		if r.Type() != Undefined {
			return r
		}
	}
	return UNDEFINED_RESULT
}
