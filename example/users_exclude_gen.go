// Code generated by "go-gen-lo"; DO NOT EDIT.
// Based on information from https://github.com/snamiki1212/go-gen-lo

package main

import "github.com/samber/lo"

// Map
func (xs UserExclude) Map(iteratee func(item User, index int) User) UserExclude {
	return lo.Map(xs, iteratee)
}

// Find
func (xs UserExclude) Find(predicate func(User) bool) (User, bool) {
	return lo.Find(xs, predicate)
}

// FindByUserID
func (xs UserExclude) FindByUserID(field string) (User, bool) {
	return lo.Find(xs, func(entity User) bool {
		return entity.UserID == field
	})
}

// FindByInt
func (xs UserExclude) FindByInt(field int) (User, bool) {
	return lo.Find(xs, func(entity User) bool {
		return entity.Int == field
	})
}

// FindByIntPtr
func (xs UserExclude) FindByIntPtr(field *int) (User, bool) {
	return lo.Find(xs, func(entity User) bool {
		return entity.IntPtr == field
	})
}

// FindByBool
func (xs UserExclude) FindByBool(field bool) (User, bool) {
	return lo.Find(xs, func(entity User) bool {
		return entity.Bool == field
	})
}

// FindByBoolPtr
func (xs UserExclude) FindByBoolPtr(field *bool) (User, bool) {
	return lo.Find(xs, func(entity User) bool {
		return entity.BoolPtr == field
	})
}

// FindByStr
func (xs UserExclude) FindByStr(field string) (User, bool) {
	return lo.Find(xs, func(entity User) bool {
		return entity.Str == field
	})
}

// FindByStrPtr
func (xs UserExclude) FindByStrPtr(field *string) (User, bool) {
	return lo.Find(xs, func(entity User) bool {
		return entity.StrPtr == field
	})
}

// FindByStruct0
func (xs UserExclude) FindByStruct0(field DefinedStruct0) (User, bool) {
	return lo.Find(xs, func(entity User) bool {
		return entity.Struct0 == field
	})
}

// FindByStructPtr0
func (xs UserExclude) FindByStructPtr0(field *DefinedStruct0) (User, bool) {
	return lo.Find(xs, func(entity User) bool {
		return entity.StructPtr0 == field
	})
}

// FindByStruct1
func (xs UserExclude) FindByStruct1(field DefinedStruct1) (User, bool) {
	return lo.Find(xs, func(entity User) bool {
		return entity.Struct1 == field
	})
}

// FindByStructPtr1
func (xs UserExclude) FindByStructPtr1(field *DefinedStruct1) (User, bool) {
	return lo.Find(xs, func(entity User) bool {
		return entity.StructPtr1 == field
	})
}

// FindByChanSend0
func (xs UserExclude) FindByChanSend0(field chan<- int) (User, bool) {
	return lo.Find(xs, func(entity User) bool {
		return entity.ChanSend0 == field
	})
}

// FindByChanSendPtr0
func (xs UserExclude) FindByChanSendPtr0(field *chan<- int) (User, bool) {
	return lo.Find(xs, func(entity User) bool {
		return entity.ChanSendPtr0 == field
	})
}
