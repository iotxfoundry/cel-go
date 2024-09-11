package cel_test

import (
	"bytes"
	"testing"

	"github.com/google/cel-go/cel"
	compute "github.com/iotxfoundry/cel-go"
)

func TestBytes2Double(tt *testing.T) {
	tests := map[string]struct {
		buff   []byte
		source string
		result float64
	}{
		"buff.tof(32)": {
			buff:   []byte{0x43, 0xAC, 0x8F, 0x5C},
			source: "buff.tof(32)",
			result: 345.12,
		},
		"buff.tof(64)": {
			buff:   []byte{0x40, 0x75, 0x91, 0xEB, 0x85, 0x1E, 0xB8, 0x52},
			source: "buff.tof(64)",
			result: 345.12,
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
				"buff": v.buff,
			})
			if err != nil {
				t.Errorf("Evaluation error: %v\n", err)
				t.FailNow()
			}

			t.Log(out, v.result)
		})
	}
}

func TestDouble2Bytes(tt *testing.T) {
	tests := map[string]struct {
		buff   float64
		source string
		result []byte
	}{
		"buff.to_bytes(32)": {
			buff:   345.12,
			source: "buff.to_bytes(32)",
			result: []byte{0x43, 0xAC, 0x8F, 0x5C},
		},
		"buff.to_bytes(64)": {
			buff:   345.12,
			source: "buff.to_bytes(64)",
			result: []byte{0x40, 0x75, 0x91, 0xEB, 0x85, 0x1E, 0xB8, 0x52},
		},
	}

	env, err := cel.NewEnv(
		cel.Variable("buff", cel.DoubleType),
		compute.ComputeLib(),
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
				"buff": v.buff,
			})
			if err != nil {
				t.Errorf("Evaluation error: %v\n", err)
				t.FailNow()
			}

			t.Log(out, v.result)
		})
	}
}

func TestBytes2Uint(tt *testing.T) {
	tests := map[string]struct {
		buff   []byte
		source string
		result uint64
	}{
		"buff.toui(8)": {
			buff:   []byte{0xFF},
			source: "buff.toui(8)",
			result: 0xFF,
		},
		"buff.toui(16)": {
			buff:   []byte{0xFF, 0xFF},
			source: "buff.toui(16)",
			result: 0xFFFF,
		},
		"buff.toui(32)": {
			buff:   []byte{0xFF, 0xFF, 0xFF, 0xFF},
			source: "buff.toui(32)",
			result: 0xFFFFFFFF,
		},
		"buff.toui(64)": {
			buff:   []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF},
			source: "buff.toui(64)",
			result: 0xFFFFFFFFFFFFFFFF,
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
				"buff": v.buff,
			})
			if err != nil {
				t.Errorf("Evaluation error: %v\n", err)
				t.FailNow()
			}

			if out.Value().(uint64) != v.result {
				t.Logf("%d", out.Value())
				t.Logf("%d", v.result)
				t.FailNow()
			}
		})
	}
}

func TestUint2Bytes(tt *testing.T) {
	tests := map[string]struct {
		buff   uint64
		source string
		result []byte
	}{
		"buff.to_bytes(8)": {
			buff:   0xFF,
			source: "buff.to_bytes(8)",
			result: []byte{0xFF},
		},
		"buff.to_bytes(16)": {
			buff:   0xFFFF,
			source: "buff.to_bytes(16)",
			result: []byte{0xFF, 0xFF},
		},
		"buff.to_bytes(32)": {
			buff:   0xFFFFFFFF,
			source: "buff.to_bytes(32)",
			result: []byte{0xFF, 0xFF, 0xFF, 0xFF},
		},
		"buff.to_bytes(64)": {
			buff:   0xFFFFFFFFFFFFFFFF,
			source: "buff.to_bytes(64)",
			result: []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF},
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
				"buff": v.buff,
			})
			if err != nil {
				t.Errorf("Evaluation error: %v\n", err)
				t.FailNow()
			}

			if !bytes.Equal(out.Value().([]byte), v.result) {
				t.Logf("%d", out.Value())
				t.Logf("%d", v.result)
				t.FailNow()
			}
		})
	}
}

