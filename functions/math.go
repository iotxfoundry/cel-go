package functions

import (
	"math/rand"

	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/common/ast"
	"github.com/google/cel-go/common/types"
	"github.com/google/cel-go/common/types/ref"
	"github.com/iotxfoundry/cel-go/overloads"
)

// randFunctions math rand functions
//
// Examples:
//
//	math.randf()      // [0, 1)
//	math.randf(32)      // [0, 1) float32
//	math.randf(64)      // [0, 1) float64
//
// Examples:
//
//	math.randi() // [0, MaxInt64]
//	math.randi(32) // [0, MaxInt32]
//	math.randi(64) // [0, MaxInt64]
//	math.randi(32, 10) // [0, 10)
//	math.randi(64, 10) // [0, 10)
//
// Examples:
//
//	math.randui() // [0, MaxUint64]
//	math.randui(32) // [0, MaxUint32]
//	math.randui(64) // [0, MaxUint64]
var randFunctions = []cel.EnvOption{
	cel.Macros(
		// math.randf()
		cel.ReceiverVarArgMacro(overloads.MathRandDoubleMicro, mathRandDouble),
		// math.randi()
		cel.ReceiverVarArgMacro(overloads.MathRandIntMicro, mathRandInt),
		// math.randui()
		cel.ReceiverVarArgMacro(overloads.MathRandUintMicro, mathRandUint),
	),
	cel.Function(
		overloads.MathRandDouble,
		// math.randf() -> double
		cel.Overload(
			overloads.MathRandDoubleNone,
			[]*cel.Type{},
			cel.DoubleType,
			cel.FunctionBinding(
				func(values ...ref.Val) ref.Val {
					if len(values) != 0 {
						return types.NewErr("no such overload")
					}
					return types.Double(rand.Float64())
				},
			),
		),
		// math.randf(int) -> double
		cel.Overload(
			overloads.MathRandDoubleInt,
			[]*cel.Type{cel.IntType},
			cel.DoubleType,
			cel.UnaryBinding(
				func(value ref.Val) ref.Val {
					base, ok := value.(types.Int)
					if !ok {
						return types.ValOrErr(value, "no such overload")
					}
					var ret ref.Val
					switch base {
					case 32:
						ret = types.Double(rand.Float32())
					case 64:
						ret = types.Double(rand.Float64())
					default:
						return types.NewErr("base '%d' out of float32, float64 size", base)
					}
					return ret
				},
			),
		),
		// math.randf(uint) -> double
		cel.Overload(
			overloads.MathRandDoubleUint,
			[]*cel.Type{cel.UintType},
			cel.DoubleType,
			cel.UnaryBinding(
				func(value ref.Val) ref.Val {
					base, ok := value.(types.Uint)
					if !ok {
						return types.ValOrErr(value, "no such overload")
					}
					var ret ref.Val
					switch base {
					case 32:
						ret = types.Double(rand.Float32())
					case 64:
						ret = types.Double(rand.Float64())
					default:
						return types.NewErr("base '%d' out of float32, float64 size", base)
					}
					return ret
				},
			),
		),
		// math.randf(double) -> double
		cel.Overload(
			overloads.MathRandDoubleDouble,
			[]*cel.Type{cel.DoubleType},
			cel.DoubleType,
			cel.UnaryBinding(
				func(value ref.Val) ref.Val {
					base, ok := value.(types.Double)
					if !ok {
						return types.ValOrErr(value, "no such overload")
					}
					var ret ref.Val
					switch int(base) {
					case 32:
						ret = types.Double(rand.Float32())
					case 64:
						ret = types.Double(rand.Float64())
					default:
						return types.NewErr("base '%d' out of float32, float64 size", int(base))
					}
					return ret
				},
			),
		),
	),
	cel.Function(
		overloads.MathRandInt,
		// math.randi() -> int
		cel.Overload(
			overloads.MathRandIntNone,
			[]*cel.Type{},
			cel.IntType,
			cel.FunctionBinding(
				func(values ...ref.Val) ref.Val {
					if len(values) != 0 {
						return types.NewErr("no such overload")
					}
					return types.Int(rand.Int63())
				},
			),
		),
		// math.randi(int) -> int
		cel.Overload(
			overloads.MathRandIntInt,
			[]*cel.Type{cel.IntType},
			cel.IntType,
			cel.UnaryBinding(
				func(value ref.Val) ref.Val {
					base, ok := value.(types.Int)
					if !ok {
						return types.ValOrErr(value, "no such overload")
					}
					var ret ref.Val
					switch base {
					case 32:
						ret = types.Int(rand.Int31())
					case 64:
						ret = types.Int(rand.Int63())
					default:
						return types.NewErr("base '%d' out of int32, int64 size", base)
					}
					return ret
				},
			),
		),
		// math.randi(int, int) -> int
		cel.Overload(
			overloads.MathRandIntIntInt,
			[]*cel.Type{cel.IntType, cel.IntType},
			cel.IntType,
			cel.BinaryBinding(
				func(lhs ref.Val, rhs ref.Val) ref.Val {
					base, ok := lhs.(types.Int)
					if !ok {
						return types.ValOrErr(lhs, "no such overload")
					}
					n, ok := rhs.(types.Int)
					if !ok {
						return types.ValOrErr(rhs, "no such overload")
					}
					var ret ref.Val
					switch base {
					case 32:
						ret = types.Int(rand.Int31n(int32(n)))
					case 64:
						ret = types.Int(rand.Int63n(int64(n)))
					default:
						return types.NewErr("base '%d' out of int32, int64 size", base)
					}
					return ret
				},
			),
		),
		// math.randi(int, uint) -> int
		cel.Overload(
			overloads.MathRandIntIntUint,
			[]*cel.Type{cel.IntType, cel.UintType},
			cel.IntType,
			cel.BinaryBinding(
				func(lhs ref.Val, rhs ref.Val) ref.Val {
					base, ok := lhs.(types.Int)
					if !ok {
						return types.ValOrErr(lhs, "no such overload")
					}
					n, ok := rhs.(types.Uint)
					if !ok {
						return types.ValOrErr(rhs, "no such overload")
					}
					var ret ref.Val
					switch base {
					case 32:
						ret = types.Int(rand.Int31n(int32(n)))
					case 64:
						ret = types.Int(rand.Int63n(int64(n)))
					default:
						return types.NewErr("base '%d' out of int32, int64 size", base)
					}
					return ret
				},
			),
		),
		// math.randi(int, double) -> int
		cel.Overload(
			overloads.MathRandIntIntDouble,
			[]*cel.Type{cel.IntType, cel.DoubleType},
			cel.IntType,
			cel.BinaryBinding(
				func(lhs ref.Val, rhs ref.Val) ref.Val {
					base, ok := lhs.(types.Int)
					if !ok {
						return types.ValOrErr(lhs, "no such overload")
					}
					n, ok := rhs.(types.Double)
					if !ok {
						return types.ValOrErr(rhs, "no such overload")
					}
					var ret ref.Val
					switch base {
					case 32:
						ret = types.Int(rand.Int31n(int32(n)))
					case 64:
						ret = types.Int(rand.Int63n(int64(n)))
					default:
						return types.NewErr("base '%d' out of int32, int64 size", base)
					}
					return ret
				},
			),
		),
		// math.randi(uint) -> int
		cel.Overload(
			overloads.MathRandIntUint,
			[]*cel.Type{cel.UintType},
			cel.IntType,
			cel.UnaryBinding(
				func(value ref.Val) ref.Val {
					base, ok := value.(types.Uint)
					if !ok {
						return types.ValOrErr(value, "no such overload")
					}
					var ret ref.Val
					switch base {
					case 32:
						ret = types.Int(rand.Int31())
					case 64:
						ret = types.Int(rand.Int63())
					default:
						return types.NewErr("base '%d' out of int32, int64 size", base)
					}
					return ret
				},
			),
		),

		// math.randi(uint, int) -> int
		cel.Overload(
			overloads.MathRandIntUintInt,
			[]*cel.Type{cel.UintType, cel.IntType},
			cel.IntType,
			cel.BinaryBinding(
				func(lhs ref.Val, rhs ref.Val) ref.Val {
					base, ok := lhs.(types.Uint)
					if !ok {
						return types.ValOrErr(lhs, "no such overload")
					}
					n, ok := rhs.(types.Int)
					if !ok {
						return types.ValOrErr(rhs, "no such overload")
					}
					var ret ref.Val
					switch base {
					case 32:
						ret = types.Int(rand.Int31n(int32(n)))
					case 64:
						ret = types.Int(rand.Int63n(int64(n)))
					default:
						return types.NewErr("base '%d' out of int32, int64 size", base)
					}
					return ret
				},
			),
		),
		// math.randi(uint, uint) -> int
		cel.Overload(
			overloads.MathRandIntUintUint,
			[]*cel.Type{cel.UintType, cel.UintType},
			cel.IntType,
			cel.BinaryBinding(
				func(lhs ref.Val, rhs ref.Val) ref.Val {
					base, ok := lhs.(types.Uint)
					if !ok {
						return types.ValOrErr(lhs, "no such overload")
					}
					n, ok := rhs.(types.Uint)
					if !ok {
						return types.ValOrErr(rhs, "no such overload")
					}
					var ret ref.Val
					switch base {
					case 32:
						ret = types.Int(rand.Int31n(int32(n)))
					case 64:
						ret = types.Int(rand.Int63n(int64(n)))
					default:
						return types.NewErr("base '%d' out of int32, int64 size", base)
					}
					return ret
				},
			),
		),
		// math.randi(uint, double) -> int
		cel.Overload(
			overloads.MathRandIntUintDouble,
			[]*cel.Type{cel.UintType, cel.DoubleType},
			cel.IntType,
			cel.BinaryBinding(
				func(lhs ref.Val, rhs ref.Val) ref.Val {
					base, ok := lhs.(types.Uint)
					if !ok {
						return types.ValOrErr(lhs, "no such overload")
					}
					n, ok := rhs.(types.Double)
					if !ok {
						return types.ValOrErr(rhs, "no such overload")
					}
					var ret ref.Val
					switch base {
					case 32:
						ret = types.Int(rand.Int31n(int32(n)))
					case 64:
						ret = types.Int(rand.Int63n(int64(n)))
					default:
						return types.NewErr("base '%d' out of int32, int64 size", base)
					}
					return ret
				},
			),
		),

		// math.randi(double) -> int
		cel.Overload(
			overloads.MathRandIntDouble,
			[]*cel.Type{cel.DoubleType},
			cel.IntType,
			cel.UnaryBinding(
				func(value ref.Val) ref.Val {
					base, ok := value.(types.Double)
					if !ok {
						return types.ValOrErr(value, "no such overload")
					}
					var ret ref.Val
					switch int(base) {
					case 32:
						ret = types.Int(rand.Int31())
					case 64:
						ret = types.Int(rand.Int63())
					default:
						return types.NewErr("base '%d' out of int32, int64 size", int(base))
					}
					return ret
				},
			),
		),

		// math.randi(double, int) -> int
		cel.Overload(
			overloads.MathRandIntDoubleInt,
			[]*cel.Type{cel.DoubleType, cel.IntType},
			cel.IntType,
			cel.BinaryBinding(
				func(lhs ref.Val, rhs ref.Val) ref.Val {
					base, ok := lhs.(types.Double)
					if !ok {
						return types.ValOrErr(lhs, "no such overload")
					}
					n, ok := rhs.(types.Int)
					if !ok {
						return types.ValOrErr(rhs, "no such overload")
					}
					var ret ref.Val
					switch base {
					case 32:
						ret = types.Int(rand.Int31n(int32(n)))
					case 64:
						ret = types.Int(rand.Int63n(int64(n)))
					default:
						return types.NewErr("base '%d' out of int32, int64 size", int(base))
					}
					return ret
				},
			),
		),
		// math.randi(double, uint) -> int
		cel.Overload(
			overloads.MathRandIntDoubleUint,
			[]*cel.Type{cel.DoubleType, cel.UintType},
			cel.IntType,
			cel.BinaryBinding(
				func(lhs ref.Val, rhs ref.Val) ref.Val {
					base, ok := lhs.(types.Double)
					if !ok {
						return types.ValOrErr(lhs, "no such overload")
					}
					n, ok := rhs.(types.Uint)
					if !ok {
						return types.ValOrErr(rhs, "no such overload")
					}
					var ret ref.Val
					switch base {
					case 32:
						ret = types.Int(rand.Int31n(int32(n)))
					case 64:
						ret = types.Int(rand.Int63n(int64(n)))
					default:
						return types.NewErr("base '%d' out of int32, int64 size", int(base))
					}
					return ret
				},
			),
		),
		// math.randi(double, double) -> int
		cel.Overload(
			overloads.MathRandIntDoubleDouble,
			[]*cel.Type{cel.DoubleType, cel.DoubleType},
			cel.IntType,
			cel.BinaryBinding(
				func(lhs ref.Val, rhs ref.Val) ref.Val {
					base, ok := lhs.(types.Double)
					if !ok {
						return types.ValOrErr(lhs, "no such overload")
					}
					n, ok := rhs.(types.Double)
					if !ok {
						return types.ValOrErr(rhs, "no such overload")
					}
					var ret ref.Val
					switch base {
					case 32:
						ret = types.Int(rand.Int31n(int32(n)))
					case 64:
						ret = types.Int(rand.Int63n(int64(n)))
					default:
						return types.NewErr("base '%d' out of int32, int64 size", int(base))
					}
					return ret
				},
			),
		),
	),
	cel.Function(
		overloads.MathRandUint,
		// math.randui() -> uint
		cel.Overload(
			overloads.MathRandUintNone,
			[]*cel.Type{},
			cel.UintType,
			cel.FunctionBinding(
				func(values ...ref.Val) ref.Val {
					if len(values) != 0 {
						return types.NewErr("no such overload")
					}
					return types.Uint(rand.Uint64())
				},
			),
		),
		// math.randui(int) -> uint
		cel.Overload(
			overloads.MathRandUintInt,
			[]*cel.Type{cel.IntType},
			cel.UintType,
			cel.UnaryBinding(
				func(value ref.Val) ref.Val {
					base, ok := value.(types.Int)
					if !ok {
						return types.ValOrErr(value, "no such overload")
					}
					var ret ref.Val
					switch base {
					case 32:
						ret = types.Uint(rand.Uint32())
					case 64:
						ret = types.Uint(rand.Uint64())
					default:
						return types.NewErr("base '%d' out of uint32, uint64 size", base)
					}
					return ret
				},
			),
		),
		// math.randui(uint) -> uint
		cel.Overload(
			overloads.MathRandUintUint,
			[]*cel.Type{cel.UintType},
			cel.UintType,
			cel.UnaryBinding(
				func(value ref.Val) ref.Val {
					base, ok := value.(types.Uint)
					if !ok {
						return types.ValOrErr(value, "no such overload")
					}
					var ret ref.Val
					switch base {
					case 32:
						ret = types.Uint(rand.Uint32())
					case 64:
						ret = types.Uint(rand.Uint64())
					default:
						return types.NewErr("base '%d' out of uint32, uint64 size", base)
					}
					return ret
				},
			),
		),

		// math.randui(double) -> uint
		cel.Overload(
			overloads.MathRandUintDouble,
			[]*cel.Type{cel.DoubleType},
			cel.UintType,
			cel.UnaryBinding(
				func(value ref.Val) ref.Val {
					base, ok := value.(types.Double)
					if !ok {
						return types.ValOrErr(value, "no such overload")
					}
					var ret ref.Val
					switch int(base) {
					case 32:
						ret = types.Uint(rand.Uint32())
					case 64:
						ret = types.Uint(rand.Uint64())
					default:
						return types.NewErr("base '%d' out of uint32, uint64 size", int(base))
					}
					return ret
				},
			),
		),
	),
}

