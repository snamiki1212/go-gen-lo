package generator

type LoEveryBy struct{}

func NewLoEveryBy() LoEveryBy { return LoEveryBy{} }

func (l LoEveryBy) Skip(_ loStdTemplateMapper) bool { return false }

func (l LoEveryBy) StdName() string { return "EveryBy" }

func (l LoEveryBy) ExtendName() string { return "EveryBy" }

func (l LoEveryBy) StdTemplate() (string, bool) {
	return `
// {{ .Method }}
func (xs {{ .Slice }}) {{ .Method }}(predicate func(item {{ .Entity }}) bool) bool {
	return lo.EveryBy(xs, predicate)
}
`, true
}

func (l LoEveryBy) ExtendTemplate() (string, bool) {
	return ``, false
}
