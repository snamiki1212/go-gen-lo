package internal

import (
	"fmt"
	"regexp"
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

	// Rename flag that is mapping field name to accessor name
	RawRename      []string
	renameInfoList []renameInfo

	// Raw entity
	RawEntity string

	// Excluded lo methods
	RawLoMethodsToExclude []string
	LoMethodsToExclude    []regexp.Regexp

	// Included lo methods
	RawLoMethodsToInclude []string
	LoMethodsToInclude    []regexp.Regexp
}

var Args = Arguments{}

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

	// Load exclude
	if err := a.loadLoMethodsToExclude(); err != nil {
		container = append(container, fmt.Errorf("load exclude error: %w", err))
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

type renameInfo struct {
	pattern *regexp.Regexp
	templ   string
}

// Rename
func (a Arguments) Rename(src string) (string, bool) {
	for _, info := range a.renameInfoList {
		if info.pattern.MatchString(src) {
			return info.pattern.ReplaceAllString(src, info.templ), true
		}
	}
	return src, false
}

// Exclude with regex
func (a Arguments) IsExcluded(method string) bool {
	for _, re := range a.LoMethodsToExclude {
		if re.MatchString(method) {
			return true
		}
	}
	return false
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

		pattern, templ := pair[0], pair[1]

		// NOTE: workaround for escape character
		// dollar might not be supported for spf13/cobra using spf13/pflag.
		// so, replace all backslash to dollar.
		templ = strings.ReplaceAll(templ, "\\", "$")

		a.renameInfoList = append(a.renameInfoList, renameInfo{
			pattern: regexp.MustCompile(pattern),
			templ:   templ,
		})
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

// load LoMethodsToExclude flag
func (a *Arguments) loadLoMethodsToExclude() error {
	container := make([]error, 0)
	for _, raw := range a.RawLoMethodsToExclude {
		re, err := regexp.Compile(raw)
		if err != nil {
			container = append(container, fmt.Errorf("invalid regex: %s", raw))
			continue
		}
		a.LoMethodsToExclude = append(a.LoMethodsToExclude, *re)
	}
	if len(container) != 0 {
		return fmt.Errorf("%v", container)
	}
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