func mathRandUint(meh cel.MacroExprFactory, target ast.Expr, args []ast.Expr) (ast.Expr, *cel.Error) {
	if !macroTargetMatchesNamespace(overloads.MathNamespace, target) {
		return nil, nil
	}
	switch len(args) {
	case 0:
		return meh.NewCall(overloads.MathRandUint), nil
	case 1:
		if isValidArgType(args[0]) {
			return meh.NewCall(overloads.MathRandUint, args[0]), nil
		}
		return nil, meh.NewError(args[0].ID(), "math.randi() invalid single argument value")
	default:
		return nil, meh.NewError(target.ID(), "math.randi() requires at zero or one argument")
	}
}

func mathRandInt(meh cel.MacroExprFactory, target ast.Expr, args []ast.Expr) (ast.Expr, *cel.Error) {
	if !macroTargetMatchesNamespace(overloads.MathNamespace, target) {
		return nil, nil
	}
	switch len(args) {
	case 0:
		return meh.NewCall(overloads.MathRandInt), nil
	case 1:
		if isValidArgType(args[0]) {
			return meh.NewCall(overloads.MathRandInt, args[0]), nil
		}
		return nil, meh.NewError(args[0].ID(), "math.randi() invalid single argument value")
	case 2:
		if !isValidArgType(args[0]) {
			return nil, meh.NewError(args[0].ID(), "math.randi() invalid argument value")
		}
		if !isValidArgType(args[1]) {
			return nil, meh.NewError(args[1].ID(), "math.randi() invalid argument value")
		}
		return meh.NewCall(overloads.MathRandInt, args[0], args[1]), nil
	default:
		return nil, meh.NewError(target.ID(), "math.randi() requires at zero or one argument")
	}
}

