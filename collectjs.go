package tdtl

import (
	"errors"
	"fmt"
	"strings"

	"github.com/tkeel-io/tdtl/pkg/json/jsonparser"
)

var EmptyBytes = []byte("")

type Collect = JSONNode

func New(raw string) *Collect {
	return newCollect(Byte(raw))
}

func ByteNew(raw []byte) *Collect {
	return newCollect(raw)
}

func newCollect(data []byte) *Collect {
	collect := &Collect{}
	value := make([]byte, len(data))
	copy(value, data)
	collect.path = ""
	collect.value = value
	if _, jtype, _, err := jsonparser.Get(data); err == nil {
		collect.datatype = datetype(jtype)
	} else {
		collect.err = err
	}
	return collect
}

func newCollectFromGjsonResult(ret Result) *Collect {
	collect := &Collect{}
	collect.path = ""
	collect.datatype = datetype(ret)
	if collect.datatype == String {
		collect.value = []byte(ret.Str)
	} else {
		collect.value = []byte(ret.Raw)
	}

	return collect
}

func newCollectFromJsonparserResult(dataType jsonparser.ValueType, value []byte) *Collect {
	collect := &Collect{}
	collect.path = ""
	collect.value = []byte(value)
	collect.datatype = datetype(dataType)
	return collect
}

// GetError returns collect error.
func (cc *Collect) GetError() error {
	return cc.err
}

func (cc *Collect) Get(path ...string) *Collect {
	absPath := strings.Join(path, ".")
	if absPath == "" {
		return cc
	}
	ret := get(cc.value, absPath)
	return ret
}

func (cc *Collect) Set(path string, value Node) {
	cc.value, cc.err = set(cc.value, path, value.Raw())
}

func (cc *Collect) Append(path string, value Node) {
	cc.value, cc.err = add(cc.value, path, value.Raw())
}

func (cc *Collect) Del(path ...string) {
	cc.value = del(cc.value, path...)
}

func (cc *Collect) Copy() *JSONNode {
	return newCollect(cc.value)
}

func (cc *Collect) Foreach(fn ForeachHandle) {
	cc.value = forEach(cc.value, cc.datatype, fn)
}

func (cc *Collect) Map(handle MapHandle) {
	ret := cc.Copy()
	cc.Foreach(func(key []byte, value *Collect) {
		newValue := handle(key, value)
		ret.Set(string(key), newValue)
	})
	cc.value, cc.datatype = ret.value, ret.datatype
}

func (cc *Collect) GroupBy(path string) *Collect {
	if cc.datatype != Array {
		cc.err = fmt.Errorf("datatype is not array")
		return cc
	}

	ret := New("{}")
	cc.Foreach(func(key []byte, value *Collect) {
		keyValue := get(value.Raw(), path).String()
		if len(keyValue) == 0 {
			return
		}
		keyValue = strings.Replace(keyValue, ".", "_", -1)
		ret.Append(keyValue, value)
	})
	return ret
}

func (c *Collect) MergeBy(paths ...string) *Collect {
	if c.datatype != Array {
		c.err = fmt.Errorf("MergeBy: datatype is not array")
		return c
	}

	var err error
	ret := New("{}")
	c.Foreach(func(key []byte, value *Collect) {
		keys := make([]string, 0, len(paths))
		for _, path := range paths {
			keyValueRaw := value.Get(path).String()
			if len(keyValueRaw) == 0 {
				break
			}
			keys = append(keys, keyValueRaw)
		}

		if len(keys) == 0 {
			return
		}

		keyValue := strings.Join(keys, "+")
		keyValue = strings.Replace(keyValue, ".", "_", -1)

		nv := ret.Get(keyValue)
		nv = nv.Merge(value)
		if nv.err != nil {
			ret.err = err
		}
		ret.Set(keyValue, nv)
	})
	return ret
}

func (cc *Collect) SortBy(fn func(p1 *Collect, p2 *Collect) bool) {
	if cc.datatype != Array && cc.datatype != Object {
		cc.err = errors.New("SortBy:datatype is not array or object")
		return
	}
	carr := make([]*Collect, 0)
	cc.Foreach(func(key []byte, value *Collect) {
		carr = append(carr, value)
	})
	By(fn).Sort(carr)

	ret := New("[]")
	for _, c := range carr {
		ret.Append("", c)
	}
	cc.value = ret.value
	cc.datatype = ret.datatype
}

func (c *Collect) KeyBy(path string) *Collect {
	if c.datatype != Array {
		c.err = errors.New("KeyBy:datatype is not array")
	}

	ret := New("{}")
	c.Foreach(func(key []byte, value *Collect) {
		keyValue := value.Get(path)
		ret.Set(keyValue.String(), value)
	})

	return ret
}

func (cc *Collect) Merge(mc *Collect) *Collect {
	if cc.datatype != Object && mc.datatype != Object {
		cc.err = errors.New("datatype is not object")
		return cc
	}
	if cc.datatype == Null {
		return mc
	}

	mc.Foreach(func(key []byte, value *Collect) {
		cc.Set(string(key), value)
	})

	return cc
}

//
//func (c *Collect) Sort( path string) ([]byte, error) {
//	if c.datatype != Array {
//		c.err = errors.New("SortBy:datatype is not array")
//		return nil, errors.New("datatype is not array")
//	}
//
//	var err error
//	ret := []byte("[]")
//	c.Foreach(func(key []byte, value *Collect) {
//		keyValue := get(value.Raw(), path).Raw()
//		if keyValue[0] == '"' && keyValue[len(keyValue)-1] == '"' {
//			keyValue = keyValue[1 : len(keyValue)-1]
//		}
//		if ret, err = jsonparser.Append(ret, value.Raw(), string(keyValue)); nil != err {
//			c.err = err
//		}
//	})
//
//	return ret, c.err
//}

func Combine(cKey *Collect, cValue *Collect) ([]byte, error) {
	if cKey.datatype != Array {
		return nil, errors.New("datatype is not array")
	} else if cValue.datatype != Array {
		return nil, errors.New("datatype is not array")
	}

	var (
		idx int
		err error
		ret = []byte("{}")
	)

	cKey.Foreach(func(key []byte, value *Collect) {
		if ret, err = jsonparser.Set(ret, get(cValue.value, fmt.Sprintf("[%d]", idx)).Raw(), value.String()); nil != err {
			cKey.err = err
		}
		idx++
	})
	return ret, cKey.err
}