package tdtl

import (
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/tkeel-io/tdtl/parser"
)

// ---------------------------------------------------------------------------
// 1 & 2. evalBinary with *JSONNode as LHS/RHS
// ---------------------------------------------------------------------------

func TestEvalBinaryJSONNode(t *testing.T) {
	ctx := MutilContext{DefaultValue, NewJSONContext(`{}`)}

	Convey("evalBinary *JSONNode LHS with boolean op returns BoolNode(false)", t, func() {
		// *JSONNode LHS + isBooleanOP → BoolNode(false)
		lhs := UNDEFINED_RESULT // *JSONNode
		result := evalBinary(parser.TDTLParserEQ, lhs, IntNode(1))
		So(result, ShouldEqual, BoolNode(false))
	})

	Convey("evalBinary *JSONNode LHS with logic op delegates to evalBinary(op, BoolNode(false), rhs)", t, func() {
		// *JSONNode LHS + isLogicOP(AND) → evalBinary(AND, false, true) = false
		lhs := UNDEFINED_RESULT // *JSONNode
		result := evalBinary(parser.TDTLParserAND, lhs, BoolNode(true))
		So(result, ShouldEqual, BoolNode(false))
	})

	Convey("eval BinaryExpr with JSONNode LHS via eval()", t, func() {
		// Using a JSONPathExpr that resolves to UNDEFINED_RESULT (*JSONNode)
		// so lhs will be *JSONNode when evalBinaryExpr is called.
		expr := &BinaryExpr{
			Op:  parser.TDTLParserEQ,
			LHS: &JSONPathExpr{val: "nonexistent"},
			RHS: IntNode(1),
		}
		result := eval(ctx, expr)
		// *JSONNode lhs + isBooleanOP → BoolNode(false)
		So(result, ShouldEqual, BoolNode(false))
	})

	Convey("BoolNode(true) ADD *JSONNode (UNDEFINED_RESULT) - non-boolean, non-logic op returns UNDEFINED_RESULT", t, func() {
		// BoolNode case: rhs is *JSONNode, op is ADD (not boolean, not logic)
		// → falls through to return UNDEFINED_RESULT (line 278 in evaluator.go)
		result := evalBinary(parser.TDTLParserADD, BoolNode(true), UNDEFINED_RESULT)
		So(result, ShouldEqual, UNDEFINED_RESULT)
	})
}

// ---------------------------------------------------------------------------
// 3, 4, 5. evalFilter paths
// ---------------------------------------------------------------------------

func TestEvalFilterPaths(t *testing.T) {
	ctx := MutilContext{DefaultValue, NewJSONContext(`{}`)}

	Convey("evalFilter with FieldsExpr returns false", t, func() {
		fields := FieldsExpr{}
		So(evalFilter(ctx, fields), ShouldBeFalse)
	})

	Convey("evalFilter with *SelectStatementExpr nil filter returns true", t, func() {
		stmt := &SelectStatementExpr{filter: nil}
		So(evalFilter(ctx, stmt), ShouldBeTrue)
	})

	Convey("evalFilter with *SelectStatementExpr with filter but nil exp returns true", t, func() {
		stmt := &SelectStatementExpr{filter: &FilterExpr{exp: nil}}
		So(evalFilter(ctx, stmt), ShouldBeTrue)
	})

	Convey("evalFilter with *SelectStatementExpr BoolNode(true) filter returns true", t, func() {
		stmt := &SelectStatementExpr{filter: &FilterExpr{exp: BoolNode(true)}}
		So(evalFilter(ctx, stmt), ShouldBeTrue)
	})

	Convey("evalFilter with *SelectStatementExpr BoolNode(false) filter returns false", t, func() {
		stmt := &SelectStatementExpr{filter: &FilterExpr{exp: BoolNode(false)}}
		So(evalFilter(ctx, stmt), ShouldBeFalse)
	})
}

// ---------------------------------------------------------------------------
// 6. EvalSelect and evalSelect
// ---------------------------------------------------------------------------

func TestEvalSelectAPI(t *testing.T) {
	ctx := NewJSONContext(`{}`)

	Convey("EvalSelect builds JSON result from fields", t, func() {
		stmt := &SelectStatementExpr{
			fields: FieldsExpr{
				{exp: IntNode(42), alias: "val"},
			},
		}
		result := EvalSelect(ctx, stmt)
		So(result, ShouldNotBeNil)
		So(result.Type(), ShouldNotEqual, Undefined)
		// The result should contain "val":42
		raw := string(result.Raw())
		So(strings.Contains(raw, "val"), ShouldBeTrue)
		So(strings.Contains(raw, "42"), ShouldBeTrue)
	})

	Convey("evalSelect with nil returns UNDEFINED_RESULT", t, func() {
		result := evalSelect(MutilContext{DefaultValue, ctx}, nil)
		So(result, ShouldEqual, UNDEFINED_RESULT)
	})
}

