package cel_test

import (
	"encoding/binary"
	"math"
	"testing"
	"time"

	"github.com/google/cel-go/common/types"
	"github.com/google/cel-go/common/types/ref"
	compute "github.com/iotxfoundry/cel-go"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func int64Bytes(v uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, v)
	return b
}

func TestVal2Bytes(tt *testing.T) {
	reg, err := types.NewRegistry()
	if err != nil {
		tt.Fatal(err)
	}

	negOne := types.Int(-1).Value().(int64)
	ts := time.Unix(0, 0).UTC()
	tests := map[string]struct {
		val    ref.Val
		result []byte
	}{
		"bytes":      {types.Bytes([]byte{0x01, 0x02, 0x03}), []byte{0x01, 0x02, 0x03}},
		"bool_true":  {types.Bool(true), []byte{1}},
		"bool_false": {types.Bool(false), []byte{0}},
		"double":     {types.Double(3.14), int64Bytes(math.Float64bits(3.14))},
		"int_pos":    {types.Int(42), int64Bytes(42)},
		"int_neg":    {types.Int(-1), int64Bytes(uint64(negOne))},
		"uint":       {types.Uint(42), int64Bytes(42)},
		"string":     {types.String("hello"), []byte("hello")},
		"null":       {types.NullValue, []byte{}},
		"duration":   {types.Duration{Duration: 3 * time.Minute}, []byte("3m0s")},
		"timestamp":  {types.Timestamp{Time: ts}, []byte(ts.String())},
		"error":      {types.NoSuchOverloadErr(), []byte("no such overload")},
		"list":       {types.NewDynamicList(reg, []float64{0.12}), []byte("[0.12]")},
		"map":        {types.NewDynamicMap(reg, map[string]int64{"a": 1}), []byte(`{"a":1}`)},
		"type_type":  {types.IntType, []byte("int")},
	}

	for k, v := range tests {
		tt.Run(k, func(t *testing.T) {
			out, err := compute.Val2Bytes(v.val)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if string(out) != string(v.result) {
				t.Fatalf("got %v, want %v", out, v.result)
			}
		})
	}
}

func TestVal2Pb(tt *testing.T) {
	reg, err := types.NewRegistry()
	if err != nil {
		tt.Fatal(err)
	}
	ts := time.Unix(0, 0).UTC()

	tests := []struct {
		name   string
		val    ref.Val
		expect func(t *testing.T, v interface{})
	}{
		{"bytes", types.Bytes([]byte{0x01, 0x02, 0x03}), func(t *testing.T, v interface{}) {
			s, ok := v.(string)
			// structpb encodes bytes as base64
			if !ok || s != "AQID" {
				t.Errorf("got %v (%T)", v, v)
			}
		}},
		{"bool_true", types.Bool(true), func(t *testing.T, v interface{}) {
			b, ok := v.(bool)
			if !ok || !b {
				t.Errorf("got %v (%T)", v, v)
			}
		}},
		{"bool_false", types.Bool(false), func(t *testing.T, v interface{}) {
			b, ok := v.(bool)
			if !ok || b {
				t.Errorf("got %v (%T)", v, v)
			}
		}},
		{"double", types.Double(3.14), func(t *testing.T, v interface{}) {
			f, ok := v.(float64)
			if !ok || f != 3.14 {
				t.Errorf("got %v (%T)", v, v)
			}
		}},
		{"int_neg", types.Int(-42), func(t *testing.T, v interface{}) {
			f, ok := v.(float64)
			if !ok || f != -42 {
				t.Errorf("got %v (%T)", v, v)
			}
		}},
		{"uint", types.Uint(42), func(t *testing.T, v interface{}) {
			f, ok := v.(float64)
			if !ok || f != 42 {
				t.Errorf("got %v (%T)", v, v)
			}
		}},
		{"string", types.String("hello"), func(t *testing.T, v interface{}) {
			s, ok := v.(string)
			if !ok || s != "hello" {
				t.Errorf("got %v (%T)", v, v)
			}
		}},
		{"null", types.NullValue, func(t *testing.T, v interface{}) {
			if v != nil {
				t.Errorf("got %v, want nil", v)
			}
		}},
		{"duration", types.Duration{Duration: 3 * time.Minute}, func(t *testing.T, v interface{}) {
			s, ok := v.(string)
			if !ok || s != "3m0s" {
				t.Errorf("got %v (%T)", v, v)
			}
		}},
		{"timestamp", types.Timestamp{Time: ts}, func(t *testing.T, v interface{}) {
			s, ok := v.(string)
			if !ok || s != ts.String() {
				t.Errorf("got %q, want %q", s, ts.String())
			}
		}},
		{"error", types.NoSuchOverloadErr(), func(t *testing.T, v interface{}) {
			s, ok := v.(string)
			if !ok || s != "no such overload" {
				t.Errorf("got %v (%T)", v, v)
			}
		}},
		{"type_type", types.BoolType, func(t *testing.T, v interface{}) {
			s, ok := v.(string)
			if !ok || s != "bool" {
				t.Errorf("got %v (%T)", v, v)
			}
		}},
		{"list", types.NewDynamicList(reg, []float64{0.12}), func(t *testing.T, v interface{}) {
			l, ok := v.([]interface{})
			if !ok || len(l) != 1 {
				t.Errorf("got %v (%T)", v, v)
			}
		}},
		{"map", types.NewDynamicMap(reg, map[string]int64{"a": 1}), func(t *testing.T, v interface{}) {
			m, ok := v.(map[string]interface{})
			if !ok || m["a"].(float64) != 1 {
				t.Errorf("got %v (%T)", v, v)
			}
		}},
	}

	for _, tc := range tests {
		tt.Run(tc.name, func(t *testing.T) {
			out, err := compute.Val2Pb(tc.val)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			tc.expect(t, out.AsInterface())
		})
	}
}

func TestVal2PbProtoAny(tt *testing.T) {
	reg, err := types.NewRegistry()
	if err != nil {
		tt.Fatal(err)
	}
	a, err := anypb.New(wrapperspb.Int64(5))
	if err != nil {
		tt.Fatal(err)
	}
	val := reg.NativeToValue(a)
	out, err := compute.Val2Pb(val)
	if err != nil {
		tt.Fatal(err)
	}
	if out.AsInterface() == nil {
		tt.Error("got nil for proto Any")
	}
}

func TestVal2BytesProtoAny(tt *testing.T) {
	reg, err := types.NewRegistry()
	if err != nil {
		tt.Fatal(err)
	}
	a, err := anypb.New(wrapperspb.Int64(5))
	if err != nil {
		tt.Fatal(err)
	}
	val := reg.NativeToValue(a)
	out, err := compute.Val2Bytes(val)
	if err != nil {
		tt.Fatal(err)
	}
	if len(out) == 0 {
		tt.Error("got empty bytes for Any")
	}
}

func TestVal2StringMissingTypes(tt *testing.T) {
	tests := map[string]struct {
		val    ref.Val
		result string
	}{
		"type_type": {types.BoolType, "bool"},
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

func TestVal2StringProtoAny(tt *testing.T) {
	reg, err := types.NewRegistry()
	if err != nil {
		tt.Fatal(err)
	}
	a, err := anypb.New(wrapperspb.Int64(5))
	if err != nil {
		tt.Fatal(err)
	}
	val := reg.NativeToValue(a)
	out, err := compute.Val2String(val)
	if err != nil {
		tt.Fatal(err)
	}
	if out == "" {
		tt.Error("got empty string for Any")
	}
}