func TestBytes2Int(tt *testing.T) {
	tests := map[string]struct {
		buff   []byte
		source string
		result int64
	}{
		"buff.toi(8)": {
			buff:   []byte{0xFF},
			source: "buff.toi(8)",
			result: -1,
		},
		"buff.toi(16)": {
			buff:   []byte{0xFF, 0xFF},
			source: "buff.toi(16)",
			result: -1,
		},
		"buff.toi(32)": {
			buff:   []byte{0xFF, 0xFF, 0xFF, 0xFF},
			source: "buff.toi(32)",
			result: -1,
		},
		"buff.toi(64)": {
			buff:   []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF},
			source: "buff.toi(64)",
			result: -1,
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
				"buff": v.buff,
			})
			if err != nil {
				t.Errorf("Evaluation error: %v\n", err)
				t.FailNow()
			}

			if out.Value().(int64) != v.result {
				t.Logf("%d", out.Value())
				t.Logf("%d", v.result)
				t.FailNow()
			}
		})
	}
}

func TestInt2Bytes(tt *testing.T) {
	tests := map[string]struct {
		buff   int64
		source string
		result []byte
	}{
		"buff.to_bytes(8)": {
			buff:   -1,
			source: "buff.to_bytes(8)",
			result: []byte{0xFF},
		},
		"buff.to_bytes(16)": {
			buff:   -1,
			source: "buff.to_bytes(16)",
			result: []byte{0xFF, 0xFF},
		},
		"buff.to_bytes(32)": {
			buff:   -1,
			source: "buff.to_bytes(32)",
			result: []byte{0xFF, 0xFF, 0xFF, 0xFF},
		},
		"buff.to_bytes(64)": {
			buff:   -1,
			source: "buff.to_bytes(64)",
			result: []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF},
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
				"buff": v.buff,
			})
			if err != nil {
				t.Errorf("Evaluation error: %v\n", err)
				t.FailNow()
			}

			if !bytes.Equal(out.Value().([]byte), v.result) {
				t.Logf("%d", out.Value())
				t.Logf("%d", v.result)
				t.FailNow()
			}
		})
	}
}

