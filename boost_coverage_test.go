package tdtl

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/tkeel-io/tdtl/parser"
)

// TestBuiltinFunctionEdgeCases covers abs/base64/count/agg UNDEFINED_RESULT paths.
func TestBuiltinFunctionEdgeCases(t *testing.T) {
	Convey("absFunc string argument", t, func() {
		result := absFunc(StringNode("-5.5"))
		So(result, ShouldNotBeNil)
	})
	Convey("absFunc bool argument returns UNDEFINED_RESULT", t, func() {
		result := absFunc(BoolNode(true))
		So(result.Type(), ShouldEqual, Undefined)
	})
	Convey("absFunc no args returns UNDEFINED_RESULT", t, func() {
		result := absFunc()
		So(result.Type(), ShouldEqual, Undefined)
	})
	Convey("absFuncHandle with non-numeric type returns UNDEFINED_RESULT", t, func() {
		result := absFuncHandle(StringNode("hello"))
		So(result.Type(), ShouldEqual, Undefined)
	})
	Convey("base64Func no args returns UNDEFINED_RESULT", t, func() {
		result := base64Func()
		So(result.Type(), ShouldEqual, Undefined)
	})
	Convey("base64Func with StringNode returns encoded string", t, func() {
		result := base64Func(StringNode("hello"))
		So(result.Type(), ShouldEqual, String)
	})
	Convey("base64Func with JSONNode returns encoded string", t, func() {
		jn := JSONNode{value: []byte(`{"a":1}`), datatype: JSON}
		result := base64Func(jn)
		So(result.Type(), ShouldEqual, String)
	})
	Convey("base64Func with IntNode (default type) returns UNDEFINED_RESULT", t, func() {
		result := base64Func(IntNode(42))
		So(result.Type(), ShouldEqual, Undefined)
	})
	Convey("base64Func with empty string returns UNDEFINED_RESULT", t, func() {
		result := base64Func(StringNode(""))
		So(result.Type(), ShouldEqual, Undefined)
	})
	Convey("countFunc no args returns UNDEFINED_RESULT", t, func() {
		result := countFunc()
		So(result.Type(), ShouldEqual, Undefined)
	})
	Convey("countFunc with non-array returns IntNode(1)", t, func() {
		result := countFunc(IntNode(5))
		So(result, ShouldEqual, IntNode(1))
	})
	Convey("aggFuncHandle with no args returns UNDEFINED_RESULT", t, func() {
		fn := aggFuncHandle(sumAgg)
		result := fn()
		So(result.Type(), ShouldEqual, Undefined)
	})
	Convey("aggFuncHandle with non-numeric-only array returns UNDEFINED_RESULT", t, func() {
		fn := aggFuncHandle(sumAgg)
		jn := JSONNode{value: []byte(`["a","b","c"]`), datatype: Array}
		result := fn(jn)
		So(result.Type(), ShouldEqual, Undefined)
	})
	Convey("minAgg with decreasing values updates minimum", t, func() {
		result := minAgg([]float64{3.0, 1.0, 2.0})
		So(result, ShouldEqual, FloatNode(1.0))
	})
}

// TestValueNilFunctions covers value.Call when functions map is nil.
func TestValueNilFunctions(t *testing.T) {
	Convey("value.Call with nil functions returns UNDEFINED_RESULT", t, func() {
		v := &value{functions: nil}
		callExpr := &CallExpr{key: "any"}
		result := v.Call(callExpr, []Node{})
		So(result.Type(), ShouldEqual, Undefined)
	})
}

// TestEvalBinaryTypeCombinations covers uncovered type-combination branches in evalBinary.
func TestEvalBinaryTypeCombinations(t *testing.T) {
	Convey("evalBinary StringNode+BoolNode delegates via To(Bool)", t, func() {
		result := evalBinary(parser.TDTLLexerEQ, StringNode("true"), BoolNode(true))
		So(result, ShouldResemble, BoolNode(true))
	})
	Convey("evalBinary StringNode+*JSONNode returns UNDEFINED_RESULT", t, func() {
		result := evalBinary(parser.TDTLLexerEQ, StringNode("hello"), UNDEFINED_RESULT)
		So(result.Type(), ShouldEqual, Undefined)
	})
	Convey("evalBinary FloatNode+StringNode coerces via To(Float)", t, func() {
		result := evalBinary(parser.TDTLLexerEQ, FloatNode(3.14), StringNode("3.14"))
		So(result, ShouldResemble, BoolNode(true))
	})
	Convey("evalBinary FloatNode+*JSONNode returns UNDEFINED_RESULT", t, func() {
		result := evalBinary(parser.TDTLLexerEQ, FloatNode(3.14), UNDEFINED_RESULT)
		So(result.Type(), ShouldEqual, Undefined)
	})
	Convey("evalBinary BoolNode+StringNode delegates via To(Bool)", t, func() {
		result := evalBinary(parser.TDTLLexerEQ, BoolNode(true), StringNode("true"))
		So(result, ShouldResemble, BoolNode(true))
	})
	Convey("evalBinary JSONNode (value type) as LHS hits default case", t, func() {
		jn := JSONNode{value: []byte(`{"a":1}`), datatype: JSON}
		result := evalBinary(parser.TDTLLexerEQ, jn, IntNode(1))
		So(result.Type(), ShouldEqual, Undefined)
	})
	Convey("evalBinary *JSONNode LHS with ADD op (non-bool non-logic) returns UNDEFINED_RESULT", t, func() {
		result := evalBinary(parser.TDTLParserADD, UNDEFINED_RESULT, IntNode(1))
		So(result.Type(), ShouldEqual, Undefined)
	})
}

