package generator

type LoFilter struct{}

func NewLoFilter() LoFilter { return LoFilter{} }

func (l LoFilter) Skip(_ loStdTemplateMapper) bool { return false }

func (l LoFilter) StdName() string { return "Filter" }

func (l LoFilter) ExtendName() string { return "FilterBy" }

func (l LoFilter) StdTemplate() (string, bool) {
	return `
// {{ .Method }}
func (xs {{ .Slice }}) {{ .Method }}(predicate func({{ .Entity }}, int) bool) {{ .Slice }} {
	return lo.Filter(xs, predicate)
}
`, true
}

func (l LoFilter) ExtendTemplate() (string, bool) {
	return `
// {{ .Method }}
func (xs {{ .Slice }}) {{ .Method }}(_{{ .Field }} {{ .Type }}) {{ .Slice }} {
	return lo.Filter(xs, func(entity {{ .Entity }}, index int) bool {
		return entity.{{ .Field }} == _{{ .Field }}
	})
}
`, true
}
