package internal

// Replace variable from key to value in template.
type loMethodTemplateMapper struct {
	Slice  string // Slice name for target struct (ex. Users).
	Entity string // Entity name for target struct (ex. User / *User).
}

// Replace variable from key to value in template.
type loMethodExtendTemplateMapper struct {
	Slice  string // Slice name for target struct (ex. Users).
	Entity string // Entity name for target struct (ex. User / *User).
	Type   string // Type name of field (ex. string).
	Field  string // Field name of struct (ex. UserID).
}
