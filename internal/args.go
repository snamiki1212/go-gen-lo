package internal

import (
	"fmt"
	"regexp"
	"slices"
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

	// Included lo methods
	RawLoMethodsToInclude []string
	LoMethodsToInclude    []regexp.Regexp
}

var Args = Arguments{
	RenameMap: map[string]string{},
}

func (a *Arguments) Load() error {
	container := make([]error, 0)

	// Load rename
	if err := a.loadRename(a.RawRename); err != nil {
		container = append(container, fmt.Errorf("load rename error: %w", err))
	}

	// Load entity
	if err := a.loadEntity(a.RawEntity); err != nil {
		container = append(container, fmt.Errorf("load entity error: %w", err))
	}

	// Load include
	if err := a.loadLoMethodsToInclude(); err != nil {
		container = append(container, fmt.Errorf("load include error: %w", err))
	}

	if len(container) != 0 {
		return fmt.Errorf("%v", container)
	}

	return nil
}

func (a Arguments) DisplayEntity() string {
	if a.IsPtrEntity {
		return "*" + a.Entity
	}
	return a.Entity
}

// Exclude
// TODO: support regex
func (a Arguments) IsExcluded(method string) bool {
	return slices.Contains(a.LoMethodsToExclude, method)
}

// Include with regex
func (a Arguments) IsIncluded(method string) bool {
	if len(a.LoMethodsToInclude) == 0 { // default: include all
		return true
	}
	for _, re := range a.LoMethodsToInclude {
		if re.MatchString(method) {
			return true
		}
	}
	return false
}

// load Rename flag
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

// load Entity flag
func (a *Arguments) loadEntity(e string) error {
	a.IsPtrEntity = strings.HasPrefix(e, "*")
	if a.IsPtrEntity {
		e = strings.TrimPrefix(e, "*")
	}
	a.Entity = e
	return nil
}

// load LoMethodsToInclude flag
func (a *Arguments) loadLoMethodsToInclude() error {
	container := make([]error, 0)
	for _, raw := range a.RawLoMethodsToInclude {
		re, err := regexp.Compile(raw)
		if err != nil {
			container = append(container, fmt.Errorf("invalid regex: %s", raw))
			continue
		}
		a.LoMethodsToInclude = append(a.LoMethodsToInclude, *re)
	}
	if len(container) != 0 {
		return fmt.Errorf("%v", container)
	}
	return nil
}
