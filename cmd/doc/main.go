package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/common"
	"github.com/google/cel-go/ext"
	compute "github.com/iotxfoundry/cel-go"
)

func main() {
	file, err := os.OpenFile("expr.md", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open file: %v\n", err)
		return
	}
	defer file.Close()
	fmt.Fprintln(file, "# CEL Standard Library Functions")
	fmt.Fprintln(file, "")
	fmt.Fprintln(file, "This document lists the standard functions available in the CEL standard library.")
	fmt.Fprintln(file, "Each function is represented with its name, ID, expression signature, deprecation status, and examples.")
	fmt.Fprintln(file, "The table below provides a comprehensive overview of these functions.")
	fmt.Fprintln(file, "")
	fmt.Fprintln(file, "- Standard Functions")
	fmt.Fprintln(file, "")
	fmt.Fprintln(file, "| name | id | expr | example |")
	fmt.Fprintln(file, "|------|----|------| ------- |")

	env, err := cel.NewEnv(
		cel.OptionalTypes(cel.OptionalTypesVersion(3)),
		compute.ComputeLib(),
		ext.Strings(ext.StringsVersion(3)),
		ext.Sets(ext.SetsVersion(3)),
		ext.Regex(ext.RegexVersion(3)),
		ext.Protos(ext.ProtosVersion(3)),
		ext.NativeTypes(ext.NativeTypesVersion(3)),
		ext.Math(ext.MathVersion(3)),
		ext.Lists(ext.ListsVersion(3)),
		ext.Encoders(ext.EncodersVersion(3)),
		ext.TwoVarComprehensions(ext.TwoVarComprehensionsVersion(3)),
		ext.Bindings(ext.BindingsVersion(3)),
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create env: %v\n", err)
		return
	}
	decls := env.Functions()
	for _, decl := range decls {
		if decl.IsDeclarationDisabled() {
			continue
		}
		for _, od := range decl.OverloadDecls() {
			args := []string{}
			for _, at := range od.ArgTypes() {
				args = append(args, at.TypeName())
			}
			example := strings.Join(od.Examples(), "\n")
			example = escapeMarkdown(example)
			name := decl.Name()
			name = escapeMarkdown(name)
			resultType := od.ResultType().String()
			resultType = escapeMarkdown(resultType)
			fmt.Fprintf(file, "|`%s`|%s|(%s) -> %s|%s|\n", name, od.ID(), strings.Join(args, ","), resultType, example)
		}
	}

	fmt.Fprintln(file, "- Standard Macros")
	fmt.Fprintln(file, "")
	fmt.Fprintln(file, "| name |  expr | example |")
	fmt.Fprintln(file, "|------|-------|---------|")
	micros := env.Macros()
	for _, micro := range micros {
		doc := micro.(common.Documentor)
		descs := []string{}
		for _, child := range doc.Documentation().Children {
			descs = append(descs, child.Description)
		}
		key := micro.MacroKey()
		key = escapeMarkdown(key)
		example := strings.Join(descs, "\n")
		example = escapeMarkdown(example)
		fmt.Fprintf(file, "|`%s`|`%s`|%s|\n", micro.Function(), key, example)
	}
}

func escapeMarkdown(s string) string {
	s = strings.ReplaceAll(s, "|", "\\|") // escape pipes for markdown
	s = strings.ReplaceAll(s, "<", "\\<") // escape less than for markdown
	s = strings.ReplaceAll(s, ">", "\\>") // escape greater than for markdown
	s = strings.ReplaceAll(s, "\n", "<br>")
	return s
}
