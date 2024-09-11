package cel_test

import (
	"testing"

	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/common/types"
)

func TestType(tt *testing.T) {

	tests := map[string]struct {
		in     int64
		source string
		out    *types.Type
	}{
		"type(6u)": {
			in:     6,
			source: "type(6u)",
			out:    types.UintType,
		},
		"type(6)": {
			in:     6,
			source: "type(6)",
			out:    types.IntType,
		},
		"type(6.0)": {
			in:     6,
			source: "type(6.0)",
			out:    types.DoubleType,
		},
		"type(6e0)": {
			in:     6,
			source: "type(6e0)",
			out:    types.DoubleType,
		},
		"type('abcd')": {
			in:     6,
			source: "type('abcd')",
			out:    types.StringType,
		},
		`type("abcd")`: {
			in:     6,
			source: `type("abcd")`,
			out:    types.StringType,
		},
		`type(r"abcd")`: {
			in:     6,
			source: `type(r"abcd")`,
			out:    types.StringType,
		},
		"type({'a': 6, 'b': 'bbb'})": {
			in:     6,
			source: "type({'a': 6, 'b': 'bbb'})",
			out:    types.MapType,
		},
		"type([1,2,3,4])": {
			in:     6,
			source: "type([1,2,3,4])",
			out:    types.ListType,
		},
		"type(true)": {
			in:     6,
			source: "type(true)",
			out:    types.BoolType,
		},
		"type(false)": {
			in:     6,
			source: "type(false)",
			out:    types.BoolType,
		},
		`type(b"\x31\x32\x33")`: {
			in:     6,
			source: `type(b"\x31\x32\x33")`,
			out:    types.BytesType,
		},
		`type(b"\303\277")`: {
			in:     6,
			source: `type(b"\303\277")`,
			out:    types.BytesType,
		},
		`type("\303\277")`: {
			in:     6,
			source: `type("\303\277")`,
			out:    types.StringType,
		},
		`type(int(6u))`: {
			in:     6,
			source: `type(int(6u))`,
			out:    types.IntType,
		},
		`type(uint(6))`: {
			in:     6,
			source: `type(uint(6))`,
			out:    types.UintType,
		},
		`type(int(6.0))`: {
			in:     6,
			source: `type(int(6.0))`,
			out:    types.IntType,
		},
		`type(uint(6.0))`: {
			in:     6,
			source: `type(uint(6.0))`,
			out:    types.UintType,
		},
		`type(double(6u))`: {
			in:     6,
			source: `type(double(6u))`,
			out:    types.DoubleType,
		},
		`type(double(6))`: {
			in:     6,
			source: `type(double(6))`,
			out:    types.DoubleType,
		},
		`type(string(6))`: {
			in:     6,
			source: `type(string(6))`,
			out:    types.StringType,
		},
		`type(int("6"))`: {
			in:     6,
			source: `type(int("6"))`,
			out:    types.IntType,
		},
		`type(string(6u))`: {
			in:     6,
			source: `type(string(6u))`,
			out:    types.StringType,
		},
		`type(uint("6"))`: {
			in:     6,
			source: `type(uint("6"))`,
			out:    types.UintType,
		},
		`type(string(6.0))`: {
			in:     6,
			source: `type(string(6.0))`,
			out:    types.StringType,
		},
		`type(double("6.0"))`: {
			in:     6,
			source: `type(double("6.0"))`,
			out:    types.DoubleType,
		},
		`type(string(true))`: {
			in:     6,
			source: `type(string(true))`,
			out:    types.StringType,
		},
		`type(bool("true"))`: {
			in:     6,
			source: `type(bool("true"))`,
			out:    types.BoolType,
		},
		`type(string(b"\x31\x32"))`: {
			in:     6,
			source: `type(string(b"\x31\x32"))`,
			out:    types.StringType,
		},
		`type(bytes("\x31\x32"))`: {
			in:     6,
			source: `type(bytes("\x31\x32"))`,
			out:    types.BytesType,
		},
	}

	env, err := cel.NewEnv(
		cel.Variable("buff", cel.IntType),
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

			out, _, err := prg.Eval(map[string]interface{}{
				"buff": v.in,
			})
			if err != nil {
				t.Errorf("Evaluation error: %v\n", err)
				t.FailNow()
			}

			t.Logf("%s -> %v", v.source, out)
			if out != v.out {
				t.Error(out, v.out)
				t.FailNow()
			}
		})
	}
}
