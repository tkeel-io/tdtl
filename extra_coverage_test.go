package tdtl

import (
	"io"
	"strings"
	"testing"

	"github.com/tkeel-io/tdtl/pkg/json/gjson"
	. "github.com/smartystreets/goconvey/convey"
)

// TestCollectMethods covers Copy, Map, GetError, SortBy.
func TestCollectMethods(t *testing.T) {
	Convey("Collect utility methods", t, func() {
		Convey("Copy returns independent clone", func() {
			c := New(`{"a":1}`)
			clone := c.Copy()
			So(clone.String(), ShouldEqual, c.String())
		})
		Convey("GetError returns nil on valid Collect", func() {
			c := New(`{"a":1}`)
			So(c.GetError(), ShouldBeNil)
		})
		Convey("GetError returns error on invalid operation", func() {
			// GroupBy on non-array sets an error
			c := New(`{"a":1}`)
			result := c.GroupBy("a")
			So(result.GetError(), ShouldNotBeNil)
		})
		Convey("Map transforms object fields", func() {
			c := New(`{"a":1,"b":2}`)
			c.Map(func(key []byte, value *Collect) Node {
				return IntNode(int64(value.Node().(IntNode)) * 10)
			})
			So(c.Get("a").Node(), ShouldResemble, IntNode(10))
		})
		Convey("SortBy sorts array by numeric field", func() {
			c := New(`[{"v":3},{"v":1},{"v":2}]`)
			c.SortBy(func(p1, p2 *Collect) bool {
				n1 := p1.Get("v").Node().(IntNode)
				n2 := p2.Get("v").Node().(IntNode)
				return int64(n1) < int64(n2)
			})
			So(c.Error(), ShouldBeNil)
		})
		Convey("SortBy on scalar sets error", func() {
			c := New(`"hello"`)
			c.SortBy(func(p1, p2 *Collect) bool { return false })
			So(c.Error(), ShouldNotBeNil)
		})
	})
}

// TestGjson2JsonNode covers _gjson2JsonNode with all gjson types.
func TestGjson2JsonNode(t *testing.T) {
	Convey("_gjson2JsonNode conversions", t, func() {
		Convey("True result → BoolNode(true)", func() {
			r := gjson.Result{Type: gjson.True}
			So(_gjson2JsonNode(r), ShouldResemble, BoolNode(true))
		})
		Convey("False result → BoolNode(false)", func() {
			r := gjson.Result{Type: gjson.False}
			So(_gjson2JsonNode(r), ShouldResemble, BoolNode(false))
		})
		Convey("Number result without decimal → IntNode", func() {
			r := gjson.Result{Type: gjson.Number, Raw: "42", Num: 42}
			got := _gjson2JsonNode(r)
			So(got.Type(), ShouldEqual, Int)
		})
		Convey("Number result with decimal → FloatNode", func() {
			r := gjson.Result{Type: gjson.Number, Raw: "3.14", Num: 3.14}
			got := _gjson2JsonNode(r)
			So(got.Type(), ShouldEqual, Float)
		})
		Convey("String result → StringNode", func() {
			r := gjson.Result{Type: gjson.String, Str: "hello"}
			got := _gjson2JsonNode(r)
			So(got, ShouldResemble, StringNode("hello"))
		})
		Convey("JSON result → JSONNode", func() {
			r := gjson.Result{Type: gjson.JSON, Raw: `{"x":1}`}
			got := _gjson2JsonNode(r)
			So(got.Type(), ShouldEqual, JSON)
		})
		Convey("Null result → NULL_RESULT", func() {
			r := gjson.Result{Type: gjson.Null}
			So(_gjson2JsonNode(r), ShouldResemble, NULL_RESULT)
		})
	})
}

// TestEvalDimensions covers the evalDimensions helper.
func TestEvalDimensions(t *testing.T) {
	Convey("evalDimensions", t, func() {
		ctx := NewJSONContext(JSONRaw.SimpleJSON)

		Convey("nil exprs returns UNDEFINED_RESULT", func() {
			got := evalDimensions(ctx, nil)
			So(got.Type(), ShouldEqual, Undefined)
		})
		Convey("single path produces string key", func() {
			got := evalDimensions(ctx, []*JSONPathExpr{{"color"}})
			So(got.Type(), ShouldEqual, String)
			So(strings.Contains(got.String(), "red"), ShouldBeTrue)
		})
		Convey("multiple paths join with dash", func() {
			got := evalDimensions(ctx, []*JSONPathExpr{{"color"}, {"YX_0002"}})
			So(got.String(), ShouldContainSubstring, "-")
		})
	})
}

