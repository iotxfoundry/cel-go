package cel_test

import (
	"testing"

	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/common/types"
	"github.com/google/cel-go/common/types/ref"
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
		"3u.add(4u)": types.Uint(7),
		"3u.add(4)":  types.Int(7),
		"3u.add(-4)": types.Int(-1),
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
		"10u.sub(4u)": types.Uint(6),
		"3u.sub(10)":  types.Int(-7),
		"10u.sub(3.5)": types.Double(6.5),
		// sub: double receiver
		"10.5.sub(3.0)": types.Double(7.5),
		"10.5.sub(3)":   types.Double(7.5),
		"3.5.sub(1u)":   types.Double(2.5),

		// mul: int receiver
		"3.mul(4)":    types.Int(12),
		"3.mul(4u)":   types.Int(12),
		"-3.mul(4u)":  types.Int(-12),
		"3.mul(4.5)":  types.Double(13.5),
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
		"10u.div(4u)": types.Uint(2),
		"10u.div(4)":  types.Int(2),
		"10u.div(4.0)": types.Double(2.5),
		// div: double receiver
		"10.0.div(4.0)": types.Double(2.5),
		"10.0.div(4)":   types.Double(2.5),
		"10.0.div(4u)":  types.Double(2.5),

		// mod: int receiver
		"10.mod(4)":    types.Int(2),
		"10.mod(4u)":   types.Int(2),
		"-10.mod(4u)":  types.Int(-2),
		// mod: uint receiver
		"10u.mod(4u)": types.Uint(2),
		"10u.mod(4)":  types.Int(2),
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
