package tdtl

import (
	"fmt"
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestEval(t *testing.T) {
	got := (float64(1.1) + float64(IntNode(1))) * float64(IntNode(1))
	want := float64(2.1)
	fmt.Println(reflect.DeepEqual(got, want))
}

func TestDefaultValue_Eval(t *testing.T) {
	tests := []struct {
		name string
		json string
		expr string
		want Node
	}{
		{"number", JSONRaw.SimpleJSON, `(YX_0002 * 2 + YX_0003)`, IntNode(4)},
		{"number", JSONRaw.SimpleJSON, `1`, IntNode(1)},
		{"number", JSONRaw.SimpleJSON, `1 + 2 `, IntNode(3)},
		{"number", JSONRaw.SimpleJSON, `1 + 2 * 3 + 2 * 3`, IntNode(13)},
		{"number", JSONRaw.SimpleJSON, `1 + 2 * (3 + 2) * 3`, IntNode(31)},
		{"number", JSONRaw.SimpleJSON, `1 + 2 * (3 + 2) * 3`, IntNode(31)},
		{"number", JSONRaw.SimpleJSON, `1 + 1.1`, FloatNode(2.1)},     //@TODO precision
		{"number", JSONRaw.SimpleJSON, `1.1 + 1`, FloatNode(2.1)},     //@TODO precision
		{"number", JSONRaw.SimpleJSON, `(1.1 + 1)*1`, FloatNode(2.1)}, //@TODO precision
		{"number", JSONRaw.SimpleJSON, `(1 + 1)*1.1`, FloatNode(2.2)}, //@TODO precision
		{"number", JSONRaw.SimpleJSON, `1 / 0.1`, FloatNode(10)},
		{"number", JSONRaw.SimpleJSON, `1 / 0`, UNDEFINED_RESULT},
		{"number", JSONRaw.SimpleJSON, `1.0 / 0`, UNDEFINED_RESULT},
		{"number", JSONRaw.SimpleJSON, `5 % 2`, IntNode(1)},
		{"number", JSONRaw.SimpleJSON, `5 % 2.0`, FloatNode(1)},
		{"string2", JSONRaw.SimpleJSON, `'aaa' + 'bbb'`, StringNode("aaabbb")},
		{"string", JSONRaw.SimpleJSON, `'aaa' + 11111`, StringNode("aaa11111")},
		{"string", JSONRaw.SimpleJSON, `'aaa' - 11111`, UNDEFINED_RESULT},
		{"string", JSONRaw.SimpleJSON, `'aaa' * 11111`, UNDEFINED_RESULT},
		{"string", JSONRaw.SimpleJSON, `'aaa' / 11111`, UNDEFINED_RESULT},
		{"string", JSONRaw.SimpleJSON, `'aaa' % 11111`, UNDEFINED_RESULT},
		{"string", JSONRaw.SimpleJSON, `11111 + 'bbb'`, StringNode("11111bbb")},
		{"string", JSONRaw.SimpleJSON, `11111 - 'bbb'`, UNDEFINED_RESULT},
		{"string", JSONRaw.SimpleJSON, `11111 * 'bbb'`, UNDEFINED_RESULT},
		{"string", JSONRaw.SimpleJSON, `11111 / 'b'bb'`, UNDEFINED_RESULT},
		{"string", JSONRaw.SimpleJSON, `11111 % 'bbb'`, UNDEFINED_RESULT},
		{"string", JSONRaw.SimpleJSON, `'111' + 11`, StringNode("11111")},
		{"string", JSONRaw.SimpleJSON, `'111' - 11`, IntNode(100)},
		{"string", JSONRaw.SimpleJSON, `'111' * 11`, IntNode(1221)},
		{"string", JSONRaw.SimpleJSON, `'111' / 11`, IntNode(10)},
		{"string", JSONRaw.SimpleJSON, `'111' % 11`, IntNode(1)},
		{"string", JSONRaw.SimpleJSON, `'11' = 111`, BoolNode(false)},
		{"string", JSONRaw.SimpleJSON, `'11' > 111`, BoolNode(false)},
		{"string", JSONRaw.SimpleJSON, `'111' = 111`, BoolNode(true)},
		{"string", JSONRaw.SimpleJSON, `111 + '11'`, StringNode("11111")},
		{"string", JSONRaw.SimpleJSON, `111 - '11'`, IntNode(100)},
		{"string", JSONRaw.SimpleJSON, `111 * '11'`, IntNode(1221)},
		{"string", JSONRaw.SimpleJSON, `111 / '11'`, IntNode(10)},
		{"string", JSONRaw.SimpleJSON, `111 % '11'`, IntNode(1)},
		{"string", JSONRaw.SimpleJSON, `111 > '11'`, BoolNode(true)},
		{"string", JSONRaw.SimpleJSON, `111 = '11'`, BoolNode(false)},
		{"string", JSONRaw.SimpleJSON, `111 = '111'`, BoolNode(true)},
		{"string", JSONRaw.SimpleJSON, `'111' + '11'`, StringNode("11111")},
		{"string", JSONRaw.SimpleJSON, `'111' - '11'`, IntNode(100)},
		{"string", JSONRaw.SimpleJSON, `'111' * '11'`, IntNode(1221)},
		{"string", JSONRaw.SimpleJSON, `'111' / '11'`, IntNode(10)},
		{"string", JSONRaw.SimpleJSON, `'111' % '11'`, IntNode(1)},
		{"string", JSONRaw.SimpleJSON, `'111' > '11'`, BoolNode(true)},
		{"string", JSONRaw.SimpleJSON, `'111' = '11'`, BoolNode(false)},
		{"string", JSONRaw.SimpleJSON, `'111' = '111'`, BoolNode(true)},
		{"string", JSONRaw.SimpleJSON, `'1' + 1`, StringNode("11")},
		{"string", JSONRaw.SimpleJSON, `1 + '1'`, StringNode("11")},
		{"string", JSONRaw.SimpleJSON, `'1' - 1`, IntNode(0)},
		{"string", JSONRaw.SimpleJSON, `1 - '1'`, IntNode(0)},
		{"string", JSONRaw.SimpleJSON, `'1.0' - 1`, FloatNode(0)},
		{"string", JSONRaw.SimpleJSON, `1 - '1.0'`, FloatNode(0)},
		{"json", JSONRaw.JSON, "temperature", IntNode(50)},
		{"json", JSONRaw.JSON, "temperature + 2.0 * (temperature + 2) * 3", FloatNode(362)},
		{"json", JSONRaw.JSON, "temperature + 2.0 * (temperature + 2) * 3 > 362", BoolNode(false)},
		{"json", JSONRaw.JSON, "temperature + 2.0 * (temperature + 2) * 3 < 363", BoolNode(true)},
		{"json", JSONRaw.JSON, "friends[0]", New("{\"first\": \"Dale\", \"last\": \"Murphy\", \"age\": 44}")},
		{"json", JSONRaw.JSON, "friends[#]", IntNode(3)},
		{"json", JSONRaw.JSON, "friends[#].first", New("[\"Dale\",\"Roger\",\"Jane\"]")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := NewJSONContext(tt.json)
			expr, _ := ParseExpr(tt.expr) // Field expr
			//print()
			got := eval(v, expr)
			if got.Type() == tt.want.Type() && !reflect.DeepEqual(got.Raw(), tt.want.Raw()) {
				t.Errorf("expr %v", expr)
				t.Errorf("defaultEvalContext.Eval() \ngot=  %v \nwant= %v", got, tt.want)
			}
		})
	}
}

