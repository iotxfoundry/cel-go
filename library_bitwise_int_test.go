package cel_test

import (
	"testing"

	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/common/types"
	"github.com/google/cel-go/common/types/ref"
	compute "github.com/iotxfoundry/cel-go"
)

func TestIntBitwise(tt *testing.T) {
	tests := map[string]struct {
		source string
		out    ref.Val
	}{
		// bitwise_and
		"int.bitwise_and": {
			source: `5.bitwise_and(3)`,
			out:    types.Int(1),
		},
		"int.bitwise_and_neg": {
			source: `-1.bitwise_and(7)`,
			out:    types.Int(7),
		},
		// bitwise_or
		"int.bitwise_or": {
			source: `5.bitwise_or(3)`,
			out:    types.Int(7),
		},
		// bitwise_xor
		"int.bitwise_xor": {
			source: `5.bitwise_xor(3)`,
			out:    types.Int(6),
		},
		"int.bitwise_xor_self": {
			source: `-1.bitwise_xor(-1)`,
			out:    types.Int(0),
		},
		// bitwise_clear
		"int.bitwise_clear": {
			source: `7.bitwise_clear(1)`,
			out:    types.Int(6),
		},
		"int.bitwise_clear_none": {
			source: `-1.bitwise_clear(0)`,
			out:    types.Int(-1),
		},
		// bitwise_shr
		"int.bitwise_shr": {
			source: `8.bitwise_shr(2)`,
			out:    types.Int(2),
		},
		"int.bitwise_shr_neg": {
			source: `-8.bitwise_shr(2)`,
			out:    types.Int(-2),
		},
		"int.bitwise_shr_neg_shift": {
			source: `2.bitwise_shr(-2)`,
			out:    types.Int(8),
		},
		// bitwise_shl
		"int.bitwise_shl": {
			source: `2.bitwise_shl(2)`,
			out:    types.Int(8),
		},
		"int.bitwise_shl_neg_shift": {
			source: `8.bitwise_shl(-2)`,
			out:    types.Int(2),
		},
		// bitwise_not
		"int.bitwise_not": {
			source: `0.bitwise_not()`,
			out:    types.Int(-1),
		},
		"int.bitwise_not_neg": {
			source: `(-1).bitwise_not()`,
			out:    types.Int(0),
		},
		// bitwise_index
		"int.bitwise_index_1": {
			source: `10.bitwise_index(1) == b"\x01"`,
			out:    types.Bool(true),
		},
		"int.bitwise_index_0": {
			source: `10.bitwise_index(0) == b"\x00"`,
			out:    types.Bool(true),
		},
		"int.bitwise_index_63": {
			source: `(-9223372036854775808).bitwise_index(63) == b"\x01"`,
			out:    types.Bool(true),
		},
		// shift counts beyond the 64-bit width and MinInt64 magnitudes must
		// evaluate without panicking (Go full-width shift semantics)
		"int.bitwise_shl_huge": {
			source: `1.bitwise_shl(64)`,
			out:    types.Int(0),
		},
		"int.bitwise_shr_huge": {
			source: `1.bitwise_shr(64)`,
			out:    types.Int(0),
		},
		"int.bitwise_shr_minint_shift": {
			source: `1.bitwise_shr(-9223372036854775808)`,
			out:    types.Int(0),
		},
		"int.bitwise_shl_minint_shift": {
			source: `(-1).bitwise_shl(-9223372036854775808)`,
			out:    types.Int(-1),
		},
	}

	env, err := cel.NewEnv(
		cel.Variable("buff", cel.IntType),
		compute.ComputeLib(),
	)
	if err != nil {
		tt.Errorf("environment creation error: %v\n", err)
		tt.FailNow()
	}

	for k, v := range tests {
		tt.Run(k, func(t *testing.T) {
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
				t.Errorf("got %v, want %v", out, v.out)
				t.FailNow()
			}
		})
	}
}

