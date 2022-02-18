package tdtl

import "testing"

func Test_thePath(t *testing.T) {
	tests := []struct {
		name string
		path string
		want string
	}{
		{"", "", ""},
		{"", "a", "a"},
		{"", "a.b", "a.b"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := thePath(tt.path); got != tt.want {
				t.Errorf("thePath() = %v, want %v", got, tt.want)
			}
		})
	}
}
