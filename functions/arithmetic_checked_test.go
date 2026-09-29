package functions

import (
	"math"
	"math/big"
	"testing"

	"github.com/google/cel-go/common/types"
	"github.com/google/cel-go/common/types/ref"
)

// boundary values for exhaustive cross-type arithmetic verification
var boundaryInts = []int64{
	math.MinInt64, math.MinInt64 + 1, -(1 << 62) - 1, -(1 << 62), -(1 << 62) + 1,
	-3, -2, -1, 0, 1, 2, 3,
	(1 << 62) - 1, 1 << 62, (1 << 62) + 1,
	math.MaxInt64 - 1, math.MaxInt64,
}

var boundaryUints = []uint64{
	0, 1, 2, 3,
	(1 << 62) - 1, 1 << 62, (1 << 62) + 1,
	1<<63 - 1, 1 << 63, 1<<63 + 1,
	math.MaxUint64 - 1, math.MaxUint64,
}

// assertIntVal verifies that a checked helper result agrees with the exact
// big.Int expectation: the exact value when it fits in int64, an error value
// otherwise.
func assertIntVal(t *testing.T, op string, x, y any, got ref.Val, want *big.Int) {
	t.Helper()
	if want.IsInt64() {
		if types.IsError(got) {
			t.Fatalf("%s(%v, %v): unexpected error %v, want %d", op, x, y, got, want)
		}
		if got.Equal(types.Int(want.Int64())) != types.True {
			t.Fatalf("%s(%v, %v): got %v, want %d", op, x, y, got, want)
		}
		return
	}
	if !types.IsError(got) {
		t.Fatalf("%s(%v, %v): got %v, want integer overflow", op, x, y, got)
	}
}

// TestCheckedIntUintAgainstBig exhaustively verifies the int-receiver
// cross-type helpers against arbitrary-precision arithmetic.
func TestCheckedIntUintAgainstBig(t *testing.T) {
	bx := new(big.Int)
	by := new(big.Int)
	for _, x := range boundaryInts {
		bx.SetInt64(x)
		for _, y := range boundaryUints {
			by.SetUint64(y)

			want := new(big.Int).Add(bx, by)
			assertIntVal(t, "checkedAddInt64Uint64", x, y, checkedAddInt64Uint64(x, y), want)

			want = new(big.Int).Sub(bx, by)
			assertIntVal(t, "checkedSubInt64Uint64", x, y, checkedSubInt64Uint64(x, y), want)

			want = new(big.Int).Mul(bx, by)
			assertIntVal(t, "checkedMulInt64Uint64", x, y, checkedMulInt64Uint64(x, y), want)

			if y == 0 {
				if got := checkedDivInt64Uint64(x, y); !types.IsError(got) {
					t.Fatalf("checkedDivInt64Uint64(%d, %d): expected division by zero error", x, y)
				}
				if got := checkedModInt64Uint64(x, y); !types.IsError(got) {
					t.Fatalf("checkedModInt64Uint64(%d, %d): expected modulus by zero error", x, y)
				}
				continue
			}
			want = new(big.Int).Quo(bx, by) // truncating division, as in Go
			assertIntVal(t, "checkedDivInt64Uint64", x, y, checkedDivInt64Uint64(x, y), want)

			want = new(big.Int).Rem(bx, by) // remainder sign follows x, as in Go
			assertIntVal(t, "checkedModInt64Uint64", x, y, checkedModInt64Uint64(x, y), want)
		}
	}
}

// TestCheckedUintIntAgainstBig exhaustively verifies the uint-receiver
// cross-type helpers against arbitrary-precision arithmetic.
func TestCheckedUintIntAgainstBig(t *testing.T) {
	bx := new(big.Int)
	by := new(big.Int)
	for _, x := range boundaryUints {
		bx.SetUint64(x)
		for _, y := range boundaryInts {
			by.SetInt64(y)

			want := new(big.Int).Add(bx, by)
			assertIntVal(t, "checkedAddUint64Int64", x, y, checkedAddUint64Int64(x, y), want)

			want = new(big.Int).Sub(bx, by)
			assertIntVal(t, "checkedSubUint64Int64", x, y, checkedSubUint64Int64(x, y), want)

			want = new(big.Int).Mul(bx, by)
			assertIntVal(t, "checkedMulUint64Int64", x, y, checkedMulUint64Int64(x, y), want)

			if y == 0 {
				if got := checkedDivUint64Int64(x, y); !types.IsError(got) {
					t.Fatalf("checkedDivUint64Int64(%d, %d): expected division by zero error", x, y)
				}
				if got := checkedModUint64Int64(x, y); !types.IsError(got) {
					t.Fatalf("checkedModUint64Int64(%d, %d): expected modulus by zero error", x, y)
				}
				continue
			}
			want = new(big.Int).Quo(bx, by)
			assertIntVal(t, "checkedDivUint64Int64", x, y, checkedDivUint64Int64(x, y), want)

			want = new(big.Int).Rem(bx, by)
			assertIntVal(t, "checkedModUint64Int64", x, y, checkedModUint64Int64(x, y), want)
		}
	}
}

// TestShiftBytesAgainstBig verifies shiftBytes against arbitrary-precision
// big-endian shifts over every direction, sub-byte and whole-byte magnitude,
// and the MinInt64/MaxInt64 extremes.
func TestShiftBytesAgainstBig(t *testing.T) {
	srcs := [][]byte{
		{},
		{0x00},
		{0xFF},
		{0xF0},
		{0x01, 0x02},
		{0xFF, 0xFF},
		{0x80, 0x00},
		{0x80, 0x00, 0x00, 0x00},
		{0xDE, 0xAD, 0xBE, 0xEF, 0xCA, 0xFE, 0xBA, 0xBE, 0x01},
	}
	shifts := []int64{
		0, 1, 3, 7, 8, 9, 15, 16, 17, 63, 64, 65, 127, 128,
		-1, -3, -7, -8, -9, -15, -16, -17, -63, -64, -65, -127, -128,
		math.MinInt64, math.MinInt64 + 1, math.MaxInt64,
	}

	v := new(big.Int)
	shifted := new(big.Int)
	for _, src := range srcs {
		v.SetBytes(src) // big-endian
		n := len(src)
		for _, s := range shifts {
			for _, left := range []bool{false, true} {
				// left=true means toward the most significant byte (Lsh).
				// Clamp the magnitude: once every bit has shifted out the
				// result is all zeros, and big.Int cannot shift by 2^63.
				mag := s
				if mag == math.MinInt64 {
					mag = math.MaxInt64
				} else if mag < 0 {
					mag = -mag
				}
				if uint64(mag) > uint64(8*n) {
					mag = int64(8 * n)
				}
				if left {
					// Left shifts grow past the buffer; the buffer keeps a
					// fixed length, so truncate to 8n bits.
					shifted.Lsh(v, uint(mag))
					if n > 0 {
						mask := new(big.Int).Lsh(big.NewInt(1), uint(8*n))
						mask.Sub(mask, big.NewInt(1))
						shifted.And(shifted, mask)
					}
				} else {
					shifted.Rsh(v, uint(mag))
				}

				got := shiftBytes(src, s, left)
				want := make([]byte, n)
				b := shifted.Bytes()
				copy(want[n-len(b):], b)
				if len(got) != n {
					t.Fatalf("shiftBytes(% X, %d, %v): got len %d, want %d", src, s, left, len(got), n)
				}
				for i := range want {
					if got[i] != want[i] {
						t.Fatalf("shiftBytes(% X, %d, %v): got % X, want % X", src, s, left, got, want)
					}
				}
			}
		}
	}
}
