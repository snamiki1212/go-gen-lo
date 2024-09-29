package generator

type LoGroupBy struct{}

func NewLoGroupBy() LoGroupBy { return LoGroupBy{} }

func (l LoGroupBy) Skip(_ loStdTemplateMapper) bool { return false }

func (l LoGroupBy) StdName() string { return "GroupBy" }

func (l LoGroupBy) ExtendName() string { return "GroupBy" }

func (l LoGroupBy) StdTemplate() (string, bool) {
	return ``, false
}

func (l LoGroupBy) ExtendTemplate() (string, bool) {
	return `
// {{ .Method }}
func (xs {{ .Slice }}) {{ .Method }}() map[{{ .Type }}]{{ .Slice }} {
	return lo.GroupBy(xs, func(entity {{ .Entity }}) {{ .Type }} {
		return entity.{{ .Field }}
	})
}
`, true
}
