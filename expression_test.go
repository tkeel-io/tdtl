package tdtl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Expression(t *testing.T) {
	tests := []struct {
		name       string
		path       string
		expression string
		params     map[string]Node
		want       string
	}{
		{
			name:       "test add",
			path:       "metrics.cpu_used",
			expression: "20 + 8",
			want:       "28",
		},
		{
			name:       "test sub",
			path:       "metrics.cpu_used",
			expression: "20 - 8",
			want:       "12",
		},
		{
			name:       "test mul",
			path:       "metrics.cpu_used",
			expression: "20 * 8",
			want:       "160",
		},
		{
			name:       "test div",
			path:       "metrics.cpu_used",
			expression: "20 / 8",
			want:       "2",
		},
		{
			name:       "test1",
			path:       "metrics.cpu_used",
			expression: "device234.metrics.cpu_used / 20",
			params: map[string]Node{
				"device234.metrics.cpu_used": NewInt64(100),
			},
			want: "5",
		},
		{
			name:       "test2",
			path:       "sysField._spacePath",
			expression: "device234.sysField._spacePath + '/tomas'",
			params: map[string]Node{
				"device234.sysField._spacePath": NewString("iotd-device123"),
			},
			want: "iotd-device123/tomas",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			exprIns, err := NewExpr(test.expression, nil)
			assert.Nil(t, err)

			result := exprIns.Eval(test.params)
			assert.Equal(t, test.want, string(result.String()))
		})
	}
}
