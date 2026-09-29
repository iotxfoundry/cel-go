package functions

import (
	"encoding/binary"

	"cel.dev/cel-go/cel"
	"cel.dev/cel-go/common/types"
	"cel.dev/cel-go/common/types/ref"
	"github.com/iotxfoundry/cel-go/overloads"
)

var uintFunctions = []cel.EnvOption{
	cel.Function(
		overloads.UintToBytes,
		cel.FunctionDocs(
			"Converts a uint to a bytes representation.",
		),
		// uint.to_bytes(int) -> bytes
		cel.MemberOverload(
			overloads.UintToBytesInt,
			[]*cel.Type{cel.UintType, cel.IntType},
			cel.BytesType,
			cel.OverloadExamples(
				`42.to_bytes(8) // b"\x2a"`,
				`42.to_bytes(16) // b"\x00\x2a"`,
				`42.to_bytes(32) // b"\x00\x00\x00\x2a"`,
				`42.to_bytes(64) // b"\x00\x00\x00\x00\x00\x00\x00\x2a"`,
				`42.to_bytes(128) // error: base '128' out of uint8, uint16, uint32, uint64 size`,
			),
			cel.BinaryBinding(
				func(lhs, rhs ref.Val) ref.Val {
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
				},
			),
		),
	),
}
