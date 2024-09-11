package cel

import (
	"encoding/base64"
	"fmt"
	"reflect"
	"time"

	structpb "github.com/golang/protobuf/ptypes/struct"
	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/common/types"
	"github.com/google/cel-go/common/types/ref"
	"github.com/iotxfoundry/cel-go/functions"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/anypb"
)

func ComputeLib() cel.EnvOption {
	return cel.Lib(computeLibrary{})
}

type computeLibrary struct{}

func (computeLibrary) CompileOptions() []cel.EnvOption {
	return functions.Functions()
}

func (computeLibrary) ProgramOptions() []cel.ProgramOption {
	return []cel.ProgramOption{}
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
				buff, err := protojson.Marshal(anyVal)
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
				buff, err := protojson.Marshal(anyVal)
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
			buff, err := protojson.Marshal(t)
			if err == nil {
				out = string(buff)
			}
		}
	case types.IteratorType.TypeName():
	}
	return
}
