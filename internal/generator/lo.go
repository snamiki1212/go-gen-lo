package generator

type Lo interface {
	// Function name of lo
	StdName() string

	// Method name of lo with extend. Return false if no need.
	ExtendName() string

	// Template for lo. Return false if not implemented.
	StdTemplate() (string, bool)

	// Template for lo with extend. Return false if not implemented.
	ExtendTemplate() (string, bool)

	// Skip generate lo method if true.
	Skip(params loStdTemplateMapper) bool
}

// Replace variable from key to value in template.
type loStdTemplateMapper struct {
	Slice      string // Slice name for target struct (ex. Users).
	Entity     string // Entity name with pointer for target struct (ex. User / *User).
	EntityName string // Entity name without pointer for target struct (ex. User ).
	Method     string // Method name of struct (ex. Filter).
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
		NewLoKeyBy(),
		NewLoGroupBy(),
		NewLoUniqBy(),
		NewLoFilterReject(),
		NewLoFind(),
		NewLoContainsBy(),
		NewLoEveryBy(),
		NewLoSomeBy(),
		NewLoToSlicePtr(),
		NewLoFromSlicePtr(),
	}
}
