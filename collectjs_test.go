package tdtl

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

var raw = Byte(`{"cpu":1,"mem": ["lo0", "eth1", "eth2"],"a":[{"v":0},{"v":1},{"v":2}],"b":[{"v":{"cv":1}},{"v":{"cv":2}},{"v":{"cv":3}}],"where": 10,"metadata": {"name": "Light1", "price": 11.05}}`)
var rawArray = Byte(`[{"v":{"cv":1}},{"v":{"cv":2}},{"v":{"cv":3}}]`)
var rawEmptyArray = Byte(`[]`)
var rawEmptyObject = Byte(`{}`)

func TestCollect_New(t *testing.T) {
	tests := []struct {
		name string
		path string
		raw  *Collect
		want string
	}{
		{"1", "", New(raw), "{}"},
		{"2", "cpu", NewString("1"), `{"cpu":"1"}`},
		{"2", "cpu", NewBool(true), `{"cpu":true}`},
		{"2", "cpu", NewInt64(3456), `{"cpu":3456}`},
		{"2", "cpu", NewFloat64(1 / 3.0), `{"cpu":0.333333}`},
	}
	for _, tt := range tests {
		got := New("{}")
		got.Set(tt.path, tt.raw)
		t.Run(tt.name, func(t *testing.T) {
			if !reflect.DeepEqual(string(got.Raw()), tt.want) {
				t.Errorf("Get() = %v, want %v", string(got.Raw()), tt.want)
			}
		})
	}
}
func TestCollect_Get(t *testing.T) {
	tests := []struct {
		name string
		raw  []byte
		path string
		want interface{}
	}{
		{"1", raw, "", string(raw)},
		{"2", raw, "cpu", "1"},
		{"2.1", raw, "mem[0]", "\"lo0\""},
		{"3", raw, "a", `[{"v":0},{"v":1},{"v":2}]`},
		{"4", raw, "a[0]", `{"v":0}`},
		{"5", raw, "a[1]", `{"v":1}`},
		{"6", raw, "a[2]", `{"v":2}`},
		{"7", raw, "a[#]", `3`}, // count
		{"8", raw, "a[#].v", `[0,1,2]`},
		{"9", raw, "b[0].v", `{"cv":1}`},
		{"10", raw, "b[1].v", `{"cv":2}`},
		{"11", raw, "b[2].v", `{"cv":3}`},
		{"12", raw, "b[#].v", `[{"cv":1},{"cv":2},{"cv":3}]`},
		{"13", raw, "b[1].v.cv", `2`},
		{"14", raw, "b[#].v.cv", `[1,2,3]`},
		{"14", rawArray, "[0]", `{"v":{"cv":1}}`},
	}
	for _, tt := range tests {
		cc := newCollect(tt.raw)
		t.Run(tt.name, func(t *testing.T) {
			if got := cc.Get(tt.path); !reflect.DeepEqual(string(got.Raw()), tt.want) {
				t.Errorf("Get() = %v, want %v", string(got.Raw()), tt.want)
			}
		})
	}
}

func TestCollect_Set(t *testing.T) {
	tests := []struct {
		name  string
		path  string
		value *Collect
		want  interface{}
	}{
		{"1", "metadata.name", New(`"abc"`), `{"cpu":2,"mem": ["lo0", "eth1", "eth2"],"a":[{"v":0},{"v":1},{"v":2}],"b":[{"v":{"cv":1}},{"v":{"cv":2}},{"v":{"cv":3}}],"where": 10,"metadata": {"name": "Light1", "price": 11.05}}`},
		{"2", "cpu", New("2"), `{"cpu":2,"mem": ["lo0", "eth1", "eth2"],"a":[{"v":0},{"v":1},{"v":2}],"b":[{"v":{"cv":1}},{"v":{"cv":2}},{"v":{"cv":3}}],"where": 10,"metadata": {"name": "Light1", "price": 11.05}}`},
		{"3", "a", New(`{"v":0}`), `{"cpu":1,"mem": ["lo0", "eth1", "eth2"],"a":{"v":0},"b":[{"v":{"cv":1}},{"v":{"cv":2}},{"v":{"cv":3}}],"where": 10,"metadata": {"name": "Light1", "price": 11.05}}`},
		{"4", "a[0]", New(`0`), `{"cpu":1,"mem": ["lo0", "eth1", "eth2"],"a":[0,{"v":1},{"v":2}],"b":[{"v":{"cv":1}},{"v":{"cv":2}},{"v":{"cv":3}}],"where": 10,"metadata": {"name": "Light1", "price": 11.05}}`},
		{"5", "a[0].v", New(`{"v":0}`), `{"cpu":1,"mem": ["lo0", "eth1", "eth2"],"a":[{"v":{"v":0}},{"v":1},{"v":2}],"b":[{"v":{"cv":1}},{"v":{"cv":2}},{"v":{"cv":3}}],"where": 10,"metadata": {"name": "Light1", "price": 11.05}}`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cc := newCollect(raw)
			cc.Set(tt.path, tt.value)
			if got := cc.Raw(); !reflect.DeepEqual(string(got), tt.want) {
				t.Errorf("\nGet()  = %v\nWant() = %v", string(got), tt.want)
			}
		})
	}
}

