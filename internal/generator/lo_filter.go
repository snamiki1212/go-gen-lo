package generator

type LoFilter struct{}

func NewLoFilter() LoFilter { return LoFilter{} }

func (l LoFilter) Kind() string { return "Filter" }

func (l LoFilter) ExtendMethodName() (string, bool) { return "FilterBy", true }

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
// {{ .Method }}{{ .Field }}
func (xs {{ .Slice }}) {{ .Method }}{{ .Field }}(_{{ .Field }} {{ .Type }}) {{ .Slice }} {
	return lo.Filter(xs, func(entity {{ .Entity }}, index int) bool {
		return entity.{{ .Field }} == _{{ .Field }}
	})
}
`, true
}