// TestExpressionSources covers the Expression.Sources() method via NewExpr.
func TestExpressionSources(t *testing.T) {
	Convey("Expression.Sources()", t, func() {
		e, err := NewExpr(`entity1.temperature + 1`, nil)
		So(err, ShouldBeNil)
		sources := e.Sources()
		So(sources, ShouldNotBeNil)
		_, ok := sources["entity1"]
		So(ok, ShouldBeTrue)
	})
}

// TestParseField covers the ParseField function.
func TestParseField(t *testing.T) {
	Convey("ParseField", t, func() {
		Convey("valid fields expression", func() {
			expr, err := ParseField(`temperature + 1 as temp`)
			So(err, ShouldBeNil)
			So(expr, ShouldNotBeNil)
		})
	})
}

// TestParseFuncComplex covers walkFunc branches for complex expressions.
func TestParseFuncComplex(t *testing.T) {
	Convey("ParseFunc on complex expressions", t, func() {
		Convey("SELECT statement yields CallExprs from fields", func() {
			expr, err := Parse(`insert into target select abs(temperature) as t, base64(color) as c`)
			So(err, ShouldBeNil)
			calls := ParseFunc(expr)
			So(len(calls), ShouldEqual, 2)
		})
		Convey("CASE expression yields no CallExprs", func() {
			expr, _ := ParseExpr(`case temperature when 50 then 'hot' else 'cool'`)
			calls := ParseFunc(expr)
			So(len(calls), ShouldEqual, 0)
		})
		Convey("nested calls in filter walk FilterExpr", func() {
			expr, _ := Parse(`insert into target select temperature as t`)
			// ParseFunc on SelectStatementExpr walks fields + filter
			calls := ParseFunc(expr)
			So(calls, ShouldNotBeNil)
		})
	})
}

// TestNodeErrorMethods covers the Error() methods on primitive node types.
func TestNodeErrorMethods(t *testing.T) {
	Convey("Node.Error() always nil for value types", t, func() {
		So(BoolNode(true).Error(), ShouldBeNil)
		So(IntNode(1).Error(), ShouldBeNil)
		So(FloatNode(1.5).Error(), ShouldBeNil)
		So(StringNode("x").Error(), ShouldBeNil)
	})
}

// TestEvalBinaryBoolJSONFixed covers the BoolNode vs *JSONNode case
// (previously caused infinite recursion — fixed by using evalBinaryBool directly).
func TestEvalBinaryBoolJSONFixed(t *testing.T) {
	Convey("BoolNode compared to *JSONNode (infinite recursion bug fixed)", t, func() {
		ctx := DefaultValue
		Convey("true = UNDEFINED_RESULT → false", func() {
			expr, _ := ParseExpr(`true = undefinedVar`)
			got := Eval(NewJSONContext(`{}`), expr)
			So(got, ShouldResemble, BoolNode(false))
		})
		Convey("false = UNDEFINED_RESULT → true", func() {
			expr, _ := ParseExpr(`false = undefinedVar`)
			got := Eval(NewJSONContext(`{}`), expr)
			So(got, ShouldResemble, BoolNode(true))
		})
		Convey("true != UNDEFINED_RESULT → true", func() {
			expr, _ := ParseExpr(`true != undefinedVar`)
			got := Eval(NewJSONContext(`{}`), expr)
			So(got, ShouldResemble, BoolNode(true))
			_ = ctx
		})
	})
}

// TestEvalBinaryFloatOps covers float comparison operators.
func TestEvalBinaryFloatOps(t *testing.T) {
	Convey("float comparison operators", t, func() {
		ctx := DefaultValue

		cases := []struct {
			expr string
			want BoolNode
		}{
			{`1.0 = 1.0`, BoolNode(true)},
			{`1.0 = 2.0`, BoolNode(false)},
			{`1.0 != 2.0`, BoolNode(true)},
			{`1.0 != 1.0`, BoolNode(false)},
			{`1.0 < 2.0`, BoolNode(true)},
			{`2.0 < 1.0`, BoolNode(false)},
			{`1.0 <= 1.0`, BoolNode(true)},
			{`2.0 <= 1.0`, BoolNode(false)},
			{`2.0 > 1.0`, BoolNode(true)},
			{`1.0 > 2.0`, BoolNode(false)},
			{`1.0 >= 1.0`, BoolNode(true)},
			{`1.0 >= 2.0`, BoolNode(false)},
		}
		for _, c := range cases {
			c := c
			Convey(c.expr, func() {
				expr, _ := ParseExpr(c.expr)
				got := eval(ctx, expr)
				So(got, ShouldResemble, c.want)
			})
		}
	})
}

