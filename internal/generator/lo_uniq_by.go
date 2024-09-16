package generator

type LoUniqBy struct{}

func NewLoUniqBy() LoUniqBy { return LoUniqBy{} }

func (l LoUniqBy) Name() string { return "UniqBy" }

func (l LoUniqBy) ExtendMethodName() (string, bool) { return "", false }

func (l LoUniqBy) StdTemplate() (string, bool) {
	return ``, false
}

func (l LoUniqBy) ExtendTemplate() (string, bool) {
	return `
// {{ .Method }}{{ .Field }}
func (xs {{ .Slice }}) {{ .Method }}{{ .Field }}() {{ .Slice }} {
	return lo.UniqBy(xs, func(entity {{ .Entity }}) {{ .Type }} {
		return entity.{{ .Field }}
	})
}
`, true
}
