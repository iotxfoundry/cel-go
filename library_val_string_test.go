package cel_test

import (
	"testing"
	"time"

	"cel.dev/cel-go/common/types"
	"cel.dev/cel-go/common/types/ref"
	compute "github.com/iotxfoundry/cel-go"
)

// TestVal2String exhaustively verifies string conversion for every supported
// ref.Val type, mirroring TestVal2Bytes and TestVal2Pb.
func TestVal2String(tt *testing.T) {
	reg, err := types.NewRegistry()
	if err != nil {
		tt.Fatal(err)
	}
	ts := time.Unix(0, 0).UTC()

	tests := map[string]struct {
		val    ref.Val
		result string
	}{
		"bytes":      {types.Bytes([]byte{0x01, 0x02, 0x03}), "AQID"},
		"bool_true":  {types.Bool(true), "true"},
		"bool_false": {types.Bool(false), "false"},
		"double":     {types.Double(3.14), "3.140"},
		"int":        {types.Int(-42), "-42"},
		"uint":       {types.Uint(42), "42"},
		"string":     {types.String("hello"), "hello"},
		"null":       {types.NullValue, ""},
		"duration":   {types.Duration{Duration: 3 * time.Minute}, "3m0s"},
		"timestamp":  {types.Timestamp{Time: ts}, ts.String()},
		"error":      {types.NoSuchOverloadErr(), "no such overload"},
		"list":       {types.NewDynamicList(reg, []float64{0.12}), "[0.12]"},
		"map":        {types.NewDynamicMap(reg, map[string]int64{"a": 1}), `{"a":1}`},
		"type_type":  {types.IntType, "int"},
	}

	for k, v := range tests {
		tt.Run(k, func(t *testing.T) {
			out, err := compute.Val2String(v.val)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if out != v.result {
				t.Fatalf("got %q, want %q", out, v.result)
			}
		})
	}
}

// TestVal2StringUnknown verifies unknown values convert to the empty string.
func TestVal2StringUnknown(tt *testing.T) {
	out, err := compute.Val2String(types.NewUnknown(1, nil))
	if err != nil {
		tt.Fatalf("unexpected error: %v", err)
	}
	if out != "" {
		tt.Fatalf("got %q, want empty string", out)
	}
}