func TestCollect_Set2(t *testing.T) {
	tests := []struct {
		name  string
		path  string
		value *Collect
		want  interface{}
	}{
		{"1", "metadata._name", New(`"abc"`), `{"metadata":{"_name":"abc"}}`},
		{"2", "metadata._name", New(`[20]`), `{"metadata":{"_name":[20]}}`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cc := newCollect(rawEmptyObject)
			cc.Set(tt.path, tt.value)
			if got := cc.Raw(); !reflect.DeepEqual(string(got), tt.want) {
				t.Errorf("\nGet()  = %v\nWant() = %v", string(got), tt.want)
			}
		})
	}
}

func TestCollect_Append(t *testing.T) {
	tests := []struct {
		name    string
		raw     []byte
		path    string
		value   *Collect
		wantErr interface{}
		want    interface{}
	}{
		{"1", raw, "cpu", New("2"), `Unknown value type`, raw},
		{"2", raw, "mem", New("2"), nil, `{"cpu":1,"mem": ["lo0", "eth1", "eth2",2],"a":[{"v":0},{"v":1},{"v":2}],"b":[{"v":{"cv":1}},{"v":{"cv":2}},{"v":{"cv":3}}],"where": 10,"metadata": {"name": "Light1", "price": 11.05}}`},
		{"3", raw, "a", New(`{"v":11}`), nil, `{"cpu":1,"mem": ["lo0", "eth1", "eth2"],"a":[{"v":0},{"v":1},{"v":2},{"v":11}],"b":[{"v":{"cv":1}},{"v":{"cv":2}},{"v":{"cv":3}}],"where": 10,"metadata": {"name": "Light1", "price": 11.05}}`},
		{"4", raw, "a[0]", New(`0`), `Unknown value type`, raw},
		{"5", raw, "a[0].v", New(`{"v":0}`), `Unknown value type`, raw},
		{"5", rawEmptyArray, "", New(`{"v":0}`), nil, `[{"v":0}]`},
		{"6", []byte(`{}`), "metrics.cpus", New(`20`), nil, `{"metrics":{"cpus":[20]}}`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cc := newCollect(tt.raw)
			cc.Append(tt.path, tt.value)
			if tt.wantErr != nil {
				if !reflect.DeepEqual(cc.Error().Error(), tt.wantErr) {
					t.Errorf("Get(2) = %v, want %v", cc.Error(), tt.wantErr)
				}
			} else {
				if cc.Error() != nil {
					t.Errorf("Get(1) = %v, want %v", cc.Error(), tt.wantErr)
				}
				if !reflect.DeepEqual(cc.String(), tt.want) {
					t.Errorf("Get() = %v, want %v", cc.String(), tt.want)
				}
			}
		})
	}
}

func TestCollect_Append2(t *testing.T) {
	cc := New([]byte(`{}`))
	cc.Append("metrics.cpus", New(`20`))
	assert.Equal(t, `{"metrics":{"cpus":[20]}}`, cc.String())
}

