// Package cel provides an extended CEL (Common Expression Language)
// environment: cross-type arithmetic member functions, bitwise operations
// for bytes/int/uint, bytes utilities, and random-number functions. Build an
// environment with cel.NewEnv(cel.Lib(ComputeLib())).
//
// The expr.md reference document is generated from the library
// declarations; regenerate it after changing any function definitions:
//
//	//go:generate go run ./cmd/doc
//
//go:generate go run ./cmd/doc
package cel
