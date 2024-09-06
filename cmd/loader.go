package cmd

import (
	"fmt"
	"strings"
)

// TODO:
// cmd --entity=User  --slice=UserVals --lo=Filter,Map --rename=Map:Loop    --no-extend
// cmd --entity=*User --slice=UserPtrs --lo=Filter,Map --rename=Map:Loop    --no-extend
type arguments struct {
	// Target entity name
	entity string

	// Entitty is pointer or not
	isPtrEntity bool

	// Target slice name
	slice string

	// Input file name
	input string

	// Output file name
	output string

	// Method name of lo to generate
	// lo []string

	// Mapping field name to accessor name
	rename map[string]string // key: lo method name, value: generated method name.

	// noExtend []string
}

var argRename []string
var argEntity string

var args = arguments{
	rename: map[string]string{},
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
		args.rename[src] = dst
	}
	if len(container) != 0 {
		return fmt.Errorf("%v", container)
	}
	return nil
}

func (a *arguments) loadEntity(e string) error {
	a.isPtrEntity = strings.HasPrefix(e, "*")
	if a.isPtrEntity {
		e = strings.TrimPrefix(e, "*")
	}
	a.entity = e
	return nil
}

func (a *arguments) load() error {
	// Load arguments
	if err := args.loadRename(argRename); err != nil {
		return fmt.Errorf("load accessor error: %w", err)
	}

	// Load entity
	if err := args.loadEntity(argEntity); err != nil {
		return fmt.Errorf("load entity error: %w", err)
	}
	return nil
}
