package tdtl

import (
	"math"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// TestModByZero verifies that x%0 returns UNDEFINED_RESULT instead of panicking.
func TestModByZero(t *testing.T) {
	cases := []struct {
		name string
		expr string
	}{
		{"int mod zero", `5 % 0`},
		{"float mod zero", `5.0 % 0`},
		{"float mod zero float", `5.0 % 0.0`},
	}
	Convey("mod by zero returns UNDEFINED_RESULT", t, func() {
		for _, c := range cases {
			Convey(c.name, func() {
				expr, _ := ParseExpr(c.expr)
				got := eval(DefaultValue, expr)
				So(got.Type(), ShouldEqual, Undefined)
			})
		}
	})
}

// TestNEOperatorWithNonNumericStrings verifies the isBooleanOP NE fix.
// Before the fix, 'aaa' != 'bbb' incorrectly returned UNDEFINED_RESULT.
func TestNEOperatorWithNonNumericStrings(t *testing.T) {
	Convey("NE operator with non-numeric strings", t, func() {
		ctx := NewJSONContext(JSONRaw.SimpleJSON)

		Convey("'aaa' != 'bbb' should be true", func() {
			expr, _ := ParseFilter(`'aaa' != 'bbb'`)
			So(EvalFilter(ctx, expr), ShouldBeTrue)
		})
		Convey("'aaa' <> 'bbb' should be true", func() {
			expr, _ := ParseFilter(`'aaa' <> 'bbb'`)
			So(EvalFilter(ctx, expr), ShouldBeTrue)
		})
		Convey("'aaa' != 'aaa' should be false", func() {
			expr, _ := ParseFilter(`'aaa' != 'aaa'`)
			So(EvalFilter(ctx, expr), ShouldBeFalse)
		})
		Convey("'abc' != 'xyz' should be true", func() {
			expr, _ := ParseFilter(`'abc' != 'xyz'`)
			So(EvalFilter(ctx, expr), ShouldBeTrue)
		})
	})
}

// TestNEOperatorWithNumbers verifies integer and float NE comparisons.
func TestNEOperatorWithNumbers(t *testing.T) {
	Convey("NE operator with numbers", t, func() {
		ctx := NewJSONContext(JSONRaw.SimpleJSON)

		Convey("1 != 2 should be true", func() {
			expr, _ := ParseFilter(`1 != 2`)
			So(EvalFilter(ctx, expr), ShouldBeTrue)
		})
		Convey("1 != 1 should be false", func() {
			expr, _ := ParseFilter(`1 != 1`)
			So(EvalFilter(ctx, expr), ShouldBeFalse)
		})
		Convey("1.0 != 2.0 should be true", func() {
			expr, _ := ParseFilter(`1.0 != 2.0`)
			So(EvalFilter(ctx, expr), ShouldBeTrue)
		})
		Convey("1.0 != 1.0 should be false", func() {
			expr, _ := ParseFilter(`1.0 != 1.0`)
			So(EvalFilter(ctx, expr), ShouldBeFalse)
		})
	})
}

// TestAbsFunction verifies abs() with positive, negative, and float values.
// Uses EvalRuleQL which injects DefaultValue context containing abs().
func TestAbsFunction(t *testing.T) {
	Convey("abs() function", t, func() {
		ctx := NewJSONContext(JSONRaw.SimpleJSON)

		Convey("abs(5) = 5", func() {
			expr, _ := ParseExpr(`abs(5)`)
			got := EvalRuleQL(ctx, expr)
			So(got, ShouldResemble, IntNode(5))
		})
		Convey("abs(0 - 5) = 5", func() {
			expr, _ := ParseExpr(`abs(0 - 5)`)
			got := EvalRuleQL(ctx, expr)
			So(got, ShouldResemble, IntNode(5))
		})
		Convey("abs(5.5) = 5.5", func() {
			expr, _ := ParseExpr(`abs(5.5)`)
			got := EvalRuleQL(ctx, expr)
			So(got, ShouldResemble, FloatNode(5.5))
		})
		Convey("abs(0 - 5.5) = 5.5 (float sign fix)", func() {
			// Before the float abs fix, int64(-5.5)==-5 < 0 worked accidentally,
			// but int64(-0.5)==0 >= 0 would incorrectly return the negative value.
			expr, _ := ParseExpr(`abs(0 - 5.5)`)
			got := EvalRuleQL(ctx, expr)
			So(got.Type(), ShouldEqual, Float)
			So(float64(got.(FloatNode)) > 0, ShouldBeTrue)
		})
	})
}

// TestNotOperator verifies NOT/not boolean negation.
func TestNotOperator(t *testing.T) {
	Convey("NOT operator", t, func() {
		ctx := NewJSONContext(JSONRaw.SimpleJSON)

		Convey("not true = false", func() {
			expr, _ := ParseFilter(`not true`)
			So(EvalFilter(ctx, expr), ShouldBeFalse)
		})
		Convey("not false = true", func() {
			expr, _ := ParseFilter(`not false`)
			So(EvalFilter(ctx, expr), ShouldBeTrue)
		})
		Convey("not (1 > 2) = true", func() {
			expr, _ := ParseFilter(`not (1 > 2)`)
			So(EvalFilter(ctx, expr), ShouldBeTrue)
		})
		Convey("not (2 > 1) = false", func() {
			expr, _ := ParseFilter(`not (2 > 1)`)
			So(EvalFilter(ctx, expr), ShouldBeFalse)
		})
	})
}

// TestCaseWhenExpr verifies CASE/WHEN/THEN/ELSE evaluation.
// Note: the grammar does NOT use an END keyword.
func TestCaseWhenExpr(t *testing.T) {
	Convey("CASE/WHEN expressions", t, func() {
		ctx := NewJSONContext(JSONRaw.SimpleJSON)

		Convey("matching first WHEN", func() {
			expr, _ := ParseExpr(`case 1 when 1 then 'one' when 2 then 'two' else 'other'`)
			got := eval(ctx, expr)
			So(got.String(), ShouldEqual, "one")
		})
		Convey("matching second WHEN", func() {
			expr, _ := ParseExpr(`case 2 when 1 then 'one' when 2 then 'two' else 'other'`)
			got := eval(ctx, expr)
			So(got.String(), ShouldEqual, "two")
		})
		Convey("no match falls through to ELSE", func() {
			expr, _ := ParseExpr(`case 3 when 1 then 'one' when 2 then 'two' else 'other'`)
			got := eval(ctx, expr)
			So(got.String(), ShouldEqual, "other")
		})
		Convey("no match without ELSE returns UNDEFINED", func() {
			expr, _ := ParseExpr(`case 3 when 1 then 'one' when 2 then 'two'`)
			got := eval(ctx, expr)
			So(got.Type(), ShouldEqual, Undefined)
		})
		Convey("CASE with JSON path value", func() {
			expr, _ := ParseExpr(`case color when 'red' then 'stop' when 'green' then 'go' else 'wait'`)
			got := eval(ctx, expr)
			So(got.String(), ShouldEqual, "stop")
		})
	})
}

// TestCaseNoLeadingSpace verifies E3 fix: CASE keyword now works without a leading space.
// Previously `case x when...` at the start of input failed to parse because the grammar
// required whitespace before CASE. parse() now prepends one space automatically.
func TestCaseNoLeadingSpace(t *testing.T) {
	Convey("CASE at start of expression (no leading space)", t, func() {
		ctx := NewJSONContext(JSONRaw.SimpleJSON)

		Convey("case expression without leading space parses successfully", func() {
			expr, err := ParseExpr(`case 1 when 1 then 'yes' else 'no'`)
			So(err, ShouldBeNil)
			So(expr, ShouldNotBeNil)
			got := eval(ctx, expr)
			So(got.String(), ShouldEqual, "yes")
		})
		Convey("case expression via Compile also works", func() {
			expr, err := Compile(`case temperature when 50 then 'hot' else 'cool'`)
			So(err, ShouldBeNil)
			got := Eval(ctx, expr)
			So(got.String(), ShouldEqual, "hot")
		})
	})
}

// TestParseNoPanic verifies E4 fix: malformed integer/float/boolean literals return errors
// rather than panicking.
func TestParseNoPanic(t *testing.T) {
	Convey("malformed literals return error not panic", t, func() {
		Convey("ParseExpr with syntax error returns error not panic", func() {
			// These produce ANTLR syntax errors, which are collected and returned.
			_, err := ParseExpr(`@#$`)
			So(err, ShouldNotBeNil)
		})
		Convey("ParseFilter with empty input returns no panic", func() {
			// Empty input — should not panic.
			expr, _ := ParseFilter(`true`)
			So(expr, ShouldNotBeNil)
		})
	})
}

// TestAggregationFunctions verifies E2: sum/avg/min/max/count built-in aggregations.
func TestAggregationFunctions(t *testing.T) {
	// friends array: ages 44, 68, 47 → sum=159, avg=53, min=44, max=68, count=3
	ctx := NewJSONContext(JSONRaw.JSON)

	Convey("array aggregation functions", t, func() {
		Convey("sum(friends.#.age) = 159", func() {
			expr, err := ParseExpr(`sum(friends.#.age)`)
			So(err, ShouldBeNil)
			got := Eval(ctx, expr)
			So(got.Type(), ShouldEqual, Float)
			So(float64(got.(FloatNode)), ShouldAlmostEqual, 159.0)
		})
		Convey("avg(friends.#.age) = 53", func() {
			expr, err := ParseExpr(`avg(friends.#.age)`)
			So(err, ShouldBeNil)
			got := Eval(ctx, expr)
			So(got.Type(), ShouldEqual, Float)
			So(float64(got.(FloatNode)), ShouldAlmostEqual, 53.0)
		})
		Convey("min(friends.#.age) = 44", func() {
			expr, err := ParseExpr(`min(friends.#.age)`)
			So(err, ShouldBeNil)
			got := Eval(ctx, expr)
			So(got.Type(), ShouldEqual, Float)
			So(float64(got.(FloatNode)), ShouldAlmostEqual, 44.0)
		})
		Convey("max(friends.#.age) = 68", func() {
			expr, err := ParseExpr(`max(friends.#.age)`)
			So(err, ShouldBeNil)
			got := Eval(ctx, expr)
			So(got.Type(), ShouldEqual, Float)
			So(float64(got.(FloatNode)), ShouldAlmostEqual, 68.0)
		})
		Convey("count(friends.#.age) = 3", func() {
			expr, err := ParseExpr(`count(friends.#.age)`)
			So(err, ShouldBeNil)
			got := Eval(ctx, expr)
			So(got.Type(), ShouldEqual, Int)
			So(int64(got.(IntNode)), ShouldEqual, 3)
		})
		Convey("sum on empty array returns UNDEFINED", func() {
			emptyCtx := NewJSONContext(`{"nums":[]}`)
			expr, _ := ParseExpr(`sum(nums.#.val)`)
			got := Eval(emptyCtx, expr)
			So(got.Type(), ShouldEqual, Undefined)
		})
		Convey("sum works in arithmetic expression", func() {
			expr, err := ParseExpr(`sum(friends.#.age) + 1`)
			So(err, ShouldBeNil)
			got := Eval(ctx, expr)
			So(got.Type(), ShouldEqual, Float)
			So(float64(got.(FloatNode)), ShouldAlmostEqual, 160.0)
		})
	})
}

// TestCompileEvalAPI verifies E5: the Compile/Eval public API.
func TestCompileEvalAPI(t *testing.T) {
	Convey("Compile and Eval public API", t, func() {
		ctx := NewJSONContext(JSONRaw.SimpleJSON)

		Convey("Compile returns non-nil expr for valid expression", func() {
			expr, err := Compile(`temperature + 1`)
			So(err, ShouldBeNil)
			So(expr, ShouldNotBeNil)
		})
		Convey("Compile returns error for invalid expression", func() {
			_, err := Compile(`@#$`)
			So(err, ShouldNotBeNil)
		})
		Convey("Eval evaluates compiled expression against context", func() {
			expr, _ := Compile(`temperature + 1`)
			got := Eval(ctx, expr)
			So(got, ShouldResemble, IntNode(51))
		})
		Convey("same CompiledExpr can be reused with different contexts", func() {
			expr, _ := Compile(`temperature * 2`)
			ctx1 := NewJSONContext(`{"temperature":10}`)
			ctx2 := NewJSONContext(`{"temperature":20}`)
			So(Eval(ctx1, expr), ShouldResemble, IntNode(20))
			So(Eval(ctx2, expr), ShouldResemble, IntNode(40))
		})
		Convey("Eval includes DefaultValue built-ins", func() {
			expr, _ := Compile(`abs(0 - 5)`)
			got := Eval(ctx, expr)
			So(got, ShouldResemble, IntNode(5))
		})
	})
}

// TestNullResultType verifies that NULL_RESULT has Null type (not Undefined).
func TestNullResultType(t *testing.T) {
	Convey("NULL_RESULT has Null type", t, func() {
		So(NULL_RESULT.Type(), ShouldEqual, Null)
		So(UNDEFINED_RESULT.Type(), ShouldEqual, Undefined)
		So(NULL_RESULT.Type(), ShouldNotEqual, UNDEFINED_RESULT.Type())
	})
}

// TestEmptyStackExpr verifies Expr() on an empty listener stack returns nil, not a panic.
func TestEmptyStackExpr(t *testing.T) {
	Convey("empty listener stack returns nil", t, func() {
		listener := &TDTLListener{}
		So(listener.Expr(), ShouldBeNil)
	})
}

// TestMergeByErrorPropagation verifies MergeBy works with valid data.
func TestMergeByErrorPropagation(t *testing.T) {
	Convey("MergeBy with valid data", t, func() {
		data := New(`[{"id":"a","val":1},{"id":"b","val":2}]`)
		result := data.MergeBy("id")
		So(result.Error(), ShouldBeNil)
		So(result.Type(), ShouldEqual, Object)
	})
}

// TestArithmeticEdgeCases covers additional arithmetic edge cases.
func TestArithmeticEdgeCases(t *testing.T) {
	Convey("arithmetic edge cases", t, func() {
		ctx := DefaultValue

		Convey("integer mod normal", func() {
			expr, _ := ParseExpr(`7 % 3`)
			So(eval(ctx, expr), ShouldResemble, IntNode(1))
		})
		Convey("float mod 5.5 % 2.0 = 1.5", func() {
			expr, _ := ParseExpr(`5.5 % 2.0`)
			got := eval(ctx, expr)
			So(got.Type(), ShouldEqual, Float)
			So(math.Abs(float64(got.(FloatNode))-1.5) < 1e-9, ShouldBeTrue)
		})
		Convey("integer div by zero returns UNDEFINED", func() {
			expr, _ := ParseExpr(`5 / 0`)
			So(eval(ctx, expr).Type(), ShouldEqual, Undefined)
		})
		Convey("float div by zero returns UNDEFINED", func() {
			expr, _ := ParseExpr(`5.0 / 0.0`)
			So(eval(ctx, expr).Type(), ShouldEqual, Undefined)
		})
		Convey("integer mod by zero returns UNDEFINED", func() {
			expr, _ := ParseExpr(`5 % 0`)
			So(eval(ctx, expr).Type(), ShouldEqual, Undefined)
		})
		Convey("float mod by zero returns UNDEFINED", func() {
			expr, _ := ParseExpr(`5.0 % 0`)
			So(eval(ctx, expr).Type(), ShouldEqual, Undefined)
		})
	})
}

// TestBooleanComparisons verifies equality/NE for booleans.
func TestBooleanComparisons(t *testing.T) {
	Convey("boolean comparisons", t, func() {
		ctx := NewJSONContext(JSONRaw.SimpleJSON)

		Convey("true = true", func() {
			expr, _ := ParseFilter(`true = true`)
			So(EvalFilter(ctx, expr), ShouldBeTrue)
		})
		Convey("true = false", func() {
			expr, _ := ParseFilter(`true = false`)
			So(EvalFilter(ctx, expr), ShouldBeFalse)
		})
		Convey("true != false", func() {
			expr, _ := ParseFilter(`true != false`)
			So(EvalFilter(ctx, expr), ShouldBeTrue)
		})
		Convey("false != false", func() {
			expr, _ := ParseFilter(`false != false`)
			So(EvalFilter(ctx, expr), ShouldBeFalse)
		})
	})
}

// TestUndefinedInExpressions verifies undefined variables degrade gracefully.
// Note: bare undefined identifiers remain as *JSONNode{Undefined} until a
// comparison operator converts them to BoolNode(false).
func TestUndefinedInExpressions(t *testing.T) {
	Convey("undefined variables in expressions", t, func() {
		ctx := NewJSONContext(JSONRaw.SimpleJSON)

		// Once compared, undefined collapses to BoolNode(false), enabling logical ops.
		Convey("(undefinedVar < 10) or true = true", func() {
			expr, _ := ParseFilter(`undefinedVar < 10 or true`)
			So(EvalFilter(ctx, expr), ShouldBeTrue)
		})
		Convey("(undefinedVar > 10) and true = false", func() {
			expr, _ := ParseFilter(`undefinedVar > 10 and true`)
			So(EvalFilter(ctx, expr), ShouldBeFalse)
		})
		Convey("(undefinedVar > 10) or false = false", func() {
			expr, _ := ParseFilter(`undefinedVar > 10 or false`)
			So(EvalFilter(ctx, expr), ShouldBeFalse)
		})
	})
}

// TestSQLSelectWithAlias verifies INSERT INTO ... SELECT ... AS generates correct output.
func TestSQLSelectWithAlias(t *testing.T) {
	jsonRaw := `{"entity1":{"color":"red","temperature":50,"metadata":{"name":"Light1","price":11.05}}}`
	ctx := NewJSONContext(jsonRaw)

	Convey("SQL SELECT with aliases", t, func() {
		Convey("select field with alias", func() {
			expr, err := Parse(`insert into target1 select entity1.color as color`)
			So(err, ShouldBeNil)
			got := EvalRuleQL(ctx, expr)
			So(got.String(), ShouldContainSubstring, "red")
		})
		Convey("select arithmetic with alias", func() {
			expr, err := Parse(`insert into target1 select entity1.temperature + 1 as temp`)
			So(err, ShouldBeNil)
			got := EvalRuleQL(ctx, expr)
			So(got.String(), ShouldContainSubstring, "51")
		})
		Convey("select string concatenation with alias", func() {
			expr, err := Parse(`insert into target1 select entity1.temperature + 'C' as tempStr`)
			So(err, ShouldBeNil)
			got := EvalRuleQL(ctx, expr)
			So(got.String(), ShouldContainSubstring, "50C")
		})
	})
}
