/*
 * Copyright (C) 2019 Yunify, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this work except in compliance with the License.
 * You may obtain a copy of the License in the LICENSE file, or at:
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package ruleql

var (
	JSONRaw = struct {
		SimpleJSON string
		JSON       string
		EventJSON  string
		EmptyJSON  string
	}{
		SimpleJSON: `{
			"YX_0002":1, "YX_0003":2,
			"color":"red", "temperature":50, 
			"metadata": {"name": "Light1", "price": 11.05}
		}`,
		JSON: `{
  			"name": {"first": "Tom", "last": "Anderson"},
  			"age":37, "temperature":50, 
  			"children": ["Sara","Alex","Jack"],
  			"fav.movie": "Deer Hunter",
  			"movie": {"1111": [{"1111": "Tom", "last": "Anderson"}], "last": "Anderson"},
  			"friends": [
  				{"first": "Dale", "last": "Murphy", "age": 44},
  				{"first": "Roger", "last": "Craig", "age": 68},
  				{"first": "Jane", "last": "Murphy", "age": 47}
  			]
		}`,
		EventJSON: `{
  			"id": "10000000000000000000000000DF0EF6",
  			"version": "1.0",
  			"params": {
	    		"Power": {
    	  		  "id": "10000000000000000000000000000001", //32
      			  "type": "text",
      			  "value": "on",
      		  	"time": 1524448722000
    			},
    			"Color": {
  	    		  "id": "10000000000000000000000000000001", //32
    	  		  "type": "text",
      			  "value": "red",
      			  "time": 1524448722000
    			},
    			"Temperature": {
  	    		  "id": "10000000000000000000000000000002", //32
    	  		  "type": "float",
      			  "value": 50.1,
      			  "time": 1524448722000
    			}
  			}
    	}`,
		EmptyJSON: `{}`,
	}
)
