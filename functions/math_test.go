package functions

import (
	"math"
	"strings"
	"testing"

	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/common/types"
	"github.com/google/cel-go/common/types/ref"
)

// testLib adapts Functions() to the cel.Library interface so that every test
// exercises the public entry point rather than an ad-hoc option list.
type testLib struct{}

func (testLib) LibraryName() string                 { return "functions.test" }
func (testLib) CompileOptions() []cel.EnvOption     { return Functions() }
func (testLib) ProgramOptions() []cel.ProgramOption { return nil }

// newTestEnv builds a CEL environment exposing every option from Functions().
func newTestEnv(t *testing.T) *cel.Env {
	t.Helper()
	env, err := cel.NewEnv(cel.Lib(testLib{}))
	if err != nil {
		t.Fatalf("cel.NewEnv: %v", err)
	}
	return env
}

// newTestEnvWithVars builds a CEL environment from Functions() plus typed
// dynamic variables for exercising the non-literal macro argument paths.
func newTestEnvWithVars(t *testing.T) *cel.Env {
	t.Helper()
	env, err := cel.NewEnv(
		cel.Lib(testLib{}),
		cel.Variable("x", cel.IntType),
		cel.Variable("n", cel.IntType),
		cel.Variable("m", cel.MapType(cel.StringType, cel.DynType)),
	)
	if err != nil {
		t.Fatalf("cel.NewEnv: %v", err)
	}
	return env
}

// evalCompilesAndRuns compiles src and evaluates it once, returning the result.
func evalOnce(t *testing.T, env *cel.Env, src string, vars map[string]any) (ref.Val, error) {
	t.Helper()
	ast, iss := env.Compile(src)
	if iss.Err() != nil {
		t.Fatalf("Compile(%q): %v", src, iss.Err())
	}
	prg, err := env.Program(ast)
	if err != nil {
		t.Fatalf("Program(%q): %v", src, err)
	}
	out, _, err := prg.Eval(vars)
	return out, err
}

// evalCase is a happy-path case whose random result is verified by check.
type randCase struct {
	name   string
	source string
	vars   map[string]any
	check  func(t *testing.T, out ref.Val)
}

// checkDoubleInRange asserts out is a types.Double within [lo, hi).
func checkDoubleInRange(lo, hi float64) func(*testing.T, ref.Val) {
	return func(t *testing.T, out ref.Val) {
		t.Helper()
		d, ok := out.(types.Double)
		if !ok {
			t.Fatalf("got %T (%v), want types.Double", out, out)
		}
		if float64(d) < lo || float64(d) >= hi {
			t.Fatalf("got %v, want value in [%v, %v)", float64(d), lo, hi)
		}
	}
}

// checkIntInRange asserts out is a types.Int within [lo, hi] (both inclusive,
// matching math/rand's Int63/Int31 contracts and the library docs).
func checkIntInRange(lo, hi int64) func(*testing.T, ref.Val) {
	return func(t *testing.T, out ref.Val) {
		t.Helper()
		i, ok := out.(types.Int)
		if !ok {
			t.Fatalf("got %T (%v), want types.Int", out, out)
		}
		if int64(i) < lo || int64(i) > hi {
			t.Fatalf("got %d, want value in [%d, %d]", int64(i), lo, hi)
		}
	}
}

// checkUintInRange asserts out is a types.Uint within [lo, hi] (inclusive).
func checkUintInRange(lo, hi uint64) func(*testing.T, ref.Val) {
	return func(t *testing.T, out ref.Val) {
		t.Helper()
		u, ok := out.(types.Uint)
		if !ok {
			t.Fatalf("got %T (%v), want types.Uint", out, out)
		}
		if uint64(u) < lo || uint64(u) > hi {
			t.Fatalf("got %d, want value in [%d, %d]", uint64(u), lo, hi)
		}
	}
}

