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
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestType(t *testing.T) {
	{
		first := StringNode("Collect")
		last := StringNode("JS")
		fullname := first + last
		assert.Equal(t, fullname.String(), "CollectJS", "The two words should be the same.")
	}

	{
		a := FloatNode(1.1)
		b := FloatNode(0.1)
		c := a / b
		assert.Equal(t, c.String(), "11.000000", "The two words should be the same.")
	}

	{
		a := IntNode(11)
		b := IntNode(4)
		c := a % b
		assert.Equal(t, c.String(), "3", "The two words should be the same.")
	}
}

func BenchmarkStringNodeToNumber(b *testing.B) {
	s := StringNode("123.4")
	for i := 0; i < b.N; i++ {
		s.To(Number)
	}
}

func BenchmarkStringNodeToInt(b *testing.B) {
	s := StringNode("123.4")
	for i := 0; i < b.N; i++ {
		s.To(Int)
	}
}

func BenchmarkStringNodeToString(b *testing.B) {
	s := StringNode("123.4")
	for i := 0; i < b.N; i++ {
		s.To(String)
	}
}

func BenchmarkIntNodeToInt(b *testing.B) {
	s := IntNode(23)
	for i := 0; i < b.N; i++ {
		s.To(Int)
	}
}
