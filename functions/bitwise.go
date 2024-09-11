package functions

import (
	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/common/types"
	"github.com/google/cel-go/common/types/ref"
	"github.com/iotxfoundry/cel-go/overloads"
)

var bitwiseFunctions = []cel.EnvOption{
	cel.Function(overloads.BitwiseShiftRight,
		// bytes.bitwise_shr(int) -> bytes
		cel.MemberOverload(overloads.BitwiseShiftRightInt64,
			[]*cel.Type{cel.BytesType, cel.IntType},
			cel.BytesType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				data, ok := lhs.(types.Bytes)
				if !ok {
					return types.ValOrErr(lhs, "no such overload")
				}
				bits, ok := rhs.(types.Int)
				if !ok {
					return types.ValOrErr(rhs, "no such overload")
				}
				n := len(data)
				if bits < 0 {
					bits = -bits
					for i := 0; i < n-1; i++ {
						data[i] = data[i]<<bits | data[i+1]>>(8-bits)
					}
					data[n-1] <<= bits
				} else {
					for i := n - 1; i > 0; i-- {
						data[i] = data[i]>>bits | data[i-1]<<(8-bits)
					}
					data[0] >>= bits
				}
				return data
			}),
		),
	),
	cel.Function(overloads.BitwiseShiftLeft,
		// bytes.bitwise_shl(int) -> bytes
		cel.MemberOverload(overloads.BitwiseShiftLeftInt64,
			[]*cel.Type{cel.BytesType, cel.IntType},
			cel.BytesType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				data, ok := lhs.(types.Bytes)
				if !ok {
					return types.ValOrErr(lhs, "no such overload")
				}
				bits, ok := rhs.(types.Int)
				if !ok {
					return types.ValOrErr(rhs, "no such overload")
				}
				n := len(data)
				if bits < 0 {
					bits = -bits
					for i := n - 1; i > 0; i-- {
						data[i] = data[i]>>bits | data[i-1]<<(8-bits)
					}
					data[0] >>= bits
				} else {
					for i := 0; i < n-1; i++ {
						data[i] = data[i]<<bits | data[i+1]>>(8-bits)
					}
					data[n-1] <<= bits
				}
				return data
			}),
		),
	),
	cel.Function(overloads.BitwiseAnd,
		// bytes.bitwise_and(bytes) -> bytes
		cel.MemberOverload(overloads.BitwiseAndBytes,
			[]*cel.Type{cel.BytesType, cel.BytesType},
			cel.BytesType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				buff, ok := lhs.(types.Bytes)
				if !ok {
					return types.ValOrErr(lhs, "no such overload")
				}
				temp, ok := rhs.(types.Bytes)
				if !ok {
					return types.ValOrErr(rhs, "no such overload")
				}
				for k := range buff {
					if k >= len(temp) {
						break
					}
					buff[k] &= temp[k]
				}
				return buff
			}),
		),
	),

	cel.Function(overloads.BitwiseOr,
		// bytes.bitwise_or(bytes) -> bytes
		cel.MemberOverload(overloads.BitwiseOrBytes,
			[]*cel.Type{cel.BytesType, cel.BytesType},
			cel.BytesType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				buff, ok := lhs.(types.Bytes)
				if !ok {
					return types.ValOrErr(lhs, "no such overload")
				}
				temp, ok := rhs.(types.Bytes)
				if !ok {
					return types.ValOrErr(rhs, "no such overload")
				}
				for k := range buff {
					if k >= len(temp) {
						break
					}
					buff[k] |= temp[k]
				}
				return buff
			}),
		),
	),
	cel.Function(overloads.BitwiseXor,
		// bytes.bitwise_xor(bytes) -> bytes
		cel.MemberOverload(overloads.BitwiseXorBytes,
			[]*cel.Type{cel.BytesType, cel.BytesType},
			cel.BytesType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				buff, ok := lhs.(types.Bytes)
				if !ok {
					return types.ValOrErr(lhs, "no such overload")
				}
				temp, ok := rhs.(types.Bytes)
				if !ok {
					return types.ValOrErr(rhs, "no such overload")
				}
				for k := range buff {
					if k >= len(temp) {
						break
					}
					buff[k] ^= temp[k]
				}
				return buff
			}),
		),
	),
	cel.Function(overloads.BitwiseClear,
		// bytes.bitwise_clear(bytes) -> bytes
		cel.MemberOverload(overloads.BitwiseClearBytes,
			[]*cel.Type{cel.BytesType, cel.BytesType},
			cel.BytesType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				buff, ok := lhs.(types.Bytes)
				if !ok {
					return types.ValOrErr(lhs, "no such overload")
				}
				temp, ok := rhs.(types.Bytes)
				if !ok {
					return types.ValOrErr(rhs, "no such overload")
				}
				for k := range buff {
					if k >= len(temp) {
						break
					}
					buff[k] &^= temp[k]
				}
				return buff
			}),
		),
	),

	cel.Function(overloads.BitwiseIndex,
		// bytes.bitwise_index(int) -> bytes
		cel.MemberOverload(overloads.BitwiseIndexInt64,
			[]*cel.Type{cel.BytesType, cel.IntType},
			cel.BytesType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				buff, ok := lhs.(types.Bytes)
				if !ok {
					return types.ValOrErr(lhs, "no such overload")
				}
				index, ok := rhs.(types.Int)
				if !ok {
					return types.ValOrErr(rhs, "no such overload")
				}
				if int(index) >= len(buff)*8 || int(index) < 0 {
					return types.NewErr("index '%d' out of range in bitwise size '%d'", index, len(buff)*8)
				}
				remainder := int(index) % 8
				ret := []byte{}
				for k, v := range buff {
					if k*8 == int(index)-remainder {
						for i := 0; i < remainder; i++ {
							v = v >> 1
						}
						v &= 0x01
						ret = []byte{v}
						break
					}
				}
				return types.Bytes(ret)
			}),
		),
	),
}
