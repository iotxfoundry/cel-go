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
	BitwiseIndex      = "bitwise_index"
	BitwiseIndexInt64 = "bitwise_index_int64"

	// BitwiseShiftRight 按位右移
	BitwiseShiftRight      = "bitwise_shr"
	BitwiseShiftRightInt64 = "bitwise_shr_int64"

	// BitwiseShiftLeft 按位左移
	BitwiseShiftLeft      = "bitwise_shl"
	BitwiseShiftLeftInt64 = "bitwise_shl_int64"
)

// Bytes overloads
const (
	BytesIndex      = "index"
	BytesIndexInt64 = "index_int64"

	BytesSwap           = "swap"
	BytesSwapInt64Int64 = "swap_int64_int64"

	BytesDelete           = "delete"
	BytesDeleteInt64      = "delete_int64"
	BytesDeleteInt64Int64 = "delete_int64_int64"

	BytesSlice           = "slice"
	BytesSliceInt64Int64 = "slice_int64_int64"

	BytesToInt      = "toi"
	BytesToIntInt64 = "toi_int64"

	BytesToUint      = "toui"
	BytesToUintInt64 = "toui_int64"

	BytesToDouble      = "tof"
	BytesToDoubleInt64 = "tof_int64"
)

const (
	IntToBytes      = "to_bytes"
	IntToBytesInt64 = "int64_to_bytes_int64"
)

const (
	UintToBytes      = "to_bytes"
	UintToBytesInt64 = "uint64_to_bytes_int64"
)

const (
	DoubleToBytes      = "to_bytes"
	DoubleToBytesInt64 = "float64_to_bytes_int64"
)
