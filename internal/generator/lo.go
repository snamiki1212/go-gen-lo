package generator

type Lo interface {
	Kind() string
	StdTemplate() (string, bool)
	ExtendTemplate() (string, bool)
	ExtendMethodName() (string, bool)
}

// Replace variable from key to value in template.
type loStdTemplateMapper struct {
	Slice  string // Slice name for target struct (ex. Users).
	Entity string // Entity name for target struct (ex. User / *User).
	Method string // Method name of struct (ex. Filter).
}

// Replace variable from key to value in template.
type loExtendTemplateMapper struct {
	Slice  string // Slice name for target struct (ex. Users).
	Entity string // Entity name for target struct (ex. User / *User).
	Type   string // Type name of field (ex. string).
	Field  string // Field name of struct (ex. UserID).
	Method string // Method name of struct (ex. Filter).
}

func NewAllLoList() []Lo {
	return []Lo{
		NewLoFilter(),
		NewLoMap(),
		NewLoKeyBy(),
	}
}
