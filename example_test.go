package cel_test

import (
	"encoding/json"
	"fmt"

	"github.com/google/cel-go/common/types"
	compute "github.com/iotxfoundry/cel-go"
)

func ExampleVal2String() {
	reg, err := types.NewRegistry()
	if err != nil {
		return
	}

	val := types.NewDynamicMap(reg, map[string]int64{"10000": 1})
	out, err := compute.Val2String(val)
	if err != nil {
		return
	}
	fmt.Println(out)
	// Output: {"10000":1}
}

func ExampleVal2Pb() {
	reg, err := types.NewRegistry()
	if err != nil {
		return
	}

	val := types.NewDynamicMap(reg, map[string]int64{"10000": 1})
	out, err := compute.Val2Pb(val)
	if err != nil {
		return
	}
	buff, err := json.Marshal(out)
	if err != nil {
		return
	}
	fmt.Println(string(buff))
	// Output: {"10000":1}
}

func ExampleVal2Bytes() {
	reg, err := types.NewRegistry()
	if err != nil {
		return
	}

	val := types.NewDynamicMap(reg, map[string]int64{"10000": 1})
	out, err := compute.Val2Bytes(val)
	if err != nil {
		return
	}
	fmt.Printf("% 02X\n", out)
	fmt.Println(string(out))
	// Output: 7B 22 31 30 30 30 30 22 3A 31 7D
	// {"10000":1}
}
