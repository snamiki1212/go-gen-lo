package generator

type LoContainsBy struct{}

func NewLoContainsBy() LoContainsBy { return LoContainsBy{} }

func (l LoContainsBy) StdName() string { return "ContainsBy" }

func (l LoContainsBy) ExtendName() (string, bool) { return "", false }

func (l LoContainsBy) StdTemplate() (string, bool) {
	return `
// {{ .Method }}
func (xs {{ .Slice }}) {{ .Method }}(predicate func({{ .Entity }}) bool) bool {
	return lo.ContainsBy(xs, predicate)
}
`, true
}

func (l LoContainsBy) ExtendTemplate() (string, bool) {
	return `
// {{ .Method }}{{ .Field }}
func (xs {{ .Slice }}) {{ .Method }}{{ .Field }}(_{{ .Field }} {{ .Type }}) bool {
	return lo.ContainsBy(xs, func(entity {{ .Entity }}) bool {
		return entity.{{ .Field }} == _{{ .Field }}
	})
}
`, true
}
