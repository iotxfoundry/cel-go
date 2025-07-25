package functions

import (
	"encoding/binary"
	"math"

	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/common/types"
	"github.com/google/cel-go/common/types/ref"
	"github.com/iotxfoundry/cel-go/overloads"
)

var bytesFunctions = []cel.EnvOption{
	cel.Function(
		overloads.BytesToDouble,
		cel.FunctionDocs(
			"Converts the bytes to a double-precision floating-point number. "+
				"The conversion is based on the specified base, which can be 32 or 64. "+
				"Bytes are interpreted as big-endian, and the result is a double value.",
		),
		// bytes.tof(int) -> int
		cel.MemberOverload(
			overloads.BytesToDoubleInt,
			[]*cel.Type{cel.BytesType, cel.IntType},
			cel.DoubleType,
			cel.OverloadExamples(
				`b"\x40\x49\x0f\xdb".tof(32) // 3.141592653589793`,
				`b"\x40\x09\x21\xfb\x54\x44\x2d\x18".tof(64) // 3.141592653589793`,
			),
			cel.BinaryBinding(
				func(lhs, rhs ref.Val) ref.Val {
					buff, ok := lhs.(types.Bytes)
					if !ok {
						return types.ValOrErr(lhs, "no such overload")
					}
					base, ok := rhs.(types.Int)
					if !ok {
						return types.ValOrErr(rhs, "no such overload")
					}
					length := len(buff)
					if int(base)%8 != 0 {
						return types.NewErr("base '%d' out of float32, float64 size", base)
					}
					baseLen := int(base) / 8
					if length < baseLen {
						temp := make([]byte, baseLen-length)
						buff = append(temp, buff...)
					} else {
						buff = buff[:baseLen]
					}
					ret := types.Double(0)
					switch base {
					case 32:
						ret = types.Double(math.Float32frombits(binary.BigEndian.Uint32(buff)))
					case 64:
						ret = types.Double(math.Float64frombits(binary.BigEndian.Uint64(buff)))
					default:
						return types.NewErr("base '%d' out of float32, float64 size", base)
					}
					return ret
				},
			),
		),
	),
	cel.Function(
		overloads.BytesToUint,
		cel.FunctionDocs(
			"Converts the bytes to an unsigned integer. "+
				"The conversion is based on the specified base, which can be 8, 16, 32, or 64. "+
				"Bytes are interpreted as big-endian, and the result is an unsigned integer value.",
		),
		// bytes.toui(int) -> int
		cel.MemberOverload(
			overloads.BytesToUintInt,
			[]*cel.Type{cel.BytesType, cel.IntType},
			cel.UintType,
			cel.OverloadExamples(
				`b"\x00\x00\x00\x01".tou(8) // 1`,
				`b"\x00\x00\x01\x02".tou(16) // 258`,
				`b"\x00\x01\x02\x03".tou(32) // 16909060`,
				`b"\x01\x02\x03\x04".tou(64) // 72623859790382856`,
			),
			cel.BinaryBinding(
				func(lhs, rhs ref.Val) ref.Val {
					buff, ok := lhs.(types.Bytes)
					if !ok {
						return types.ValOrErr(lhs, "no such overload")
					}
					base, ok := rhs.(types.Int)
					if !ok {
						return types.ValOrErr(rhs, "no such overload")
					}
					length := len(buff)
					if int(base)%8 != 0 {
						return types.NewErr("base '%d' out of uint8, uint16, uint32, uint64 size", base)
					}
					baseLen := int(base) / 8
					if length < baseLen {
						temp := make([]byte, baseLen-length)
						buff = append(temp, buff...)
					} else {
						buff = buff[:baseLen]
					}
					ret := types.Uint(0)
					switch base {
					case 8:
						ret = types.Uint(uint8(buff[0]))
					case 16:
						ret = types.Uint(binary.BigEndian.Uint16(buff))
					case 32:
						ret = types.Uint(binary.BigEndian.Uint32(buff))
					case 64:
						ret = types.Uint(binary.BigEndian.Uint64(buff))
					default:
						return types.NewErr("base '%d' out of uint8, uint16, uint32, uint64 size", base)
					}
					return ret
				},
			),
		),
	),
	cel.Function(
		overloads.BytesToInt,
		cel.FunctionDocs(
			"Converts the bytes to a signed integer. "+
				"The conversion is based on the specified base, which can be 8, 16, 32, or 64. "+
				"Bytes are interpreted as big-endian, and the result is a signed integer value.",
		),
		// bytes.toi(int) -> bytes
		cel.MemberOverload(
			overloads.BytesToIntInt,
			[]*cel.Type{cel.BytesType, cel.IntType},
			cel.IntType,
			cel.OverloadExamples(
				`b"\x00\x00\x00\x01".toi(8) // 1`,
				`b"\x00\x00\x01\x02".toi(16) // 258`,
				`b"\x00\x01\x02\x03".toi(32) // 16909060`,
				`b"\x01\x02\x03\x04".toi(64) // 72623859790382856`,
			),
			cel.BinaryBinding(
				func(lhs, rhs ref.Val) ref.Val {
					buff, ok := lhs.(types.Bytes)
					if !ok {
						return types.ValOrErr(lhs, "no such overload")
					}
					base, ok := rhs.(types.Int)
					if !ok {
						return types.ValOrErr(rhs, "no such overload")
					}
					length := len(buff)
					if int(base)%8 != 0 {
						return types.NewErr("base '%d' out of int8, int16, int32, int64 size", base)
					}
					baseLen := int(base) / 8
					if length < baseLen {
						temp := make([]byte, baseLen-length)
						buff = append(temp, buff...)
					} else {
						buff = buff[:baseLen]
					}
					ret := types.Int(0)
					switch base {
					case 8:
						ret = types.Int(int8(buff[0]))
					case 16:
						ret = types.Int(int16(binary.BigEndian.Uint16(buff)))
					case 32:
						ret = types.Int(int32(binary.BigEndian.Uint32(buff)))
					case 64:
						ret = types.Int(int64(binary.BigEndian.Uint64(buff)))
					default:
						return types.NewErr("base '%d' out of int8, int16, int32, int64 size", base)
					}
					return ret
				},
			),
		),
	),
	cel.Function(
		overloads.BytesSlice,
		cel.FunctionDocs(
			"Returns a slice of the byte sequence from the specified start index to the end index. "+
				"The start index is inclusive, and the end index is exclusive. "+
				"If the end index is less than or equal to the start index, an empty byte sequence is returned. "+
				"If the start index is out of range, it is adjusted to the beginning of the byte sequence. "+
				"If the end index is out of range, it is adjusted to the end of the byte sequence.",
		),
		// bytes.slice(int, int) -> bytes
		cel.MemberOverload(
			overloads.BytesSliceIntInt,
			[]*cel.Type{cel.BytesType, cel.IntType, cel.IntType},
			cel.BytesType,
			cel.OverloadExamples(
				`b"\x01\x02\x03\x04".slice(1, 3) // b"\x02\x03"`,
				`b"\x01\x02\x03\x04".slice(2, 2) // b""`,
				`b"\x01\x02\x03\x04".slice(0, 5) // b"\x01\x02\x03\x04"`,
				`b"\x01\x02\x03\x04".slice(2, 1) // b""`,
			),
			cel.FunctionBinding(
				func(values ...ref.Val) ref.Val {
					if len(values) != 3 {
						return types.NewErr("values length not equal 3")
					}
					buff, ok := values[0].(types.Bytes)
					if !ok {
						return types.ValOrErr(values[0], "no such overload")
					}
					start, ok := values[1].(types.Int)
					if !ok {
						return types.ValOrErr(values[1], "no such overload")
					}
					if start < 0 {
						start = 0
					}
					if int(start) > len(buff) {
						return types.NewErr("index '%d' out of range in bytes size '%d'", start, len(buff))
					}
					end, ok := values[2].(types.Int)
					if !ok {
						return types.ValOrErr(values[2], "no such overload")
					}
					if end < 0 {
						return types.NewErr("index '%d' out of range in bytes size '%d'", end, len(buff))
					}

					if int(end) > len(buff) {
						end = types.Int(len(buff))
					}

					if end < start {
						start, end = end, start
					}
					buff = buff[start:end]
					return buff
				},
			),
		),
	),
	cel.Function(
		overloads.BytesDelete,
		cel.FunctionDocs(
			"Deletes a byte or a range of bytes from the byte sequence. "+
				"If the index is out of range, an error is returned. "+
				"If the start index is greater than the end index, the indices are swapped.",
		),
		// bytes.delete(int, int) -> bytes
		cel.MemberOverload(
			overloads.BytesDeleteIntInt,
			[]*cel.Type{cel.BytesType, cel.IntType, cel.IntType},
			cel.BytesType,
			cel.OverloadExamples(
				`b"\x01\x02\x03\x04".delete(1, 3) // b"\x01\x04"`,
				`b"\x01\x02\x03\x04".delete(2, 2) // b"\x01\x02\x03\x04"`,
				`b"\x01\x02\x03\x04".delete(0, 5) // b""`,
				`b"\x01\x02\x03\x04".delete(2, 1) // b"\x01\x02\x04"`,
			),
			cel.FunctionBinding(
				func(values ...ref.Val) ref.Val {
					if len(values) != 3 {
						return types.NewErr("values length not equal 3")
					}
					buff, ok := values[0].(types.Bytes)
					if !ok {
						return types.ValOrErr(values[0], "no such overload")
					}
					start, ok := values[1].(types.Int)
					if !ok {
						return types.ValOrErr(values[1], "no such overload")
					}
					if start < 0 || int(start) > len(buff) {
						return types.NewErr("index '%d' out of range in bytes size '%d'", start, len(buff))
					}
					end, ok := values[2].(types.Int)
					if !ok {
						return types.ValOrErr(values[2], "no such overload")
					}
					if end < 0 {
						return types.NewErr("index '%d' out of range in bytes size '%d'", end, len(buff))
					}

					if int(end) >= len(buff) {
						end = types.Int(len(buff))
					}

					if end < start {
						start, end = end, start
					}
					buff = append(buff[:start], buff[end:]...)
					return buff
				},
			),
		),
		// bytes.delete(int) -> bytes
		cel.MemberOverload(
			overloads.BytesDeleteInt,
			[]*cel.Type{cel.BytesType, cel.IntType},
			cel.BytesType,
			cel.OverloadExamples(
				`b"\x01\x02\x03\x04".delete(1) // b"\x01\x03\x04"`,
				`b"\x01\x02\x03\x04".delete(2) // b"\x01\x02\x04"`,
				`b"\x01\x02\x03\x04".delete(0) // b"\x02\x03\x04"`,
				`b"\x01\x02\x03\x04".delete(3) // b"\x01\x02\x03"`,
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
					if index < 0 {
						return types.NewErr("index '%d' out of range in bytes size '%d'", index, len(buff))
					}
					if int(index) >= len(buff) {
						return buff
					}
					buff = append(buff[:index], buff[index+1:]...)
					return buff
				},
			),
		),
	),
	cel.Function(
		overloads.BytesSwap,
		cel.FunctionDocs(
			"Swaps two bytes in the byte sequence at the specified indices. "+
				"If either index is out of range, an error is returned. "+
				"If the indices are equal, the byte remains unchanged.",
		),
		// bytes.swap(int, int) -> bytes
		cel.MemberOverload(
			overloads.BytesSwapIntInt,
			[]*cel.Type{cel.BytesType, cel.IntType, cel.IntType},
			cel.BytesType,
			cel.OverloadExamples(
				`b"\x01\x02\x03\x04".swap(1, 3) // b"\x01\x04\x03\x02"`,
				`b"\x01\x02\x03\x04".swap(2, 2) // b"\x01\x02\x03\x04"`,
				`b"\x01\x02\x03\x04".swap(0, 5) // error: index '5' out of range in bytes size '4'`,
				`b"\x01\x02\x03\x04".swap(2, 1) // b"\x01\x02\x03\x04"`,
			),
			cel.FunctionBinding(
				func(values ...ref.Val) ref.Val {
					if len(values) != 3 {
						return types.NewErr("values length not equal 3")
					}
					buff, ok := values[0].(types.Bytes)
					if !ok {
						return types.ValOrErr(values[0], "no such overload")
					}
					before, ok := values[1].(types.Int)
					if !ok {
						return types.ValOrErr(values[1], "no such overload")
					}
					if before < 0 || int(before) >= len(buff) {
						return types.NewErr("index '%d' out of range in bytes size '%d'", before, len(buff))
					}
					after, ok := values[2].(types.Int)
					if !ok {
						return types.ValOrErr(values[2], "no such overload")
					}
					if after < 0 || int(after) >= len(buff) {
						return types.NewErr("index '%d' out of range in bytes size '%d'", after, len(buff))
					}
					buff[before], buff[after] = buff[after], buff[before]
					return buff
				},
			),
		),
	),
	cel.Function(
		overloads.BytesIndex,
		cel.FunctionDocs(
			"Returns the byte at the specified index in the byte sequence. "+
				"The index is zero-based. "+
				"If the index is out of range, an error is returned.",
		),
		// bytes.index(int) -> bytes
		cel.MemberOverload(
			overloads.BytesIndexInt,
			[]*cel.Type{cel.BytesType, cel.IntType},
			cel.BytesType,
			cel.OverloadExamples(
				`b"\x01\x02\x03\x04".index(1) // b"\x02"`,
				`b"\x01\x02\x03\x04".index(2) // b"\x03"`,
				`b"\x01\x02\x03\x04".index(0) // b"\x01"`,
				`b"\x01\x02\x03\x04".index(3) // b"\x04"`,
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
					if index < 0 || int(index) >= len(buff) {
						return types.NewErr("index '%d' out of range in bytes size '%d'", index, len(buff))
					}
					return types.Bytes([]byte{
						buff[index],
					})
				},
			),
		),
	),
}