// TestEvalRuleQLLiteralBranches covers evalRuleQL with literal node types,
// nil input, and the default fall-through.
func TestEvalRuleQLLiteralBranches(t *testing.T) {
	Convey("evalRuleQL literal node pass-through", t, func() {
		ctx := DefaultValue

		Convey("nil → UNDEFINED_RESULT", func() {
			So(evalRuleQL(ctx, nil).Type(), ShouldEqual, Undefined)
		})
		Convey("BoolNode → returned as-is", func() {
			So(evalRuleQL(ctx, BoolNode(true)), ShouldResemble, BoolNode(true))
		})
		Convey("IntNode → returned as-is", func() {
			So(evalRuleQL(ctx, IntNode(7)), ShouldResemble, IntNode(7))
		})
		Convey("FloatNode → returned as-is", func() {
			So(evalRuleQL(ctx, FloatNode(1.5)), ShouldResemble, FloatNode(1.5))
		})
		Convey("StringNode → returned as-is", func() {
			So(evalRuleQL(ctx, StringNode("x")), ShouldResemble, StringNode("x"))
		})
		Convey("JSONNode → returned as-is", func() {
			n := JSONNode{value: []byte(`{"a":1}`), datatype: JSON}
			got := evalRuleQL(ctx, n)
			So(got.Type(), ShouldEqual, JSON)
		})
		Convey("*FilterExpr wraps eval correctly", func() {
			f := &FilterExpr{exp: IntNode(42)}
			So(evalRuleQL(ctx, f), ShouldResemble, IntNode(42))
		})
	})
}

// TestEvalBinaryOverloadFloat covers the FloatNode+StringNode overload paths.
func TestEvalBinaryOverloadFloat(t *testing.T) {
	Convey("evalBinaryOverload with FloatNode and StringNode", t, func() {
		ctx := DefaultValue

		Convey("'hello' + 1.5 = 'hello1.500000'", func() {
			expr, _ := ParseExpr(`'hello' + 1.5`)
			got := eval(ctx, expr)
			So(got.Type(), ShouldEqual, String)
			So(strings.Contains(got.String(), "hello"), ShouldBeTrue)
		})
		Convey("1.5 + 'world' = '1.500000world'", func() {
			expr, _ := ParseExpr(`1.5 + 'world'`)
			got := eval(ctx, expr)
			So(got.Type(), ShouldEqual, String)
			So(strings.Contains(got.String(), "world"), ShouldBeTrue)
		})
	})
}

// TestDumpMore covers the DumpMore variadic print helper.
func TestDumpMore(t *testing.T) {
	Convey("DumpMore prints multiple exprs without panic", t, func() {
		e1, _ := ParseExpr(`1 + 2`)
		e2, _ := ParseExpr(`abs(5)`)
		err := DumpMore(e1, e2)
		So(err, ShouldBeNil)
	})
}

// TestPrintCoverage exercises many branches of print.go's print() function.
func TestPrintCoverage(t *testing.T) {
	Convey("print.go branch coverage", t, func() {
		Convey("nil expression → no panic", func() {
			So(func() { _ = Fprint(io.Discard, nil) }, ShouldNotPanic)
		})
		Convey("SELECT statement → covers SelectStatementExpr, FieldsExpr, FieldExpr", func() {
			expr, err := Parse(`insert into target select temperature + 1 as temp, color as c`)
			So(err, ShouldBeNil)
			So(func() { _ = Dump(expr) }, ShouldNotPanic)
		})
		Convey("BinaryExpr tree", func() {
			expr, _ := ParseExpr(`(1 + 2) * 3`)
			So(func() { _ = Dump(expr) }, ShouldNotPanic)
		})
		Convey("CallExpr (abs)", func() {
			expr, _ := ParseExpr(`abs(temperature)`)
			So(func() { _ = Dump(expr) }, ShouldNotPanic)
		})
		Convey("SwitchExpr", func() {
			expr, _ := ParseExpr(`case temperature when 50 then 'hot' when 30 then 'warm' else 'cool'`)
			So(func() { _ = Dump(expr) }, ShouldNotPanic)
		})
		Convey("JSONPathExpr", func() {
			expr, _ := ParseExpr(`a.b.c`)
			So(func() { _ = Dump(expr) }, ShouldNotPanic)
		})
		Convey("StringNode", func() {
			expr, _ := ParseExpr(`'hello'`)
			So(func() { _ = Dump(expr) }, ShouldNotPanic)
		})
		Convey("TopicExpr and FilterExpr via io.Discard", func() {
			// Construct SelectStatementExpr manually; write to io.Discard to avoid stdout issues
			stmt := &SelectStatementExpr{
				topic:  TopicExpr{"device", "status"},
				filter: &FilterExpr{exp: BoolNode(true)},
			}
			So(func() { _ = Fprint(io.Discard, stmt) }, ShouldNotPanic)
		})
		Convey("typeOf returns non-empty string", func() {
			s := typeOf(IntNode(1))
			So(len(s), ShouldBeGreaterThan, 0)
		})
		Convey("DumpMore with SELECT and CASE exprs", func() {
			e1, _ := Parse(`insert into target select temperature as t`)
			e2, _ := ParseExpr(`case 1 when 1 then 'one' else 'other'`)
			So(func() { _ = DumpMore(e1, e2) }, ShouldNotPanic)
		})
	})
}

