package functions

import (
	"encoding/binary"

	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/common/types"
	"github.com/google/cel-go/common/types/ref"
	"github.com/iotxfoundry/cel-go/overloads"
)

var uintFunctions = []cel.EnvOption{
	cel.Function(overloads.UintToBytes,
		// uint.to_bytes(int) -> bytes
		cel.MemberOverload(overloads.UintToBytesInt64,
			[]*cel.Type{cel.UintType, cel.IntType},
			cel.BytesType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				in, ok := lhs.(types.Uint)
				if !ok {
					return types.ValOrErr(lhs, "no such overload")
				}
				base, ok := rhs.(types.Int)
				if !ok {
					return types.ValOrErr(rhs, "no such overload")
				}
				buff := make([]byte, 8)
				binary.BigEndian.PutUint64(buff, uint64(in))
				if int(base)%8 != 0 {
					return types.NewErr("base '%d' out of uint8, uint16, uint32, uint64 size", base)
				}
				var ret []byte
				switch base {
				case 8:
					ret = []byte{buff[7]}
				case 16:
					ret = []byte{buff[6], buff[7]}
				case 32:
					ret = []byte(buff[4:8])
				case 64:
					ret = []byte(buff)
				default:
					return types.NewErr("base '%d' out of uint8, uint16, uint32, uint64 size", base)
				}
				return types.Bytes(ret)
			}),
		),
	),
}
