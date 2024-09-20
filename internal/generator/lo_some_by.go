package generator

type LoSomeBy struct{}

func NewLoSomeBy() LoSomeBy { return LoSomeBy{} }

func (l LoSomeBy) StdName() string { return "SomeBy" }

func (l LoSomeBy) ExtendName() string { return "SomeBy" }

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