func TestUintBitwise(tt *testing.T) {
	tests := map[string]struct {
		source string
		out    ref.Val
	}{
		// bitwise_and
		"uint.bitwise_and": {
			source: `5u.bitwise_and(3u)`,
			out:    types.Uint(1),
		},
		// bitwise_or
		"uint.bitwise_or": {
			source: `5u.bitwise_or(3u)`,
			out:    types.Uint(7),
		},
		// bitwise_xor
		"uint.bitwise_xor": {
			source: `5u.bitwise_xor(3u)`,
			out:    types.Uint(6),
		},
		"uint.bitwise_xor_self": {
			source: `255u.bitwise_xor(255u)`,
			out:    types.Uint(0),
		},
		// bitwise_clear
		"uint.bitwise_clear": {
			source: `7u.bitwise_clear(1u)`,
			out:    types.Uint(6),
		},
		// bitwise_shr
		"uint.bitwise_shr": {
			source: `8u.bitwise_shr(2)`,
			out:    types.Uint(2),
		},
		"uint.bitwise_shr_neg_shift": {
			source: `2u.bitwise_shr(-2)`,
			out:    types.Uint(8),
		},
		// bitwise_shl
		"uint.bitwise_shl": {
			source: `2u.bitwise_shl(2)`,
			out:    types.Uint(8),
		},
		"uint.bitwise_shl_neg_shift": {
			source: `8u.bitwise_shl(-2)`,
			out:    types.Uint(2),
		},
		// bitwise_not
		"uint.bitwise_not": {
			source: `0u.bitwise_not() == 18446744073709551615u`,
			out:    types.Bool(true),
		},
		"uint.bitwise_not_one": {
			source: `1u.bitwise_not() == 18446744073709551614u`,
			out:    types.Bool(true),
		},
		// bitwise_index
		"uint.bitwise_index_1": {
			source: `10u.bitwise_index(1) == b"\x01"`,
			out:    types.Bool(true),
		},
		"uint.bitwise_index_0": {
			source: `10u.bitwise_index(0) == b"\x00"`,
			out:    types.Bool(true),
		},
		"uint.bitwise_index_63": {
			source: `(18446744073709551615u).bitwise_index(63) == b"\x01"`,
			out:    types.Bool(true),
		},
		// shift counts beyond the 64-bit width and MinInt64 magnitudes must
		// evaluate without panicking (Go full-width shift semantics)
		"uint.bitwise_shl_huge": {
			source: `1u.bitwise_shl(64)`,
			out:    types.Uint(0),
		},
		"uint.bitwise_shr_huge": {
			source: `1u.bitwise_shr(64)`,
			out:    types.Uint(0),
		},
		"uint.bitwise_shr_minint_shift": {
			source: `1u.bitwise_shr(-9223372036854775808)`,
			out:    types.Uint(0),
		},
		"uint.bitwise_shl_minint_shift": {
			source: `18446744073709551615u.bitwise_shl(-9223372036854775808)`,
			out:    types.Uint(0),
		},
	}

	env, err := cel.NewEnv(
		cel.Variable("buff", cel.UintType),
		compute.ComputeLib(),
	)
	if err != nil {
		tt.Errorf("environment creation error: %v\n", err)
		tt.FailNow()
	}

	for k, v := range tests {
		tt.Run(k, func(t *testing.T) {
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
				t.Errorf("got %v, want %v", out, v.out)
				t.FailNow()
			}
		})
	}
}

func TestBytesPopcnt(tt *testing.T) {
	tests := map[string]struct {
		buff   []byte
		source string
		result int64
	}{
		`b"\xf0".popcnt==4`: {
			buff:   []byte{0xF0},
			source: `buff.bitwise_popcnt()`,
			result: 4,
		},
		`b"\xff\x00".popcnt==8`: {
			buff:   []byte{0xFF, 0x00},
			source: `buff.bitwise_popcnt()`,
			result: 8,
		},
		`b"\x00".popcnt==0`: {
			buff:   []byte{0x00},
			source: `buff.bitwise_popcnt()`,
			result: 0,
		},
	}

	env, err := cel.NewEnv(
		cel.Variable("buff", cel.BytesType),
		compute.ComputeLib(),
	)
	if err != nil {
		tt.Errorf("environment creation error: %v\n", err)
		tt.FailNow()
	}

	for k, v := range tests {
		tt.Run(k, func(t *testing.T) {
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

			out, _, err := prg.Eval(map[string]any{
				"buff": v.buff,
			})
			if err != nil {
				t.Errorf("Evaluation error: %v\n", err)
				t.FailNow()
			}

			t.Logf("%s -> %v", v.source, out)
			if out.Value().(int64) != v.result {
				t.Errorf("got %v, want %v", out.Value(), v.result)
				t.FailNow()
			}
		})
	}
}
