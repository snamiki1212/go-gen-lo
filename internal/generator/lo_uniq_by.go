package generator

type LoUniqBy struct{}

func NewLoUniqBy() LoUniqBy { return LoUniqBy{} }

func (l LoUniqBy) StdName() string { return "UniqBy" }

func (l LoUniqBy) ExtendName() (string, bool) { return "", false }

func (l LoUniqBy) StdTemplate() (string, bool) {
	return ``, false
}

func (l LoUniqBy) ExtendTemplate() (string, bool) {
	return `
// {{ .Method }}
func (xs {{ .Slice }}) {{ .Method }}() {{ .Slice }} {
	return lo.UniqBy(xs, func(entity {{ .Entity }}) {{ .Type }} {
		return entity.{{ .Field }}
	})
}
`, true
}
