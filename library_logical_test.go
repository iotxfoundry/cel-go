package cel_test

import (
	"testing"

	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/common/types"
)

func TestLogical(tt *testing.T) {

	tests := map[string]struct {
		source string
		out    types.Bool
	}{
		"6 == 6": {
			source: "6 == 6",
			out:    true,
		},
		"6u != 7u": {
			source: "6u != 7u",
			out:    true,
		},
		"6.0 == 6.0": {
			source: "6.0 == 6.0",
			out:    true,
		},
		"6.0 != 6.0": {
			source: "6.0 != 6.0",
			out:    false,
		},
		`"abc" != "cba"`: {
			source: `"abc" != "cba"`,
			out:    true,
		},
		`b"\x31" == b"\x32"`: {
			source: `b"\x31" == b"\x32"`,
			out:    false,
		},
		`5 > 6 ? true : false`: {
			source: `5 > 6 ? true : false`,
			out:    false,
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

			out, _, err := prg.Eval(map[string]interface{}{})
			if err != nil {
				t.Errorf("Evaluation error: %v\n", err)
				t.FailNow()
			}

			t.Logf("%s -> %v", v.source, out)
			if out.(types.Bool) != v.out {
				t.FailNow()
			}
		})
	}
}
