package functions

import (
	"bytes"
	"math"
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
			"Performs a bitwise right shift on the bytes, treated as an unsigned big-endian integer, "+
				"by the specified number of bit positions. "+
				"Negative values shift to the left, and positive values shift to the right. "+
				"Bits that are shifted out of the byte are discarded, and new bits are filled with zeros. "+
				"Any shift magnitude is accepted; shifts of eight or more move whole bytes.",
		),
		// bytes.bitwise_shr(int) -> bytes
		cel.MemberOverload(
			overloads.BitwiseShiftRightInt64,
			[]*cel.Type{cel.BytesType, cel.IntType},
			cel.BytesType,
			cel.OverloadExamples(
				`b"\xf0".bitwise_shr(4) // b"\x0f"`,
				`b"\xff\xff".bitwise_shr(8) // b"\x00\xff"`,
			),
			cel.BinaryBinding(
				func(lhs, rhs ref.Val) ref.Val {
					src, ok := lhs.(types.Bytes)
					if !ok {
						return types.ValOrErr(lhs, "no such overload")
					}
					shift, ok := rhs.(types.Int)
					if !ok {
						return types.ValOrErr(rhs, "no such overload")
					}
					return shiftBytes(src, int64(shift), shift < 0)
				},
			),
		),
	),

	cel.Function(
		overloads.BitwiseShiftLeft,
		cel.FunctionDocs(
			"Performs a bitwise left shift on the bytes, treated as an unsigned big-endian integer, "+
				"by the specified number of bit positions. "+
				"Negative values shift to the right, and positive values shift to the left. "+
				"Bits that are shifted out of the byte are discarded, and new bits are filled with zeros. "+
				"Any shift magnitude is accepted; shifts of eight or more move whole bytes.",
		),
		// bytes.bitwise_shl(int) -> bytes
		cel.MemberOverload(
			overloads.BitwiseShiftLeftInt64,
			[]*cel.Type{cel.BytesType, cel.IntType},
			cel.BytesType,
			cel.OverloadExamples(
				`b"\xf0".bitwise_shl(4) // b"\x00"`,
				`b"\x00\xff".bitwise_shl(8) // b"\xff\x00"`,
			),
			cel.BinaryBinding(
				func(lhs, rhs ref.Val) ref.Val {
					src, ok := lhs.(types.Bytes)
					if !ok {
						return types.ValOrErr(lhs, "no such overload")
					}
					shift, ok := rhs.(types.Int)
					if !ok {
						return types.ValOrErr(rhs, "no such overload")
					}
					return shiftBytes(src, int64(shift), shift >= 0)
				},
			),
		),
	),

	cel.Function(
		overloads.BitwiseAnd,
		cel.FunctionDocs(
			"Performs a bitwise AND operation on the bytes, combining the bits of two byte sequences. "+
				"Each bit in the result is set to 1 if both corresponding bits in the input bytes are 1, otherwise it is set to 0. "+
				"The result has the length of the left operand; bytes beyond the length of the right operand are unchanged.",
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
					buff := bytes.Clone([]byte(src))
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
				"Each bit in the result is set to 1 if at least one of the corresponding bits in the input bytes is 1, otherwise it is set to 0. "+
				"The result has the length of the left operand; bytes beyond the length of the right operand are unchanged.",
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
					buff := bytes.Clone([]byte(src))
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
				"Each bit in the result is set to 1 if the corresponding bits in the input bytes are different, otherwise it is set to 0. "+
				"The result has the length of the left operand; bytes beyond the length of the right operand are unchanged.",
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
					buff := bytes.Clone([]byte(src))
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
				"Each bit in the result is set to 0 if the corresponding bit in the second byte sequence is 1, otherwise it retains the value from the first byte sequence. "+
				"The result has the length of the left operand; bytes beyond the length of the right operand are unchanged.",
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
					buff := bytes.Clone([]byte(src))
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
					idx := int(index)
					if idx < 0 || idx >= len(buff)*8 {
						return types.NewErr("index '%d' out of range in bitwise size '%d'", index, len(buff)*8)
					}
					return types.Bytes{buff[idx/8] >> uint(idx%8) & 1}
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
					buff := bytes.Clone([]byte(src))
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

// shiftBytes returns a copy of src shifted, as an unsigned big-endian integer
// (src[0] is the most significant byte), by the given number of bit
// positions. When left is true the bits move toward the most significant
// byte, otherwise toward the least significant byte. Bits shifted out are
// discarded and zeros are shifted in, so every shift magnitude is valid:
// the shift is decomposed into a whole-byte move plus a sub-byte (0-7 bit)
// rotate-free shift, and magnitudes of eight or more bytes zero the buffer.
func shiftBytes(src types.Bytes, shift int64, left bool) types.Bytes {
	data := bytes.Clone([]byte(src))
	n := len(data)
	if n == 0 {
		return types.Bytes(data)
	}
	// Normalize the magnitude; MinInt64 negation would overflow, and its
	// magnitude clamped to MaxInt64 still shifts every bit out.
	mag := shift
	if mag == math.MinInt64 {
		mag = math.MaxInt64
	} else if mag < 0 {
		mag = -mag
	}
	byteShift := int(mag / 8)
	bitShift := uint(mag % 8)
	if byteShift >= n {
		clear(data) // every bit shifted out
		return types.Bytes(data)
	}
	if left {
		// Move bytes toward the most significant end, zero-filling the tail.
		copy(data, data[byteShift:])
		clear(data[n-byteShift:])
	} else {
		// Move bytes toward the least significant end, zero-filling the head.
		copy(data[byteShift:], data)
		clear(data[:byteShift])
	}
	if bitShift > 0 {
		if left {
			for i := 0; i < n-1; i++ {
				data[i] = data[i]<<bitShift | data[i+1]>>(8-bitShift)
			}
			data[n-1] <<= bitShift
		} else {
			for i := n - 1; i > 0; i-- {
				data[i] = data[i]>>bitShift | data[i-1]<<(8-bitShift)
			}
			data[0] >>= bitShift
		}
	}
	return types.Bytes(data)
}