func TestBoolExpr(t *testing.T) {
	tests := []struct {
		name string
		ctx  Context
		expr string
		want interface{}
	}{
		{"bool", NewJSONContext(JSONRaw.SimpleJSON), `true`, true},
		{"bool", NewJSONContext(JSONRaw.SimpleJSON), `false`, false},
		{"bool", NewJSONContext(JSONRaw.SimpleJSON), `false and false`, false},
		{"bool", NewJSONContext(JSONRaw.SimpleJSON), `true  and false`, false},
		{"bool", NewJSONContext(JSONRaw.SimpleJSON), `true  and true`, true},
		{"bool", NewJSONContext(JSONRaw.SimpleJSON), `true  and true and false`, false},
		{"bool", NewJSONContext(JSONRaw.SimpleJSON), `true  and true and true`, true},
		{"bool", NewJSONContext(JSONRaw.SimpleJSON), `false or  false`, false},
		{"bool1", NewJSONContext(JSONRaw.SimpleJSON), `true  or  false`, true},
		{"bool", NewJSONContext(JSONRaw.SimpleJSON), `true  or  true`, true},
		{"bool", NewJSONContext(JSONRaw.SimpleJSON), `false or  false  or true`, true},
		{"bool", NewJSONContext(JSONRaw.SimpleJSON), `false or  false  or false`, false},
		{"bool", NewJSONContext(JSONRaw.SimpleJSON), `true  or  true  or true`, true},
		{"bool", NewJSONContext(JSONRaw.SimpleJSON), `true  or  true  or false`, true},
		{"bool", NewJSONContext(JSONRaw.SimpleJSON), `1 > 1`, false},
		{"bool", NewJSONContext(JSONRaw.SimpleJSON), `1 > 1.0`, false},
		{"bool", NewJSONContext(JSONRaw.SimpleJSON), `2 > 1`, true},
		{"bool", NewJSONContext(JSONRaw.SimpleJSON), `1 + 2 > 2.0`, true},
		{"bool", NewJSONContext(JSONRaw.SimpleJSON), `1 + 2 > 2 and false`, false},
		{"bool", NewJSONContext(JSONRaw.SimpleJSON), `1 + 2 > 2 and true`, true},
		{"bool", NewJSONContext(JSONRaw.SimpleJSON), `color = 'red'`, true},
		{"bool", NewJSONContext(JSONRaw.SimpleJSON), `'2' =  '2'`, true},
		{"bool", NewJSONContext(JSONRaw.SimpleJSON), `'2' >  '2'`, false},
		{"bool", NewJSONContext(JSONRaw.SimpleJSON), `'2' >= '2'`, true},
		{"bool", NewJSONContext(JSONRaw.SimpleJSON), `'2' <  '2'`, false},
		{"bool", NewJSONContext(JSONRaw.SimpleJSON), `'2' <= '2'`, true},
		{"bool", NewJSONContext(JSONRaw.SimpleJSON), `'2' <> '2'`, false},
		{"bool", NewJSONContext(JSONRaw.SimpleJSON), `'2' != '2'`, false},
		{"bool", NewJSONContext(JSONRaw.SimpleJSON), `'20" = '3'`, false},
		{"bool", NewJSONContext(JSONRaw.SimpleJSON), `'20' > '3'`, false},
		{"bool", NewJSONContext(JSONRaw.SimpleJSON), `'20' >= '3'`, false},
		{"bool", NewJSONContext(JSONRaw.SimpleJSON), `'20' <  '3'`, true},
		{"bool", NewJSONContext(JSONRaw.SimpleJSON), `'20' <= '3'`, true},
		{"bool", NewJSONContext(JSONRaw.SimpleJSON), `'20' <> '3'`, true},
		{"bool", NewJSONContext(JSONRaw.SimpleJSON), `'20' != '3'`, true},
		{"bool", NewJSONContext(JSONRaw.SimpleJSON), `'20' = '3'`, false},
		{"bool", NewJSONContext(JSONRaw.SimpleJSON), `'20' > 3`, true},
		{"bool", NewJSONContext(JSONRaw.SimpleJSON), `'20' >= 3`, true},
		{"bool", NewJSONContext(JSONRaw.SimpleJSON), `'20' < 3`, false},
		{"bool", NewJSONContext(JSONRaw.SimpleJSON), `'20' <= 3`, false},
		{"bool", NewJSONContext(JSONRaw.SimpleJSON), `'20' <> 3`, true},
		{"bool", NewJSONContext(JSONRaw.SimpleJSON), `'20' != 3`, true},
		{"bool", NewJSONContext(JSONRaw.SimpleJSON), `11 < 10 or true`, true},
		{"bool", NewJSONContext(JSONRaw.SimpleJSON), `u < 10 or true`, true},
		{"bool", NewJSONContext(JSONRaw.SimpleJSON), `u > 10 and true`, false},
	}
	Convey("run test", t, func() {
		for idx, tt := range tests {
			Convey(fmt.Sprintf("%s-%d", tt.name, idx), func() {
				expr, _ := ParseFilter(tt.expr) // Filter expr
				So(EvalFilter(tt.ctx, expr), ShouldEqual, tt.want)
			})
		}
	})
}

