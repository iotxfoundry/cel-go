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
	cel.Function(overloads.BytesToDouble,
		// bytes.tof(int) -> int
		cel.MemberOverload(overloads.BytesToDoubleInt64,
			[]*cel.Type{cel.BytesType, cel.IntType},
			cel.DoubleType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
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
			}),
		),
	),
	cel.Function(overloads.BytesToUint,
		// bytes.toui(int) -> int
		cel.MemberOverload(overloads.BytesToUintInt64,
			[]*cel.Type{cel.BytesType, cel.IntType},
			cel.UintType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
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
			}),
		),
	),
	cel.Function(overloads.BytesToInt,
		// bytes.toi(int) -> bytes
		cel.MemberOverload(overloads.BytesToIntInt64,
			[]*cel.Type{cel.BytesType, cel.IntType},
			cel.IntType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
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
			}),
		),
	),
	cel.Function(overloads.BytesSlice,
		// bytes.slice(int, int) -> bytes
		cel.MemberOverload(overloads.BytesSliceInt64Int64,
			[]*cel.Type{cel.BytesType, cel.IntType, cel.IntType},
			cel.BytesType,
			cel.FunctionBinding(func(values ...ref.Val) ref.Val {
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
			}),
		),
	),
	cel.Function(overloads.BytesDelete,
		// bytes.delete(int, int) -> bytes
		cel.MemberOverload(overloads.BytesDeleteInt64Int64,
			[]*cel.Type{cel.BytesType, cel.IntType, cel.IntType},
			cel.BytesType,
			cel.FunctionBinding(func(values ...ref.Val) ref.Val {
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
			}),
		),
		// bytes.delete(int) -> bytes
		cel.MemberOverload(overloads.BytesDeleteInt64,
			[]*cel.Type{cel.BytesType, cel.IntType},
			cel.BytesType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
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
			}),
		),
	),
	cel.Function(overloads.BytesSwap,
		// bytes.swap(int, int) -> bytes
		cel.MemberOverload(overloads.BytesSwapInt64Int64,
			[]*cel.Type{cel.BytesType, cel.IntType, cel.IntType},
			cel.BytesType,
			cel.FunctionBinding(func(values ...ref.Val) ref.Val {
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
			}),
		),
	),
	cel.Function(overloads.BytesIndex,
		// bytes.index(int) -> bytes
		cel.MemberOverload(overloads.BytesIndexInt64,
			[]*cel.Type{cel.BytesType, cel.IntType},
			cel.BytesType,
			cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
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
			}),
		),
	),
}
