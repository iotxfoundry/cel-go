package functions

import (
	"math"

	"cel.dev/cel-go/cel"
	"cel.dev/cel-go/common/types"
	"cel.dev/cel-go/common/types/ref"
)

// arithmeticFunctions provides cross-type arithmetic member functions
// for int, uint, and double types.
//
// CEL's built-in operators (+, -, *, /, %) only support same-type arithmetic
// via trait-based dispatch and cannot be extended with new overloads.
// These member functions fill the gap:
//
//	3.add(4u)   // int + uint → int
//	3u.add(4.5) // uint + double → double
//
// Type promotion rules:
//   - int/uint mixes return int (signed type needed for possible negative results)
//   - any mix with double returns double (double has the widest range)
//
// Error semantics mirror CEL's built-in operators:
//   - add/sub/mul return an "integer overflow" error instead of wrapping
//   - div/mod by zero return a "division by zero" / "modulus by zero" error
//   - MinInt64 divided or multiplied by -1 returns an "integer overflow" error
//
// mod is intentionally limited to int/uint receivers and arguments, mirroring
// CEL's % operator which has no double support. Use div for double operands.
var arithmeticFunctions = []cel.EnvOption{

	// --- add ---
	cel.Function("add",
		// int receiver
		cel.MemberOverload("arith_add_int64_int64", []*cel.Type{cel.IntType, cel.IntType}, cel.IntType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				i, ok := lhs.(types.Int)
				if !ok {
					return types.ValOrErr(lhs, "no such overload")
				}
				j, ok := rhs.(types.Int)
				if !ok {
					return types.ValOrErr(rhs, "no such overload")
				}
				return i.Add(j)
			})),
		cel.MemberOverload("arith_add_int64_uint64", []*cel.Type{cel.IntType, cel.UintType}, cel.IntType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				i, ok := lhs.(types.Int)
				if !ok {
					return types.ValOrErr(lhs, "no such overload")
				}
				u, ok := rhs.(types.Uint)
				if !ok {
					return types.ValOrErr(rhs, "no such overload")
				}
				return checkedAddInt64Uint64(int64(i), uint64(u))
			})),
		cel.MemberOverload("arith_add_int64_double", []*cel.Type{cel.IntType, cel.DoubleType}, cel.DoubleType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				i, ok := lhs.(types.Int)
				if !ok {
					return types.ValOrErr(lhs, "no such overload")
				}
				d, ok := rhs.(types.Double)
				if !ok {
					return types.ValOrErr(rhs, "no such overload")
				}
				return types.Double(float64(i) + float64(d))
			})),
		// uint receiver
		cel.MemberOverload("arith_add_uint64_uint64", []*cel.Type{cel.UintType, cel.UintType}, cel.UintType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				u, ok := lhs.(types.Uint)
				if !ok {
					return types.ValOrErr(lhs, "no such overload")
				}
				v, ok := rhs.(types.Uint)
				if !ok {
					return types.ValOrErr(rhs, "no such overload")
				}
				return u.Add(v)
			})),
		cel.MemberOverload("arith_add_uint64_int64", []*cel.Type{cel.UintType, cel.IntType}, cel.IntType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				u, ok := lhs.(types.Uint)
				if !ok {
					return types.ValOrErr(lhs, "no such overload")
				}
				i, ok := rhs.(types.Int)
				if !ok {
					return types.ValOrErr(rhs, "no such overload")
				}
				return checkedAddUint64Int64(uint64(u), int64(i))
			})),
		cel.MemberOverload("arith_add_uint64_double", []*cel.Type{cel.UintType, cel.DoubleType}, cel.DoubleType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				u, ok := lhs.(types.Uint)
				if !ok {
					return types.ValOrErr(lhs, "no such overload")
				}
				d, ok := rhs.(types.Double)
				if !ok {
					return types.ValOrErr(rhs, "no such overload")
				}
				return types.Double(float64(u) + float64(d))
			})),
		// double receiver
		cel.MemberOverload("arith_add_double_double", []*cel.Type{cel.DoubleType, cel.DoubleType}, cel.DoubleType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				d, ok := lhs.(types.Double)
				if !ok {
					return types.ValOrErr(lhs, "no such overload")
				}
				v, ok := rhs.(types.Double)
				if !ok {
					return types.ValOrErr(rhs, "no such overload")
				}
				return d + v
			})),
		cel.MemberOverload("arith_add_double_int64", []*cel.Type{cel.DoubleType, cel.IntType}, cel.DoubleType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				d, ok := lhs.(types.Double)
				if !ok {
					return types.ValOrErr(lhs, "no such overload")
				}
				i, ok := rhs.(types.Int)
				if !ok {
					return types.ValOrErr(rhs, "no such overload")
				}
				return types.Double(float64(d) + float64(i))
			})),
		cel.MemberOverload("arith_add_double_uint64", []*cel.Type{cel.DoubleType, cel.UintType}, cel.DoubleType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				d, ok := lhs.(types.Double)
				if !ok {
					return types.ValOrErr(lhs, "no such overload")
				}
				u, ok := rhs.(types.Uint)
				if !ok {
					return types.ValOrErr(rhs, "no such overload")
				}
				return types.Double(float64(d) + float64(u))
			})),
	),

	// --- sub ---
	cel.Function("sub",
		// int receiver
		cel.MemberOverload("arith_sub_int64_int64", []*cel.Type{cel.IntType, cel.IntType}, cel.IntType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				i, ok := lhs.(types.Int)
				if !ok {
					return types.ValOrErr(lhs, "no such overload")
				}
				j, ok := rhs.(types.Int)
				if !ok {
					return types.ValOrErr(rhs, "no such overload")
				}
				return i.Subtract(j)
			})),
		cel.MemberOverload("arith_sub_int64_uint64", []*cel.Type{cel.IntType, cel.UintType}, cel.IntType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				i, ok := lhs.(types.Int)
				if !ok {
					return types.ValOrErr(lhs, "no such overload")
				}
				u, ok := rhs.(types.Uint)
				if !ok {
					return types.ValOrErr(rhs, "no such overload")
				}
				return checkedSubInt64Uint64(int64(i), uint64(u))
			})),
		cel.MemberOverload("arith_sub_int64_double", []*cel.Type{cel.IntType, cel.DoubleType}, cel.DoubleType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				i, ok := lhs.(types.Int)
				if !ok {
					return types.ValOrErr(lhs, "no such overload")
				}
				d, ok := rhs.(types.Double)
				if !ok {
					return types.ValOrErr(rhs, "no such overload")
				}
				return types.Double(float64(i) - float64(d))
			})),
		// uint receiver
		cel.MemberOverload("arith_sub_uint64_uint64", []*cel.Type{cel.UintType, cel.UintType}, cel.UintType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				u, ok := lhs.(types.Uint)
				if !ok {
					return types.ValOrErr(lhs, "no such overload")
				}
				v, ok := rhs.(types.Uint)
				if !ok {
					return types.ValOrErr(rhs, "no such overload")
				}
				return u.Subtract(v)
			})),
		cel.MemberOverload("arith_sub_uint64_int64", []*cel.Type{cel.UintType, cel.IntType}, cel.IntType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				u, ok := lhs.(types.Uint)
				if !ok {
					return types.ValOrErr(lhs, "no such overload")
				}
				i, ok := rhs.(types.Int)
				if !ok {
					return types.ValOrErr(rhs, "no such overload")
				}
				return checkedSubUint64Int64(uint64(u), int64(i))
			})),
		cel.MemberOverload("arith_sub_uint64_double", []*cel.Type{cel.UintType, cel.DoubleType}, cel.DoubleType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				u, ok := lhs.(types.Uint)
				if !ok {
					return types.ValOrErr(lhs, "no such overload")
				}
				d, ok := rhs.(types.Double)
				if !ok {
					return types.ValOrErr(rhs, "no such overload")
				}
				return types.Double(float64(u) - float64(d))
			})),
		// double receiver
		cel.MemberOverload("arith_sub_double_double", []*cel.Type{cel.DoubleType, cel.DoubleType}, cel.DoubleType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				d, ok := lhs.(types.Double)
				if !ok {
					return types.ValOrErr(lhs, "no such overload")
				}
				v, ok := rhs.(types.Double)
				if !ok {
					return types.ValOrErr(rhs, "no such overload")
				}
				return d - v
			})),
		cel.MemberOverload("arith_sub_double_int64", []*cel.Type{cel.DoubleType, cel.IntType}, cel.DoubleType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				d, ok := lhs.(types.Double)
				if !ok {
					return types.ValOrErr(lhs, "no such overload")
				}
				i, ok := rhs.(types.Int)
				if !ok {
					return types.ValOrErr(rhs, "no such overload")
				}
				return types.Double(float64(d) - float64(i))
			})),
		cel.MemberOverload("arith_sub_double_uint64", []*cel.Type{cel.DoubleType, cel.UintType}, cel.DoubleType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				d, ok := lhs.(types.Double)
				if !ok {
					return types.ValOrErr(lhs, "no such overload")
				}
				u, ok := rhs.(types.Uint)
				if !ok {
					return types.ValOrErr(rhs, "no such overload")
				}
				return types.Double(float64(d) - float64(u))
			})),
	),

	// --- mul ---
	cel.Function("mul",
		// int receiver
		cel.MemberOverload("arith_mul_int64_int64", []*cel.Type{cel.IntType, cel.IntType}, cel.IntType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				i, ok := lhs.(types.Int)
				if !ok {
					return types.ValOrErr(lhs, "no such overload")
				}
				j, ok := rhs.(types.Int)
				if !ok {
					return types.ValOrErr(rhs, "no such overload")
				}
				return i.Multiply(j)
			})),
		cel.MemberOverload("arith_mul_int64_uint64", []*cel.Type{cel.IntType, cel.UintType}, cel.IntType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				i, ok := lhs.(types.Int)
				if !ok {
					return types.ValOrErr(lhs, "no such overload")
				}
				u, ok := rhs.(types.Uint)
				if !ok {
					return types.ValOrErr(rhs, "no such overload")
				}
				return checkedMulInt64Uint64(int64(i), uint64(u))
			})),
		cel.MemberOverload("arith_mul_int64_double", []*cel.Type{cel.IntType, cel.DoubleType}, cel.DoubleType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				i, ok := lhs.(types.Int)
				if !ok {
					return types.ValOrErr(lhs, "no such overload")
				}
				d, ok := rhs.(types.Double)
				if !ok {
					return types.ValOrErr(rhs, "no such overload")
				}
				return types.Double(float64(i) * float64(d))
			})),
		// uint receiver
		cel.MemberOverload("arith_mul_uint64_uint64", []*cel.Type{cel.UintType, cel.UintType}, cel.UintType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				u, ok := lhs.(types.Uint)
				if !ok {
					return types.ValOrErr(lhs, "no such overload")
				}
				v, ok := rhs.(types.Uint)
				if !ok {
					return types.ValOrErr(rhs, "no such overload")
				}
				return u.Multiply(v)
			})),
		cel.MemberOverload("arith_mul_uint64_int64", []*cel.Type{cel.UintType, cel.IntType}, cel.IntType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				u, ok := lhs.(types.Uint)
				if !ok {
					return types.ValOrErr(lhs, "no such overload")
				}
				i, ok := rhs.(types.Int)
				if !ok {
					return types.ValOrErr(rhs, "no such overload")
				}
				return checkedMulUint64Int64(uint64(u), int64(i))
			})),
		cel.MemberOverload("arith_mul_uint64_double", []*cel.Type{cel.UintType, cel.DoubleType}, cel.DoubleType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				u, ok := lhs.(types.Uint)
				if !ok {
					return types.ValOrErr(lhs, "no such overload")
				}
				d, ok := rhs.(types.Double)
				if !ok {
					return types.ValOrErr(rhs, "no such overload")
				}
				return types.Double(float64(u) * float64(d))
			})),
		// double receiver
		cel.MemberOverload("arith_mul_double_double", []*cel.Type{cel.DoubleType, cel.DoubleType}, cel.DoubleType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				d, ok := lhs.(types.Double)
				if !ok {
					return types.ValOrErr(lhs, "no such overload")
				}
				v, ok := rhs.(types.Double)
				if !ok {
					return types.ValOrErr(rhs, "no such overload")
				}
				return d * v
			})),
		cel.MemberOverload("arith_mul_double_int64", []*cel.Type{cel.DoubleType, cel.IntType}, cel.DoubleType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				d, ok := lhs.(types.Double)
				if !ok {
					return types.ValOrErr(lhs, "no such overload")
				}
				i, ok := rhs.(types.Int)
				if !ok {
					return types.ValOrErr(rhs, "no such overload")
				}
				return types.Double(float64(d) * float64(i))
			})),
		cel.MemberOverload("arith_mul_double_uint64", []*cel.Type{cel.DoubleType, cel.UintType}, cel.DoubleType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				d, ok := lhs.(types.Double)
				if !ok {
					return types.ValOrErr(lhs, "no such overload")
				}
				u, ok := rhs.(types.Uint)
				if !ok {
					return types.ValOrErr(rhs, "no such overload")
				}
				return types.Double(float64(d) * float64(u))
			})),
	),

	// --- div ---
	cel.Function("div",
		// int receiver
		cel.MemberOverload("arith_div_int64_int64", []*cel.Type{cel.IntType, cel.IntType}, cel.IntType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				i, ok := lhs.(types.Int)
				if !ok {
					return types.ValOrErr(lhs, "no such overload")
				}
				j, ok := rhs.(types.Int)
				if !ok {
					return types.ValOrErr(rhs, "no such overload")
				}
				return i.Divide(j)
			})),
		cel.MemberOverload("arith_div_int64_uint64", []*cel.Type{cel.IntType, cel.UintType}, cel.IntType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				i, ok := lhs.(types.Int)
				if !ok {
					return types.ValOrErr(lhs, "no such overload")
				}
				u, ok := rhs.(types.Uint)
				if !ok {
					return types.ValOrErr(rhs, "no such overload")
				}
				return checkedDivInt64Uint64(int64(i), uint64(u))
			})),
		cel.MemberOverload("arith_div_int64_double", []*cel.Type{cel.IntType, cel.DoubleType}, cel.DoubleType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				i, ok := lhs.(types.Int)
				if !ok {
					return types.ValOrErr(lhs, "no such overload")
				}
				d, ok := rhs.(types.Double)
				if !ok {
					return types.ValOrErr(rhs, "no such overload")
				}
				return types.Double(float64(i) / float64(d))
			})),
		// uint receiver
		cel.MemberOverload("arith_div_uint64_uint64", []*cel.Type{cel.UintType, cel.UintType}, cel.UintType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				u, ok := lhs.(types.Uint)
				if !ok {
					return types.ValOrErr(lhs, "no such overload")
				}
				v, ok := rhs.(types.Uint)
				if !ok {
					return types.ValOrErr(rhs, "no such overload")
				}
				return u.Divide(v)
			})),
		cel.MemberOverload("arith_div_uint64_int64", []*cel.Type{cel.UintType, cel.IntType}, cel.IntType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				u, ok := lhs.(types.Uint)
				if !ok {
					return types.ValOrErr(lhs, "no such overload")
				}
				i, ok := rhs.(types.Int)
				if !ok {
					return types.ValOrErr(rhs, "no such overload")
				}
				return checkedDivUint64Int64(uint64(u), int64(i))
			})),
		cel.MemberOverload("arith_div_uint64_double", []*cel.Type{cel.UintType, cel.DoubleType}, cel.DoubleType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				u, ok := lhs.(types.Uint)
				if !ok {
					return types.ValOrErr(lhs, "no such overload")
				}
				d, ok := rhs.(types.Double)
				if !ok {
					return types.ValOrErr(rhs, "no such overload")
				}
				return types.Double(float64(u) / float64(d))
			})),
		// double receiver
		cel.MemberOverload("arith_div_double_double", []*cel.Type{cel.DoubleType, cel.DoubleType}, cel.DoubleType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				d, ok := lhs.(types.Double)
				if !ok {
					return types.ValOrErr(lhs, "no such overload")
				}
				v, ok := rhs.(types.Double)
				if !ok {
					return types.ValOrErr(rhs, "no such overload")
				}
				return d / v
			})),
		cel.MemberOverload("arith_div_double_int64", []*cel.Type{cel.DoubleType, cel.IntType}, cel.DoubleType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				d, ok := lhs.(types.Double)
				if !ok {
					return types.ValOrErr(lhs, "no such overload")
				}
				i, ok := rhs.(types.Int)
				if !ok {
					return types.ValOrErr(rhs, "no such overload")
				}
				return types.Double(float64(d) / float64(i))
			})),
		cel.MemberOverload("arith_div_double_uint64", []*cel.Type{cel.DoubleType, cel.UintType}, cel.DoubleType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				d, ok := lhs.(types.Double)
				if !ok {
					return types.ValOrErr(lhs, "no such overload")
				}
				u, ok := rhs.(types.Uint)
				if !ok {
					return types.ValOrErr(rhs, "no such overload")
				}
				return types.Double(float64(d) / float64(u))
			})),
	),

	// --- mod ---
	// mod has no double overloads on purpose: CEL's % operator does not
	// support double operands either. Use div for floating-point operands.
	cel.Function("mod",
		// int receiver
		cel.MemberOverload("arith_mod_int64_int64", []*cel.Type{cel.IntType, cel.IntType}, cel.IntType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				i, ok := lhs.(types.Int)
				if !ok {
					return types.ValOrErr(lhs, "no such overload")
				}
				j, ok := rhs.(types.Int)
				if !ok {
					return types.ValOrErr(rhs, "no such overload")
				}
				return i.Modulo(j)
			})),
		cel.MemberOverload("arith_mod_int64_uint64", []*cel.Type{cel.IntType, cel.UintType}, cel.IntType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				i, ok := lhs.(types.Int)
				if !ok {
					return types.ValOrErr(lhs, "no such overload")
				}
				u, ok := rhs.(types.Uint)
				if !ok {
					return types.ValOrErr(rhs, "no such overload")
				}
				return checkedModInt64Uint64(int64(i), uint64(u))
			})),
		// uint receiver
		cel.MemberOverload("arith_mod_uint64_uint64", []*cel.Type{cel.UintType, cel.UintType}, cel.UintType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				u, ok := lhs.(types.Uint)
				if !ok {
					return types.ValOrErr(lhs, "no such overload")
				}
				v, ok := rhs.(types.Uint)
				if !ok {
					return types.ValOrErr(rhs, "no such overload")
				}
				return u.Modulo(v)
			})),
		cel.MemberOverload("arith_mod_uint64_int64", []*cel.Type{cel.UintType, cel.IntType}, cel.IntType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				u, ok := lhs.(types.Uint)
				if !ok {
					return types.ValOrErr(lhs, "no such overload")
				}
				i, ok := rhs.(types.Int)
				if !ok {
					return types.ValOrErr(rhs, "no such overload")
				}
				return checkedModUint64Int64(uint64(u), int64(i))
			})),
	),
}