func TestBytes(tt *testing.T) {
	tests := map[string]struct {
		buff   []byte
		source string
		result []byte
	}{
		"bytes.slice(0,3)": {
			buff:   []byte{0x01, 0x02, 0x03, 0x04, 0x05},
			source: "buff.slice(0,3)",
			result: []byte{0x01, 0x02, 0x03},
		},
		"bytes.index(0)": {
			buff:   []byte{0x01, 0x02, 0x03, 0x04, 0x05},
			source: "buff.index(0)",
			result: []byte{0x01},
		},
		"bytes.delete(0)": {
			buff:   []byte{0x01, 0x02, 0x03, 0x04, 0x05},
			source: "buff.delete(0)",
			result: []byte{0x02, 0x03, 0x04, 0x05},
		},
		"bytes.index(1)": {
			buff:   []byte{0x01, 0x02, 0x03, 0x04, 0x05},
			source: "buff.index(1)",
			result: []byte{0x02},
		},
		"bytes.delete(1)": {
			buff:   []byte{0x01, 0x02, 0x03, 0x04, 0x05},
			source: "buff.delete(1)",
			result: []byte{0x01, 0x03, 0x04, 0x05},
		},
		"bytes.delete(4)": {
			buff:   []byte{0x01, 0x02, 0x03, 0x04, 0x05},
			source: "buff.delete(4)",
			result: []byte{0x01, 0x02, 0x03, 0x04},
		},
		"bytes.delete(10)": {
			buff:   []byte{0x01, 0x02, 0x03, 0x04, 0x05},
			source: "buff.delete(10)",
			result: []byte{0x01, 0x02, 0x03, 0x04, 0x05},
		},
		"bytes.delete(1,2)": {
			buff:   []byte{0x01, 0x02, 0x03, 0x04, 0x05},
			source: "buff.delete(1,2)",
			result: []byte{0x01, 0x03, 0x04, 0x05},
		},
		"bytes.delete(1,10)": {
			buff:   []byte{0x01, 0x02, 0x03, 0x04, 0x05},
			source: "buff.delete(1,10)",
			result: []byte{0x01},
		},
		"bytes.index(2)": {
			buff:   []byte{0x01, 0x02, 0x03, 004, 0x05},
			source: "buff.index(2)",
			result: []byte{0x03},
		},
		"bytes.swap(2,3)": {
			buff:   []byte{0x01, 0x02, 0x03, 0x04, 0x05},
			source: "buff.swap(2,3)",
			result: []byte{0x01, 0x02, 0x04, 0x03, 0x05},
		},
		"bytes.index(3)": {
			buff:   []byte{0x01, 0x02, 0x03, 0x04, 0x05},
			source: "buff.index(3)",
			result: []byte{0x04},
		},
		"bytes.index(4)": {
			buff:   []byte{0x01, 0x02, 0x03, 0x04, 0x05},
			source: "buff.index(4)",
			result: []byte{0x05},
		},
		`buff.bitwise_shr(4)`: {
			buff:   []byte{0xF0},
			source: `buff.bitwise_shr(4)`,
			result: []byte{0x0F},
		},
		`buff.bitwise_shl(4)`: {
			buff:   []byte{0xF0},
			source: `buff.bitwise_shl(4)`,
			result: []byte{0x00},
		},
		`buff.bitwise_and(b"\x01")`: {
			buff:   []byte{0xFF},
			source: `buff.bitwise_and(b"\x01")`,
			result: []byte{0x01},
		},
		`buff.bitwise_or(b"\xFF")`: {
			buff:   []byte{0x01},
			source: `buff.bitwise_or(b"\xFF")`,
			result: []byte{0xFF},
		},
		`buff.bitwise_xor(b"\x0F")`: {
			buff:   []byte{0xF0},
			source: `buff.bitwise_xor(b"\x0F")`,
			result: []byte{0xFF},
		},
		`buff.bitwise_clear(b"\x0F")`: {
			buff:   []byte{0xFF},
			source: `buff.bitwise_clear(b"\x0F")`,
			result: []byte{0xF0},
		},
		`buff.bitwise_index(3)`: {
			buff:   []byte{0b0000_1000},
			source: `buff.bitwise_index(3)`,
			result: []byte{0x01},
		},
		`buff.bitwise_index(2)`: {
			buff:   []byte{0b0000_1000},
			source: `buff.bitwise_index(2)`,
			result: []byte{0x00},
		},
		`buff.bitwise_index(1)`: {
			buff:   []byte{0b0000_1000},
			source: `buff.bitwise_index(1)`,
			result: []byte{0x00},
		},
		`buff.bitwise_index(0)`: {
			buff:   []byte{0b0000_1000},
			source: `buff.bitwise_index(0)`,
			result: []byte{0x00},
		},
		`buff.bitwise_index(13)`: {
			buff:   []byte{0b0000_1000, 0b0010_0000},
			source: `buff.bitwise_index(13)`,
			result: []byte{0x01},
		},
		"compute1": {
			buff:   []byte{0x01, 0x02, 0x03, 0x04, 0x05},
			source: `buff.index(0).bitwise_and(b"\x01").bitwise_or(b"\xFF").bitwise_shl(7).bitwise_shr(7)`,
			result: []byte{0x01},
		},
		"compute2": {
			buff:   []byte{0x01, 0x02, 0x03, 0x04, 0x05},
			source: `buff`,
			result: []byte{0x01, 0x02, 0x03, 0x04, 0x05},
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
				"buff": v.buff,
			})
			if err != nil {
				t.Errorf("Evaluation error: %v\n", err)
				t.FailNow()
			}

			if !bytes.Equal(out.Value().([]byte), v.result) {
				t.Logf("% 02X", out.Value())
				t.Logf("% 02X", v.result)
				t.FailNow()
			}
		})
	}
}