func TestCollect_Del(t *testing.T) {
	tests := []struct {
		name string
		path []string
		want interface{}
	}{
		{"2", []string{"cpu"}, `{"mem": ["lo0", "eth1", "eth2"],"a":[{"v":0},{"v":1},{"v":2}],"b":[{"v":{"cv":1}},{"v":{"cv":2}},{"v":{"cv":3}}],"where": 10,"metadata": {"name": "Light1", "price": 11.05}}`},
		{"3", []string{"a", "b", "cpu"}, `{"mem": ["lo0", "eth1", "eth2"],"where": 10,"metadata": {"name": "Light1", "price": 11.05}}`},
		{"4", []string{"a[0]"}, `{"cpu":1,"mem": ["lo0", "eth1", "eth2"],"a":[{"v":1},{"v":2}],"b":[{"v":{"cv":1}},{"v":{"cv":2}},{"v":{"cv":3}}],"where": 10,"metadata": {"name": "Light1", "price": 11.05}}`},
		{"5", []string{"a[0].v"}, `{"cpu":1,"mem": ["lo0", "eth1", "eth2"],"a":[{},{"v":1},{"v":2}],"b":[{"v":{"cv":1}},{"v":{"cv":2}},{"v":{"cv":3}}],"where": 10,"metadata": {"name": "Light1", "price": 11.05}}`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cc := newCollect(raw)
			cc.Del(tt.path...)
			if got := cc.Raw(); !reflect.DeepEqual(string(got), tt.want) {
				t.Errorf("Get() = %v, want %v", string(got), tt.want)
			}
		})
	}
}

func Example_Combine() {
	collection := New(`["name", "number"]`)
	collection2 := New(`["Mohamed Salah", 11]`)
	combine, _ := Combine(collection, collection2)
	fmt.Println(string(combine))

	// Output:
	// {"name":"Mohamed Salah","number":11}
}

var rawGroup = Byte(`[
	{"count": "1","product": "Chair","manufacturer": "IKEA"},
	{"sum": "10","product": "Desk","manufacturer": "IKEA"},
	{"product": "Chair","manufacturer": "Herman Miller"}
]`)

func Example_GroupBy() {
	cc := New(rawGroup)
	ret := cc.GroupBy("manufacturer") //node_memory_MemTotal_bytes
	fmt.Println(ret.String(), ret.Error())

	// Output:
	// {"IKEA":[{"count": "1","product": "Chair","manufacturer": "IKEA"},{"sum": "10","product": "Desk","manufacturer": "IKEA"}],"Herman Miller":[{"product": "Chair","manufacturer": "Herman Miller"}]} <nil>
}

func Example_MergeBy() {
	cc := New(rawGroup)
	ret := cc.MergeBy("product", "manufacturer") //node_memory_MemTotal_bytes
	fmt.Println(ret.String(), ret.Error())

	// Output:
	// {"Chair+IKEA":{"count": "1","product": "Chair","manufacturer": "IKEA"},"Desk+IKEA":{"sum": "10","product": "Desk","manufacturer": "IKEA"},"Chair+Herman Miller":{"product": "Chair","manufacturer": "Herman Miller"}} <nil>
}

func Example_KeyBy() {
	cc := New(rawGroup)
	ret := cc.KeyBy("manufacturer") //node_memory_MemTotal_bytes
	fmt.Println(ret.String(), ret.Error())

	// Output:
	// {"IKEA":{"sum": "10","product": "Desk","manufacturer": "IKEA"},"Herman Miller":{"product": "Chair","manufacturer": "Herman Miller"}} <nil>
}

func Example_Merge() {
	var rawObject1 = New(`{"id": 1,"price": 29}`)
	var rawObject2 = New(`{"price": "229","discount": false}`)
	ret := rawObject1.Merge(rawObject2)
	fmt.Println(rawObject1.String(), rawObject1.Error())
	fmt.Println(ret.String(), ret.Error())

	var rawObject3 = rawObject1.Get("__not_exist")
	ret = rawObject3.Merge(rawObject2)
	fmt.Println(rawObject3.Error())
	fmt.Println(ret.String(), ret.Error())

	var rawObject4 = rawObject1.Get("__not_exist")
	ret = rawObject2.Merge(rawObject4)
	fmt.Println(rawObject2.String(), rawObject2.Error())
	fmt.Println(ret.String(), ret.Error())

	// Output:
	//{"id": 1,"price": "229","discount":false} <nil>
	//{"id": 1,"price": "229","discount":false} <nil>
	//<nil>
	//{"price": "229","discount": false} <nil>
	//{"price": "229","discount": false} <nil>
	//{"price": "229","discount": false} <nil>
}

