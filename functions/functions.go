package functions

import "github.com/google/cel-go/cel"

func Functions() (fns []cel.EnvOption) {
	fns = append(fns, bytesFunctions...)
	fns = append(fns, bitwiseFunctions...)
	fns = append(fns, intFunctions...)
	fns = append(fns, uintFunctions...)
	fns = append(fns, doubleFunctions...)
	return fns
}
