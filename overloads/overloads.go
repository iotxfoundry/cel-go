package overloads

// bitwisewise overloads
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

	// BitwiseNot 按位取反
	BitwiseNot = "bitwise_not"

	// BitwiseClear 按位清零
	BitwiseClear      = "bitwise_clear"
	BitwiseClearBytes = "bitwise_clear_bytes"

	// BitwiseIndex 按位取bit
	BitwiseIndex    = "bitwise_index"
	BitwiseIndexInt = "bitwise_index_int"

	// BitwiseShiftRight 按位右移
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
