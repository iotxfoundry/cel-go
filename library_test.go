package cel_test

import (
	"strings"
	"testing"

	"github.com/google/cel-go/common/stdlib"
)

func TestStandardLibrary(t *testing.T) {
	decls := stdlib.Functions()
	for _, decl := range decls {
		for _, od := range decl.OverloadDecls() {
			args := []string{}
			for _, at := range od.ArgTypes() {
				args = append(args, at.TypeName())
			}
			t.Logf("|`%s`|%s|(%s) -> %s|", decl.Name(), od.ID(), strings.Join(args, ","), od.ResultType().String())
		}
	}
}
