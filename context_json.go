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
	"github.com/tkeel-io/tdtl/pkg/json/gjson"
	"regexp"
	"strings"
)

var (
	pattern  = regexp.MustCompile(`\[(?P<key>\w+)\]`)
	template = `.$1`
)

type jsonContext struct {
	raw string
}

//NewJSONContext new context from json
func NewJSONContext(jsonRaw string) Context {
	return &jsonContext{
		raw: jsonRaw,
	}
}

//Value get value from context
func (c *jsonContext) Value(path string) Node {
	if path == "" {
		return New(c.raw)
	}
	ret := gjson.Get(c.raw, thePath(path))
	switch ret.Type {
	case gjson.True:
		return BoolNode(true)
	case gjson.False:
		return BoolNode(false)
	case gjson.String:
		return StringNode(ret.Str)
	case gjson.Number:
		if strings.Index(ret.Raw, ".") != -1 {
			return FloatNode(ret.Num)
		}
		return IntNode(ret.Num)
	case gjson.JSON:
		return newCollectFromGjsonResult(ret)
	case gjson.Null:
		return UNDEFINED_RESULT
	}
	return UNDEFINED_RESULT
}

//Call call function from context
func (c *jsonContext) Call(expr *CallExpr, args []Node) Node {
	return UNDEFINED_RESULT
}

func thePath(path string) string {
	//return pattern.ReplaceAllString(path, template)
	path = strings.ReplaceAll(path, "[", ".")
	path = strings.ReplaceAll(path, "]", "")
	return path
}

//Value get value from context
func (c *jsonContext) Range(path string) Node {
	ret := gjson.Get(c.raw, thePath(path))
	return _gjson2JsonNode(ret)
}