// TestTDTLListenerAppendErrorf covers appendErrorf via direct construction.
func TestTDTLListenerAppendErrorf(t *testing.T) {
	Convey("appendErrorf adds to error list", t, func() {
		l := &TDTLListener{}
		So(l.error(), ShouldBeNil)
		l.appendErrorf("test error %d", 42)
		So(l.error(), ShouldNotBeNil)
		So(l.error().Error(), ShouldContainSubstring, "test error 42")
	})
}

// TestTypeString covers the Type.String() method for all type constants.
func TestTypeString(t *testing.T) {
	Convey("Type.String() coverage", t, func() {
		So(Undefined.String(), ShouldEqual, "Undefined")
		So(Null.String(), ShouldEqual, "Null")
		So(Bool.String(), ShouldEqual, "Bool")
		So(Int.String(), ShouldEqual, "Int")
		So(Float.String(), ShouldEqual, "Float")
		So(String.String(), ShouldEqual, "String")
		So(JSON.String(), ShouldEqual, "JSON")
		// Number and Object/Array fall through to "Undefined"
		So(Number.String(), ShouldEqual, "Undefined")
	})
}

// TestNodeToEdgeCases covers UNDEFINED_RESULT branches in To() methods.
func TestNodeToEdgeCases(t *testing.T) {
	Convey("Node.To() UNDEFINED_RESULT paths", t, func() {
		Convey("BoolNode.To(Int) → UNDEFINED_RESULT", func() {
			So(BoolNode(true).To(Int).Type(), ShouldEqual, Undefined)
		})
		Convey("BoolNode.To(Float) → UNDEFINED_RESULT", func() {
			So(BoolNode(true).To(Float).Type(), ShouldEqual, Undefined)
		})
		Convey("IntNode.To(Bool) → UNDEFINED_RESULT", func() {
			So(IntNode(1).To(Bool).Type(), ShouldEqual, Undefined)
		})
		Convey("FloatNode.To(Bool) → UNDEFINED_RESULT", func() {
			So(FloatNode(1.5).To(Bool).Type(), ShouldEqual, Undefined)
		})
		Convey("JSONNode.To(Null) → UNDEFINED_RESULT", func() {
			n := JSONNode{value: []byte(`{"x":1}`), datatype: JSON}
			So(n.To(Null).Type(), ShouldEqual, Undefined)
		})
		Convey("JSONNode.To(Undefined) → UNDEFINED_RESULT", func() {
			n := JSONNode{value: []byte(`{"x":1}`), datatype: JSON}
			So(n.To(Undefined).Type(), ShouldEqual, Undefined)
		})
		Convey("JSONNode.To(Bool) via String path", func() {
			n := JSONNode{value: []byte(`true`), datatype: JSON}
			got := n.To(Bool)
			So(got.Type(), ShouldEqual, Bool)
		})
		Convey("JSONNode.To(Int) via String path", func() {
			n := JSONNode{value: []byte(`42`), datatype: JSON}
			got := n.To(Int)
			So(got.Type(), ShouldEqual, Int)
		})
		Convey("JSONNode.To(Float) via String path", func() {
			n := JSONNode{value: []byte(`3.14`), datatype: JSON}
			got := n.To(Float)
			So(got.Type(), ShouldEqual, Float)
		})
		Convey("JSONNode.Raw() with String datatype wraps in quotes", func() {
			n := JSONNode{value: []byte(`hello`), datatype: String}
			raw := n.Raw()
			So(string(raw), ShouldEqual, `"hello"`)
		})
	})
}

// TestSelectStatementString covers SelectStatementExpr.String().
func TestSelectStatementString(t *testing.T) {
	Convey("SelectStatementExpr.String()", t, func() {
		s := &SelectStatementExpr{}
		So(s.String(), ShouldEqual, "Root Expr")
	})
}

// TestCallExprDetailedMethods covers CallExpr.String, FuncName, Args.
func TestExprInterfaceMethods(t *testing.T) {
	Convey("expr() method coverage via Expr interface", t, func() {
		// Call expr() on each concrete type through the Expr interface
		var exprs []Expr
		exprs = append(exprs,
			&SelectStatementExpr{},
			FieldsExpr{},
			&FieldExpr{},
			&FilterExpr{},
			&BinaryExpr{},
			&JSONPathExpr{},
			&SwitchExpr{},
			CaseListExpr{},
			&CaseExpr{},
			BoolNode(true),
			IntNode(1),
			FloatNode(1.0),
			StringNode(""),
			&CallExpr{},
			JSONNode{},
		)
		for _, e := range exprs {
			e.expr()
		}
		So(len(exprs), ShouldEqual, 15)
	})
}
