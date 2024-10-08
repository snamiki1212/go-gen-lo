package generator

type LoFilterReject struct{}

func NewLoFilterReject() LoFilterReject { return LoFilterReject{} }

func (l LoFilterReject) Skip(_ loStdTemplateMapper) bool { return false }

func (l LoFilterReject) StdName() string { return "FilterReject" }

func (l LoFilterReject) ExtendName() string { return "FilterRejectBy" }

func (l LoFilterReject) StdTemplate() (string, bool) {
	return `
// {{ .Method }}
func (xs {{ .Slice }}) {{ .Method }}(predicate func({{ .Entity }}, int) bool) ({{ .Slice }}, {{ .Slice }}) {
	return lo.FilterReject(xs, predicate)
}
`, true
}

func (l LoFilterReject) ExtendTemplate() (string, bool) {
	return `
// {{ .Method }}
func (xs {{ .Slice }}) {{ .Method }}(_{{ .Field }} {{ .Type }}) (kept {{ .Slice }}, rejected {{ .Slice }}) {
	return lo.FilterReject(xs, func(entity {{ .Entity }}, index int) bool {
		return entity.{{ .Field }} == _{{ .Field }}
	})
}
`, true
}
