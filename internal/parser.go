package internal

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"slices"
	"strings"
)

// Parse sorce code to own struct.
func Parse(args arguments, reader func(path string) (*ast.File, error)) (data, error) {
	// Convert source code to ast
	file, err := reader(args.Input)
	if err != nil {
		return data{}, fmt.Errorf("parse error: %w", err)
	}

	// Parse ast
	fields, err := parseFile(file, args)
	if err != nil {
		return data{}, err
	}

	// Convert ast to own struct
	fs := newFields(fields)

	// Transform data
	fs = fs.excludeUncomparable()

	return data{
		fields:    fs,
		pkgName:   getPackageNameFromFile(file),
		sliceName: args.Slice,
	}, nil
}

// Read source code from file.
func Reader(path string) (*ast.File, error) {
	fset := token.NewFileSet()
	return parser.ParseFile(fset, path, nil, parser.AllErrors)
}

// Get package name.
func getPackageNameFromFile(node *ast.File) string { return node.Name.Name }

// Parse file.
func parseFile(node *ast.File, args arguments) ([]*ast.Field, error) {
	// Find entity object
	obj, ok := node.Scope.Objects[args.Entity]
	if !ok {
		return nil, fmt.Errorf("entity not found: %s", args.Entity)
	}

	// Find entity
	entity, ok := obj.Decl.(*ast.TypeSpec)
	if !ok {
		return nil, fmt.Errorf("invalid entity: %s", args.Entity)
	}

	// Find fields
	str, ok := entity.Type.(*ast.StructType)
	if !ok {
		return nil, fmt.Errorf("invalid type: %T", str)
	}
	fs := str.Fields.List

	return fs, nil
}

type (
	// Data from parsed source code and will be used in code generation.
	data struct {
		fields    fields
		pkgName   string
		sliceName string
	}
	fields []field

	// Struct field from entity in source code.
	field struct {
		Name string // field name
		Type string // field type like string,int64...
	}
)

// Parse expression.
func parseExpr(x ast.Expr) string {
	switch tt := x.(type) {
	case *ast.Ident:
		return parseIdent(tt)
	case *ast.StarExpr:
		return parseStarExpr(tt)
	case *ast.FuncType:
		return parseFuncType(tt)
	case *ast.Ellipsis:
		return parseEllipsis(tt)
	}
	log.Println("parse error: unknown type")
	return "any"
}

// Parse identifier.
func parseIdent(x *ast.Ident) string {
	return x.Name
}

// Parse star expression.
func parseStarExpr(x *ast.StarExpr) string {
	switch tt := x.X.(type) {
	case *ast.Ident:
		return "*" + tt.Name
	default:
		log.Println("parseStarExpr: parse error: unknown type")
		return "any"
	}
}

// Parse function type.
func parseFuncType(x *ast.FuncType) string {
	params := func() string {
		if x != nil && x.Params != nil && x.Params.List != nil {
			return newFields(x.Params.List).display()
		}
		return ""
	}()
	results := func() string {
		if x != nil && x.Results != nil && x.Results.List != nil {
			return newFields(x.Results.List).display()
		}
		return ""
	}()
	return fmt.Sprintf("func(%s) (%s)", params, results)
}

// Parse ellipsis.
func parseEllipsis(x *ast.Ellipsis) string {
	str := parseExpr(x.Elt)
	switch str {
	case "any":
		log.Println("parseEllipsis: parse error: unknown type")
		return str
	default:
		return "..." + str
	}
}

// Constructor for fields.
func newFields(raws []*ast.Field) fields {
	fs := make(fields, 0, len(raws))
	for _, raw := range raws {
		fs = append(fs, newField(raw))
	}
	return fs
}

// Get field name safely.
func safeGetNameFromField(raw *ast.Field) string {
	if len(raw.Names) == 0 {
		return ""
	}
	return raw.Names[0].Name
}

// ----------------
// Field
// ----------------

// Constructor for field.
func newField(raw *ast.Field) field {
	name := safeGetNameFromField(raw)
	ty := parseExpr(raw.Type)
	return field{
		Name: name,
		Type: ty,
	}
}

// Display fields.
func (fs fields) display() string {
	if len(fs) == 0 {
		return ""
	}
	var pairs []string
	for _, f := range fs {
		pairs = append(pairs, f.display())
	}
	return strings.Join(pairs, ", ")
}

// Display field.
func (f field) display() string {
	if f.Name == "" {
		return f.Type
	}
	return fmt.Sprintf("%s %s", f.Name, f.Type)
}

// Exclude fields by name.
// func (fs fields) exclude(targets []string) fields {
// 	return slices.DeleteFunc(fs, func(f field) bool {
// 		return slices.Contains(targets, f.Name)
// 	})
// }

// excludeNoncomparable
func (fs fields) excludeUncomparable() fields {
	return slices.DeleteFunc(fs, func(f field) bool {
		return f.isUncomparable()
	})
}

// isComparable
func (f field) isUncomparable() bool {
	NGs := []string{
		"func(",
		"map[",
		"chan ",
		"interface{",
	}
	for _, ng := range NGs {
		if strings.Contains(f.Type, ng) {
			return true
		}
	}
	return false
}
