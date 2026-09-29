package overloads

// bitwise overloads
const (
	// BitwiseAnd 按位与
	BitwiseAnd      = "bitwise_and"
	BitwiseAndBytes = "bitwise_and_bytes"

	// BitwiseOr 按位或
	BitwiseOr      = "bitwise_or"
	BitwiseOrBytes = "bitwise_or_bytes"

	// BitwiseXor 按位异或
	BitwiseXor      = "bitwise_xor"
	BitwiseXorBytes = "bitwise_xor_bytes"

	// BitwiseClear 按位清零
	BitwiseClear      = "bitwise_clear"
	BitwiseClearBytes = "bitwise_clear_bytes"

	// BitwiseIndex 按位取bit
	BitwiseIndex    = "bitwise_index"
	BitwiseIndexInt = "bitwise_index_int"

	// BitwiseShiftRight 按位右移。ID 中的 int 指移位量参数类型；Go 常量名
	// 沿用 Int64 后缀，与 ID 字符串 "bitwise_shr_int" 并不完全一致。
	BitwiseShiftRight      = "bitwise_shr"
	BitwiseShiftRightInt64 = "bitwise_shr_int"

	// BitwiseShiftLeft 按位左移
	BitwiseShiftLeft      = "bitwise_shl"
	BitwiseShiftLeftInt64 = "bitwise_shl_int"
)

// Bytes overloads
const (
	BytesIndex    = "index"
	BytesIndexInt = "index_int"

	BytesSwap       = "swap"
	BytesSwapIntInt = "swap_int_int"

	BytesDelete       = "delete"
	BytesDeleteInt    = "delete_int"
	BytesDeleteIntInt = "delete_int_int"

	BytesSlice       = "slice"
	BytesSliceIntInt = "slice_int_int"

	BytesToInt    = "toi"
	BytesToIntInt = "toi_int"

	BytesToUint    = "toui"
	BytesToUintInt = "toui_int"

	BytesToDouble    = "tof"
	BytesToDoubleInt = "tof_int"
)

const (
	IntToBytes    = "to_bytes"
	IntToBytesInt = "int_to_bytes_int"
)

const (
	UintToBytes    = "to_bytes"
	UintToBytesInt = "uint_to_bytes_int"
)

const (
	DoubleToBytes    = "to_bytes"
	DoubleToBytesInt = "float64_to_bytes_int"
)

const (
	MathNamespace       = "math"
	MathRandDoubleMicro = "randf"
	MathRandIntMicro    = "randi"
	MathRandUintMicro   = "randui"

	MathRandDouble       = "math_randf"
	MathRandDoubleNone   = "math_randf_none"
	MathRandDoubleInt    = "math_randf_int"
	MathRandDoubleUint   = "math_randf_uint"
	MathRandDoubleDouble = "math_randf_double"

	MathRandInt     = "math_randi"
	MathRandIntNone = "math_randi_none"

	MathRandIntInt       = "math_randi_int"
	MathRandIntIntInt    = "math_randi_int_int"
	MathRandIntIntUint   = "math_randi_int_uint"
	MathRandIntIntDouble = "math_randi_int_double"

	MathRandIntUint       = "math_randi_uint"
	MathRandIntUintInt    = "math_randi_uint_int"
	MathRandIntUintUint   = "math_randi_uint_uint"
	MathRandIntUintDouble = "math_randi_uint_double"

	MathRandIntDouble       = "math_randi_double"
	MathRandIntDoubleInt    = "math_randi_double_int"
	MathRandIntDoubleUint   = "math_randi_double_uint"
	MathRandIntDoubleDouble = "math_randi_double_double"

	MathRandUint       = "math_randui"
	MathRandUintNone   = "math_randui_none"
	MathRandUintInt    = "math_randui_int"
	MathRandUintUint   = "math_randui_uint"
	MathRandUintDouble = "math_randui_double"
)

// int bitwise overloads
const (
	// IntBitwiseAnd int 按位与 — function name is "bitwise_and" (same as bytes),
	// overload ID is unique
	IntBitwiseAndIntInt = "int_bitwise_and_int_int"

	// IntBitwiseOr int 按位或
	IntBitwiseOrIntInt = "int_bitwise_or_int_int"

	// IntBitwiseXor int 按位异或
	IntBitwiseXorIntInt = "int_bitwise_xor_int_int"

	// IntBitwiseClear int 按位清零
	IntBitwiseClearIntInt = "int_bitwise_clear_int_int"

	// IntBitwiseShiftRight int 按位右移
	IntBitwiseShiftRightIntInt = "int_bitwise_shr_int_int"

	// IntBitwiseShiftLeft int 按位左移
	IntBitwiseShiftLeftIntInt = "int_bitwise_shl_int_int"

	// IntBitwiseNot int 按位取反 — function name is "bitwise_not" (same as bytes)
	IntBitwiseNotInt = "int_bitwise_not_int"

	// IntBitwiseIndex int 按位取bit — function name is "bitwise_index" (same as bytes)
	IntBitwiseIndexIntInt = "int_bitwise_index_int_int"
)

// uint bitwise overloads
const (
	// UintBitwiseAnd uint 按位与 — function name is "bitwise_and" (same as bytes)
	UintBitwiseAndUintUint = "uint_bitwise_and_uint_uint"

	// UintBitwiseOr uint 按位或
	UintBitwiseOrUintUint = "uint_bitwise_or_uint_uint"

	// UintBitwiseXor uint 按位异或
	UintBitwiseXorUintUint = "uint_bitwise_xor_uint_uint"

	// UintBitwiseClear uint 按位清零
	UintBitwiseClearUintUint = "uint_bitwise_clear_uint_uint"

	// UintBitwiseShiftRight uint 按位右移
	UintBitwiseShiftRightUintInt = "uint_bitwise_shr_uint_int"

	// UintBitwiseShiftLeft uint 按位左移
	UintBitwiseShiftLeftUintInt = "uint_bitwise_shl_uint_int"

	// UintBitwiseNot uint 按位取反 — function name is "bitwise_not" (same as bytes)
	UintBitwiseNotUint = "uint_bitwise_not_uint"

	// UintBitwiseIndex uint 按位取bit — function name is "bitwise_index" (same as bytes)
	UintBitwiseIndexUintInt = "uint_bitwise_index_uint_int"
)

// bytes bitwise unary and utility overloads
const (
	// BitwiseNot bytes 按位取反
	BitwiseNot      = "bitwise_not"
	BitwiseNotBytes = "bitwise_not_bytes"

	// BitwisePopcnt bytes 统计置1位数
	BitwisePopcnt      = "bitwise_popcnt"
	BitwisePopcntBytes = "bitwise_popcnt_bytes"
)
