package generator

import "strings"

type LoFromSlicePtr struct{}

func NewLoFromSlicePtr() LoFromSlicePtr { return LoFromSlicePtr{} }

func (l LoFromSlicePtr) Skip(params loStdTemplateMapper) bool {
	isPtr := strings.HasPrefix(params.Entity, "*")
	return !isPtr // run for only value slices
}

func (l LoFromSlicePtr) StdName() string { return "FromSlicePtr" }

func (l LoFromSlicePtr) ExtendName() string { return "" }

func (l LoFromSlicePtr) StdTemplate() (string, bool) {
	return `
// {{ .Method }}
func (xs {{ .Slice }}) {{ .Method }}() []{{ .EntityName }} {
	return lo.FromSlicePtr(xs)
}
`, true
}

func (l LoFromSlicePtr) ExtendTemplate() (string, bool) {
	return ``, false
}
