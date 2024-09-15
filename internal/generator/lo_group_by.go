package generator

type LoGroupBy struct{}

func NewLoGroupBy() LoGroupBy { return LoGroupBy{} }

func (l LoGroupBy) Name() string { return "GroupBy" }

func (l LoGroupBy) ExtendMethodName() (string, bool) { return "", false }

func (l LoGroupBy) StdTemplate() (string, bool) {
	return ``, false
}

func (l LoGroupBy) ExtendTemplate() (string, bool) {
	return `
// {{ .Method }}{{ .Field }}
func (xs {{ .Slice }}) {{ .Method }}{{ .Field }}() map[{{ .Type }}]{{ .Slice }} {
	return lo.GroupBy(xs, func(entity {{ .Entity }}) {{ .Type }} {
		return entity.{{ .Field }}
	})
}
`, true
}