// TestFunctionsEntryPoint verifies Functions() assembles a usable library.
func TestFunctionsEntryPoint(t *testing.T) {
	opts := Functions()
	if len(opts) < 9 {
		t.Fatalf("Functions() returned %d options, want at least 9", len(opts))
	}
	env := newTestEnv(t)

	// One representative expression per function group: bytes, bytes bitwise,
	// int bitwise, uint bitwise, int members, uint members, double members,
	// cross-type arithmetic, math rand.
	tests := []struct {
		source string
		want   ref.Val
	}{
		{`b"\x01\x02".index(1)`, types.Bytes{0x02}},
		{`b"\x0f".bitwise_and(b"\x03")`, types.Bytes{0x03}},
		{`5.bitwise_or(3)`, types.Int(7)},
		{`5u.bitwise_xor(3u)`, types.Uint(6)},
		{`42.to_bytes(8)`, types.Bytes{0x2a}},
		{`42u.to_bytes(16)`, types.Bytes{0x00, 0x2a}},
		{`3.14.to_bytes(32)`, types.Bytes{0x40, 0x48, 0xf5, 0xc3}},
		{`3.add(4u)`, types.Int(7)},
		{`3u.add(-4)`, types.Int(-1)},
	}
	for _, tt := range tests {
		t.Run(tt.source, func(t *testing.T) {
			out, err := evalOnce(t, env, tt.source, nil)
			if err != nil {
				t.Fatalf("Eval(%q): %v", tt.source, err)
			}
			if out.Equal(tt.want) != types.True {
				t.Errorf("Eval(%q) = %v, want %v", tt.source, out, tt.want)
			}
		})
	}
}

// TestMathRandFunctions verifies value ranges of every rand overload across
// int, uint, and double literal arguments. Each case is evaluated repeatedly
// so that boundary violations surface with high probability.
func TestMathRandFunctions(t *testing.T) {
	const iterations = 200
	tests := []randCase{
		// --- math.randf ---
		{name: "randf()", source: `math.randf()`, check: checkDoubleInRange(0.0, 1.0)},
		{name: "randf(32) int", source: `math.randf(32)`, check: checkDoubleInRange(0.0, 1.0)},
		{name: "randf(64) int", source: `math.randf(64)`, check: checkDoubleInRange(0.0, 1.0)},
		{name: "randf(32u) uint", source: `math.randf(32u)`, check: checkDoubleInRange(0.0, 1.0)},
		{name: "randf(64u) uint", source: `math.randf(64u)`, check: checkDoubleInRange(0.0, 1.0)},
		{name: "randf(32.0) double", source: `math.randf(32.0)`, check: checkDoubleInRange(0.0, 1.0)},
		{name: "randf(64.0) double", source: `math.randf(64.0)`, check: checkDoubleInRange(0.0, 1.0)},
		// dynamic (non-literal) arguments are accepted by the macro
		{name: "randf(x) dynamic", source: `math.randf(x)`, vars: map[string]any{"x": int64(64)},
			check: checkDoubleInRange(0.0, 1.0)},

		// --- math.randi ---
		{name: "randi()", source: `math.randi()`, check: checkIntInRange(0, math.MaxInt64)},
		{name: "randi(32)", source: `math.randi(32)`, check: checkIntInRange(0, math.MaxInt32)},
		{name: "randi(64)", source: `math.randi(64)`, check: checkIntInRange(0, math.MaxInt64)},
		{name: "randi(32u)", source: `math.randi(32u)`, check: checkIntInRange(0, math.MaxInt32)},
		{name: "randi(64u)", source: `math.randi(64u)`, check: checkIntInRange(0, math.MaxInt64)},
		{name: "randi(32.0)", source: `math.randi(32.0)`, check: checkIntInRange(0, math.MaxInt32)},
		{name: "randi(64.0)", source: `math.randi(64.0)`, check: checkIntInRange(0, math.MaxInt64)},
		{name: "randi(32, 10)", source: `math.randi(32, 10)`, check: checkIntInRange(0, 9)},
		{name: "randi(64, 10)", source: `math.randi(64, 10)`, check: checkIntInRange(0, 9)},
		{name: "randi(32u, 10)", source: `math.randi(32u, 10)`, check: checkIntInRange(0, 9)},
		{name: "randi(64u, 10u)", source: `math.randi(64u, 10u)`, check: checkIntInRange(0, 9)},
		{name: "randi(32.0, 10.0)", source: `math.randi(32.0, 10.0)`, check: checkIntInRange(0, 9)},
		{name: "randi(64.0, 10)", source: `math.randi(64.0, 10)`, check: checkIntInRange(0, 9)},
		{name: "randi(x, n) dynamic", source: `math.randi(x, n)`,
			vars:  map[string]any{"x": int64(32), "n": int64(100)},
			check: checkIntInRange(0, 99)},

		// --- math.randui ---
		{name: "randui()", source: `math.randui()`, check: checkUintInRange(0, math.MaxUint64)},
		{name: "randui(32)", source: `math.randui(32)`, check: checkUintInRange(0, math.MaxUint32)},
		{name: "randui(64)", source: `math.randui(64)`, check: checkUintInRange(0, math.MaxUint64)},
		{name: "randui(32u)", source: `math.randui(32u)`, check: checkUintInRange(0, math.MaxUint32)},
		{name: "randui(64u)", source: `math.randui(64u)`, check: checkUintInRange(0, math.MaxUint64)},
		{name: "randui(32.0)", source: `math.randui(32.0)`, check: checkUintInRange(0, math.MaxUint32)},
		{name: "randui(64.0)", source: `math.randui(64.0)`, check: checkUintInRange(0, math.MaxUint64)},
		{name: "randui(x) dynamic", source: `math.randui(x)`, vars: map[string]any{"x": int64(32)},
			check: checkUintInRange(0, math.MaxUint32)},
	}

	env := newTestEnvWithVars(t)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for range iterations {
				out, err := evalOnce(t, env, tt.source, tt.vars)
				if err != nil {
					t.Fatalf("Eval: %v", err)
				}
				tt.check(t, out)
			}
		})
	}
}

