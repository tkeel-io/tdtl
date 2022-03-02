package tdtl

import (
	"fmt"
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestParseError(t *testing.T) {

	tests := []struct {
		expr   string
		hasErr bool
		err    interface{}
	}{
		{"insert into entity3 select entity1.aaa as aaa", false, nil},
		{"insert into entity3 select entity1.aaa + 'aaa' as aaa", false, nil},
		{"insert into entity3 select entity1.aaa", false, ""},
		{"insert into entity3 select entity1.*", false, ""},
		{"insert into entity3 select entity1.* as ccc", true, "[1:36]mismatched input ' as ' expecting <EOF>"},
		{"insert into entity3 select entity1.abc as aaa, entity2.aaa as aa", false, nil},
		{"insert into entity3 select entity1.ccc as aaa, entity2.aaa + 1 as aa", false, nil},
		{"insert into entity3 select 0entity1.eee as aaa, entity2.aaa + 1 + '/AAA' as aa", false, nil},
	}
	idx := 1
	for _, tt := range tests {
		t.Run("tt.name", func(t *testing.T) {
			_, err := Parse(tt.expr)
			fmt.Println("================================")
			fmt.Println(tt.expr)
			fmt.Println("--------------------------------")
			fmt.Println(err)
			fmt.Println("================================")
			if (err != nil) != tt.hasErr {
				t.Errorf("[%d] has error(%v), want (%v)", idx, err, tt.err)
			}
			if tt.hasErr && err != nil && err.Error() != tt.err {
				t.Errorf("[%d] %v, want %v", idx, err, tt.err)
			}
			idx++
		})
	}
}

func TestParseNumber(t *testing.T) {
	tests := []struct {
		expr string
		want interface{}
	}{
		// calc
		{"1", IntNode(1)},
		{"0", IntNode(0)},
		{"-1", IntNode(-1)},
		{"1.0", FloatNode(1)},
		{"0.0", FloatNode(0)},
		{".0", FloatNode(0)},
		{"0.", FloatNode(0)},
		{"-1.0", FloatNode(-1)},
		{"1 + 1 ", IntNode(2)},
		{"1 - 1", IntNode(0)},
		{"7 * 3", IntNode(21)},
		{"7 / 3", IntNode(2)},
		{"1 - 1.0", FloatNode(0)},
		{"1.0 - 1", FloatNode(0)},
		{"3 * 7.0", FloatNode(21)},
		{"7.0 / 3.0", FloatNode(2.3333333333333335)},
	}

	for i, tt := range tests {
		expr, _ := ParseExpr(tt.expr)
		if got := eval(DefaultValue, expr); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("[%d] %v, want %v", i, got, tt.want)
		}
	}
}

func TestParseBool(t *testing.T) {
	tests := []struct {
		expr string
		want bool
	}{
		// bool
		{"tRUE", true},
		{"fALSE", false},
		{"true", true},
		{"false", false},
		{"True", true},
		{"False", false},
		{"TRUE", true},
		{"FALSE", false},
	}
	for i, tt := range tests {
		expr, _ := ParseExpr(tt.expr)
		if got := eval(DefaultValue, expr); !reflect.DeepEqual(got, BoolNode(tt.want)) {
			t.Errorf("[%d] %v, want %v", i, got, tt.want)
		}
	}
}

func TestJson(t *testing.T) {
	tests := []struct {
		name    string
		context Context
		args    string
		want    Node
	}{
		//{"", NewJSONContext(JSONRaw.JSON), `age`, IntNode(37)},
		{"", NewJSONContext(JSONRaw.JSON), `age + 1`, IntNode(38)},
		{"", NewJSONContext(JSONRaw.JSON), `(age + 1) * 2`, IntNode(76)},
		{"", NewJSONContext(JSONRaw.JSON), `(1 + 1) * age`, IntNode(74)},
		{"", NewJSONContext(JSONRaw.JSON), `name.first`, StringNode("Tom")},
		{"1", NewJSONContext(JSONRaw.JSON), `movie.1111`, New("[{\"1111\": \"Tom\", \"last\": \"Anderson\"},{\"first\": \"Neo\"}]")},
		{"1", NewJSONContext(JSONRaw.JSON), `movie.1111[0].1111`, StringNode("Tom")},
		{"1", NewJSONContext(JSONRaw.JSON), `movie.1111[#].1111`, New("[\"Tom\"]")},
		{"1", NewJSONContext(JSONRaw.JSON), `movie.1111[#]`, IntNode(2)},
	}
	for idx, tt := range tests {
		Convey(fmt.Sprintf("Test Float [%d]%s", idx, tt.name), t, func() {
			expr, err := ParseExpr(tt.args)
			So(err, ShouldBeNil)
			So(string(eval(tt.context, expr).Raw()), ShouldEqual, string(tt.want.Raw()))
		})
	}
}
