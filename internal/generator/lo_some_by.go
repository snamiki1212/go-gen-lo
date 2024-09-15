package generator

type LoSomeBy struct{}

func NewLoSomeBy() LoSomeBy { return LoSomeBy{} }

func (l LoSomeBy) Name() string { return "SomeBy" }

func (l LoSomeBy) ExtendMethodName() (string, bool) { return "", false }

func (l LoSomeBy) StdTemplate() (string, bool) {
	return `
// {{ .Method }}
func (xs {{ .Slice }}) {{ .Method }}(predicate func(item {{ .Entity }}) bool) bool {
	return lo.SomeBy(xs, predicate)
}
`, true
}

func (l LoSomeBy) ExtendTemplate() (string, bool) {
	return ``, false
}
