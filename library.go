package cel

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"reflect"
	"time"

	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/common/types"
	"github.com/google/cel-go/common/types/ref"
	"github.com/iotxfoundry/cel-go/functions"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/structpb"
)

func ComputeLib() cel.EnvOption {
	return cel.Lib(computeLib{})
}

type computeLib struct{}

func (computeLib) LibraryName() string {
	return "cel.lib.ext.compute"
}

func (computeLib) CompileOptions() []cel.EnvOption {
	return functions.Functions()
}

func (computeLib) ProgramOptions() []cel.ProgramOption {
	return []cel.ProgramOption{}
}

func Val2Bytes(val ref.Val) (out []byte, err error) {
	buffer := &bytes.Buffer{}
	switch val.Type().TypeName() {
	case types.BytesType.TypeName():
		t, ok := val.Value().([]byte)
		if ok {
			_, err = buffer.Write(t)
		}
	case types.BoolType.TypeName():
		t, ok := val.Value().(bool)
		if ok {
			if t {
				_, err = buffer.Write([]byte{1})
			} else {
				_, err = buffer.Write([]byte{0})
			}
		}
	case types.DoubleType.TypeName():
		t, ok := val.Value().(float64)
		if ok {
			err = binary.Write(buffer, binary.BigEndian, t)
		}
	case types.IntType.TypeName():
		t, ok := val.Value().(int64)
		if ok {
			err = binary.Write(buffer, binary.BigEndian, t)
		}
	case types.UintType.TypeName():
		t, ok := val.Value().(uint64)
		if ok {
			err = binary.Write(buffer, binary.BigEndian, t)
		}
	case types.StringType.TypeName():
		t, ok := val.Value().(string)
		if ok {
			_, err = buffer.Write([]byte(t))
		}
	case types.NullType.TypeName():
		_, err = buffer.Write([]byte{})
	case types.DurationType.TypeName():
		t, ok := val.Value().(time.Duration)
		if ok {
			_, err = buffer.Write([]byte(t.String()))
		}
	case types.TimestampType.TypeName():
		t, ok := val.Value().(time.Time)
		if ok {
			_, err = buffer.Write([]byte(t.String()))
		}
	case types.ErrType.TypeName():
		t, ok := val.Value().(error)
		if ok {
			_, err = buffer.Write([]byte(t.Error()))
		}
	case types.ListType.TypeName():
		t, err := val.ConvertToNative(reflect.TypeOf(&structpb.Value{}))
		if err == nil {
			anyVal, ok := t.(*structpb.Value)
			if ok {
				buff, err := json.Marshal(anyVal)
				if err == nil {
					_, err = buffer.Write(buff)
					if err != nil {
						// pass
					}
				}
			}
		}
	case types.UnknownType.TypeName():
		_, err = buffer.Write([]byte{})
	case types.MapType.TypeName():
		t, err := val.ConvertToNative(reflect.TypeOf(&structpb.Value{}))
		if err == nil {
			anyVal, ok := t.(*structpb.Value)
			if ok {
				buff, err := json.Marshal(anyVal)
				if err == nil {
					_, err = buffer.Write(buff)
					if err != nil {
						// pass
					}
				}
			}
		}
	case types.TypeType.TypeName():
		t, ok := val.Value().(string)
		if ok {
			_, err = buffer.Write([]byte(t))
		}
	case "google.protobuf.Any":
		t, ok := val.Value().(*anypb.Any)
		if ok {
			buff, err := json.Marshal(t)
			if err == nil {
				_, err = buffer.Write(buff)
				if err != nil {
					// pass
				}
			}
		}
	case types.IteratorType.TypeName():
	}

	if err != nil {
		return
	}
	out = buffer.Bytes()
	return
}