func TestSQL(t *testing.T) {
	jsonRaw := `{"entity1":{"color":"red", "temperature":50,{name": "Light1", "price": 11.05},"params": {"OPCUA#Lu1_Bottom_Waice_Temp": {"value":123}}}`
	ctx := NewJSONContext(jsonRaw)
	tests := []struct {
		name string
		ctx  Context
		expr string
		want *Collect
	}{
		{"sql", ctx, `insert into target1 select entity1.color`, New(`{}`)},
		{"sql", ctx, `insert into target1 select entity1.*`, New(`{}`)},
		{"sql", ctx, `insert into target1 select entity1.color as aaa`, New(`{"aaa":"red"}`)},
		{"sql", ctx, `insert into target1 select entity1.temperature + 1 AS temp`, New(`{"temp":51}`)},
		{"sql", ctx, `insert into target1 select entity1.temperature - 1 AS temp`, New(`{"temp":49}`)},
		{"sql", ctx, `insert into target1 select entity1.temperature + '1' AS temp`, New(`{"temp":"501"}`)},
		{"sql", ctx, `insert into target1  select entity1.temperature - '1' AS temp`, New(`{"temp":49}`)},
		//{"sql", ctx, `insert into select entity1.params.OPCUA#Lu1_Bottom_Waice_Temp.value - 20 AS temp`, JSONNode(`{"temp":103}`)},
		//{"sql", ctx,
		//	`insert into select * from a/b`,
		//	JSONNode(jsonRaw)},
		//{"sql", ctx,
		//	`select '' as a from a/b`,
		//	JSONNode(`{"a":""}`)},
		//{"sql", ctx,
		//	`insert into select
		//			* ,
		//			temperature - '1' AS temp
		//         from a/b`,
		//	JSONNode(`{"color":"red", "temperature":50,"metadata": {"name": "Light1", "price": 11.05},"params": {"OPCUA#Lu1_Bottom_Waice_Temp": {"value":123}},"temp":49}`)},
		//{"sql", ctx,
		//	`insert into select
		//			temperature - '1' AS temp,
		//			*
		//         from a/b`,
		//	JSONNode(jsonRaw)},
	}
	Convey("run test", t, func() {
		for _, tt := range tests {
			expr, err := Parse(tt.expr) // Filter expr
			So(err, ShouldBeNil)
			So(EvalRuleQL(tt.ctx, expr).String(), ShouldEqual, tt.want.String())
		}
	})
}

