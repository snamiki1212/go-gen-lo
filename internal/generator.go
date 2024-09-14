package internal

import (
	"bytes"
	"fmt"
	"text/template"
)

type loMethods string

const (
	loMethodsFilter loMethods = "Filter"
	loMethodsMap    loMethods = "Map"
)

var loMethodAll = []loMethods{loMethodsFilter, loMethodsMap}

var loMethodTemplates = map[loMethods]string{
	loMethodsFilter: `
// Filter
func (xs {{ .Slice }}) Filter(predicate func(item {{ .Entity }}, index int) bool) {{ .Slice }} {
	return lo.Filter(xs, predicate)
}
`,
	loMethodsMap: `
// Map
func (xs {{ .Slice }}) Map(iteratee func(item {{ .Entity }}, index int) {{ .Entity }}) {{ .Slice }} {
	return lo.Map(xs, iteratee)
}
`,
}

type loMethodsExtend string

const (
	loMethodsExtendFilter loMethodsExtend = "FilterBy"
	loMethodsExtendKeyBy  loMethodsExtend = "KeyBy"
)

var loMethodExtendAll = []loMethodsExtend{loMethodsExtendFilter, loMethodsExtendKeyBy}

var loMethodExtendTemplates = map[loMethodsExtend]string{
	loMethodsExtendFilter: `
// FilterBy{{ .Field }}
func (xs {{ .Slice }}) FilterBy{{ .Field }}({{ .Field }} {{ .Type }}) {{ .Slice }} {
	return lo.Filter(xs, func(item {{ .Entity }}, index int) bool {
		return item.{{ .Field }} == {{ .Field }}
	})
}
`,
	loMethodsExtendKeyBy: `
// KeyBy{{ .Field }}
func (xs {{ .Slice }}) KeyBy{{ .Field }}() map[{{ .Type }}]{{ .Entity }} {
	return lo.KeyBy(xs, func(item {{ .Entity }}) {{ .Type }} {
		return item.{{ .Field }}
	})
}
`,
}

// Replace variable from key to value in template.
type loMethodTemplateMapper struct {
	Slice  string // Slice name for target struct (ex. Users).
	Entity string // Entity name for target struct (ex. User / *User).
}

// Replace variable from key to value in template.
type loMethodExtendTemplateMapper struct {
	Slice  string // Slice name for target struct (ex. Users).
	Entity string // Entity name for target struct (ex. User / *User).
	Type   string // Type name of field (ex. string).
	Field  string // Field name of struct (ex. UserID).
}

// Generate code
func Generate(args arguments, data data) (string, error) {
	pkgName := data.pkgName
	sliceName := data.sliceName
	fields := data.fields

	var txt string

	// append header
	txt += "// Code generated by go generate DO NOT EDIT.\n\n"
	txt += "package " + pkgName + "\n\n"
	txt += `import "github.com/samber/lo"` + "\n\n"

	// append lo
	lo, err := generateLo(args, sliceName)
	if err != nil {
		return "", fmt.Errorf("generate lo error: %w", err)
	}
	txt += lo

	// append extend
	lo, err = generateLoExtend(args, sliceName, fields)
	if err != nil {
		return "", fmt.Errorf("generate lo extend error: %w", err)
	}
	txt += lo

	return txt, nil
}

func generateLo(args arguments, sliceName string) (string, error) {
	// append loMethodTemplates
	var doc bytes.Buffer

	for _, method := range loMethodAll {
		// Get template src
		rawTp, ok := loMethodTemplates[method]
		if !ok {
			return "", fmt.Errorf("template not found: %s", method)
		}

		// New Template
		tp, err := template.New("").Parse(rawTp)
		if err != nil {
			return "", fmt.Errorf("template parse error: %w", err)
		}

		// Generate txt from template
		data := &loMethodTemplateMapper{Slice: sliceName, Entity: args.DisplayEntity()}
		if err = tp.Execute(&doc, data); err != nil {
			return "", fmt.Errorf("template execute error: %w", err)
		}
	}
	return doc.String(), nil
}

func generateLoExtend(args arguments, sliceName string, fields fields) (string, error) {
	if len(fields) == 0 {
		return "", nil
	}
	var doc bytes.Buffer
	for _, method := range loMethodExtendAll {
		// Get template src
		rawTp, ok := loMethodExtendTemplates[method]
		if !ok {
			return "", fmt.Errorf("template not found: %s", method)
		}

		// New Template
		tp, err := template.New("").Parse(rawTp)
		if err != nil {
			return "", fmt.Errorf("template parse error: %w", err)
		}

		// Generate txt from template
		for _, field := range fields {
			data := &loMethodExtendTemplateMapper{Slice: sliceName, Entity: args.DisplayEntity(), Type: field.Type, Field: field.Name}
			err = tp.Execute(&doc, data)
			if err != nil {
				return "", fmt.Errorf("template execute error: %w", err)
			}
		}
	}
	return doc.String(), nil
}
