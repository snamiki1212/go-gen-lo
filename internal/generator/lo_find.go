package generator

type LoFind struct{}

func NewLoFind() LoFind { return LoFind{} }

func (l LoFind) Skip(_ loStdTemplateMapper) bool { return false }

func (l LoFind) StdName() string { return "Find" }

func (l LoFind) ExtendName() string { return "FindBy" }

func (l LoFind) StdTemplate() (string, bool) {
	return `
// {{ .Method }}
func (xs {{ .Slice }}) {{ .Method }}(predicate func({{ .Entity }}) bool) ({{ .Entity }}, bool) {
	return lo.Find(xs, predicate)
}
`, true
}

func (l LoFind) ExtendTemplate() (string, bool) {
	return `
// {{ .Method }}
func (xs {{ .Slice }}) {{ .Method }}(_{{ .Field }} {{ .Type }}) ({{ .Entity }}, bool) {
	return lo.Find(xs, func(entity {{ .Entity }}) bool {
		return entity.{{ .Field }} == _{{ .Field }}
	})
}
`, true
}
