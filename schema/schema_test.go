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
	"reflect"
	"testing"
)

func Test_parseReference(t *testing.T) {
	ret := JSONSchemaPropsOrBool{}
	err := json.Unmarshal(TKeelSchema, &ret)
	t.Log(ret, err)
	currentSchema := NewSchema(ret.Schema)
	tests := []struct {
		name    string
		refStr  string
		want    string
		wantErr bool
	}{
		{"1", "http://tkeel-io/tdtl/v1/schema.json#/definitions/event"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseReference(currentSchema.Collect, tt.refStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseReference() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseReference() got = %v, want %v", got, tt.want)
			}
		})
	}
}