// errIntOverflow is returned by the checked cross-type arithmetic helpers when
// the result cannot be represented in the destination type. The message
// matches CEL's built-in operator overflow error.
const errIntOverflow = "integer overflow"

// checkedAddInt64Uint64 computes x + y as an int64, returning a CEL
// integer overflow error when the mathematical result does not fit in int64.
func checkedAddInt64Uint64(x int64, y uint64) ref.Val {
	// x + y <= MaxInt64  ⇔  y <= MaxInt64 - x. The subtraction is computed in
	// uint64 space, where the two's-complement wraparound of negative x yields
	// the correct non-negative bound.
	if y > uint64(math.MaxInt64)-uint64(x) {
		return types.NewErr(errIntOverflow)
	}
	return types.Int(x + int64(y))
}

// checkedSubInt64Uint64 computes x - y as an int64, returning a CEL
// integer overflow error when the mathematical result does not fit in int64.
func checkedSubInt64Uint64(x int64, y uint64) ref.Val {
	// x - y >= MinInt64  ⇔  y <= x - MinInt64 = x + 2^63, computed in uint64
	// space where the wraparound is exact.
	if y > uint64(x)+1<<63 {
		return types.NewErr(errIntOverflow)
	}
	return types.Int(x - int64(y))
}

// checkedMulInt64Uint64 computes x * y as an int64, returning a CEL
// integer overflow error when the mathematical result does not fit in int64.
func checkedMulInt64Uint64(x int64, y uint64) ref.Val {
	switch {
	case x == 0 || y == 0:
		return types.Int(0)
	case x > 0:
		if uint64(x) > uint64(math.MaxInt64)/y {
			return types.NewErr(errIntOverflow)
		}
	case x < 0:
		// The product is negative; it fits iff |x| * y <= 2^63. The boundary
		// case |x| * y == 2^63 is exactly MinInt64 and is allowed.
		if uint64(-x) > 1<<63/y {
			return types.NewErr(errIntOverflow)
		}
	}
	return types.Int(int64(uint64(x) * y))
}

