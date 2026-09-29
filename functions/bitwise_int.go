package functions

import (
	"math"

	"cel.dev/cel-go/cel"
	"cel.dev/cel-go/common/types"
	"cel.dev/cel-go/common/types/ref"
	"github.com/iotxfoundry/cel-go/overloads"
)

var bitwiseIntFunctions = []cel.EnvOption{
	cel.Function(
		overloads.BitwiseAnd,
		cel.FunctionDocs(
			"Performs a bitwise AND operation on the int values. "+
				"Each bit in the result is set to 1 if both corresponding bits in the input values are 1, otherwise it is set to 0.",
		),
		cel.MemberOverload(
			overloads.IntBitwiseAndIntInt,
			[]*cel.Type{cel.IntType, cel.IntType},
			cel.IntType,
			cel.OverloadExamples(
				`5.bitwise_and(3) // 1`,
				`-1.bitwise_and(7) // 7`,
			),
			cel.BinaryBinding(
				func(lhs, rhs ref.Val) ref.Val {
					l, ok := lhs.(types.Int)
					if !ok {
						return types.ValOrErr(lhs, "no such overload")
					}
					r, ok := rhs.(types.Int)
					if !ok {
						return types.ValOrErr(rhs, "no such overload")
					}
					return l & r
				},
			),
		),
	),

	cel.Function(
		overloads.BitwiseOr,
		cel.FunctionDocs(
			"Performs a bitwise OR operation on the int values. "+
				"Each bit in the result is set to 1 if at least one of the corresponding bits in the input values is 1, otherwise it is set to 0.",
		),
		cel.MemberOverload(
			overloads.IntBitwiseOrIntInt,
			[]*cel.Type{cel.IntType, cel.IntType},
			cel.IntType,
			cel.OverloadExamples(
				`5.bitwise_or(3) // 7`,
				`0.bitwise_or(1) // 1`,
			),
			cel.BinaryBinding(
				func(lhs, rhs ref.Val) ref.Val {
					l, ok := lhs.(types.Int)
					if !ok {
						return types.ValOrErr(lhs, "no such overload")
					}
					r, ok := rhs.(types.Int)
					if !ok {
						return types.ValOrErr(rhs, "no such overload")
					}
					return l | r
				},
			),
		),
	),

	cel.Function(
		overloads.BitwiseXor,
		cel.FunctionDocs(
			"Performs a bitwise XOR operation on the int values. "+
				"Each bit in the result is set to 1 if the corresponding bits in the input values are different, otherwise it is set to 0.",
		),
		cel.MemberOverload(
			overloads.IntBitwiseXorIntInt,
			[]*cel.Type{cel.IntType, cel.IntType},
			cel.IntType,
			cel.OverloadExamples(
				`5.bitwise_xor(3) // 6`,
				`-1.bitwise_xor(-1) // 0`,
			),
			cel.BinaryBinding(
				func(lhs, rhs ref.Val) ref.Val {
					l, ok := lhs.(types.Int)
					if !ok {
						return types.ValOrErr(lhs, "no such overload")
					}
					r, ok := rhs.(types.Int)
					if !ok {
						return types.ValOrErr(rhs, "no such overload")
					}
					return l ^ r
				},
			),
		),
	),

	cel.Function(
		overloads.BitwiseClear,
		cel.FunctionDocs(
			"Performs a bitwise clear operation on the int values. "+
				"Each bit in the result is set to 0 if the corresponding bit in the second value is 1, otherwise it retains the value from the first value.",
		),
		cel.MemberOverload(
			overloads.IntBitwiseClearIntInt,
			[]*cel.Type{cel.IntType, cel.IntType},
			cel.IntType,
			cel.OverloadExamples(
				`7.bitwise_clear(1) // 6`,
				`-1.bitwise_clear(0) // -1`,
			),
			cel.BinaryBinding(
				func(lhs, rhs ref.Val) ref.Val {
					l, ok := lhs.(types.Int)
					if !ok {
						return types.ValOrErr(lhs, "no such overload")
					}
					r, ok := rhs.(types.Int)
					if !ok {
						return types.ValOrErr(rhs, "no such overload")
					}
					return l &^ r
				},
			),
		),
	),

	cel.Function(
		overloads.BitwiseShiftRight,
		cel.FunctionDocs(
			"Performs a bitwise right shift on the int value by the specified number of bits. "+
				"Negative shift values shift to the left. Bits shifted out are discarded.",
		),
		cel.MemberOverload(
			overloads.IntBitwiseShiftRightIntInt,
			[]*cel.Type{cel.IntType, cel.IntType},
			cel.IntType,
			cel.OverloadExamples(
				`8.bitwise_shr(2) // 2`,
				`-8.bitwise_shr(2) // -2`,
			),
			cel.BinaryBinding(
				func(lhs, rhs ref.Val) ref.Val {
					l, ok := lhs.(types.Int)
					if !ok {
						return types.ValOrErr(lhs, "no such overload")
					}
					r, ok := rhs.(types.Int)
					if !ok {
						return types.ValOrErr(rhs, "no such overload")
					}
					// Normalize the shift magnitude before negating: negating
					// MinInt64 would overflow back to a negative count and
					// panic. MaxInt64 behaves identically (a full-width shift).
					mag := r
					if mag == math.MinInt64 {
						mag = math.MaxInt64
					} else if mag < 0 {
						mag = -mag
					}
					if r < 0 {
						return l << mag
					}
					return l >> mag
				},
			),
		),
	),

	cel.Function(
		overloads.BitwiseShiftLeft,
		cel.FunctionDocs(
			"Performs a bitwise left shift on the int value by the specified number of bits. "+
				"Negative shift values shift to the right. Bits shifted out are discarded.",
		),
		cel.MemberOverload(
			overloads.IntBitwiseShiftLeftIntInt,
			[]*cel.Type{cel.IntType, cel.IntType},
			cel.IntType,
			cel.OverloadExamples(
				`2.bitwise_shl(2) // 8`,
				`8.bitwise_shl(-2) // 2`,
			),
			cel.BinaryBinding(
				func(lhs, rhs ref.Val) ref.Val {
					l, ok := lhs.(types.Int)
					if !ok {
						return types.ValOrErr(lhs, "no such overload")
					}
					r, ok := rhs.(types.Int)
					if !ok {
						return types.ValOrErr(rhs, "no such overload")
					}
					// Normalize the shift magnitude before negating: negating
					// MinInt64 would overflow back to a negative count and
					// panic. MaxInt64 behaves identically (a full-width shift).
					mag := r
					if mag == math.MinInt64 {
						mag = math.MaxInt64
					} else if mag < 0 {
						mag = -mag
					}
					if r < 0 {
						return l >> mag
					}
					return l << mag
				},
			),
		),
	),

	cel.Function(
		overloads.BitwiseNot,
		cel.FunctionDocs(
			"Performs a bitwise NOT (ones-complement) on the int value. "+
				"Each bit in the result is the inverse of the corresponding bit in the input.",
		),
		cel.MemberOverload(
			overloads.IntBitwiseNotInt,
			[]*cel.Type{cel.IntType},
			cel.IntType,
			cel.OverloadExamples(
				`0.bitwise_not() // -1`,
				`-1.bitwise_not() // 0`,
			),
			cel.UnaryBinding(
				func(val ref.Val) ref.Val {
					v, ok := val.(types.Int)
					if !ok {
						return types.ValOrErr(val, "no such overload")
					}
					return ^v
				},
			),
		),
	),

	cel.Function(
		overloads.BitwiseIndex,
		cel.FunctionDocs(
			"Returns the bit at the specified index in the int value as a single-byte bytes value. "+
				"The index is zero-based, and the bits are counted from the least significant bit (rightmost). "+
				"If the index is out of range, an error is returned.",
		),
		cel.MemberOverload(
			overloads.IntBitwiseIndexIntInt,
			[]*cel.Type{cel.IntType, cel.IntType},
			cel.BytesType,
			cel.OverloadExamples(
				`10.bitwise_index(1) // b"\x01"`,
				`10.bitwise_index(0) // b"\x00"`,
			),
			cel.BinaryBinding(
				func(lhs, rhs ref.Val) ref.Val {
					v, ok := lhs.(types.Int)
					if !ok {
						return types.ValOrErr(lhs, "no such overload")
					}
					index, ok := rhs.(types.Int)
					if !ok {
						return types.ValOrErr(rhs, "no such overload")
					}
					// int64 is always 64 bits wide, independent of the
					// platform's native int size.
					if index < 0 || index >= 64 {
						return types.NewErr("index '%d' out of range in int size '%d'", index, 64)
					}
					return types.Bytes{byte((uint64(v) >> uint64(index)) & 1)}
				},
			),
		),
	),
}
