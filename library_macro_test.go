package cel_test

import (
	"testing"

	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/common/types"
	"github.com/google/cel-go/common/types/ref"
)

func TestMacro(tt *testing.T) {

	reg, err := types.NewRegistry()
	if err != nil {
		tt.FailNow()
	}
	tests := map[string]struct {
		source string
		out    ref.Val
	}{
		"exists": {
			source: `[1, 2, 3, 4, 5u, 1.0].exists(e, type(e) == uint)`,
			out:    types.Bool(true),
		},
		"exists_one": {
			source: `[1, 2, 3, 4u, 5u, 1.0].exists_one(e, type(e) == uint)`,
			out:    types.Bool(false),
		},
		"exists_one2": {
			source: `[1, 2, 3, 4, 5u, 1.0].exists_one(e, type(e) == uint)`,
			out:    types.Bool(true),
		},
		"all": {
			source: `["a", "b"].all(x, x in ["a", "b"])`,
			out:    types.Bool(true),
		},
		"filter": {
			source: `[-1, 0, 1, 2, 3].filter(x, x > 0)`,
			out:    types.NewDynamicList(reg, []int{1, 2, 3}),
		},
		"has": {
			source: `has({'a':1}.a)`,
			out:    types.Bool(true),
		},
		"map": {
			source: `[1, 2, 3].map(x, x * 2)`,
			out:    types.NewDynamicList(reg, []int{2, 4, 6}),
		},
		"map2": {
			source: `[-1, 0, 1, 2, 3].map(x, x > 0, x * 2)`,
			out:    types.NewDynamicList(reg, []int{2, 4, 6}),
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

			tt.Logf("%s -> %v", v.source, out.Value())
			if !out.Equal(v.out).(types.Bool) {
				t.FailNow()
			}
		})
	}
}
