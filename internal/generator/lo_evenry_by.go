package generator

type LoEveryBy struct{}

func NewLoEveryBy() LoEveryBy { return LoEveryBy{} }

func (l LoEveryBy) Name() string { return "EveryBy" }

func (l LoEveryBy) ExtendMethodName() (string, bool) { return "", false }

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
