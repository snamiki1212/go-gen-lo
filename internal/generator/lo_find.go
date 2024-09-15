package generator

type LoFind struct{}

func NewLoFind() LoFind { return LoFind{} }

func (l LoFind) Kind() string { return "Find" }

func (l LoFind) ExtendMethodName() (string, bool) { return "FindBy", true }

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
// {{ .Method }}{{ .Field }}
func (xs {{ .Slice }}) {{ .Method }}{{ .Field }}(_{{ .Field }} {{ .Type }}) ({{ .Entity }}, bool) {
	return lo.Find(xs, func(entity {{ .Entity }}) bool {
		return entity.{{ .Field }} == _{{ .Field }}
	})
}
`, true
}