// TestMathRandInvalidBaseErrors verifies invalid base arguments produce
// domain errors at evaluation time rather than values or internal panics.
func TestMathRandInvalidBaseErrors(t *testing.T) {
	env := newTestEnv(t)
	tests := []struct {
		name        string
		source      string
		errContains string
	}{
		{"randf 16 int", `math.randf(16)`, "base '16' out of float32, float64 size"},
		{"randf 128 int", `math.randf(128)`, "base '128' out of float32, float64 size"},
		{"randf 16 uint", `math.randf(16u)`, "base '16' out of float32, float64 size"},
		{"randf 16 double", `math.randf(16.0)`, "base '16' out of float32, float64 size"},
		{"randi 16 int", `math.randi(16)`, "base '16' out of int32, int64 size"},
		{"randi 16 uint", `math.randi(16u)`, "base '16' out of int32, int64 size"},
		{"randi 16 double", `math.randi(16.0)`, "base '16' out of int32, int64 size"},
		{"randi 16,10 int", `math.randi(16, 10)`, "base '16' out of int32, int64 size"},
		{"randi 16,10 uint", `math.randi(16u, 10)`, "base '16' out of int32, int64 size"},
		{"randi 16,10 double", `math.randi(16.0, 10)`, "base '16' out of int32, int64 size"},
		{"randui 16 int", `math.randui(16)`, "base '16' out of uint32, uint64 size"},
		{"randui 16 uint", `math.randui(16u)`, "base '16' out of uint32, uint64 size"},
		{"randui 16 double", `math.randui(16.0)`, "base '16' out of uint32, uint64 size"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := evalOnce(t, env, tt.source, nil)
			if err == nil {
				t.Fatalf("Eval(%q) = %v, want error %q", tt.source, out, tt.errContains)
			}
			if !strings.Contains(err.Error(), tt.errContains) {
				t.Errorf("Eval(%q) error = %q, want containing %q", tt.source, err.Error(), tt.errContains)
			}
		})
	}
}

// TestMathRandBoundErrors verifies the bound argument n of randi(base, n):
// non-positive bounds and bounds exceeding the selected integer width must
// produce domain errors, never internal "invalid argument" panics from
// math/rand, and never silently truncated bounds.
func TestMathRandBoundErrors(t *testing.T) {
	env := newTestEnv(t)
	tests := []struct {
		name        string
		source      string
		errContains string
	}{
		// zero bound
		{"randi 32,0", `math.randi(32, 0)`, "bound"},
		{"randi 64,0", `math.randi(64, 0)`, "bound"},
		{"randi 32u,0", `math.randi(32u, 0)`, "bound"},
		{"randi 32.0,0", `math.randi(32.0, 0.0)`, "bound"},
		// negative bound
		{"randi 32,-1", `math.randi(32, -1)`, "bound"},
		{"randi 64,-5", `math.randi(64, -5)`, "bound"},
		{"randi 32u,-1", `math.randi(32u, -1)`, "bound"},
		// bound exceeding int32 width for base 32
		{"randi 32,2^32", `math.randi(32, 4294967296)`, "bound"},
		{"randi 32,2^63", `math.randi(32, 9223372036854775808u)`, "bound"},
		// negative bound via uint cannot occur, but an oversized one can
		{"randi 64u,2^63", `math.randi(64u, 9223372036854775808u)`, "bound"},
		{"randi 32u,2^63", `math.randi(32u, 9223372036854775808u)`, "bound"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := evalOnce(t, env, tt.source, nil)
			if err == nil {
				t.Fatalf("Eval(%q) = %v, want error containing %q", tt.source, out, tt.errContains)
			}
			if !strings.Contains(err.Error(), tt.errContains) {
				t.Errorf("Eval(%q) error = %q, want containing %q", tt.source, err.Error(), tt.errContains)
			}
			if strings.Contains(err.Error(), "internal error") {
				t.Errorf("Eval(%q) leaked internal panic text: %q", tt.source, err.Error())
			}
		})
	}
}