func Example_Demo() {
	collection1 := New(`{"status":"success","data":{"resultType":"vector","result":[{"metric":{"__name__":"node_memory_MemAvailable_bytes","instance":"192.168.14.102:9100","job":"linux"},"value":[1620999810.899,"6519189504"]},{"metric":{"__name__":"node_memory_MemAvailable_bytes","instance":"192.168.14.146:9100","job":"linux"},"value":[1620999810.899,"1787977728"]},{"metric":{"__name__":"node_memory_MemAvailable_bytes","instance":"192.168.21.163:9100","job":"linux"},"value":[1620999810.899,"5775802368"]},{"metric":{"__name__":"node_memory_MemAvailable_bytes","instance":"192.168.21.174:9100","job":"linux"},"value":[1620999810.899,"19626115072"]},{"metric":{"__name__":"node_memory_MemAvailable_bytes","instance":"localhost:9100","job":"linux"},"value":[1620999810.899,"3252543488"]},{"metric":{"__name__":"node_memory_MemTotal_bytes","instance":"192.168.14.102:9100","job":"linux"},"value":[1620999810.899,"8203091968"]},{"metric":{"__name__":"node_memory_MemTotal_bytes","instance":"192.168.14.146:9100","job":"linux"},"value":[1620999810.899,"8203091968"]},{"metric":{"__name__":"node_memory_MemTotal_bytes","instance":"192.168.21.163:9100","job":"linux"},"value":[1620999810.899,"8202657792"]},{"metric":{"__name__":"node_memory_MemTotal_bytes","instance":"192.168.21.174:9100","job":"linux"},"value":[1620999810.899,"25112969216"]},{"metric":{"__name__":"node_memory_MemTotal_bytes","instance":"localhost:9100","job":"linux"},"value":[1620999810.899,"3972988928"]}]}}`)
	result := collection1.Get("data.result")
	result.Map(func(key []byte, c *Collect) Node {
		val := New("{}")
		val.Set("timestamp", c.Get("value[0]"))
		val.Set("value", c.Get("value[1]"))
		ret := New("{}")
		ret.Set(c.Get("metric.__name__").String(), val)
		ret.Set("instance", c.Get("metric.instance"))
		return ret
	})
	ret := result.GroupBy("instance") //node_memory_MemTotal_bytes

	metricValue := func(p1, p2 *Collect) bool {
		return bytes.Compare(p1.Get("[0]").Raw(), p2.Get("[0]").Raw()) > 0
	}

	ret.SortBy(metricValue)
	fmt.Println(string(result.Raw()))

	// Output:
	// [{"node_memory_MemAvailable_bytes":{"timestamp":1620999810.899,"value":"6519189504"},"instance":"192.168.14.102:9100"},{"node_memory_MemAvailable_bytes":{"timestamp":1620999810.899,"value":"1787977728"},"instance":"192.168.14.146:9100"},{"node_memory_MemAvailable_bytes":{"timestamp":1620999810.899,"value":"5775802368"},"instance":"192.168.21.163:9100"},{"node_memory_MemAvailable_bytes":{"timestamp":1620999810.899,"value":"19626115072"},"instance":"192.168.21.174:9100"},{"node_memory_MemAvailable_bytes":{"timestamp":1620999810.899,"value":"3252543488"},"instance":"localhost:9100"},{"node_memory_MemTotal_bytes":{"timestamp":1620999810.899,"value":"8203091968"},"instance":"192.168.14.102:9100"},{"node_memory_MemTotal_bytes":{"timestamp":1620999810.899,"value":"8203091968"},"instance":"192.168.14.146:9100"},{"node_memory_MemTotal_bytes":{"timestamp":1620999810.899,"value":"8202657792"},"instance":"192.168.21.163:9100"},{"node_memory_MemTotal_bytes":{"timestamp":1620999810.899,"value":"25112969216"},"instance":"192.168.21.174:9100"},{"node_memory_MemTotal_bytes":{"timestamp":1620999810.899,"value":"3972988928"},"instance":"localhost:9100"}]
}

