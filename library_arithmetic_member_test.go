package cel_test

import (
	"strings"
	"testing"

	"cel.dev/cel-go/cel"
	"cel.dev/cel-go/common/types"
	"cel.dev/cel-go/common/types/ref"
	compute "github.com/iotxfoundry/cel-go"
)

func TestArithmeticMember(tt *testing.T) {
	env, err := cel.NewEnv(compute.ComputeLib())
	if err != nil {
		tt.Fatal(err)
	}

	tests := map[string]ref.Val{
		// add: int receiver
		"3.add(4)":   types.Int(7),
		"3.add(4u)":  types.Int(7),
		"-3.add(4u)": types.Int(1),
		"3.add(4.5)": types.Double(7.5),
		// add: uint receiver
		"3u.add(4u)":  types.Uint(7),
		"3u.add(4)":   types.Int(7),
		"3u.add(-4)":  types.Int(-1),
		"3u.add(4.5)": types.Double(7.5),
		// add: double receiver
		"4.5.add(3.0)": types.Double(7.5),
		"4.5.add(3)":   types.Double(7.5),
		"4.5.add(3u)":  types.Double(7.5),

		// sub: int receiver
		"10.sub(4)":   types.Int(6),
		"10.sub(4u)":  types.Int(6),
		"-5.sub(4u)":  types.Int(-9),
		"10.sub(3.5)": types.Double(6.5),
		// sub: uint receiver
		"10u.sub(4u)":  types.Uint(6),
		"3u.sub(10)":   types.Int(-7),
		"10u.sub(3.5)": types.Double(6.5),
		// sub: double receiver
		"10.5.sub(3.0)": types.Double(7.5),
		"10.5.sub(3)":   types.Double(7.5),
		"3.5.sub(1u)":   types.Double(2.5),

		// mul: int receiver
		"3.mul(4)":   types.Int(12),
		"3.mul(4u)":  types.Int(12),
		"-3.mul(4u)": types.Int(-12),
		"3.mul(4.5)": types.Double(13.5),
		// mul: uint receiver
		"3u.mul(4u)":  types.Uint(12),
		"3u.mul(4)":   types.Int(12),
		"3u.mul(-4)":  types.Int(-12),
		"3u.mul(4.5)": types.Double(13.5),
		// mul: double receiver
		"4.5.mul(3.0)": types.Double(13.5),
		"4.5.mul(3)":   types.Double(13.5),
		"4.5.mul(3u)":  types.Double(13.5),

		// div: int receiver
		"10.div(4)":   types.Int(2),
		"10.div(4u)":  types.Int(2),
		"-10.div(4u)": types.Int(-2),
		"10.div(4.0)": types.Double(2.5),
		// div: uint receiver
		"10u.div(4u)":  types.Uint(2),
		"10u.div(4)":   types.Int(2),
		"10u.div(4.0)": types.Double(2.5),
		// div: double receiver
		"10.0.div(4.0)": types.Double(2.5),
		"10.0.div(4)":   types.Double(2.5),
		"10.0.div(4u)":  types.Double(2.5),

		// mod: int receiver
		"10.mod(4)":   types.Int(2),
		"10.mod(4u)":  types.Int(2),
		"-10.mod(4u)": types.Int(-2),
		// mod: uint receiver
		"10u.mod(4u)": types.Uint(2),
		"10u.mod(4)":  types.Int(2),

		// cross-type boundary values that still fit in the result type
		"(-9223372036854775808).add(18446744073709551615u)": types.Int(9223372036854775807),
		"0.sub(9223372036854775808u)":                       types.Int(-9223372036854775808),
		"(-1).add(9223372036854775808u)":                    types.Int(9223372036854775807),
		"(-9223372036854775808).div(9223372036854775808u)":  types.Int(-1),
		"(-2).mul(4611686018427387904u)":                    types.Int(-9223372036854775808),
		"9223372036854775808u.mul(-1)":                      types.Int(-9223372036854775808),
		"18446744073709551615u.div(3)":                      types.Int(6148914691236517205),
		"18446744073709551615u.mod(3)":                      types.Int(0),
		"9223372036854775808u.div(-1)":                      types.Int(-9223372036854775808),
		"18446744073709551615u.mod(9223372036854775808u)":   types.Int(9223372036854775807),
		"18446744073709551615u.mod(-3)":                     types.Int(0),
		"3.mod(18446744073709551615u)":                      types.Int(3),
		"3.div(18446744073709551615u)":                      types.Int(0),
		"(-3).div(18446744073709551615u)":                   types.Int(0),
		"3u.mul(-3)":                                        types.Int(-9),
	}

	for expr, want := range tests {
		tt.Run(expr, func(t *testing.T) {
			ast, iss := env.Compile(expr)
			if iss.Err() != nil {
				t.Error(iss.Err())
				return
			}
			prg, err := env.Program(ast)
			if err != nil {
				t.Error(err)
				return
			}
			out, _, err := prg.Eval(map[string]any{})
			if err != nil {
				t.Error(err)
				return
			}
			t.Logf("%s → %v", expr, out)
			result := out.Equal(want)
			if !result.Value().(bool) {
				t.Errorf("got %v, want %v", out, want)
			}
		})
	}
}