////577351 ns/op
//func BenchmarkParse(b *testing.B) {
//	b.ResetTimer()
//	b.ReportAllocs()
//	expr := `1 + (1 + 2 * (3 + 2)) * 3`
//	Convey("BenchmarkExpr", b, func() {
//		for i := 0; i < b.N; i++ {
//			ParseExpr(expr)
//		}
//	})
//}
//
////274 ns/op
//func BenchmarkExpr(b *testing.B) {
//	b.ResetTimer()
//	b.ReportAllocs()
//	v := DefaultValue
//	rule := `1 + (1 + 2 * (3 + 2)) * 3`
//	expr, _ := ParseExpr(rule)
//	Convey("BenchmarkExpr", b, func() {
//		for i := 0; i < b.N; i++ {
//			eval(v, expr)
//		}
//	})
//}
//
//// 14.3 ns/op
//func BenchmarkFloatExpr(b *testing.B) {
//	b.ResetTimer()
//	b.ReportAllocs()
//	v := DefaultValue
//	rule := `3.1`
//	expr, _ := ParseExpr(rule)
//	Convey("BenchmarkExpr", b, func() {
//		for i := 0; i < b.N; i++ {
//			eval(v, expr)
//		}
//	})
//}
//
//// 78.9 ns/op
//func BenchmarkFloat2Expr(b *testing.B) {
//	b.ResetTimer()
//	b.ReportAllocs()
//	v := DefaultValue
//	rule := `2 * (3.1 + 2)`
//	expr, _ := ParseExpr(rule)
//	Convey("BenchmarkExpr", b, func() {
//		for i := 0; i < b.N; i++ {
//			eval(v, expr)
//		}
//	})
//}
//
//// 230 ns/op
//func BenchmarkJsonExpr(b *testing.B) {
//	b.ResetTimer()
//	b.ReportAllocs()
//	v := NewJSONContext(`{
//  "name": {"first": "Tom", "last": "Anderson", "age": 44},
//  "age":37,
//  "children": ["Sara","Alex","Jack"],
//  "movie": "Deer Hunter",
//  "friends": [
//    {"first": "Dale", "last": "Murphy", "age": 44},
//    {"first": "Roger", "last": "Craig", "age": 68},
//    {"first": "Jane", "last": "Murphy", "age": 47}
//  ]
//}`)
//	rule := `2 * (3.1 + age)`
//	expr, _ := ParseExpr(rule)
//	Convey("BenchmarkExpr", b, func() {
//		for i := 0; i < b.N; i++ {
//			eval(v, expr)
//		}
//	})
//}
//
//// 78.9 ns/op
//
//func BenchmarkSQL(b *testing.B) {
//	json, sql := getData()
//	expr, _ := Parse(sql)
//	b.ResetTimer()
//	b.ReportAllocs()
//	v := MutilContext{
//		NewMapContext(map[string]Node{
//			"user.id": StringNode("1.1"),
//		}, nil),
//		NewJSONContext(json),
//	}
//	Convey("BenchmarkExpr", b, func() {
//		for i := 0; i < b.N; i++ {
//			EvalFilter(v, expr)
//			EvalRuleQL(v, expr)
//		}
//	})
//}
//
//func getData() (json, sql string) {
//	return `{
//	         "color":"red", "temperature":50,
//	         "ports": ["lo0", "eth1", "eth2"],
//             "metadata": {"name": "Light1", "price": 11.05}
//         }`,
//		`select
//			base64(*) as payload,
//			temperature + '°F' as temperature.f,
//			((temperature - 32) * 5 / 9.0) + '°C' as temperature.c,
//			color as metadata.color,
//			case "color"
//			when 'red' then '红色'
//			when 'green' then '绿色'
//			else 'on' AS color_zh,
//			ports[2] AS dev,
//			metadata.name
//			from sys/+/+/thing/event/+/post
//			where
//				color = 'red' and
//				temperature > 49`
//}
