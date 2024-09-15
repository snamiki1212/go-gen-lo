package internal

type LoFilter struct{}

func NewLoFilter() LoFilter { return LoFilter{} }

func (l LoFilter) Kind() string { return "Filter" }

func (l LoFilter) StdTemplate() (string, bool) {
	return `
// Filter
func (xs {{ .Slice }}) Filter(predicate func(item {{ .Entity }}, index int) bool) {{ .Slice }} {
	return lo.Filter(xs, predicate)
}
`, true
}

func (l LoFilter) ExtendTemplate() (string, bool) {
	return `
// FilterBy{{ .Field }}
func (xs {{ .Slice }}) FilterBy{{ .Field }}(field {{ .Type }}) {{ .Slice }} {
	return lo.Filter(xs, func(entity {{ .Entity }}, index int) bool {
		return entity.{{ .Field }} == field
	})
}
`, true
}