// checkedDivInt64Uint64 computes x / y as an int64, returning a CEL
// "division by zero" error. No further overflow is possible: the only
// overflowing int64 division is MinInt64 / -1, and no uint64 equals -1.
func checkedDivInt64Uint64(x int64, y uint64) ref.Val {
	if y == 0 {
		return types.NewErr("division by zero")
	}
	if y > math.MaxInt64 {
		// The divisor exceeds the magnitude of every int64 dividend except
		// the exact pairing MinInt64 / 2^63 = -1.
		if x == math.MinInt64 && y == 1<<63 {
			return types.Int(-1)
		}
		return types.Int(0)
	}
	return types.Int(x / int64(y))
}

// checkedModInt64Uint64 computes x % y as an int64, returning a CEL
// "modulus by zero" error.
func checkedModInt64Uint64(x int64, y uint64) ref.Val {
	if y == 0 {
		return types.NewErr("modulus by zero")
	}
	if y > math.MaxInt64 {
		// The divisor magnitude exceeds (or exactly matches) that of every
		// int64 dividend, so the quotient is 0 and the remainder is the
		// dividend itself -- except the exact pairing MinInt64 % 2^63 = 0.
		if x == math.MinInt64 && y == 1<<63 {
			return types.Int(0)
		}
		return types.Int(x)
	}
	return types.Int(x % int64(y))
}

