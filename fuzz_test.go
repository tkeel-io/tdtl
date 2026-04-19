package tdtl

import "testing"

// FuzzParseExpr verifies that arbitrary input to ParseExpr never panics.
func FuzzParseExpr(f *testing.F) {
	seeds := []string{
		`1 + 2`,
		`temperature * 2`,
		`case 1 when 1 then 'a' else 'b'`,
		`abs(0 - 5)`,
		`sum(arr.#.val)`,
		`not true`,
		`'hello' + ' world'`,
		`1 / 0`,
		``,
		`@#$%`,
	}
	for _, s := range seeds {
		f.Add(s)
	}
	f.Fuzz(func(t *testing.T, input string) {
		// Must not panic regardless of input.
		expr, _ := ParseExpr(input)
		if expr != nil {
			ctx := NewJSONContext(`{}`)
			Eval(ctx, expr)
		}
	})
}
