package functions

import (
	"encoding/binary"

	"cel.dev/cel-go/cel"
	"cel.dev/cel-go/common/types"
	"cel.dev/cel-go/common/types/ref"
	"github.com/iotxfoundry/cel-go/overloads"
)

var intFunctions = []cel.EnvOption{
	cel.Function(
		overloads.IntToBytes,
		cel.FunctionDocs(
			"Converts an integer value to a byte sequence. "+
				"The size of the byte sequence depends on the specified base. "+
				"Valid bases are 8, 16, 32, and 64, corresponding to int8, int16, int32, and int64 types, respectively. "+
				"If the base is not 8, 16, 32, or 64, an error is returned.",
		),
		// int.to_bytes(int) -> bytes
		cel.MemberOverload(
			overloads.IntToBytesInt,
			[]*cel.Type{cel.IntType, cel.IntType},
			cel.BytesType,
			cel.OverloadExamples(
				`42.to_bytes(8) // b"\x2a"`,
				`42.to_bytes(16) // b"\x00\x2a"`,
				`42.to_bytes(32) // b"\x00\x00\x00\x2a"`,
				`42.to_bytes(64) // b"\x00\x00\x00\x00\x00\x00\x00\x2a"`,
				`42.to_bytes(128) // error: base '128' out of int8, int16, int32, int64 size`,
			),
			cel.BinaryBinding(
				func(lhs, rhs ref.Val) ref.Val {
					in, ok := lhs.(types.Int)
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
						return types.NewErr("base '%d' out of int8, int16, int32, int64 size", base)
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
						return types.NewErr("base '%d' out of int8, int16, int32, int64 size", base)
					}
					return types.Bytes(ret)
				},
			),
		),
	),
}