// ---------------------------------------------------------------------------
// 7. MapContext.Call
// ---------------------------------------------------------------------------

func TestMapContextCall(t *testing.T) {
	Convey("MapContext calls custom function", t, func() {
		doubleFunc := func(args ...Node) Node {
			if len(args) == 1 {
				if n, ok := args[0].(IntNode); ok {
					return IntNode(int64(n) * 2)
				}
			}
			return UNDEFINED_RESULT
		}
		ctx := NewMapContext(nil, map[string]ContextFunc{
			"double": doubleFunc,
		})
		expr, err := ParseExpr("double(5)")
		So(err, ShouldBeNil)
		result := eval(MutilContext{DefaultValue, ctx}, expr)
		So(result, ShouldEqual, IntNode(10))
	})
}

// ---------------------------------------------------------------------------
// 8. Node.To() conversions
// ---------------------------------------------------------------------------

func TestNodeToConversions(t *testing.T) {
	Convey("BoolNode conversions", t, func() {
		Convey("BoolNode(true).To(String) = StringNode(\"true\")", func() {
			So(BoolNode(true).To(String), ShouldEqual, StringNode("true"))
		})
		Convey("BoolNode(true).To(Float) = UNDEFINED_RESULT", func() {
			So(BoolNode(true).To(Float), ShouldEqual, UNDEFINED_RESULT)
		})
	})

	Convey("IntNode conversions", t, func() {
		Convey("IntNode(5).To(Float) = FloatNode(5.0)", func() {
			So(IntNode(5).To(Float), ShouldEqual, FloatNode(5.0))
		})
		Convey("IntNode(5).To(String) = StringNode(\"5\")", func() {
			So(IntNode(5).To(String), ShouldEqual, StringNode("5"))
		})
	})

	Convey("FloatNode conversions", t, func() {
		Convey("FloatNode(3.14).To(Int) = IntNode(3)", func() {
			So(FloatNode(3.14).To(Int), ShouldEqual, IntNode(3))
		})
		Convey("FloatNode(3.14).To(String) starts with \"3.14\"", func() {
			result := FloatNode(3.14).To(String)
			So(result, ShouldNotBeNil)
			s, ok := result.(StringNode)
			So(ok, ShouldBeTrue)
			So(strings.HasPrefix(string(s), "3.14"), ShouldBeTrue)
		})
	})

	Convey("StringNode conversions", t, func() {
		Convey("StringNode(\"true\").To(Bool) = BoolNode(true)", func() {
			So(StringNode("true").To(Bool), ShouldEqual, BoolNode(true))
		})
		Convey("StringNode(\"not-a-bool\").To(Bool) = UNDEFINED_RESULT", func() {
			So(StringNode("not-a-bool").To(Bool), ShouldEqual, UNDEFINED_RESULT)
		})
		Convey("StringNode(\"42\").To(Int) = IntNode(42)", func() {
			So(StringNode("42").To(Int), ShouldEqual, IntNode(42))
		})
		Convey("StringNode(\"3.14\").To(Float) = FloatNode(3.14)", func() {
			So(StringNode("3.14").To(Float), ShouldEqual, FloatNode(3.14))
		})
	})

	Convey("JSONNode conversions", t, func() {
		Convey("JSONNode String datatype To(String) returns StringNode", func() {
			jn := JSONNode{value: []byte("hello"), datatype: String}
			result := jn.To(String)
			So(result, ShouldNotBeNil)
			s, ok := result.(StringNode)
			So(ok, ShouldBeTrue)
			So(string(s), ShouldEqual, "hello")
		})
		Convey("JSONNode Array datatype To(Array) returns self (JSONNode)", func() {
			jn := JSONNode{value: []byte("[1,2]"), datatype: Array}
			result := jn.To(Array)
			So(result, ShouldNotBeNil)
			jnResult, ok := result.(JSONNode)
			So(ok, ShouldBeTrue)
			So(string(jnResult.value), ShouldEqual, "[1,2]")
		})
	})
}

// ---------------------------------------------------------------------------
// 9. GetTopic, GetWindow, HasDimensions
// ---------------------------------------------------------------------------