// TestArithmeticMemberErrors asserts that boundary violations surface as
// normal CEL evaluation errors instead of runtime panics or silent wrapping.
func TestArithmeticMemberErrors(tt *testing.T) {
	env, err := cel.NewEnv(compute.ComputeLib())
	if err != nil {
		tt.Fatal(err)
	}

	tests := map[string]string{
		// div/mod by zero
		"div_int_zero":      `10.div(0)`,
		"div_int_uint_zero": `10.div(0u)`,
		"div_uint_zero":     `10u.div(0u)`,
		"div_uint_int_zero": `10u.div(0)`,
		"mod_int_zero":      `10.mod(0)`,
		"mod_int_uint_zero": `10.mod(0u)`,
		"mod_uint_zero":     `10u.mod(0u)`,
		"mod_uint_int_zero": `10u.mod(0)`,

		// int receiver overflow
		"add_int_overflow":  `9223372036854775807.add(1)`,
		"add_int_uint_over": `9223372036854775807.add(1u)`,
		"sub_int_overflow":  `(-9223372036854775808).sub(1)`,
		"sub_int_uint_over": `(-9223372036854775808).sub(1u)`,
		"mul_int_overflow":  `9223372036854775807.mul(2)`,
		"mul_int_uint_over": `4611686018427387904.mul(2u)`,
		"div_int_min_over":  `(-9223372036854775808).div(-1)`,
		"mod_int_min_over":  `(-9223372036854775808).mod(-1)`,

		// uint receiver overflow
		"add_uint_overflow":     `18446744073709551615u.add(1u)`,
		"add_uint_int_over":     `18446744073709551615u.add(-1)`,
		"sub_uint_int_under":    `0u.sub(18446744073709551615u)`,
		"mul_uint_overflow":     `18446744073709551615u.mul(2u)`,
		"mul_uint_int_over":     `18446744073709551615u.mul(-1)`,
		"div_uint_int_over":     `18446744073709551615u.div(-1)`,
		"mul_uint64_neg_minint": `9223372036854775809u.mul(-1)`,
	}

	for name, src := range tests {
		tt.Run(name, func(t *testing.T) {
			ast, iss := env.Compile(src)
			if iss.Err() != nil {
				t.Error(iss.Err())
				return
			}
			prg, err := env.Program(ast)
			if err != nil {
				t.Error(err)
				return
			}
			out, _, err := prg.Eval(map[string]any{})
			if err == nil {
				t.Errorf("expected error, got %v", out)
				return
			}
			// A well-formed CEL eval error must not be a recovered runtime
			// panic ("internal error: ...").
			if strings.Contains(err.Error(), "internal error") {
				t.Errorf("panic leaked through as internal error: %v", err)
			}
		})
	}
}
