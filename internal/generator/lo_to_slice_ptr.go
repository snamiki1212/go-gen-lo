package generator

type LoToSlicePtr struct{}

func NewLoToSlicePtr() LoToSlicePtr { return LoToSlicePtr{} }

func (l LoToSlicePtr) Skip(_ loStdTemplateMapper) bool { return false }

func (l LoToSlicePtr) StdName() string { return "ToSlicePtr" }

func (l LoToSlicePtr) ExtendName() string { return "" }

func (l LoToSlicePtr) StdTemplate() (string, bool) {
	return `
// {{ .Method }}
func (xs {{ .Slice }}) {{ .Method }}() []*{{ .Entity }} {
	return lo.ToSlicePtr(xs)
}
`, true
}

func (l LoToSlicePtr) ExtendTemplate() (string, bool) {
	return ``, false
}