func Val2String(val ref.Val) (out string, err error) {
	switch val.Type().TypeName() {
	case types.BytesType.TypeName():
		t, ok := val.Value().([]byte)
		if ok {
			out = base64.StdEncoding.EncodeToString(t)
		}
	case types.BoolType.TypeName():
		t, ok := val.Value().(bool)
		if ok {
			out = fmt.Sprintf("%t", t)
		}
	case types.DoubleType.TypeName():
		t, ok := val.Value().(float64)
		if ok {
			out = fmt.Sprintf("%.3f", t)
		}
	case types.IntType.TypeName():
		t, ok := val.Value().(int64)
		if ok {
			out = fmt.Sprintf("%d", t)
		}
	case types.UintType.TypeName():
		t, ok := val.Value().(uint64)
		if ok {
			out = fmt.Sprintf("%d", t)
		}
	case types.StringType.TypeName():
		t, ok := val.Value().(string)
		if ok {
			out = t
		}
	case types.NullType.TypeName():
		out = ""
	case types.DurationType.TypeName():
		t, ok := val.Value().(time.Duration)
		if ok {
			out = t.String()
		}
	case types.TimestampType.TypeName():
		t, ok := val.Value().(time.Time)
		if ok {
			out = t.String()
		}
	case types.ErrType.TypeName():
		t, ok := val.Value().(error)
		if ok {
			out = t.Error()
		}
	case types.ListType.TypeName():
		t, err := val.ConvertToNative(reflect.TypeOf(&structpb.Value{}))
		if err == nil {
			anyVal, ok := t.(*structpb.Value)
			if ok {
				buff, err := json.Marshal(anyVal)
				if err == nil {
					out = string(buff)
				}
			}
		}
	case types.UnknownType.TypeName():
		out = ""
	case types.MapType.TypeName():
		t, err := val.ConvertToNative(reflect.TypeOf(&structpb.Value{}))
		if err == nil {
			anyVal, ok := t.(*structpb.Value)
			if ok {
				buff, err := json.Marshal(anyVal)
				if err == nil {
					out = string(buff)
				}
			}
		}
	case types.TypeType.TypeName():
		t, ok := val.Value().(string)
		if ok {
			out = t
		}
	case "google.protobuf.Any":
		t, ok := val.Value().(*anypb.Any)
		if ok {
			buff, err := json.Marshal(t)
			if err == nil {
				out = string(buff)
			}
		}
	case types.IteratorType.TypeName():
	}
	return
}

func Val2Pb(val ref.Val) (out *structpb.Value, err error) {
	switch val.Type().TypeName() {
	case types.BytesType.TypeName():
		t, ok := val.Value().([]byte)
		if ok {
			out, err = structpb.NewValue(t)
		}
	case types.BoolType.TypeName():
		t, ok := val.Value().(bool)
		if ok {
			out, err = structpb.NewValue(t)
		}
	case types.DoubleType.TypeName():
		t, ok := val.Value().(float64)
		if ok {
			out, err = structpb.NewValue(t)
		}
	case types.IntType.TypeName():
		t, ok := val.Value().(int64)
		if ok {
			out, err = structpb.NewValue(t)
		}
	case types.UintType.TypeName():
		t, ok := val.Value().(uint64)
		if ok {
			out, err = structpb.NewValue(t)
		}
	case types.StringType.TypeName():
		t, ok := val.Value().(string)
		if ok {
			out, err = structpb.NewValue(t)
		}
	case types.NullType.TypeName():
		out = structpb.NewNullValue()
	case types.DurationType.TypeName():
		t, ok := val.Value().(time.Duration)
		if ok {
			out = structpb.NewStringValue(t.String())
		}
	case types.TimestampType.TypeName():
		t, ok := val.Value().(time.Time)
		if ok {
			out = structpb.NewStringValue(t.String())
		}
	case types.ErrType.TypeName():
		t, ok := val.Value().(error)
		if ok {
			out = structpb.NewStringValue(t.Error())
		}
	case types.ListType.TypeName():
		t, err := val.ConvertToNative(reflect.TypeOf(&structpb.Value{}))
		if err == nil {
			out, _ = t.(*structpb.Value)
		}
	case types.UnknownType.TypeName():
		out = structpb.NewNullValue()
	case types.MapType.TypeName():
		t, err := val.ConvertToNative(reflect.TypeOf(&structpb.Value{}))
		if err == nil {
			out, _ = t.(*structpb.Value)
		}
	case types.TypeType.TypeName():
		t, ok := val.Value().(string)
		if ok {
			out = structpb.NewStringValue(t)
		}
	case "google.protobuf.Any":
		t, ok := val.Value().(*anypb.Any)
		if ok {
			buff, err := json.Marshal(t)
			if err == nil {
				return structpb.NewValue(buff)
			}
		}
	case types.IteratorType.TypeName():
	}
	return
}