// TestEvalBinaryHelperUnknownOps covers the fallthrough UNDEFINED_RESULT in binary helper functions.
func TestEvalBinaryHelperUnknownOps(t *testing.T) {
	Convey("evalBinaryInt with AND op returns UNDEFINED_RESULT", t, func() {
		result := evalBinaryInt(parser.TDTLParserAND, IntNode(1), IntNode(2))
		So(result.Type(), ShouldEqual, Undefined)
	})
	Convey("evalBinaryFloat with AND op returns UNDEFINED_RESULT", t, func() {
		result := evalBinaryFloat(parser.TDTLParserAND, FloatNode(1.0), FloatNode(2.0))
		So(result.Type(), ShouldEqual, Undefined)
	})
	Convey("evalBinaryBool with ADD op returns UNDEFINED_RESULT", t, func() {
		result := evalBinaryBool(parser.TDTLParserADD, BoolNode(true), BoolNode(false))
		So(result.Type(), ShouldEqual, Undefined)
	})
}

// TestEvaluatorBranchCoverage covers HasDimensions, evalRuleQL, evalFilter, eval, evalSelect.
func TestEvaluatorBranchCoverage(t *testing.T) {
	ctx := MutilContext{DefaultValue, NewJSONContext(`{"temp": 42}`)}

	Convey("HasDimensions with *DimensionsExpr returns true", t, func() {
		So(HasDimensions(&DimensionsExpr{}), ShouldBeTrue)
	})
	Convey("HasDimensions with non-Expr type returns false", t, func() {
		So(HasDimensions(IntNode(1)), ShouldBeFalse)
	})
	Convey("evalRuleQL with *JSONPathExpr returns field value", t, func() {
		result := evalRuleQL(ctx, &JSONPathExpr{"temp"})
		So(result.Type(), ShouldNotEqual, Undefined)
	})
	Convey("evalRuleQL with unknown Expr type returns UNDEFINED_RESULT", t, func() {
		result := evalRuleQL(ctx, &FieldExpr{})
		So(result.Type(), ShouldEqual, Undefined)
	})
	Convey("evalFilter with nil returns true", t, func() {
		So(evalFilter(ctx, nil), ShouldBeTrue)
	})
	Convey("evalFilter default branch with non-bool result returns false", t, func() {
		So(evalFilter(ctx, IntNode(42)), ShouldBeFalse)
	})
	Convey("eval with JSONNode value type returns it directly", t, func() {
		jn := JSONNode{value: []byte(`{"a":1}`), datatype: JSON}
		result := eval(ctx, jn)
		So(result.Type(), ShouldEqual, JSON)
	})
	Convey("eval with unknown Expr type returns UNDEFINED_RESULT", t, func() {
		result := eval(ctx, &FieldExpr{})
		So(result.Type(), ShouldEqual, Undefined)
	})
	Convey("evalSelect with non-SelectStatementExpr returns UNDEFINED_RESULT", t, func() {
		result := evalSelect(ctx, IntNode(1))
		So(result.Type(), ShouldEqual, Undefined)
	})
	Convey("evalCallExpr with zero args calls ctx.Call directly", t, func() {
		callExpr := &CallExpr{key: "count", args: nil}
		result := eval(DefaultValue, callExpr)
		So(result.Type(), ShouldEqual, Undefined)
	})
	Convey("GetTopic with non-SelectStatementExpr returns empty string and false", t, func() {
		s, ok := GetTopic(IntNode(1))
		So(s, ShouldEqual, "")
		So(ok, ShouldBeFalse)
	})
	Convey("GetWindow with non-SelectStatementExpr returns nil", t, func() {
		So(GetWindow(IntNode(1)), ShouldBeNil)
	})
}

// TestContextMapMissingEntries covers mapContext Value and Call for absent entries.
func TestContextMapMissingEntries(t *testing.T) {
	ctx := NewMapContext(map[string]Node{}, map[string]ContextFunc{}).(mapContext)

	Convey("mapContext.Value for missing key returns UNDEFINED_RESULT", t, func() {
		result := ctx.Value("nonexistent")
		So(result.Type(), ShouldEqual, Undefined)
	})
	Convey("mapContext.Call for missing function returns UNDEFINED_RESULT", t, func() {
		result := ctx.Call(&CallExpr{key: "nonexistent"}, []Node{})
		So(result.Type(), ShouldEqual, Undefined)
	})
}

// TestNewExprAndTDTLErrors covers error return paths in NewExpr and NewTDTL.
func TestNewExprAndTDTLErrors(t *testing.T) {
	Convey("NewExpr with invalid expression returns error", t, func() {
		_, err := NewExpr("@", nil)
		So(err, ShouldNotBeNil)
	})
	Convey("NewTDTL with invalid SQL returns error", t, func() {
		_, err := NewTDTL("@", nil)
		So(err, ShouldNotBeNil)
	})
}
