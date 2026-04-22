package tdtl

import "testing"

var benchCtx = NewJSONContext(JSONRaw.JSON)
var benchSimpleCtx = NewJSONContext(JSONRaw.SimpleJSON)

// BenchmarkCompile measures ANTLR parse time for a simple expression.
func BenchmarkCompile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Compile(`temperature * 2 + 1`)
	}
}

// BenchmarkEval measures evaluation time reusing a pre-compiled expression.
func BenchmarkEval(b *testing.B) {
	expr, _ := Compile(`temperature * 2 + 1`)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Eval(benchSimpleCtx, expr)
	}
}

// BenchmarkCompileAndEval measures the full parse+eval cycle.
func BenchmarkCompileAndEval(b *testing.B) {
	for i := 0; i < b.N; i++ {
		expr, _ := Compile(`temperature * 2 + 1`)
		Eval(benchSimpleCtx, expr)
	}
}

// BenchmarkEvalComplex measures evaluation of a multi-operator expression.
func BenchmarkEvalComplex(b *testing.B) {
	expr, _ := Compile(`(temperature + 1) * 2 - age / 3 + 0`)
	ctx := NewJSONContext(`{"temperature":50,"age":30}`)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Eval(ctx, expr)
	}
}

// BenchmarkEvalAggregation measures aggregation over a JSON array.
func BenchmarkEvalAggregation(b *testing.B) {
	expr, _ := Compile(`sum(friends.#.age)`)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Eval(benchCtx, expr)
	}
}

// BenchmarkEvalCaseWhen measures CASE/WHEN expression evaluation.
func BenchmarkEvalCaseWhen(b *testing.B) {
	expr, _ := Compile(`case temperature when 50 then 'hot' when 30 then 'warm' else 'cool'`)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Eval(benchSimpleCtx, expr)
	}
}

// BenchmarkNewJSONContext measures JSONContext construction cost.
func BenchmarkNewJSONContext(b *testing.B) {
	raw := JSONRaw.SimpleJSON
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = NewJSONContext(raw)
	}
}

// BenchmarkEvalFilter measures filter expression evaluation.
func BenchmarkEvalFilter(b *testing.B) {
	expr, _ := ParseFilter(`temperature > 30 and color = 'red'`)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		EvalFilter(benchSimpleCtx, expr)
	}
}

// BenchmarkEvalMultiContext measures evaluation against a MutilContext.
func BenchmarkEvalMultiContext(b *testing.B) {
	expr, _ := Compile(`entity1.temperature + entity2.offset`)
	ctx1 := NewJSONContext(`{"entity1":{"temperature":50}}`)
	ctx2 := NewJSONContext(`{"entity2":{"offset":5}}`)
	mc := MutilContext{ctx1, ctx2}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Eval(mc, expr)
	}
}
