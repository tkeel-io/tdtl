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
	"testing"
)

var (
	TDTLSchema = []byte(`{
		    "id": "http://tkeel-io/tdtl/v1/schema.json",
		    "type": "object",
		    "definitions":{
		        "event":{
		            "type": "object",
		            "properties": {
		                "_version": {"type": "number"},
		                "ts": {"type": "number"},
		                "value": {"type": "number"},
		                "valueType": {"type": "string"},
		                "schemaType": {"type": "string","const":"event", "default": "event"}
		            },
		            "required": ["ts", "value", "valueType", "schemaType"]
		        },
		        "telemetry":{
	                "format": "telemetry",
		            "type": "object",
		            "properties": {
		                "_version": {"type": "number"},
		                "ts": {"type": "number"},
		                "value": {"type": "number"},
		                "schemaType": {"type": "string","const":"telemetry"}
		            }
		        },
		        "alarm": {
		            "type": "object",
		            "properties": {
		                "_version": {"type": "number"},
		                "ts": {"type": "number"},
		                "value": {"type": "number"},
		                "valueType": {"type": "string"},
		                "level": {"type": "string"},
		                "schemaType": {"type": "string","const":"event", "default": "event"}
		            },
		            "required": ["ts", "value", "valueType", "schemaType"]
		        }
		    }
		}`)

	TKeelSchema = []byte(`{
		    "id": "http://tkeel.io/v1/schema.json",
		    "type": "object",
		    "definitions":{
		        "event":{
		            "type": "object",
		            "properties": {
		                "_version": {"type": "number"},
		                "ts": {"type": "number"},
		                "value": {"type": "string"},
		                "schemaType": {"type": "string","const":"event", "default": "event"}
		            },
		            "required": ["ts", "value", "valueType", "schemaType"]
		        },
		        "telemetry":{
	                "format": "telemetry",
		            "type": "object",
		            "properties": {
		                "_version": {"type": "number"},
		                "ts": {"type": "number"},
		                "value": {"type": "number"},
		                "schemaType": {"type": "string","const":"telemetry"}
		            }
		        },
		        "alarm": {
		            "type": "object",
		            "properties": {
		                "_version": {"type": "number"},
		                "ts": {"type": "number"},
		                "value": {"type": "number"},
		                "valueType": {"type": "string"},
		                "level": {"type": "string"},
		                "schemaType": {"type": "string","const":"event", "default": "event"}
		            },
		            "required": ["ts", "value", "valueType", "schemaType"]
		        }
		    }
		}`)

	NameSpaceSchema = []byte(`{
		    "id": "http://tkeel.io/v1/schema/namespace1.json",
		    "type": "object",
		    "properties":{
		        "basicInfo":{
		            "type": "object",
		            "properties": {
						"selfLearn": {"type": "boolean"},
						"description": {
						  "type": "string",
						  "globalIndex": true
						},
						"name": {
						  "type": "string",
						  "globalSearch": true
						},
						"directConnection": {"type": "boolean"},
		                "_version": {"type": "number"},
		                "ts": {"type": "number"},
		                "value": {"type": "number"},
		                "valueType": {"type": "string"},
		                "schemaType": {"type": "string","const":"event", "default": "event"}
		            },
		            "required": ["ts", "value", "valueType", "schemaType"]
		        }
		    }
		}`)
	//
	// am.cpuHigh
	// telemetry.cpu
	// telemetry.modbus.aaa
	DeviceSchema = []byte(`{
		    "$id": "device.schema.json",
		    "type": "object",
			"allOf":[{"$ref": "http://tkeel.io/v1/schema/namespace1.json"}],
		    "properties": {
		        "am":{
		            "type": "object",
		            "properties": {
		                "cpuHigh": {
							"type": "object",
							"properties": {
							    "_version": {"type": "number"},
							    "ts": {"type": "number"},
							    "value": {"type": "string"},
							    "schemaType": {"type": "string","const":"event", "default": "event"}
		                        "local": {"type": "string"}
							},
		                }
		            }            
		        },
		        "telemetry": {
		          "type": "object",
		          "properties": {
		            "cpu": {
                        "allOf":[{"$ref": "/v1/schema.json#/definitions/telemetry"}],
		                "properties": {
		                    "code": {"type": "string"}
		                }
		            },
		            "mem": {
	                    "allOf":[{"$ref": "/v1/schema.json#/definitions/telemetry"}]
		            }
		          },
		          "additionalProperties": {"$ref": "http://tkeel.io/v1/schema.json#/definitions/telemetry"}
		        }
		    }
		}`)
	//
	//SimpleBaseSchema = []byte(`{
	//	    "id": "http://tkeel-io/tdtl/v1/schema.json",
	//	    "type": "object",
	//	    "definitions":{
	//	        "event":{
	//	            "type": "object",
	//	            "properties": {
	//	                "version": {"type": "number"},
	//	                "ts": {"type": "number"},
	//	                "value": {"type": "number"},
	//	                "valueType": {"type": "string"},
	//	                "schemaType": {"type": "string","const":"event", "default": "event"}
	//	            },
	//	            "required": ["ts", "value", "valueType", "schemaType"]
	//	        }
	//	    }
	//	}`)
	//SimpleSchema = []byte(`{
	//	    "id": "simple.schema.json",
	//	    "type": "object",
	//	    "properties": {
	//	        "telemetry": {
	//	          "type": "object",
	//	          "properties": {
	//	            "cpu": {
	//	                "format": "telemetry",
	//	                "allOf":[{"$ref": "http://tkeel-io/tdtl/v1/schema.json#/definitions/event"}],
	//	                "properties": {
	//	                    "desc": {"type": "string"}
	//	                }
	//	            },
	//	            "mem": {
	//	                "$ref": "http://tkeel-io/tdtl/v1/schema.json#/definitions/event",
	//					"properties": {
	//	                    "desc": {"type": "string"}
	//	                }
	//	            }
	//	          }
	//	        }
	//	    }
	//	}`)
)

func TestTelemetryChecker_IsFormat(t *testing.T) {
	ret := JSONSchemaPropsOrBool{}
	err := json.Unmarshal(TKeelSchema, &ret)
	t.Log(ret, err)
	NewSchema(ret.Schema)

	err = json.Unmarshal(NameSpaceSchema, &ret)
	t.Log(ret, err)
	nm := NewSchema(ret.Schema)
	parseSchema(nm.Collect)
	fmt.Println(nm.String())

	err = json.Unmarshal(DeviceSchema, &ret)
	t.Log(ret, err)
	sm := NewSchema(ret.Schema)
	parseSchema(sm.Collect)
	fmt.Println(sm.String())
}
//func TestTelemetryChecker_IsFormat2(t *testing.T) {
//	ret := JSONSchemaPropsOrBool{}
//	err := json.Unmarshal(SimpleBaseSchema, &ret)
//	t.Log(ret, err)
//	NewSchema(ret.Schema)
//	err = json.Unmarshal(SimpleSchema, &ret)
//	t.Log(ret, err)
//	sm := NewSchema(ret.Schema)
//	sm.parseSchema(sm.Collect)
//	fmt.Println(sm.String())
//}