func TestGetTopicWindowDimensions(t *testing.T) {
	Convey("GetTopic", t, func() {
		Convey("nil returns empty string and false", func() {
			s, ok := GetTopic(nil)
			So(s, ShouldEqual, "")
			So(ok, ShouldBeFalse)
		})
		Convey("empty SelectStatementExpr returns empty string and false", func() {
			s, ok := GetTopic(&SelectStatementExpr{})
			So(s, ShouldEqual, "")
			So(ok, ShouldBeFalse)
		})
		Convey("SelectStatementExpr with topic returns joined string and true", func() {
			s, ok := GetTopic(&SelectStatementExpr{topic: TopicExpr{"a", "b"}})
			So(s, ShouldEqual, "a/b")
			So(ok, ShouldBeTrue)
		})
	})

	Convey("HasDimensions", t, func() {
		Convey("nil returns false", func() {
			So(HasDimensions(nil), ShouldBeFalse)
		})
		Convey("empty SelectStatementExpr returns false", func() {
			So(HasDimensions(&SelectStatementExpr{}), ShouldBeFalse)
		})
		Convey("SelectStatementExpr with DimensionsExpr returns true", func() {
			So(HasDimensions(&SelectStatementExpr{dimensions: &DimensionsExpr{}}), ShouldBeTrue)
		})
	})

	Convey("GetWindow", t, func() {
		Convey("nil returns nil", func() {
			So(GetWindow(nil), ShouldBeNil)
		})
		Convey("empty SelectStatementExpr returns nil", func() {
			So(GetWindow(&SelectStatementExpr{}), ShouldBeNil)
		})
		Convey("SelectStatementExpr with window returns non-nil WindowExpr", func() {
			win := &WindowExpr{WindowType: TUMBLING_WINDOW, Length: 10}
			stmt := &SelectStatementExpr{
				dimensions: &DimensionsExpr{window: win},
			}
			result := GetWindow(stmt)
			So(result, ShouldNotBeNil)
			So(result.WindowType, ShouldEqual, TUMBLING_WINDOW)
		})
	})
}

// ---------------------------------------------------------------------------
// 10. ParseFunc and walkFunc
// ---------------------------------------------------------------------------

func TestParseFuncExtract(t *testing.T) {
	Convey("ParseFunc extracts CallExpr from abs(5)", t, func() {
		expr, err := ParseExpr("abs(5)")
		So(err, ShouldBeNil)
		calls := ParseFunc(expr)
		So(len(calls), ShouldEqual, 1)
		So(calls[0].FuncName(), ShouldEqual, "abs")
	})

	Convey("ParseFunc returns empty list from 1 + 2", t, func() {
		expr, err := ParseExpr("1 + 2")
		So(err, ShouldBeNil)
		calls := ParseFunc(expr)
		So(len(calls), ShouldEqual, 0)
	})
}

// ---------------------------------------------------------------------------
// 11. print.go - Dump
// ---------------------------------------------------------------------------

func TestPrintDump(t *testing.T) {
	Convey("Dump does not panic on a BinaryExpr", t, func() {
		expr, err := ParseExpr("1 + 2")
		So(err, ShouldBeNil)
		So(func() { _ = Dump(expr) }, ShouldNotPanic)
	})
}

// ---------------------------------------------------------------------------
// 12. CallExpr methods
// ---------------------------------------------------------------------------

func TestCallExprMethods(t *testing.T) {
	Convey("CallExpr String, FuncName, Args from abs(5)", t, func() {
		expr, err := ParseExpr("abs(5)")
		So(err, ShouldBeNil)
		callExpr, ok := expr.(*CallExpr)
		So(ok, ShouldBeTrue)

		Convey("String() returns raw expression text", func() {
			s := callExpr.String()
			So(s, ShouldNotBeEmpty)
		})
		Convey("FuncName() returns function name", func() {
			So(callExpr.FuncName(), ShouldEqual, "abs")
		})
		Convey("Args() returns non-empty slice", func() {
			args := callExpr.Args()
			So(len(args), ShouldEqual, 1)
		})
	})
}

// ---------------------------------------------------------------------------
// 13. context_json.go Value with empty path
// ---------------------------------------------------------------------------

func TestContextEdgeCases(t *testing.T) {
	Convey("JSONContext Value with empty path returns whole JSON", t, func() {
		ctx := NewJSONContext(`{"a":1}`)
		result := ctx.Value("")
		So(result, ShouldNotBeNil)
		So(result.Type(), ShouldNotEqual, Undefined)
		raw := string(result.Raw())
		So(strings.Contains(raw, "a"), ShouldBeTrue)
	})

	Convey("evalRuleQL with *FilterExpr directly returns inner expr result", t, func() {
		ctx := MutilContext{DefaultValue, NewJSONContext(`{}`)}
		filterExpr := &FilterExpr{exp: IntNode(42)}
		result := evalRuleQL(ctx, filterExpr)
		So(result, ShouldEqual, IntNode(42))
	})

	Convey("MapContext Value with JSONNode values", t, func() {
		jnValue := JSONNode{value: []byte("hello"), datatype: String}
		jnPtr := &JSONNode{value: []byte("world"), datatype: String}
		ctx := NewMapContext(map[string]Node{
			"a": jnValue,
			"b": jnPtr,
		}, nil)

		Convey("JSONNode value is unwrapped by Node()", func() {
			result := ctx.Value("a")
			So(result, ShouldNotBeNil)
			So(result.Type(), ShouldNotEqual, Undefined)
		})
		Convey("*JSONNode value is unwrapped by Node()", func() {
			result := ctx.Value("b")
			So(result, ShouldNotBeNil)
			So(result.Type(), ShouldNotEqual, Undefined)
		})
	})
}
