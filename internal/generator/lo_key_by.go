package generator

type LoKeyBy struct{}

func NewLoKeyBy() LoKeyBy { return LoKeyBy{} }

func (l LoKeyBy) StdName() string { return "KeyBy" }

func (l LoKeyBy) ExtendName() (string, bool) { return "", false }

func (l LoKeyBy) StdTemplate() (string, bool) {
	return ``, false
}

func (l LoKeyBy) ExtendTemplate() (string, bool) {
	return `
// {{ .Method }}{{ .Field }}
func (xs {{ .Slice }}) {{ .Method }}{{ .Field }}() map[{{ .Type }}]{{ .Entity }} {
	return lo.KeyBy(xs, func(entity {{ .Entity }}) {{ .Type }} {
		return entity.{{ .Field }}
	})
}
`, true
}
