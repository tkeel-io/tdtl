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
	"regexp"
)

var (
	pattern  = regexp.MustCompile(`\[(?P<key>\w+)\]`)
	template = `.$1`
)

type jsonContext struct {
	raw *Collect
}

//NewJSONContext new context from json
func NewJSONContext(jsonRaw string) Context {
	return &jsonContext{
		raw: New(jsonRaw),
	}
}

//Value get value from context
func (c *jsonContext) Value(path string) Node {
	if path == "" {
		return c.raw.Node()
	}
	return c.raw.Get(path).Node()
}

//Call call function from context
func (c *jsonContext) Call(expr *CallExpr, args []Node) Node {
	return UNDEFINED_RESULT
}
