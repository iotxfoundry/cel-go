package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestEscapeMarkdown verifies markdown table-cell escaping of pipes, angle
// brackets, and newlines.
func TestEscapeMarkdown(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{"empty", "", ""},
		{"plain", "int(int) -> int", `int(int) -\> int`},
		{"pipe", "a|b", `a\|b`},
		{"lt", "<T", `\<T`},
		{"gt", "T>", `T\>`},
		{"newline", "line1\nline2", "line1<br>line2"},
		{"combined", "map<K,|V>\nx", `map\<K,\|V\><br>x`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := escapeMarkdown(tt.in); got != tt.want {
				t.Errorf("escapeMarkdown(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}

// TestMainWritesDoc runs main() inside a temporary working directory and
// verifies it produces a populated expr.md document.
func TestMainWritesDoc(t *testing.T) {
	dir := t.TempDir()
	origWd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir(dir); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := os.Chdir(origWd); err != nil {
			t.Fatalf("restore working directory: %v", err)
		}
	})

	main()

	data, err := os.ReadFile(filepath.Join(dir, "expr.md"))
	if err != nil {
		t.Fatalf("expr.md not written: %v", err)
	}
	doc := string(data)
	for _, want := range []string{
		"# CEL Standard Library Functions",
		"| name | id | expr | example |",
		"Standard Macros",
	} {
		if !strings.Contains(doc, want) {
			t.Errorf("expr.md missing %q", want)
		}
	}
	if len(doc) < 1000 {
		t.Errorf("expr.md suspiciously small (%d bytes); function table likely empty", len(doc))
	}
}
