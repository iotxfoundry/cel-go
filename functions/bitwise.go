package functions

import (
	"math/bits"

	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/common/types"
	"github.com/google/cel-go/common/types/ref"
	"github.com/iotxfoundry/cel-go/overloads"
)

var bitwiseFunctions = []cel.EnvOption{
	cel.Function(
		overloads.BitwiseShiftRight,
		cel.FunctionDocs(
			"Performs a bitwise right shift on the bytes, shifting the bits to the right by the specified number of positions. "+
				"Negative values shift to the left, and positive values shift to the right. "+
				"Bits that are shifted out of the byte are discarded, and new bits are filled with zeros.",
		),
		// bytes.bitwise_shr(int) -> bytes
		cel.MemberOverload(
			overloads.BitwiseShiftRightInt64,
			[]*cel.Type{cel.BytesType, cel.IntType},
			cel.BytesType,
			cel.OverloadExamples(
				`b"\xf0".bitwise_shr(4) // b"\x0f"`,
			),
			cel.BinaryBinding(
				func(lhs, rhs ref.Val) ref.Val {
					src, ok := lhs.(types.Bytes)
					if !ok {
						return types.ValOrErr(lhs, "no such overload")
					}
					bits, ok := rhs.(types.Int)
					if !ok {
						return types.ValOrErr(rhs, "no such overload")
					}
					data := make([]byte, len(src))
					copy(data, src)
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
					return types.Bytes(data)
				},
			),
		),
	),

	cel.Function(
		overloads.BitwiseShiftLeft,
		cel.FunctionDocs(
			"Performs a bitwise left shift on the bytes, shifting the bits to the left by the specified number of positions. "+
				"Negative values shift to the right, and positive values shift to the left. "+
				"Bits that are shifted out of the byte are discarded, and new bits are filled with zeros.",
		),
		// bytes.bitwise_shl(int) -> bytes
		cel.MemberOverload(
			overloads.BitwiseShiftLeftInt64,
			[]*cel.Type{cel.BytesType, cel.IntType},
			cel.BytesType,
			cel.OverloadExamples(
				`b"\xf0".bitwise_shl(4) // b"\x00"`,
			),
			cel.BinaryBinding(
				func(lhs, rhs ref.Val) ref.Val {
					src, ok := lhs.(types.Bytes)
					if !ok {
						return types.ValOrErr(lhs, "no such overload")
					}
					bits, ok := rhs.(types.Int)
					if !ok {
						return types.ValOrErr(rhs, "no such overload")
					}
					data := make([]byte, len(src))
					copy(data, src)
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
					return types.Bytes(data)
				},
			),
		),
	),

	cel.Function(
		overloads.BitwiseAnd,
		cel.FunctionDocs(
			"Performs a bitwise AND operation on the bytes, combining the bits of two byte sequences. "+
				"Each bit in the result is set to 1 if both corresponding bits in the input bytes are 1, otherwise it is set to 0.",
		),
		// bytes.bitwise_and(bytes) -> bytes
		cel.MemberOverload(
			overloads.BitwiseAndBytes,
			[]*cel.Type{cel.BytesType, cel.BytesType},
			cel.BytesType,
			cel.OverloadExamples(
				`b"\xff".bitwise_and(b"\x0f") // b"\x0f"`,
			),
			cel.BinaryBinding(
				func(lhs, rhs ref.Val) ref.Val {
					src, ok := lhs.(types.Bytes)
					if !ok {
						return types.ValOrErr(lhs, "no such overload")
					}
					temp, ok := rhs.(types.Bytes)
					if !ok {
						return types.ValOrErr(rhs, "no such overload")
					}
					buff := make([]byte, len(src))
					copy(buff, src)
					for k := range buff {
						if k >= len(temp) {
							break
						}
						buff[k] &= temp[k]
					}
					return types.Bytes(buff)
				},
			),
		),
	),

	cel.Function(
		overloads.BitwiseOr,
		cel.FunctionDocs(
			"Performs a bitwise OR operation on the bytes, combining the bits of two byte sequences. "+
				"Each bit in the result is set to 1 if at least one of the corresponding bits in the input bytes is 1, otherwise it is set to 0.",
		),
		// bytes.bitwise_or(bytes) -> bytes
		cel.MemberOverload(
			overloads.BitwiseOrBytes,
			[]*cel.Type{cel.BytesType, cel.BytesType},
			cel.BytesType,
			cel.OverloadExamples(
				`b"\x0f".bitwise_or(b"\xf0") // b"\xff"`,
			),
			cel.BinaryBinding(
				func(lhs, rhs ref.Val) ref.Val {
					src, ok := lhs.(types.Bytes)
					if !ok {
						return types.ValOrErr(lhs, "no such overload")
					}
					temp, ok := rhs.(types.Bytes)
					if !ok {
						return types.ValOrErr(rhs, "no such overload")
					}
					buff := make([]byte, len(src))
					copy(buff, src)
					for k := range buff {
						if k >= len(temp) {
							break
						}
						buff[k] |= temp[k]
					}
					return types.Bytes(buff)
				},
			),
		),
	),

	cel.Function(
		overloads.BitwiseXor,
		cel.FunctionDocs(
			"Performs a bitwise XOR operation on the bytes, combining the bits of two byte sequences. "+
				"Each bit in the result is set to 1 if the corresponding bits in the input bytes are different, otherwise it is set to 0.",
		),
		// bytes.bitwise_xor(bytes) -> bytes
		cel.MemberOverload(
			overloads.BitwiseXorBytes,
			[]*cel.Type{cel.BytesType, cel.BytesType},
			cel.BytesType,
			cel.OverloadExamples(
				`b"\x0f".bitwise_xor(b"\xf0") // b"\xff"`,
			),
			cel.BinaryBinding(
				func(lhs, rhs ref.Val) ref.Val {
					src, ok := lhs.(types.Bytes)
					if !ok {
						return types.ValOrErr(lhs, "no such overload")
					}
					temp, ok := rhs.(types.Bytes)
					if !ok {
						return types.ValOrErr(rhs, "no such overload")
					}
					buff := make([]byte, len(src))
					copy(buff, src)
					for k := range buff {
						if k >= len(temp) {
							break
						}
						buff[k] ^= temp[k]
					}
					return types.Bytes(buff)
				},
			),
		),
	),

	cel.Function(
		overloads.BitwiseClear,
		cel.FunctionDocs(
			"Performs a bitwise clear operation on the bytes, clearing the bits of the first byte sequence where the second byte sequence has bits set to 1. "+
				"Each bit in the result is set to 0 if the corresponding bit in the second byte sequence is 1, otherwise it retains the value from the first byte sequence.",
		),
		// bytes.bitwise_clear(bytes) -> bytes
		cel.MemberOverload(
			overloads.BitwiseClearBytes,
			[]*cel.Type{cel.BytesType, cel.BytesType},
			cel.BytesType,
			cel.OverloadExamples(
				`b"\xff".bitwise_clear(b"\x0f") // b"\xf0"`,
			),
			cel.BinaryBinding(
				func(lhs, rhs ref.Val) ref.Val {
					src, ok := lhs.(types.Bytes)
					if !ok {
						return types.ValOrErr(lhs, "no such overload")
					}
					temp, ok := rhs.(types.Bytes)
					if !ok {
						return types.ValOrErr(rhs, "no such overload")
					}
					buff := make([]byte, len(src))
					copy(buff, src)
					for k := range buff {
						if k >= len(temp) {
							break
						}
						buff[k] &^= temp[k]
					}
					return types.Bytes(buff)
				},
			),
		),
	),

	cel.Function(
		overloads.BitwiseIndex,
		cel.FunctionDocs(
			"Returns the bit at the specified index in the byte sequence. "+
				"The index is zero-based, and the bits are counted from the least significant bit (rightmost) to the most significant bit (leftmost). "+
				"If the index is out of range, an error is returned.",
		),
		// bytes.bitwise_index(int) -> bytes
		cel.MemberOverload(
			overloads.BitwiseIndexInt,
			[]*cel.Type{cel.BytesType, cel.IntType},
			cel.BytesType,
			cel.OverloadExamples(
				`b"\x0f".bitwise_index(3) // b"\x01"`,
			),
			cel.BinaryBinding(
				func(lhs, rhs ref.Val) ref.Val {
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
				},
			),
		),
	),

	cel.Function(
		overloads.BitwiseNot,
		cel.FunctionDocs(
			"Performs a bitwise NOT (ones-complement) on each byte in the byte sequence. "+
				"Each bit in the result is the inverse of the corresponding bit in the input.",
		),
		cel.MemberOverload(
			overloads.BitwiseNotBytes,
			[]*cel.Type{cel.BytesType},
			cel.BytesType,
			cel.OverloadExamples(
				`b"\xf0".bitwise_not() // b"\x0f"`,
				`b"\x00".bitwise_not() // b"\xff"`,
			),
			cel.UnaryBinding(
				func(val ref.Val) ref.Val {
					src, ok := val.(types.Bytes)
					if !ok {
						return types.ValOrErr(val, "no such overload")
					}
					buff := make([]byte, len(src))
					copy(buff, src)
					for i := range buff {
						buff[i] = ^buff[i]
					}
					return types.Bytes(buff)
				},
			),
		),
	),

	cel.Function(
		overloads.BitwisePopcnt,
		cel.FunctionDocs(
			"Counts the number of bits set to 1 (population count) in the byte sequence. "+
				"This is the sum of the popcount across all bytes.",
		),
		cel.MemberOverload(
			overloads.BitwisePopcntBytes,
			[]*cel.Type{cel.BytesType},
			cel.IntType,
			cel.OverloadExamples(
				`b"\xf0".bitwise_popcnt() // 4`,
				`b"\xff\x00".bitwise_popcnt() // 8`,
			),
			cel.UnaryBinding(
				func(val ref.Val) ref.Val {
					src, ok := val.(types.Bytes)
					if !ok {
						return types.ValOrErr(val, "no such overload")
					}
					var cnt int64
					for _, b := range src {
						cnt += int64(bits.OnesCount8(b))
					}
					return types.Int(cnt)
				},
			),
		),
	),
}