func Example_Demo2() {
	collection1 := New(`{"status":"success","data":{"resultType":"vector","result":[{"metric":{"__name__":"node_memory_MemAvailable_bytes","instance":"192.168.14.102:9100","job":"linux"},"value":[1620999810.899,"1"]},{"metric":{"__name__":"node_memory_MemAvailable_bytes","instance":"192.168.14.146:9100","job":"linux"},"value":[1620999810.899,"3"]},{"metric":{"__name__":"node_memory_MemAvailable_bytes","instance":"192.168.21.163:9100","job":"linux"},"value":[1620999810.899,"2"]},{"metric":{"__name__":"node_memory_MemAvailable_bytes","instance":"192.168.21.174:9100","job":"linux"},"value":[1620999810.899,"19626115072"]},{"metric":{"__name__":"node_memory_MemAvailable_bytes","instance":"localhost:9100","job":"linux"},"value":[1620999810.899,"3252543488"]},{"metric":{"__name__":"node_memory_MemTotal_bytes","instance":"192.168.14.102:9100","job":"linux"},"value":[1620999810.899,"8203091968"]},{"metric":{"__name__":"node_memory_MemTotal_bytes","instance":"192.168.14.146:9100","job":"linux"},"value":[1620999810.899,"8203091968"]},{"metric":{"__name__":"node_memory_MemTotal_bytes","instance":"192.168.21.163:9100","job":"linux"},"value":[1620999810.899,"8202657792"]},{"metric":{"__name__":"node_memory_MemTotal_bytes","instance":"192.168.21.174:9100","job":"linux"},"value":[1620999810.899,"25112969216"]},{"metric":{"__name__":"node_memory_MemTotal_bytes","instance":"localhost:9100","job":"linux"},"value":[1620999810.899,"3972988928"]}]}}`)
	result := collection1.Get("data.result")
	result.Map(func(key []byte, c *Collect) Node {
		val := New("{}")
		val.Set("timestamp", c.Get("value[0]"))
		val.Set("value", c.Get("value[1]"))
		ret := New("{}")
		ret.Set(c.Get("metric.__name__").String(), val)
		ret.Set("instance", c.Get("metric.instance"))
		return ret
	})

	sorted := result.MergeBy("instance") //node_memory_MemTotal_bytes

	MemAvailable := func(p1, p2 *Collect) bool {
		return bytes.Compare(p1.Get("node_memory_MemAvailable_bytes.value").Raw(), p2.Get("node_memory_MemAvailable_bytes.value").Raw()) > 0
	}
	MemTotal := func(p1, p2 *Collect) bool {
		return bytes.Compare(p1.Get("node_memory_MemTotal_bytes.value").Raw(), p2.Get("node_memory_MemTotal_bytes.value").Raw()) < 0
	}

	sorted.SortBy(MemTotal)
	sorted.SortBy(MemAvailable)
	fmt.Println(string(sorted.Raw()))

	// Output:
	// [{"node_memory_MemAvailable_bytes":{"timestamp":1620999810.899,"value":"3252543488"},"instance":"localhost:9100","node_memory_MemTotal_bytes":{"timestamp":1620999810.899,"value":"3972988928"}},{"node_memory_MemAvailable_bytes":{"timestamp":1620999810.899,"value":"3"},"instance":"192.168.14.146:9100","node_memory_MemTotal_bytes":{"timestamp":1620999810.899,"value":"8203091968"}},{"node_memory_MemAvailable_bytes":{"timestamp":1620999810.899,"value":"2"},"instance":"192.168.21.163:9100","node_memory_MemTotal_bytes":{"timestamp":1620999810.899,"value":"8202657792"}},{"node_memory_MemAvailable_bytes":{"timestamp":1620999810.899,"value":"19626115072"},"instance":"192.168.21.174:9100","node_memory_MemTotal_bytes":{"timestamp":1620999810.899,"value":"25112969216"}},{"node_memory_MemAvailable_bytes":{"timestamp":1620999810.899,"value":"1"},"instance":"192.168.14.102:9100","node_memory_MemTotal_bytes":{"timestamp":1620999810.899,"value":"8203091968"}}]
}

//func Example_AAA() {
//
//	fmt.Println(gjson.Get(string(rawArray), "0"))
//	fmt.Println(gjson.Get(string(`["Mohamed Salah", 11]`), "0"))
//
//	// Output:
//	// {"v":{"cv":1}}
//	//Mohamed Salah
//}
