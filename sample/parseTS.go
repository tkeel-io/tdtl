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

package main

import (
	"fmt"

	"github.com/tkeel-io/collectjs"
	"github.com/tkeel-io/tdtl"
)

var input = `{
	"ts": 1483228800000,
	"values": {
		"temperature": 42,
		"humidity": 80
	}
}`

func main() {
	tqlString := `insert into entity3 select json(telemetryV1(entity1.input), 'values') as property1, entity2.property2.name as property2, entity1.property1 + entity2.property3 as property3`

	tqlInst, err := tdtl.NewTDTL(tqlString, map[string]tdtl.ContextFunc{
		"json": func(args ...tdtl.Node) tdtl.Node {
			if len(args) != 2 {
				return tdtl.NULL_RESULT
			}
			jsonRaw := args[0].String()
			path := args[1].String()

			cc := collectjs.New(jsonRaw)
			return tdtl.JSONNode(string(cc.Get(path).GetRaw()))
		},
		"telemetryV1": func(args ...tdtl.Node) tdtl.Node {
			if len(args) != 1 {
				return tdtl.NULL_RESULT
			}
			payload, ok := args[0].(tdtl.StringNode)
			if !ok {
				return tdtl.NULL_RESULT
			}

			cc := collectjs.New(payload.String())
			ts := cc.Get("ts")
			vals := cc.Get("values")
			vals.Map(func(key []byte, value []byte) []byte {
				val := collectjs.New("{}")
				val.Set("timestamp", ts.GetRaw())
				val.Set("value", value)
				return val.GetRaw()
			})
			cc.Set("values", vals.GetRaw())
			return tdtl.JSONNode(string(cc.GetRaw()))
		},
	})
	if nil != err {
		panic(err)
	}

	fmt.Println("target: ", tqlInst.Target())
	fmt.Println("sources: ", tqlInst.Entities())
	for _, tentacle := range tqlInst.Tentacles() {
		fmt.Println("tentacle: ", tentacle)
	}

	result, err := tqlInst.Exec(map[string]tdtl.Node{
		"entity1.input":          tdtl.StringNode(input),
		"entity1.property1":      tdtl.StringNode("test"),
		"entity2.property2.name": tdtl.StringNode("123"),
		"entity2.property3":      tdtl.StringNode("g123"),
	})

	fmt.Println(result)

}
