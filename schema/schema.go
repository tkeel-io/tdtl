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

package schema

import (
	"encoding/json"
	"fmt"
	//"github.com/google/cel-go/common/types/ref"
	"github.com/tkeel-io/tdtl"
	"net/url"
	"strings"
)

var _schemaLib = map[string]*tdtl.Collect{}

var _filterChain = []filterFn{
	schemaRefFilter,
	schemaAllOfFilter,
	schemaDenyFilter,
}

type filterFn func(schema *Schema, currentSchema *tdtl.Collect) (err error)

type Schema struct {
	*tdtl.Collect
	filterChain []filterFn
}

func NewSchema(schema *JSONSchemaProps) Schema {
	byt, _ := json.Marshal(schema)
	scm := Schema{
		Collect: tdtl.New(byt),
	}
	id := scm.Get("id").String()
	_schemaLib[id] = tdtl.New(byt)
	return scm
}

var denyFields = "anyOf|oneOf|additionalProperties"

func schemaDenyFilter(d *Schema, currentSchema *tdtl.Collect) (err error) {
	for _, key := range strings.Split(denyFields, "|") {
		node := currentSchema.Get(key)
		if node.Type() != tdtl.Null {
			return fmt.Errorf("find deny field:%s", key)
		}
	}
	return nil
}

func schemaAllOfFilter(d *Schema, currentSchema *tdtl.Collect) (err error) {
	ref := currentSchema.Get("allOf")
	if ref.Type() == tdtl.Array {
		ref.Foreach(func(key []byte, subSchema *tdtl.Collect) {
			currentSchema.Merge(subSchema)
		})
		currentSchema.Del("allOf")
	}
	return
}

func (d *Schema) parseReference(currentSchema *tdtl.Collect, refStr string) (*tdtl.Collect, error) {
	refURI, err := url.Parse(refStr)
	if err != nil {
		return nil, err
	}
	fullURI := fullURI(refURI)
	schema := currentSchema
	if fullURI != "" {
		ret, ok := _schemaLib[fullURI]
		if !ok {
			return nil, fmt.Errorf("can not find schema：%s,use %s", refStr, fullURI)
		}
		schema = ret
	}
	fragment := refURI.Fragment
	paths := strings.Split(fragment, "/")
	subSchema := schema.Get(paths...)
	return subSchema, nil
}

func fullURI(refURI *url.URL) string {
	if refURI.Host == "" {
		return refURI.Path
	}
	return fmt.Sprintf("%s://%s%s", refURI.Scheme, refURI.Host, refURI.Path)
}

func (d *Schema) String() string {
	return d.Collect.String()
}

func schemaRefFilter(d *Schema, currentSchema *tdtl.Collect) (err error) {
	ref := currentSchema.Get("$ref")
	if ref.Type() == tdtl.String {
		ret, err := d.parseReference(currentSchema, ref.String())
		if err != nil {
			return err
		}
		ret.Foreach(func(key []byte, value *tdtl.Collect) {
			currentSchema.Set(string(key), value)
		})
		currentSchema.Del("$ref")
	}
	return
}

func (d *Schema) parseSchema(currentSchema *tdtl.Collect) (err error) {
	currentSchema.Map(func(key []byte, value *tdtl.Collect) tdtl.Node {
		switch value.Type() {
		case tdtl.Object, tdtl.Array:
			d.parseSchema(value)
		case tdtl.JSON:
			fmt.Println("[+]JSON", string(key), value.String())
		case tdtl.Null:
			fmt.Println("[+]Null", string(key), value.String())
		}
		return value
	})
	for _, fn := range _filterChain {
		serr := fn(d, currentSchema)
		err = serr
		continue
	}

	return
}
