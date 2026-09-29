# cel-go

Extended computation library for [CEL](https://cel.dev) built on [cel-go](https://cel.dev/cel-go) (formerly [github.com/google/cel-go](https://github.com/google/cel-go)). Provides additional member functions for bytes, integer, and unsigned integer types, plus math random utilities and value conversion helpers.

## Documentation

The full function and macro reference is generated into [expr.md](expr.md), split into the CEL standard library (with official `ext.*` packages) and this library's extensions. After changing any function definition, regenerate it with:

```bash
go generate ./...
```

`go test ./...` verifies the document is up to date.

## Installation

```bash
go get github.com/iotxfoundry/cel-go
```

## Quick Start

```go
import (
    "cel.dev/cel-go/cel"
    compute "github.com/iotxfoundry/cel-go"
)

env, _ := cel.NewEnv(
    cel.Variable("buff", cel.BytesType),
    compute.ComputeLib(),
)

ast, _ := env.Compile(`buff.bitwise_and(b"\x01").bitwise_or(b"\xff")`)
prg, _ := env.Program(ast)
out, _, _ := prg.Eval(map[string]any{"buff": []byte{0x01, 0x02, 0x03}})
```

## Features

### Bytes Member Functions

#### Byte-level Operations

| Function | Signature | Description |
|----------|-----------|-------------|
| `index(int)` | `bytes → bytes` | Returns the byte at the given index |
| `slice(int, int)` | `bytes → bytes` | Returns a slice from start (inclusive) to end (exclusive) |
| `delete(int)` | `bytes → bytes` | Deletes the byte at the given index |
| `delete(int, int)` | `bytes → bytes` | Deletes a range from start to end |
| `swap(int, int)` | `bytes → bytes` | Swaps bytes at the two given indices |

#### Bitwise Operations on Bytes

| Function | Signature | Description |
|----------|-----------|-------------|
| `bitwise_and(bytes)` | `bytes → bytes` | Bitwise AND between two byte sequences |
| `bitwise_or(bytes)` | `bytes → bytes` | Bitwise OR between two byte sequences |
| `bitwise_xor(bytes)` | `bytes → bytes` | Bitwise XOR between two byte sequences |
| `bitwise_clear(bytes)` | `bytes → bytes` | Bitwise AND-NOT (clears bits where mask is 1) |
| `bitwise_shr(int)` | `bytes → bytes` | Bitwise right shift as unsigned big-endian integer (negative values shift left); any magnitude, zeros shifted in |
| `bitwise_shl(int)` | `bytes → bytes` | Bitwise left shift as unsigned big-endian integer (negative values shift right); any magnitude, zeros shifted in |
| `bitwise_not()` | `bytes → bytes` | Bitwise NOT (ones-complement) on each byte |
| `bitwise_index(int)` | `bytes → bytes` | Returns a single-bit byte at the given bit index |
| `bitwise_popcnt()` | `bytes → int` | Population count (number of bits set to 1) |

#### Bytes Conversion

| Function | Signature | Description |
|----------|-----------|-------------|
| `toi(int)` | `bytes → int` | Converts bytes to signed integer (base: 8/16/32/64) |
| `toui(int)` | `bytes → uint` | Converts bytes to unsigned integer (base: 8/16/32/64) |
| `tof(int)` | `bytes → double` | Converts bytes to floating-point (base: 32/64) |

### Integer Member Functions

| Function | Signature | Description |
|----------|-----------|-------------|
| `bitwise_and(int)` | `int → int` | Bitwise AND |
| `bitwise_or(int)` | `int → int` | Bitwise OR |
| `bitwise_xor(int)` | `int → int` | Bitwise XOR |
| `bitwise_clear(int)` | `int → int` | Bitwise AND-NOT |
| `bitwise_shr(int)` | `int → int` | Bitwise right shift (negative values shift left) |
| `bitwise_shl(int)` | `int → int` | Bitwise left shift (negative values shift right) |
| `bitwise_not()` | `int → int` | Bitwise NOT (ones-complement) |
| `bitwise_index(int)` | `int → bytes` | Returns a single-bit byte at the given bit index |
| `to_bytes(int)` | `int → bytes` | Converts int to bytes (base: 8/16/32/64) |

### Unsigned Integer Member Functions

| Function | Signature | Description |
|----------|-----------|-------------|
| `bitwise_and(uint)` | `uint → uint` | Bitwise AND |
| `bitwise_or(uint)` | `uint → uint` | Bitwise OR |
| `bitwise_xor(uint)` | `uint → uint` | Bitwise XOR |
| `bitwise_clear(uint)` | `uint → uint` | Bitwise AND-NOT |
| `bitwise_shr(int)` | `uint → uint` | Bitwise right shift (negative values shift left) |
| `bitwise_shl(int)` | `uint → uint` | Bitwise left shift (negative values shift right) |
| `bitwise_not()` | `uint → uint` | Bitwise NOT (ones-complement) |
| `bitwise_index(int)` | `uint → bytes` | Returns a single-bit byte at the given bit index |
| `to_bytes(int)` | `uint → bytes` | Converts uint to bytes (base: 8/16/32/64) |

### Double Member Functions

| Function | Signature | Description |
|----------|-----------|-------------|
| `to_bytes(int)` | `double → bytes` | Converts double to bytes (base: 32/64) |

### Cross-Type Arithmetic Member Functions

Since CEL's built-in operators (`+`, `-`, `*`, `/`, `%`) use trait-based singleton dispatch and cannot accept additional overloads, member functions are provided for cross-type arithmetic:

| Function | Description |
|----------|-------------|
| `add(int)` | Addition (int receiver) |
| `add(uint)` | Returns int |
| `add(double)` | Returns double |
| `sub(int)` | Subtraction (int receiver) |
| `sub(uint)` | Returns int |
| `sub(double)` | Returns double |
| `mul(int)` | Multiplication (int receiver) |
| `mul(uint)` | Returns int |
| `mul(double)` | Returns double |
| `div(int)` | Division (int receiver) |
| `div(uint)` | Returns int |
| `div(double)` | Returns double |
| `mod(int)` | Modulo (int receiver; no double support, like CEL's `%`) |
| `mod(uint)` | Returns int |

The same functions are available on `uint` and `double` receivers. Type promotion rules:

| Mix | Returns | Reason |
|-----|---------|--------|
| int + uint | int | signed type for possible negative results |
| any + double | double | double has widest range |

Error semantics mirror CEL's built-in operators: integer overflow returns an
`integer overflow` error instead of wrapping, division or modulo by zero
returns `division by zero` / `modulus by zero`, and `MinInt64` combined with
`-1` in div/mod/mul returns `integer overflow`. Double division follows IEEE
754 (e.g. `10.div(0.0)` is `+Inf`), matching CEL's `/` on doubles.

```cel
3.add(4u)      // 7 (int)
3u.add(-4)     // -1 (int)
3.add(4.5)     // 7.5 (double)
4.5.add(3u)    // 7.5 (double)

10.sub(4u)     // 6 (int)
3u.sub(10)     // -7 (int)

3.mul(4u)      // 12 (int)
-3.mul(4u)     // -12 (int)

10.div(4.0)    // 2.5 (double)
10u.div(4)     // 2 (int)

10.mod(4u)     // 2 (int)
-10.mod(4u)    // -2 (int)

9223372036854775807.add(1)  // error: integer overflow
10.div(0)                   // error: division by zero
```

### Math Random Functions

| Function | Signature | Description |
|----------|-----------|-------------|
| `math.randf()` | `→ double` | Random float64 in [0, 1) |
| `math.randf(int/uint/double)` | `→ double` | Random float32 (32) or float64 (64) |
| `math.randi()` | `→ int` | Random int64 in [0, MaxInt64] |
| `math.randi(int/uint/double)` | `→ int` | Random int32 (32) or int64 (64) |
| `math.randi(int/uint/double, int/uint/double)` | `→ int` | Random in [0, n) with given bit size |
| `math.randui()` | `→ uint` | Random uint64 in [0, MaxUint64] |
| `math.randui(int/uint/double)` | `→ uint` | Random uint32 (32) or uint64 (64) |

> **Note:** These functions use `math/rand` and are not suitable for security-sensitive applications. Use `crypto/rand` for cryptographic randomness.
>
> **Bound validation:** `math.randi(base, n)` requires a positive `n` that fits
> the selected width (`n <= MaxInt32` for base 32, `n <= MaxInt64` for base 64);
> violations return a `bound '<n>' out of ... size` evaluation error. Invalid
> base values (anything other than 32/64) return a `base '<b>' out of ... size`
> error.

### Value Conversion Utilities

```go
// Convert a CEL ref.Val to a byte slice
func Val2Bytes(val ref.Val) ([]byte, error)

// Convert a CEL ref.Val to a string
func Val2String(val ref.Val) (string, error)

// Convert a CEL ref.Val to a protobuf structpb.Value
func Val2Pb(val ref.Val) (*structpb.Value, error)
```

Examples:

```go
reg, _ := types.NewRegistry()
val := types.NewDynamicMap(reg, map[string]int64{"key": 1})

b, _ := compute.Val2Bytes(val)
fmt.Printf("% 02X\n", b)   // 7B 22 6B 65 79 22 3A 31 7D

s, _ := compute.Val2String(val)
fmt.Println(s)              // {"key":1}

pb, _ := compute.Val2Pb(val)
fmt.Println(pb.GetStructValue().Fields)
```

## Complete CEL Examples

```cel
// Chained bytes operations
b"\x01\x02\x03\x04\x05"
    .index(0)
    .bitwise_and(b"\x01")
    .bitwise_or(b"\xFF")
    .bitwise_shl(7)
    .bitwise_shr(7)  // b"\x01"

// Integer bitwise operations
5.bitwise_and(3)     // 1
5.bitwise_or(3)      // 7
5.bitwise_xor(3)     // 6
0.bitwise_not()      // -1

// Unsigned integer bitwise operations
5u.bitwise_and(3u)   // 1u
0u.bitwise_not()     // 18446744073709551615u

// Byte-level operations
b"\x01\x02\x03\x04".index(1)     // b"\x02"
b"\x01\x02\x03\x04".delete(1)    // b"\x01\x03\x04"
b"\x01\x02\x03\x04".swap(1, 3)   // b"\x01\x04\x03\x02"
b"\x01\x02\x03\x04".slice(1, 3)  // b"\x02\x03"

// Type conversions
42.to_bytes(8)      // b"\x2a"
42u.to_bytes(16)    // b"\x00\x2a"
3.14.to_bytes(32)   // b"\x40\x48\xf5\xc3"

b"\x00\x00\x00\x2a".toi(32)     // 42
b"\x00\x00\x00\x2a".toui(32)    // 42u
b"\x40\x48\xf5\xc3".tof(32)     // 3.14

// Math random
math.randf()        // [0.0, 1.0)
math.randi(32, 10)  // [0, 10) as int32
math.randui(64)     // [0, MaxUint64] as uint64
```

## License

Apache 2.0
