package cel_test

import (
	"testing"

	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/common/types"
	"github.com/google/cel-go/common/types/ref"
	"github.com/google/cel-go/ext"
	compute "github.com/iotxfoundry/cel-go"
)

func TestMath(tt *testing.T) {

	tests := map[string]struct {
		source string
		out    ref.Val
	}{
		// randf
		"math.randf_none": {
			source: "0.0 <= math.randf() && math.randf() < 1.0",
			out:    types.Bool(true),
		},
		"math.randf_int_32": {
			source: "0.0 <= math.randf(32) && math.randf(32) < 1.0",
			out:    types.Bool(true),
		},
		"math.randf_int_64": {
			source: "0.0 <= math.randf(64) && math.randf(64) < 1.0",
			out:    types.Bool(true),
		},
		"math.randf_uint_32": {
			source: "0.0 <= math.randf(uint(32)) && math.randf(uint(32)) < 1.0",
			out:    types.Bool(true),
		},
		"math.randf_uint_64": {
			source: "0.0 <= math.randf(uint(64)) && math.randf(uint(64)) < 1.0",
			out:    types.Bool(true),
		},
		"math.randf_double_32": {
			source: "0.0 <= math.randf(32.0) && math.randf(32.0) < 1.0",
			out:    types.Bool(true),
		},
		"math.randf_double_64": {
			source: "0.0 <= math.randf(64.0) && math.randf(64.0) < 1.0",
			out:    types.Bool(true),
		},
		// randi
		"math.randi_none": {
			source: "0 <= math.randi() && math.randi() <= 9223372036854775807",
			out:    types.Bool(true),
		},
		"math.randi_int_32": {
			source: "0 <= math.randi(32) && math.randi(32) <= 2147483647",
			out:    types.Bool(true),
		},
		"math.randi_int_32_int_10": {
			source: "0 <= math.randi(32, 10) && math.randi(32, 10) < 10",
			out:    types.Bool(true),
		},
		"math.randi_int_32_uint_10": {
			source: "0 <= math.randi(32, uint(10)) && math.randi(32, uint(10)) < 10",
			out:    types.Bool(true),
		},
		"math.randi_int_32_double_10": {
			source: "0 <= math.randi(32, double(10)) && math.randi(32, double(10)) < 10",
			out:    types.Bool(true),
		},
		"math.randi_int_64": {
			source: "0 <= math.randi(64) && math.randi(64) <= 9223372036854775807",
			out:    types.Bool(true),
		},
		"math.randi_int_64_int_10": {
			source: "0 <= math.randi(64, 10) && math.randi(64, 10) < 10",
			out:    types.Bool(true),
		},
		"math.randi_int_64_uint_10": {
			source: "0 <= math.randi(64, uint(10)) && math.randi(64, uint(10)) < 10",
			out:    types.Bool(true),
		},
		"math.randi_int_64_double_10": {
			source: "0 <= math.randi(64, double(10)) && math.randi(64, double(10)) < 10",
			out:    types.Bool(true),
		},
		"math.randi_uint_32": {
			source: "0 <= math.randi(uint(32)) && math.randi(uint(32)) <= 2147483647",
			out:    types.Bool(true),
		},
		"math.randi_uint_32_int_10": {
			source: "0 <= math.randi(uint(32), 10) && math.randi(uint(32), 10) < 10",
			out:    types.Bool(true),
		},
		"math.randi_uint_32_uint_10": {
			source: "0 <= math.randi(uint(32), uint(10)) && math.randi(uint(32), uint(10)) < 10",
			out:    types.Bool(true),
		},
		"math.randi_uint_32_double_10": {
			source: "0 <= math.randi(uint(32), double(10)) && math.randi(uint(32), double(10)) < 10",
			out:    types.Bool(true),
		},
		"math.randi_uint_64": {
			source: "0 <= math.randi(uint(64)) && math.randi(uint(64)) <= 9223372036854775807",
			out:    types.Bool(true),
		},
		"math.randi_uint_64_int_10": {
			source: "0 <= math.randi(uint(64), 10) && math.randi(uint(64), 10) < 10",
			out:    types.Bool(true),
		},
		"math.randi_uint_64_uint_10": {
			source: "0 <= math.randi(uint(64), uint(10)) && math.randi(uint(64), uint(10)) < 10",
			out:    types.Bool(true),
		},
		"math.randi_uint_64_double_10": {
			source: "0 <= math.randi(uint(64), double(10)) && math.randi(uint(64), double(10)) < 10",
			out:    types.Bool(true),
		},
		"math.randi_double_32": {
			source: "0 <= math.randi(32.0) && math.randi(32.0) <= 2147483647",
			out:    types.Bool(true),
		},
		"math.randi_double_32_int_10": {
			source: "0 <= math.randi(32.0, 10) && math.randi(32.0, 10) < 10",
			out:    types.Bool(true),
		},
		"math.randi_double_32_uint_10": {
			source: "0 <= math.randi(32.0, uint(10)) && math.randi(32.0, uint(10)) < 10",
			out:    types.Bool(true),
		},
		"math.randi_double_32_double_10": {
			source: "0 <= math.randi(32.0, double(10)) && math.randi(32.0, double(10)) < 10",
			out:    types.Bool(true),
		},
		"math.randi_double_64": {
			source: "0 <= math.randi(64.0) && math.randi(64.0) <= 9223372036854775807",
			out:    types.Bool(true),
		},
		"math.randi_double_64_int_10": {
			source: "0 <= math.randi(64.0, 10) && math.randi(64.0, 10) < 10",
			out:    types.Bool(true),
		},
		"math.randi_double_64_uint_10": {
			source: "0 <= math.randi(64.0, uint(10)) && math.randi(64.0, uint(10)) < 10",
			out:    types.Bool(true),
		},
		"math.randi_double_64_double_10": {
			source: "0 <= math.randi(64.0, double(10)) && math.randi(64.0, double(10)) < 10",
			out:    types.Bool(true),
		},
		// randui
		"math.randui_none": {
			source: "uint(0) <= math.randui()",
			out:    types.Bool(true),
		},
		"math.randui_int_32": {
			source: "uint(0) <= math.randui(32) && math.randui(32) <= uint(4294967295)",
			out:    types.Bool(true),
		},
		"math.randui_int_64": {
			source: "uint(0) <= math.randui(64)",
			out:    types.Bool(true),
		},
		"math.randui_uint_32": {
			source: "uint(0) <= math.randui(uint(32)) && math.randui(uint(32)) <= uint(4294967295)",
			out:    types.Bool(true),
		},
		"math.randui_uint_64": {
			source: "uint(0) <= math.randui(uint(64))",
			out:    types.Bool(true),
		},
		"math.randui_double_32": {
			source: "uint(0) <= math.randui(32.0) && math.randui(32.0) <= uint(4294967295)",
			out:    types.Bool(true),
		},
		"math.randui_double_64": {
			source: "uint(0) <= math.randui(64.0)",
			out:    types.Bool(true),
		},
		// least
		"math.least_int": {
			source: "math.least(32)",
			out:    types.Int(32),
		},
		"math.least_uint": {
			source: "math.least(uint(16))",
			out:    types.Uint(16),
		},
		"math.least_double": {
			source: "math.least(16.0)",
			out:    types.Double(16),
		},
	}

	env, err := cel.NewEnv(
		cel.Variable("buff", cel.IntType),
		compute.ComputeLib(),
		ext.Math(),
	)
	if err != nil {
		tt.Errorf("environment creation error: %v\n", err)
		tt.FailNow()
	}

	for k, v := range tests {
		tt.Run(k, func(t *testing.T) {
			// Check iss for error in both Parse and Check.
			ast, iss := env.Compile(v.source)
			if iss.Err() != nil {
				t.Error(iss.Err())
				t.FailNow()
			}
			prg, err := env.Program(ast)
			if err != nil {
				t.Errorf("Program creation error: %v\n", err)
				t.FailNow()
			}

			out, _, err := prg.Eval(map[string]interface{}{})
			if err != nil {
				t.Errorf("Evaluation error: %v\n", err)
				t.FailNow()
			}

			t.Logf("%s -> %v", v.source, out)
			if out != v.out {
				t.FailNow()
			}
		})
	}
}
