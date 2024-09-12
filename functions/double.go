package functions

import (
	"encoding/binary"
	"math"

	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/common/types"
	"github.com/google/cel-go/common/types/ref"
	"github.com/iotxfoundry/cel-go/overloads"
)

var doubleFunctions = []cel.EnvOption{
	cel.Function(
		overloads.DoubleToBytes,
		// double.to_bytes(int) -> bytes
		cel.MemberOverload(
			overloads.DoubleToBytesInt,
			[]*cel.Type{cel.DoubleType, cel.IntType},
			cel.BytesType,
			cel.BinaryBinding(
				func(lhs, rhs ref.Val) ref.Val {
					in, ok := lhs.(types.Double)
					if !ok {
						return types.ValOrErr(lhs, "no such overload")
					}
					base, ok := rhs.(types.Int)
					if !ok {
						return types.ValOrErr(rhs, "no such overload")
					}

					if int(base)%8 != 0 {
						return types.NewErr("base '%d' out of float32, float64 size", base)
					}
					var ret []byte
					switch base {
					case 32:
						ret = make([]byte, 4)
						binary.BigEndian.PutUint32(ret, math.Float32bits(float32(in)))
					case 64:
						ret = make([]byte, 8)
						binary.BigEndian.PutUint64(ret, math.Float64bits(float64(in)))
					default:
						return types.NewErr("base '%d' out of float32, float64 size", base)
					}
					return types.Bytes(ret)
				},
			),
		),
	),
}
