package internal

import (
	"bytes"
	"fmt"
	"text/template"
)

type loMethodsExtend string

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
	txt += "// Code generated by \"go-gen-lo\"; DO NOT EDIT.\n"
	txt += "// Based on information from https://github.com/snamiki1212/go-gen-lo\n"
	txt += "\n"
	txt += fmt.Sprintf("package %s\n\n", pkgName)
	txt += `import "github.com/samber/lo"` + "\n"

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

	list := []interface{ StdTemplate() (string, bool) }{
		NewLoFilter(),
		NewLoMap(),
		NewLoKeyBy(),
	}

	for _, elem := range list {
		rawTemp, ok := elem.StdTemplate()
		if !ok {
			continue
		}

		// New Template
		tp, err := template.New("").Parse(rawTemp)
		if err != nil {
			return "", fmt.Errorf("template parse error: %w", err)
		}

		// Generate code block from template
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

	list := []interface{ ExtendTemplate() (string, bool) }{
		NewLoFilter(),
		NewLoMap(),
		NewLoKeyBy(),
	}

	for _, elem := range list {
		// Get template src
		rawTemp, ok := elem.ExtendTemplate()
		if !ok {
			continue
		}

		// New Template
		tp, err := template.New("").Parse(rawTemp)
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
