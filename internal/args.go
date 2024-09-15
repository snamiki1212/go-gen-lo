package internal

import (
	"fmt"
	"strings"
)

type Arguments struct {
	// Target Entity name
	Entity string

	// Entitty is pointer or not
	IsPtrEntity bool

	// Target Slice name
	Slice string

	// Input file name
	Input string

	// Output file name
	Output string

	// Mapping field name to accessor name
	RenameMap map[string]string // key: lo method name, value: generated method name.

	// Raw rename strings
	RawRename []string

	// Raw entity
	RawEntity string

	// Excluded lo methods
	LoMethodsToExclude []string
}

var Args = Arguments{
	RenameMap: map[string]string{},
}

func (a *Arguments) Load() error {
	// Load rename
	if err := a.loadRename(a.RawRename); err != nil {
		return fmt.Errorf("load accessor error: %w", err)
	}

	// Load entity
	if err := a.loadEntity(a.RawEntity); err != nil {
		return fmt.Errorf("load entity error: %w", err)
	}

	return nil
}

func (a Arguments) DisplayEntity() string {
	if a.IsPtrEntity {
		return "*" + a.Entity
	}
	return a.Entity
}

func (a *Arguments) loadRename(as []string) error {
	container := make([]error, 0)
	for _, ac := range as {
		pair := strings.Split(ac, ":")
		if len(pair) != 2 {
			container = append(container, fmt.Errorf("invalid rename: %s", ac))
			continue
		}
		src, dst := pair[0], pair[1]
		a.RenameMap[src] = dst
	}
	if len(container) != 0 {
		return fmt.Errorf("%v", container)
	}
	return nil
}

func (a *Arguments) loadEntity(e string) error {
	a.IsPtrEntity = strings.HasPrefix(e, "*")
	if a.IsPtrEntity {
		e = strings.TrimPrefix(e, "*")
	}
	a.Entity = e
	return nil
}
