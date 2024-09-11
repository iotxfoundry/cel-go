package cel_test

import (
	"testing"
	"time"

	"github.com/google/cel-go/common/types"
	"github.com/google/cel-go/common/types/ref"
	compute "github.com/iotxfoundry/cel-go"
)

func Test_Any2String(tt *testing.T) {
	reg, err := types.NewRegistry()
	if err != nil {
		tt.FailNow()
	}

	tests := map[string]struct {
		val    ref.Val
		result string
	}{
		"bytes": {
			val:    types.Bytes([]byte{0x01, 0x02, 0x03, 0x04}),
			result: "AQIDBA==",
		},
		"bool": {
			val:    types.Bool(true),
			result: "true",
		},
		"double": {
			val:    types.Double(0.123456),
			result: "0.123",
		},
		"int": {
			val:    types.Int(-123456),
			result: "-123456",
		},
		"uint": {
			val:    types.Uint(123456),
			result: "123456",
		},
		"string": {
			val:    types.String("你好"),
			result: "你好",
		},
		"null": {
			val:    types.NullValue,
			result: "",
		},
		"duration": {
			val: types.Duration{
				Duration: 3 * time.Minute,
			},
			result: "3m0s",
		},
		"time": {
			val: types.Timestamp{
				Time: time.Time{},
			},
			result: "0001-01-01 00:00:00 +0000 UTC",
		},
		"error": {
			val:    types.NoSuchOverloadErr(),
			result: "no such overload",
		},
		"list": {
			val:    types.NewDynamicList(reg, []float64{0.12, 0.23, 0.34}),
			result: "[0.12,0.23,0.34]",
		},
		"map": {
			val:    types.NewDynamicMap(reg, map[string]int64{"10000": 1}),
			result: `{"10000":1}`,
		},
	}

	for k, v := range tests {
		tt.Run(k, func(t *testing.T) {
			out, err := compute.Val2String(v.val)
			if err != nil {
				t.Error(err)
				t.FailNow()
			}
			if out != v.result {
				t.Error(out)
				t.Error(v.result)
				t.FailNow()
			}
		})
	}
}