func mathRandDouble(meh cel.MacroExprFactory, target ast.Expr, args []ast.Expr) (ast.Expr, *cel.Error) {
	if !macroTargetMatchesNamespace(overloads.MathNamespace, target) {
		return nil, nil
	}
	switch len(args) {
	case 0:
		return meh.NewCall(overloads.MathRandDouble), nil
	case 1:
		if isValidArgType(args[0]) {
			return meh.NewCall(overloads.MathRandDouble, args[0]), nil
		}
		return nil, meh.NewError(args[0].ID(), "math.randf() invalid single argument value")
	default:
		return nil, meh.NewError(target.ID(), "math.randf() requires at zero or one argument")
	}
}

func macroTargetMatchesNamespace(ns string, target ast.Expr) bool {
	switch target.Kind() {
	case ast.IdentKind:
		if target.AsIdent() != ns {
			return false
		}
		return true
	default:
		return false
	}
}

func isValidArgType(arg ast.Expr) bool {
	switch arg.Kind() {
	case ast.LiteralKind:
		c := ref.Val(arg.AsLiteral())
		switch c.(type) {
		case types.Double, types.Int, types.Uint:
			return true
		default:
			return false
		}
	case ast.ListKind, ast.MapKind, ast.StructKind:
		return false
	default:
		return true
	}
}