// TestMathRandMacroErrors verifies compile-time macro validation: argument
// count limits, argument type restrictions, and receiver-namespace checks.
func TestMathRandMacroErrors(t *testing.T) {
	env := newTestEnvWithVars(t)
	tests := []struct {
		name        string
		source      string
		errContains string
	}{
		// too many arguments
		{"randf two args", `math.randf(32, 64)`, "math.randf() requires"},
		{"randi three args", `math.randi(32, 10, 5)`, "math.randi() requires"},
		{"randui two args", `math.randui(32, 64)`, "math.randui() requires"},
		// invalid literal argument types
		{"randf string arg", `math.randf("bad")`, "math.randf() invalid single argument"},
		{"randi string arg", `math.randi("bad")`, "math.randi() invalid single argument"},
		{"randi string second arg", `math.randi("bad", 10)`, "math.randi() invalid argument"},
		{"randi string bound", `math.randi(32, "bad")`, "math.randi() invalid argument"},
		{"randui string arg", `math.randui("bad")`, "math.randui() invalid single argument"},
		// unsupported aggregate literals
		{"randf list arg", `math.randf([1, 2])`, "math.randf() invalid single argument"},
		{"randui map arg", `math.randui({'a': 1})`, "math.randui() invalid single argument"},
		// wrong receiver namespace: the macro declines and the call is left
		// undeclared
		{"wrong target randf", `foo.randf()`, "undeclared"},
		{"wrong target randi", `foo.randi(32)`, "undeclared"},
		{"wrong target randui", `foo.randui(64)`, "undeclared"},
		// non-identifier targets (select expressions) are declined as well
		{"select target randf", `m.field.randf()`, "undeclared"},
		{"select target randui", `m.field.randui(32)`, "undeclared"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, iss := env.Compile(tt.source)
			if iss.Err() == nil {
				t.Fatalf("Compile(%q): expected error containing %q, got none", tt.source, tt.errContains)
			}
			if !strings.Contains(iss.Err().Error(), tt.errContains) {
				t.Errorf("Compile(%q) error = %q, want containing %q", tt.source, iss.Err().Error(), tt.errContains)
			}
		})
	}
}

// TestMathRandDeterministicBounds sweeps randi(base, n) over a spread of
// bounds and verifies every sample lands strictly inside [0, n). Base 32 is
// restricted to bounds fitting int32; larger bounds belong to base 64.
func TestMathRandDeterministicBounds(t *testing.T) {
	env := newTestEnv(t)
	bounds := []int64{1, 2, 3, 7, 100, 1 << 20, math.MaxInt32, int64(math.MaxInt32) + 1, 1 << 40}
	for _, n := range bounds {
		for _, base := range []string{"32", "64"} {
			if base == "32" && n > math.MaxInt32 {
				continue // int32 bound overflow is covered by TestMathRandBoundErrors
			}
			name := "randi(" + base + "," + itoa(n) + ")"
			t.Run(name, func(t *testing.T) {
				src := "math.randi(" + base + ", " + itoa(n) + ")"
				for range 100 {
					out, err := evalOnce(t, env, src, nil)
					if err != nil {
						t.Fatalf("Eval(%q): %v", src, err)
					}
					v, ok := out.(types.Int)
					if !ok {
						t.Fatalf("Eval(%q) = %T, want types.Int", src, out)
					}
					if int64(v) < 0 || int64(v) >= n {
						t.Fatalf("Eval(%q) = %d, want in [0, %d)", src, int64(v), n)
					}
				}
			})
		}
	}
}

// itoa avoids importing strconv for one call site.
func itoa(n int64) string {
	if n == 0 {
		return "0"
	}
	neg := n < 0
	if neg {
		n = -n
	}
	var buf [21]byte
	i := len(buf)
	for n > 0 {
		i--
		buf[i] = byte('0' + n%10)
		n /= 10
	}
	if neg {
		i--
		buf[i] = '-'
	}
	return string(buf[i:])
}
