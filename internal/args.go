package internal

import (
	"fmt"
	"strings"
)

// TODO:
// cmd --entity=User  --slice=UserVals --lo=Filter,Map --rename=Map:Loop    --no-extend
// cmd --entity=*User --slice=UserPtrs --lo=Filter,Map --rename=Map:Loop    --no-extend
type arguments struct {
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

	// Method name of lo to generate
	// lo []string

	// Mapping field name to accessor name
	Rename map[string]string // key: lo method name, value: generated method name.

	// noExtend []string
}

var ArgRename []string
var ArgEntity string

var Args = arguments{
	Rename: map[string]string{},
}

func (a arguments) DisplayEntity() string {
	if a.IsPtrEntity {
		return "*" + a.Entity
	}
	return a.Entity
}

func (a *arguments) loadRename(as []string) error {
	container := make([]error, 0)
	for _, ac := range as {
		pair := strings.Split(ac, ":")
		if len(pair) != 2 {
			container = append(container, fmt.Errorf("invalid rename: %s", ac))
			continue
		}
		src, dst := pair[0], pair[1]
		a.Rename[src] = dst
	}
	if len(container) != 0 {
		return fmt.Errorf("%v", container)
	}
	return nil
}

func (a *arguments) loadEntity(e string) error {
	a.IsPtrEntity = strings.HasPrefix(e, "*")
	if a.IsPtrEntity {
		e = strings.TrimPrefix(e, "*")
	}
	a.Entity = e
	return nil
}

func (a *arguments) Load() error {
	// Load arguments
	if err := a.loadRename(ArgRename); err != nil {
		return fmt.Errorf("load accessor error: %w", err)
	}

	// Load entity
	if err := a.loadEntity(ArgEntity); err != nil {
		return fmt.Errorf("load entity error: %w", err)
	}
	return nil
}
