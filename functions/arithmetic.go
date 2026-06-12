package functions

import (
	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/common/types"
	"github.com/google/cel-go/common/types/ref"
)

// arithmeticFunctions provides cross-type arithmetic member functions
// for int, uint, and double types.
//
// CEL's built-in operators (+, -, *, /, %) only support same-type arithmetic
// via trait-based dispatch and cannot be extended with new overloads.
// These member functions fill the gap:
//   3.add(4u)   // int + uint → int
//   3u.add(4.5) // uint + double → double
//
// Type promotion rules:
//   - int/uint mixes return int (signed type needed for possible negative results)
//   - any mix with double returns double (double has the widest range)
var arithmeticFunctions = []cel.EnvOption{

	// --- add ---
	cel.Function("add",
		// int receiver
		cel.MemberOverload("add_int64_int64", []*cel.Type{cel.IntType, cel.IntType}, cel.IntType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				return lhs.(types.Int) + rhs.(types.Int)
			})),
		cel.MemberOverload("add_int64_uint64", []*cel.Type{cel.IntType, cel.UintType}, cel.IntType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				return types.Int(int64(lhs.(types.Int)) + int64(rhs.(types.Uint)))
			})),
		cel.MemberOverload("add_int64_double", []*cel.Type{cel.IntType, cel.DoubleType}, cel.DoubleType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				return types.Double(float64(lhs.(types.Int)) + float64(rhs.(types.Double)))
			})),
		// uint receiver
		cel.MemberOverload("add_uint64_uint64", []*cel.Type{cel.UintType, cel.UintType}, cel.UintType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				return lhs.(types.Uint) + rhs.(types.Uint)
			})),
		cel.MemberOverload("add_uint64_int64", []*cel.Type{cel.UintType, cel.IntType}, cel.IntType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				return types.Int(int64(lhs.(types.Uint)) + int64(rhs.(types.Int)))
			})),
		cel.MemberOverload("add_uint64_double", []*cel.Type{cel.UintType, cel.DoubleType}, cel.DoubleType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				return types.Double(float64(lhs.(types.Uint)) + float64(rhs.(types.Double)))
			})),
		// double receiver
		cel.MemberOverload("add_double_double", []*cel.Type{cel.DoubleType, cel.DoubleType}, cel.DoubleType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				return lhs.(types.Double) + rhs.(types.Double)
			})),
		cel.MemberOverload("add_double_int64", []*cel.Type{cel.DoubleType, cel.IntType}, cel.DoubleType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				return types.Double(float64(lhs.(types.Double)) + float64(rhs.(types.Int)))
			})),
		cel.MemberOverload("add_double_uint64", []*cel.Type{cel.DoubleType, cel.UintType}, cel.DoubleType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				return types.Double(float64(lhs.(types.Double)) + float64(rhs.(types.Uint)))
			})),
	),

	// --- sub ---
	cel.Function("sub",
		// int receiver
		cel.MemberOverload("sub_int64_int64", []*cel.Type{cel.IntType, cel.IntType}, cel.IntType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				return lhs.(types.Int) - rhs.(types.Int)
			})),
		cel.MemberOverload("sub_int64_uint64", []*cel.Type{cel.IntType, cel.UintType}, cel.IntType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				return types.Int(int64(lhs.(types.Int)) - int64(rhs.(types.Uint)))
			})),
		cel.MemberOverload("sub_int64_double", []*cel.Type{cel.IntType, cel.DoubleType}, cel.DoubleType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				return types.Double(float64(lhs.(types.Int)) - float64(rhs.(types.Double)))
			})),
		// uint receiver
		cel.MemberOverload("sub_uint64_uint64", []*cel.Type{cel.UintType, cel.UintType}, cel.UintType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				return lhs.(types.Uint) - rhs.(types.Uint)
			})),
		cel.MemberOverload("sub_uint64_int64", []*cel.Type{cel.UintType, cel.IntType}, cel.IntType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				return types.Int(int64(lhs.(types.Uint)) - int64(rhs.(types.Int)))
			})),
		cel.MemberOverload("sub_uint64_double", []*cel.Type{cel.UintType, cel.DoubleType}, cel.DoubleType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				return types.Double(float64(lhs.(types.Uint)) - float64(rhs.(types.Double)))
			})),
		// double receiver
		cel.MemberOverload("sub_double_double", []*cel.Type{cel.DoubleType, cel.DoubleType}, cel.DoubleType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				return lhs.(types.Double) - rhs.(types.Double)
			})),
		cel.MemberOverload("sub_double_int64", []*cel.Type{cel.DoubleType, cel.IntType}, cel.DoubleType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				return types.Double(float64(lhs.(types.Double)) - float64(rhs.(types.Int)))
			})),
		cel.MemberOverload("sub_double_uint64", []*cel.Type{cel.DoubleType, cel.UintType}, cel.DoubleType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				return types.Double(float64(lhs.(types.Double)) - float64(rhs.(types.Uint)))
			})),
	),

	// --- mul ---
	cel.Function("mul",
		// int receiver
		cel.MemberOverload("mul_int64_int64", []*cel.Type{cel.IntType, cel.IntType}, cel.IntType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				return lhs.(types.Int) * rhs.(types.Int)
			})),
		cel.MemberOverload("mul_int64_uint64", []*cel.Type{cel.IntType, cel.UintType}, cel.IntType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				return types.Int(int64(lhs.(types.Int)) * int64(rhs.(types.Uint)))
			})),
		cel.MemberOverload("mul_int64_double", []*cel.Type{cel.IntType, cel.DoubleType}, cel.DoubleType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				return types.Double(float64(lhs.(types.Int)) * float64(rhs.(types.Double)))
			})),
		// uint receiver
		cel.MemberOverload("mul_uint64_uint64", []*cel.Type{cel.UintType, cel.UintType}, cel.UintType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				return lhs.(types.Uint) * rhs.(types.Uint)
			})),
		cel.MemberOverload("mul_uint64_int64", []*cel.Type{cel.UintType, cel.IntType}, cel.IntType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				return types.Int(int64(lhs.(types.Uint)) * int64(rhs.(types.Int)))
			})),
		cel.MemberOverload("mul_uint64_double", []*cel.Type{cel.UintType, cel.DoubleType}, cel.DoubleType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				return types.Double(float64(lhs.(types.Uint)) * float64(rhs.(types.Double)))
			})),
		// double receiver
		cel.MemberOverload("mul_double_double", []*cel.Type{cel.DoubleType, cel.DoubleType}, cel.DoubleType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				return lhs.(types.Double) * rhs.(types.Double)
			})),
		cel.MemberOverload("mul_double_int64", []*cel.Type{cel.DoubleType, cel.IntType}, cel.DoubleType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				return types.Double(float64(lhs.(types.Double)) * float64(rhs.(types.Int)))
			})),
		cel.MemberOverload("mul_double_uint64", []*cel.Type{cel.DoubleType, cel.UintType}, cel.DoubleType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				return types.Double(float64(lhs.(types.Double)) * float64(rhs.(types.Uint)))
			})),
	),

	// --- div ---
	cel.Function("div",
		// int receiver
		cel.MemberOverload("div_int64_int64", []*cel.Type{cel.IntType, cel.IntType}, cel.IntType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				return lhs.(types.Int) / rhs.(types.Int)
			})),
		cel.MemberOverload("div_int64_uint64", []*cel.Type{cel.IntType, cel.UintType}, cel.IntType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				return types.Int(int64(lhs.(types.Int)) / int64(rhs.(types.Uint)))
			})),
		cel.MemberOverload("div_int64_double", []*cel.Type{cel.IntType, cel.DoubleType}, cel.DoubleType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				return types.Double(float64(lhs.(types.Int)) / float64(rhs.(types.Double)))
			})),
		// uint receiver
		cel.MemberOverload("div_uint64_uint64", []*cel.Type{cel.UintType, cel.UintType}, cel.UintType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				return lhs.(types.Uint) / rhs.(types.Uint)
			})),
		cel.MemberOverload("div_uint64_int64", []*cel.Type{cel.UintType, cel.IntType}, cel.IntType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				return types.Int(int64(lhs.(types.Uint)) / int64(rhs.(types.Int)))
			})),
		cel.MemberOverload("div_uint64_double", []*cel.Type{cel.UintType, cel.DoubleType}, cel.DoubleType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				return types.Double(float64(lhs.(types.Uint)) / float64(rhs.(types.Double)))
			})),
		// double receiver
		cel.MemberOverload("div_double_double", []*cel.Type{cel.DoubleType, cel.DoubleType}, cel.DoubleType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				return lhs.(types.Double) / rhs.(types.Double)
			})),
		cel.MemberOverload("div_double_int64", []*cel.Type{cel.DoubleType, cel.IntType}, cel.DoubleType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				return types.Double(float64(lhs.(types.Double)) / float64(rhs.(types.Int)))
			})),
		cel.MemberOverload("div_double_uint64", []*cel.Type{cel.DoubleType, cel.UintType}, cel.DoubleType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				return types.Double(float64(lhs.(types.Double)) / float64(rhs.(types.Uint)))
			})),
	),

	// --- mod ---
	cel.Function("mod",
		// int receiver
		cel.MemberOverload("mod_int64_int64", []*cel.Type{cel.IntType, cel.IntType}, cel.IntType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				return lhs.(types.Int) % rhs.(types.Int)
			})),
		cel.MemberOverload("mod_int64_uint64", []*cel.Type{cel.IntType, cel.UintType}, cel.IntType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				return types.Int(int64(lhs.(types.Int)) % int64(rhs.(types.Uint)))
			})),
		// uint receiver
		cel.MemberOverload("mod_uint64_uint64", []*cel.Type{cel.UintType, cel.UintType}, cel.UintType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				return lhs.(types.Uint) % rhs.(types.Uint)
			})),
		cel.MemberOverload("mod_uint64_int64", []*cel.Type{cel.UintType, cel.IntType}, cel.IntType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
				return types.Int(int64(lhs.(types.Uint)) % int64(rhs.(types.Int)))
			})),
	),
}
