package internal

import (
	"fmt"
	"regexp"
	"strings"
)

type Arguments struct {
	// Filename
	Input  string // for input
	Output string // for output

	// Target Slice name
	Slice string

	// Rename flag that is mapping field name to accessor name
	RawRename      []string
	renameInfoList []renameInfo

	// Entity
	RawEntity string // with maybe pointer (ex. *User or User)
	Entity    string // without pointer (ex. User)

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

// EntityIsPtr
func (a Arguments) EntityIsPtr() bool {
	return strings.HasPrefix(a.Entity, "*")
}

func (a Arguments) DisplayEntity() string {
	return a.RawEntity
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

		// NOTE: Workaround for placeholder.
		// Placeholder with dollar might not be supported for spf13/cobra using spf13/pflag,
		// so using backslash as placeholder keyword.
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
	a.Entity = strings.TrimPrefix(e, "*")
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
