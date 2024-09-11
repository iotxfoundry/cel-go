package cel_test

import (
	"testing"

	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/common/types"
	"github.com/google/cel-go/common/types/ref"
)

func TestArithmetic(tt *testing.T) {

	tests := map[string]struct {
		source string
		out    ref.Val
	}{
		"6 + 7": {
			source: "6 + 7",
			out:    types.Int(13),
		},
		"6 - 7": {
			source: "6 - 7",
			out:    types.Int(-1),
		},
		"10 * 13": {
			source: "10 * 13",
			out:    types.Int(130),
		},
		"20 / 10": {
			source: "20 / 10",
			out:    types.Int(2),
		},
		`13 % 10`: {
			source: `13 % 10`,
			out:    types.Int(3),
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
			if out != v.out {
				t.FailNow()
			}
		})
	}
}
