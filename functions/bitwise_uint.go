package functions

import (
	"math/bits"

	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/common/types"
	"github.com/google/cel-go/common/types/ref"
	"github.com/iotxfoundry/cel-go/overloads"
)

var bitwiseUintFunctions = []cel.EnvOption{
	cel.Function(
		overloads.BitwiseAnd,
		cel.FunctionDocs(
			"Performs a bitwise AND operation on the uint values. "+
				"Each bit in the result is set to 1 if both corresponding bits in the input values are 1, otherwise it is set to 0.",
		),
		cel.MemberOverload(
			overloads.UintBitwiseAndUintUint,
			[]*cel.Type{cel.UintType, cel.UintType},
			cel.UintType,
			cel.OverloadExamples(
				`5u.bitwise_and(3u) // 1u`,
				`255u.bitwise_and(15u) // 15u`,
			),
			cel.BinaryBinding(
				func(lhs, rhs ref.Val) ref.Val {
					l, ok := lhs.(types.Uint)
					if !ok {
						return types.ValOrErr(lhs, "no such overload")
					}
					r, ok := rhs.(types.Uint)
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
			"Performs a bitwise OR operation on the uint values. "+
				"Each bit in the result is set to 1 if at least one of the corresponding bits in the input values is 1, otherwise it is set to 0.",
		),
		cel.MemberOverload(
			overloads.UintBitwiseOrUintUint,
			[]*cel.Type{cel.UintType, cel.UintType},
			cel.UintType,
			cel.OverloadExamples(
				`5u.bitwise_or(3u) // 7u`,
				`0u.bitwise_or(1u) // 1u`,
			),
			cel.BinaryBinding(
				func(lhs, rhs ref.Val) ref.Val {
					l, ok := lhs.(types.Uint)
					if !ok {
						return types.ValOrErr(lhs, "no such overload")
					}
					r, ok := rhs.(types.Uint)
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
			"Performs a bitwise XOR operation on the uint values. "+
				"Each bit in the result is set to 1 if the corresponding bits in the input values are different, otherwise it is set to 0.",
		),
		cel.MemberOverload(
			overloads.UintBitwiseXorUintUint,
			[]*cel.Type{cel.UintType, cel.UintType},
			cel.UintType,
			cel.OverloadExamples(
				`5u.bitwise_xor(3u) // 6u`,
				`255u.bitwise_xor(255u) // 0u`,
			),
			cel.BinaryBinding(
				func(lhs, rhs ref.Val) ref.Val {
					l, ok := lhs.(types.Uint)
					if !ok {
						return types.ValOrErr(lhs, "no such overload")
					}
					r, ok := rhs.(types.Uint)
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
			"Performs a bitwise clear operation on the uint values. "+
				"Each bit in the result is set to 0 if the corresponding bit in the second value is 1, otherwise it retains the value from the first value.",
		),
		cel.MemberOverload(
			overloads.UintBitwiseClearUintUint,
			[]*cel.Type{cel.UintType, cel.UintType},
			cel.UintType,
			cel.OverloadExamples(
				`7u.bitwise_clear(1u) // 6u`,
				`255u.bitwise_clear(0u) // 255u`,
			),
			cel.BinaryBinding(
				func(lhs, rhs ref.Val) ref.Val {
					l, ok := lhs.(types.Uint)
					if !ok {
						return types.ValOrErr(lhs, "no such overload")
					}
					r, ok := rhs.(types.Uint)
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
			"Performs a bitwise right shift on the uint value by the specified number of bits. "+
				"Negative shift values shift to the left. Bits shifted out are discarded.",
		),
		cel.MemberOverload(
			overloads.UintBitwiseShiftRightUintInt,
			[]*cel.Type{cel.UintType, cel.IntType},
			cel.UintType,
			cel.OverloadExamples(
				`8u.bitwise_shr(2) // 2u`,
				`2u.bitwise_shr(-2) // 8u`,
			),
			cel.BinaryBinding(
				func(lhs, rhs ref.Val) ref.Val {
					l, ok := lhs.(types.Uint)
					if !ok {
						return types.ValOrErr(lhs, "no such overload")
					}
					r, ok := rhs.(types.Int)
					if !ok {
						return types.ValOrErr(rhs, "no such overload")
					}
					if r < 0 {
						return l << types.Uint(-r)
					}
					return l >> types.Uint(r)
				},
			),
		),
	),

	cel.Function(
		overloads.BitwiseShiftLeft,
		cel.FunctionDocs(
			"Performs a bitwise left shift on the uint value by the specified number of bits. "+
				"Negative shift values shift to the right. Bits shifted out are discarded.",
		),
		cel.MemberOverload(
			overloads.UintBitwiseShiftLeftUintInt,
			[]*cel.Type{cel.UintType, cel.IntType},
			cel.UintType,
			cel.OverloadExamples(
				`2u.bitwise_shl(2) // 8u`,
				`8u.bitwise_shl(-2) // 2u`,
			),
			cel.BinaryBinding(
				func(lhs, rhs ref.Val) ref.Val {
					l, ok := lhs.(types.Uint)
					if !ok {
						return types.ValOrErr(lhs, "no such overload")
					}
					r, ok := rhs.(types.Int)
					if !ok {
						return types.ValOrErr(rhs, "no such overload")
					}
					if r < 0 {
						return l >> types.Uint(-r)
					}
					return l << types.Uint(r)
				},
			),
		),
	),

	cel.Function(
		overloads.BitwiseNot,
		cel.FunctionDocs(
			"Performs a bitwise NOT (ones-complement) on the uint value. "+
				"Each bit in the result is the inverse of the corresponding bit in the input.",
		),
		cel.MemberOverload(
			overloads.UintBitwiseNotUint,
			[]*cel.Type{cel.UintType},
			cel.UintType,
			cel.OverloadExamples(
				`0u.bitwise_not() // 18446744073709551615u`,
				`1u.bitwise_not() // 18446744073709551614u`,
			),
			cel.UnaryBinding(
				func(val ref.Val) ref.Val {
					v, ok := val.(types.Uint)
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
			"Returns the bit at the specified index in the uint value as a single-byte bytes value. "+
				"The index is zero-based, and the bits are counted from the least significant bit (rightmost). "+
				"If the index is out of range, an error is returned.",
		),
		cel.MemberOverload(
			overloads.UintBitwiseIndexUintInt,
			[]*cel.Type{cel.UintType, cel.IntType},
			cel.BytesType,
			cel.OverloadExamples(
				`10u.bitwise_index(1) // b"\x01"`,
				`10u.bitwise_index(0) // b"\x00"`,
			),
			cel.BinaryBinding(
				func(lhs, rhs ref.Val) ref.Val {
					v, ok := lhs.(types.Uint)
					if !ok {
						return types.ValOrErr(lhs, "no such overload")
					}
					index, ok := rhs.(types.Int)
					if !ok {
						return types.ValOrErr(rhs, "no such overload")
					}
					if int(index) >= bits.UintSize || int(index) < 0 {
						return types.NewErr("index '%d' out of range in uint size '%d'", index, bits.UintSize)
					}
					return types.Bytes{byte((uint64(v) >> uint64(index)) & 1)}
				},
			),
		),
	),
}