// checkedAddUint64Int64 computes x + y as an int64, returning a CEL
// integer overflow error when the mathematical result does not fit in int64.
// The lower bound never binds (x >= 0 and y >= MinInt64 imply x+y >= MinInt64),
// so only the upper bound is checked; the wraparound addition is exact.
func checkedAddUint64Int64(x uint64, y int64) ref.Val {
	if x > uint64(math.MaxInt64)-uint64(y) {
		return types.NewErr(errIntOverflow)
	}
	return types.Int(int64(x + uint64(y)))
}

// checkedSubUint64Int64 computes x - y as an int64, returning a CEL
// integer overflow error when the mathematical result does not fit in int64.
func checkedSubUint64Int64(x uint64, y int64) ref.Val {
	// The lower bound never binds: x - y >= -MaxInt64 > MinInt64 for
	// x >= 0 and y <= MaxInt64.
	if y == math.MinInt64 {
		// x - y == x + 2^63 exceeds MaxInt64 for every unsigned x. Negating
		// MinInt64 would wrap, so this needs an explicit case.
		return types.NewErr(errIntOverflow)
	}
	if y >= 0 {
		if x > uint64(math.MaxInt64)+uint64(y) {
			return types.NewErr(errIntOverflow)
		}
	} else if x > uint64(math.MaxInt64+y) { // y >= MinInt64+1 keeps the bound >= 0
		return types.NewErr(errIntOverflow)
	}
	return types.Int(int64(x - uint64(y)))
}

