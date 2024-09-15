package generator

type LoMap struct{}

func NewLoMap() LoMap { return LoMap{} }

func (l LoMap) Kind() string { return "Map" }

func (l LoMap) StdTemplate() (string, bool) {
	return `
// {{ .Method }}
func (xs {{ .Slice }}) {{ .Method }}(iteratee func(item {{ .Entity }}, index int) {{ .Entity }}) {{ .Slice }} {
	return lo.Map(xs, iteratee)
}
`, true
}

func (l LoMap) ExtendTemplate() (string, bool) {
	return ``, false
}
