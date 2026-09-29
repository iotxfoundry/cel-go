package cel_test

import (
	"strings"
	"testing"

	"github.com/google/cel-go/cel"
	compute "github.com/iotxfoundry/cel-go"
)

func TestBitwiseErrors(tt *testing.T) {
	env, err := cel.NewEnv(
		cel.Variable("buff", cel.BytesType),
		compute.ComputeLib(),
	)
	if err != nil {
		tt.Fatal(err)
	}

	tests := map[string]string{
		// bytes bitwise_index out of range
		"bitwise_index neg": `b"\x0f".bitwise_index(-1)`,
		"bitwise_index big": `b"\x0f".bitwise_index(8)`,
		// bytes.delete error
		"delete neg": `b"\x01".delete(-1)`,
		// bytes.swap error
		"swap neg a": `b"\x01\x02".swap(-1, 0)`,
		"swap big a": `b"\x01\x02".swap(2, 0)`,
		"swap neg b": `b"\x01\x02".swap(0, -1)`,
		"swap big b": `b"\x01\x02".swap(0, 2)`,
		// bytes.index error
		"index neg": `b"\x01".index(-1)`,
		"index big": `b"\x01".index(1)`,
		// int bitwise_index error
		"int_index neg": `5.bitwise_index(-1)`,
		"int_index big": `5.bitwise_index(64)`,
		// uint bitwise_index error
		"uint_index neg": `5u.bitwise_index(-1)`,
		"uint_index big": `5u.bitwise_index(64)`,
	}

	for k, src := range tests {
		tt.Run(k, func(t *testing.T) {
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
			_, _, err = prg.Eval(map[string]any{})
			if err == nil {
				t.Error("expected error, got none")
			}
		})
	}
}

func TestBytesConversionErrors(tt *testing.T) {
	env, err := cel.NewEnv(
		cel.Variable("buff", cel.BytesType),
		compute.ComputeLib(),
	)
	if err != nil {
		tt.Fatal(err)
	}

	tests := map[string]string{
		// invalid base values
		"tof bad base":  `b"\x01".tof(16)`,
		"toi bad base":  `b"\x01".toi(7)`,
		"toui bad base": `b"\x01".toui(7)`,
	}

	for k, src := range tests {
		tt.Run(k, func(t *testing.T) {
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
			_, _, err = prg.Eval(map[string]any{
				"buff": []byte{0x01},
			})
			if err == nil {
				t.Error("expected error, got none")
			}
		})
	}
}

func TestNumToBytesErrors(tt *testing.T) {
	env, err := cel.NewEnv(
		cel.Variable("buff", cel.IntType),
		compute.ComputeLib(),
	)
	if err != nil {
		tt.Fatal(err)
	}

	tests := map[string]string{
		"int_to_bytes bad":    `42.to_bytes(7)`,
		"uint_to_bytes bad":   `42u.to_bytes(7)`,
		"double_to_bytes bad": `3.14.to_bytes(16)`,
	}

	for k, src := range tests {
		tt.Run(k, func(t *testing.T) {
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
			_, _, err = prg.Eval(map[string]any{})
			if err == nil {
				t.Error("expected error, got none")
			}
		})
	}
}

func TestMathRandErrors(tt *testing.T) {
	env, err := cel.NewEnv(
		cel.Variable("buff", cel.IntType),
		compute.ComputeLib(),
	)
	if err != nil {
		tt.Fatal(err)
	}

	tests := map[string]string{
		"randf two args":  `math.randf(32, 64)`,
		"randi bad base":  `math.randi("bad")`,
		"randf bad base":  `math.randf("bad")`,
		"randui bad base": `math.randui("bad")`,
		"randi bad arg":   `math.randi("bad", 10)`,
	}

	for k, src := range tests {
		tt.Run(k, func(t *testing.T) {
			_, iss := env.Compile(src)
			if iss.Err() == nil {
				t.Error("expected compile error, got none")
				return
			}
		})
	}
}

func TestIntBitwiseTypeErrors(tt *testing.T) {
	env, err := cel.NewEnv(
		cel.Variable("buff", cel.IntType),
		compute.ComputeLib(),
	)
	if err != nil {
		tt.Fatal(err)
	}

	// int bitwise_and with uint arg should fail type check
	_, iss := env.Compile(`5.bitwise_and(3u)`)
	if iss.Err() == nil {
		tt.Error("expected type error for int.bitwise_and(uint), got none")
	}

	// uint bitwise_and with int arg should fail type check
	_, iss = env.Compile(`5u.bitwise_and(3)`)
	if iss.Err() == nil {
		tt.Error("expected type error for uint.bitwise_and(int), got none")
	}
}

func TestSliceErrors(tt *testing.T) {
	env, err := cel.NewEnv(
		cel.Variable("buff", cel.BytesType),
		compute.ComputeLib(),
	)
	if err != nil {
		tt.Fatal(err)
	}

	_, iss := env.Compile(`b"\x01\x02".slice(3, 4)`)
	if iss.Err() != nil && strings.Contains(iss.Err().Error(), "undec") {
		tt.Logf("compile note: %v", iss.Err())
	}

	ast, iss := env.Compile(`b"\x01\x02".slice(0, -1)`)
	if iss.Err() != nil {
		tt.Logf("compile note: %v", iss.Err())
	} else {
		prg, err := env.Program(ast)
		if err != nil {
			tt.Fatal(err)
		}
		_, _, err = prg.Eval(map[string]any{})
		if err == nil {
			tt.Error("expected error for negative end index")
		}
	}
}
