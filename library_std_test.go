package cel_test

import (
	"testing"
	"time"

	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/common/types"
	"github.com/google/cel-go/common/types/ref"
)

func TestStandardFunctions(tt *testing.T) {
	tests := map[string]struct {
		source string
		out    ref.Val
	}{
		"conditional": {
			source: "5 > 6 ? true : false",
			out:    types.Bool(false),
		},
		"logical_and": {
			source: "true && false",
			out:    types.Bool(false),
		},
		"logical_or": {
			source: "true || false",
			out:    types.Bool(true),
		},
		"logical_not": {
			source: "!true",
			out:    types.Bool(false),
		},
		`equals`: {
			source: `b"\x31" == b"\x32"`,
			out:    types.Bool(false),
		},
		`not_equals`: {
			source: `"abc" != "cba"`,
			out:    types.Bool(true),
		},
		`add_bytes`: {
			source: `b"\x31" + b"\x32"`,
			out:    types.Bytes{0x31, 0x32},
		},
		`add_double`: {
			source: `5.0 + 6.0`,
			out:    types.Double(11.0),
		},
		`add_duration_duration`: {
			source: `duration('1m') + duration('1s')`,
			out:    types.Duration{Duration: time.Duration(61 * time.Second)},
		},
		`add_duration_timestamp`: {
			source: `duration('24h') + timestamp('2023-01-01T00:00:00Z')`,
			out:    types.Timestamp{Time: time.Unix(1672617600, 0)},
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

			out, _, err := prg.Eval(map[string]any{})
			if err != nil {
				t.Errorf("Evaluation error: %v\n", err)
				t.FailNow()
			}

			t.Logf("%s -> %v", v.source, out)
			result := out.Equal(v.out)
			if !result.Value().(bool) {
				t.FailNow()
			}
		})
	}
}