// checkedMulUint64Int64 computes x * y as an int64, returning a CEL
// integer overflow error when the mathematical result does not fit in int64.
func checkedMulUint64Int64(x uint64, y int64) ref.Val {
	switch {
	case x == 0 || y == 0:
		return types.Int(0)
	case y > 0:
		if x > uint64(math.MaxInt64)/uint64(y) {
			return types.NewErr(errIntOverflow)
		}
		return types.Int(int64(x * uint64(y)))
	default:
		// y < 0: the product is negative; it fits iff x * |y| <= 2^63. The
		// boundary case x * |y| == 2^63 is exactly MinInt64 and is allowed.
		if x > 1<<63/uint64(-y) {
			return types.NewErr(errIntOverflow)
		}
		return types.Int(-int64(x * uint64(-y)))
	}
}

// checkedModUint64Int64 computes x % y as an int64, returning a CEL
// "modulus by zero" error. The remainder satisfies 0 <= r < |y| <= 2^63 and
// therefore always fits in int64, so no overflow check is required.
func checkedModUint64Int64(x uint64, y int64) ref.Val {
	if y == 0 {
		return types.NewErr("modulus by zero")
	}
	if y < 0 {
		// -y is in [1, 2^63] and fits uint64 exactly.
		return types.Int(int64(x % uint64(-y)))
	}
	return types.Int(int64(x % uint64(y)))
}

// checkedDivUint64Int64 computes x / y as an int64, returning a CEL
// "division by zero" or "integer overflow" error. The quotient is computed
// with unsigned division so that wraparound of x above MaxInt64 cannot
// corrupt the sign or magnitude of the result.
func checkedDivUint64Int64(x uint64, y int64) ref.Val {
	if y == 0 {
		return types.NewErr("division by zero")
	}
	if y > 0 {
		// y >= 2 caps the quotient at MaxUint64/2 == MaxInt64; only y == 1
		// can produce a quotient above MaxInt64.
		if y == 1 && x >= 1<<63 {
			return types.NewErr(errIntOverflow)
		}
		return types.Int(int64(x / uint64(y)))
	}
	// y < 0: the result is negative; it fits iff floor(x / |y|) <= 2^63,
	// which can only fail when |y| == 1.
	if y == -1 && x > 1<<63 {
		return types.NewErr(errIntOverflow)
	}
	return types.Int(-int64(x / uint64(-y)))
}
