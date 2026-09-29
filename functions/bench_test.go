package functions

import (
	"testing"

	"cel.dev/cel-go/cel"
)

// --- pure helper benchmarks ---

func BenchmarkCheckedAddInt64Uint64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		checkedAddInt64Uint64(9223372036854775000, 100)
	}
}

func BenchmarkCheckedMulInt64Uint64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		checkedMulInt64Uint64(-3037000500, 3)
	}
}

func BenchmarkCheckedDivUint64Int64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		checkedDivUint64Int64(1<<63, 7)
	}
}

func BenchmarkShiftBytesSmall(b *testing.B) {
	src := []byte{0xDE, 0xAD, 0xBE, 0xEF}
	for i := 0; i < b.N; i++ {
		shiftBytes(src, 5, true)
	}
}

// --- end-to-end CEL eval benchmarks ---

// benchLib adapts Functions() to cel.Library for benchmark environments.
type benchLib struct{}

func (benchLib) LibraryName() string                 { return "functions.bench" }
func (benchLib) CompileOptions() []cel.EnvOption     { return Functions() }
func (benchLib) ProgramOptions() []cel.ProgramOption { return nil }

// benchEval compiles src once and returns a program ready for evaluation.
func benchEval(b *testing.B, src string) cel.Program {
	b.Helper()
	env, err := cel.NewEnv(cel.Lib(benchLib{}))
	if err != nil {
		b.Fatal(err)
	}
	ast, iss := env.Compile(src)
	if iss.Err() != nil {
		b.Fatalf("Compile(%q): %v", src, iss.Err())
	}
	prg, err := env.Program(ast)
	if err != nil {
		b.Fatalf("Program(%q): %v", src, err)
	}
	return prg
}

func BenchmarkEvalAddIntUint(b *testing.B) {
	prg := benchEval(b, `3.add(4u)`)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, _, err := prg.Eval(map[string]any{}); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkEvalRandf(b *testing.B) {
	prg := benchEval(b, `math.randf()`)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, _, err := prg.Eval(map[string]any{}); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkEvalRandiBounded(b *testing.B) {
	prg := benchEval(b, `math.randi(32, 100)`)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, _, err := prg.Eval(map[string]any{}); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkEvalRandui(b *testing.B) {
	prg := benchEval(b, `math.randui(32)`)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, _, err := prg.Eval(map[string]any{}); err != nil {
			b.Fatal(err)
		}
	}
}
