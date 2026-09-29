package functions

import "cel.dev/cel-go/cel"

// Functions returns the full set of CEL environment options provided by this
// library: bytes utilities, bitwise operations for bytes/int/uint,
// cross-type arithmetic member functions, and random-number functions.
//
// Pass the result to cel.NewEnv via cel.Lib, or append individual options to
// an existing environment:
//
//	env, err := cel.NewEnv(cel.Lib(functions.Functions()))
func Functions() []cel.EnvOption {
	fns := make([]cel.EnvOption, 0, 9)
	fns = append(fns, bytesFunctions...)
	fns = append(fns, bitwiseFunctions...)
	fns = append(fns, bitwiseIntFunctions...)
	fns = append(fns, bitwiseUintFunctions...)
	fns = append(fns, intFunctions...)
	fns = append(fns, uintFunctions...)
	fns = append(fns, doubleFunctions...)
	fns = append(fns, arithmeticFunctions...)
	fns = append(fns, randFunctions...)
	return fns
}
