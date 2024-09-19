package generator

import (
	"bytes"
	"fmt"
	"io"
	"text/template"

	"github.com/snamiki1212/go-gen-lo/internal"
)

type Generator struct{ LoList []Lo }

func NewGenerator() Generator { return Generator{LoList: NewAllLoList()} }

// Generate code
func (g Generator) Generate(args internal.Arguments, data internal.Data) (string, error) {
	pkgName := data.PackageName
	sliceName := data.SliceName
	fields := data.Fields

	var txt string

	{
		var tmp string

		// append lo
		tmp, err := g.genStd(args, sliceName)
		if err != nil {
			return "", fmt.Errorf("generate lo error: %w", err)
		}
		if tmp != "" {
			txt += `
/************************************************
 ** lo basic methods
 ************************************************/
`
		}
		txt += tmp

		// append extend
		tmp, err = g.genExtend(args, sliceName, fields)
		if err != nil {
			return "", fmt.Errorf("generate lo extend error: %w", err)
		}
		if tmp != "" {
			txt += `
/************************************************
 ** lo extended methods
 ************************************************/
`
		}
		txt += tmp
	}

	// get header
	haeder := genHeader(pkgName, txt != "")

	haeder += txt

	return haeder, nil
}

func genHeader(pkgName string, withImport bool) string {
	var txt string
	txt += "// Code generated by \"go-gen-lo\"; DO NOT EDIT.\n"
	txt += "// Based on information from https://github.com/snamiki1212/go-gen-lo\n"
	txt += "\n"
	txt += fmt.Sprintf("package %s\n\n", pkgName)
	if withImport {
		txt += `import "github.com/samber/lo"` + "\n"
	}
	return txt
}

func (g Generator) genStd(args internal.Arguments, sliceName string) (string, error) {
	// append loMethodTemplates
	var doc bytes.Buffer

	for _, elem := range g.LoList {
		// Skip if method is excluded.
		if args.IsExcluded(elem.StdName()) {
			continue
		}

		// Skip if method is NOT included.
		if !args.IsIncluded(elem.StdName()) {
			continue
		}

		// Get template src
		rawTemp, ok := elem.StdTemplate()
		if !ok {
			continue
		}

		// New Template
		tp, err := template.New("").Parse(rawTemp)
		if err != nil {
			return "", fmt.Errorf("template parse error: %w", err)
		}

		// Get method name
		method, _ := args.Rename(elem.StdName())

		// Generate code block from template
		data := &loStdTemplateMapper{Slice: sliceName, Entity: args.DisplayEntity(), Method: method}
		if err = tp.Execute(&doc, data); err != nil {
			return "", fmt.Errorf("template execute error: %w", err)
		}
	}

	return doc.String(), nil
}

func (g Generator) genExtend(args internal.Arguments, sliceName string, fields internal.Fields) (string, error) {
	if len(fields) == 0 {
		return "", nil
	}
	var doc bytes.Buffer

	for _, elem := range g.LoList {
		// Skip if method is excluded.
		if args.IsExcluded(elem.StdName()) {
			continue
		}

		// Skip if method is NOT included.
		if !args.IsIncluded(elem.StdName()) {
			continue
		}

		// Get template src
		rawTemp, ok := elem.ExtendTemplate()
		if !ok {
			continue
		}

		// Add header
		_, err := io.WriteString(&doc, fmt.Sprintf("\n// -- %s ------------------------------------\n", elem.StdName()))
		if err != nil {
			return "", fmt.Errorf("write error: %w", err)
		}

		// New Template
		tp, err := template.New("").Parse(rawTemp)
		if err != nil {
			return "", fmt.Errorf("template parse error: %w", err)
		}

		// Get method name
		method := func() string {
			ki := elem.StdName()
			str, ok := args.Rename(ki)
			if ok {
				return str
			}
			me, ok := elem.ExtendName()
			if ok {
				return me
			}
			return ki
		}()

		// Generate txt from template
		for _, field := range fields {
			me := method + field.Name
			data := &loExtendTemplateMapper{Slice: sliceName, Entity: args.DisplayEntity(), Type: field.Type, Field: field.Name, Method: me}
			err = tp.Execute(&doc, data)
			if err != nil {
				return "", fmt.Errorf("template execute error: %w", err)
			}
		}
	}

	return doc.String(), nil
}
