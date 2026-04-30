package tdtl_test

import (
	"fmt"

	"github.com/tkeel-io/tdtl"
)

// ExampleCompile demonstrates the recommended parse-once, eval-many pattern.
func ExampleCompile() {
	expr, err := tdtl.Compile(`(fahrenheit - 32) * 5 / 9`)
	if err != nil {
		panic(err)
	}

	for _, f := range []float64{32, 212, 98.6} {
		ctx := tdtl.NewJSONContext(fmt.Sprintf(`{"fahrenheit":%v}`, f))
		result := tdtl.Eval(ctx, expr)
		fmt.Printf("%.1f°F → %s°C\n", f, result.String())
	}
	// Output:
	// 32.0°F → 0°C
	// 212.0°F → 100°C
	// 98.6°F → 37.000000°C
}

// ExampleEval shows basic expression evaluation against a JSON payload.
func ExampleEval() {
	expr, _ := tdtl.Compile(`temperature + 1`)
	ctx := tdtl.NewJSONContext(`{"temperature":50}`)
	fmt.Println(tdtl.Eval(ctx, expr))
	// Output: 51
}

// ExampleNewJSONContext shows how to create a JSON-backed evaluation context.
func ExampleNewJSONContext() {
	ctx := tdtl.NewJSONContext(`{"color":"red","temp":72}`)

	expr, _ := tdtl.Compile(`temp > 70`)
	fmt.Println(tdtl.Eval(ctx, expr))
	// Output: true
}

// ExampleMutilContext shows multi-entity evaluation by chaining contexts.
func ExampleMutilContext() {
	ctx1 := tdtl.NewJSONContext(`{"device1":{"temp":50}}`)
	ctx2 := tdtl.NewJSONContext(`{"device2":{"offset":5}}`)
	mc := tdtl.MutilContext{ctx1, ctx2}

	expr, _ := tdtl.Compile(`device1.temp + device2.offset`)
	fmt.Println(tdtl.Eval(mc, expr))
	// Output: 55
}

// ExampleEvalFilter demonstrates boolean filter evaluation.
func ExampleEvalFilter() {
	ctx := tdtl.NewJSONContext(`{"temperature":75,"status":"ok"}`)
	expr, _ := tdtl.ParseFilter(`temperature > 70 and status = 'ok'`)
	fmt.Println(tdtl.EvalFilter(ctx, expr))
	// Output: true
}

// ExampleNewMapContext shows how to build a context from a Go map.
func ExampleNewMapContext() {
	ctx := tdtl.NewMapContext(
		map[string]tdtl.Node{"x": tdtl.IntNode(10), "y": tdtl.IntNode(3)},
		nil,
	)
	expr, _ := tdtl.Compile(`x * y`)
	fmt.Println(tdtl.Eval(ctx, expr))
	// Output: 30
}
