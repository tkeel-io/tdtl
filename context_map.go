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
)

type mapContext struct {
	values    map[string]Node
	functions map[string]ContextFunc
}

//NewMapContext new context from map
func NewMapContext(values map[string]Node, functions map[string]ContextFunc) Context {
	if values == nil {
		values = make(map[string]Node)
	}
	if functions == nil {
		functions = make(map[string]ContextFunc)
	}
	return mapContext{
		values:    values,
		functions: functions,
	}
}

//Value get value from context
func (c mapContext) Value(key string) Node {
	if ret, ok := c.values[key]; ok {
		return ret
	}
	return UNDEFINED_RESULT
}

//Call call function from context
func (c mapContext) Call(expr *CallExpr, args []Node) Node {
	if ret, ok := c.functions[expr.key]; ok {
		return ret(args...)
	}
	return UNDEFINED_RESULT
}
