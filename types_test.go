package tdtl

import (
	"testing"
)

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
