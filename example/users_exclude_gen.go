// Code generated by "go-gen-lo"; DO NOT EDIT.
// Based on information from https://github.com/snamiki1212/go-gen-lo

package main

import "github.com/samber/lo"

/************************************************
 ** lo basic methods
 ************************************************/

// Map
func (xs UserExclude) Map(iteratee func(item User, index int) User) UserExclude {
	return lo.Map(xs, iteratee)
}

// FilterReject
func (xs UserExclude) FilterReject(predicate func(User, int) bool) (UserExclude, UserExclude) {
	return lo.FilterReject(xs, predicate)
}

// Find
func (xs UserExclude) Find(predicate func(User) bool) (User, bool) {
	return lo.Find(xs, predicate)
}

/************************************************
 ** lo extended methods
 ************************************************/

// -- GroupBy ------------------------------------
// GroupByUserID
func (xs UserExclude) GroupByUserID() map[string]UserExclude {
	return lo.GroupBy(xs, func(entity User) string {
		return entity.UserID
	})
}

// GroupByInt
func (xs UserExclude) GroupByInt() map[int]UserExclude {
	return lo.GroupBy(xs, func(entity User) int {
		return entity.Int
	})
}

// GroupByIntPtr
func (xs UserExclude) GroupByIntPtr() map[*int]UserExclude {
	return lo.GroupBy(xs, func(entity User) *int {
		return entity.IntPtr
	})
}

// GroupByBool
func (xs UserExclude) GroupByBool() map[bool]UserExclude {
	return lo.GroupBy(xs, func(entity User) bool {
		return entity.Bool
	})
}

// GroupByBoolPtr
func (xs UserExclude) GroupByBoolPtr() map[*bool]UserExclude {
	return lo.GroupBy(xs, func(entity User) *bool {
		return entity.BoolPtr
	})
}

// GroupByStr
func (xs UserExclude) GroupByStr() map[string]UserExclude {
	return lo.GroupBy(xs, func(entity User) string {
		return entity.Str
	})
}

// GroupByStrPtr
func (xs UserExclude) GroupByStrPtr() map[*string]UserExclude {
	return lo.GroupBy(xs, func(entity User) *string {
		return entity.StrPtr
	})
}

// GroupByStruct0
func (xs UserExclude) GroupByStruct0() map[DefinedStruct0]UserExclude {
	return lo.GroupBy(xs, func(entity User) DefinedStruct0 {
		return entity.Struct0
	})
}

// GroupByStructPtr0
func (xs UserExclude) GroupByStructPtr0() map[*DefinedStruct0]UserExclude {
	return lo.GroupBy(xs, func(entity User) *DefinedStruct0 {
		return entity.StructPtr0
	})
}

// GroupByStruct1
func (xs UserExclude) GroupByStruct1() map[DefinedStruct1]UserExclude {
	return lo.GroupBy(xs, func(entity User) DefinedStruct1 {
		return entity.Struct1
	})
}

// GroupByStructPtr1
func (xs UserExclude) GroupByStructPtr1() map[*DefinedStruct1]UserExclude {
	return lo.GroupBy(xs, func(entity User) *DefinedStruct1 {
		return entity.StructPtr1
	})
}

// GroupByChanSend0
func (xs UserExclude) GroupByChanSend0() map[chan<- int]UserExclude {
	return lo.GroupBy(xs, func(entity User) chan<- int {
		return entity.ChanSend0
	})
}

// GroupByChanSendPtr0
func (xs UserExclude) GroupByChanSendPtr0() map[*chan<- int]UserExclude {
	return lo.GroupBy(xs, func(entity User) *chan<- int {
		return entity.ChanSendPtr0
	})
}

// -- FilterReject ------------------------------------
// FilterRejectByUserID
func (xs UserExclude) FilterRejectByUserID(_UserID string) (kept UserExclude, rejected UserExclude) {
	return lo.FilterReject(xs, func(entity User, index int) bool {
		return entity.UserID == _UserID
	})
}

// FilterRejectByInt
func (xs UserExclude) FilterRejectByInt(_Int int) (kept UserExclude, rejected UserExclude) {
	return lo.FilterReject(xs, func(entity User, index int) bool {
		return entity.Int == _Int
	})
}

// FilterRejectByIntPtr
func (xs UserExclude) FilterRejectByIntPtr(_IntPtr *int) (kept UserExclude, rejected UserExclude) {
	return lo.FilterReject(xs, func(entity User, index int) bool {
		return entity.IntPtr == _IntPtr
	})
}

// FilterRejectByBool
func (xs UserExclude) FilterRejectByBool(_Bool bool) (kept UserExclude, rejected UserExclude) {
	return lo.FilterReject(xs, func(entity User, index int) bool {
		return entity.Bool == _Bool
	})
}

// FilterRejectByBoolPtr
func (xs UserExclude) FilterRejectByBoolPtr(_BoolPtr *bool) (kept UserExclude, rejected UserExclude) {
	return lo.FilterReject(xs, func(entity User, index int) bool {
		return entity.BoolPtr == _BoolPtr
	})
}

// FilterRejectByStr
func (xs UserExclude) FilterRejectByStr(_Str string) (kept UserExclude, rejected UserExclude) {
	return lo.FilterReject(xs, func(entity User, index int) bool {
		return entity.Str == _Str
	})
}

// FilterRejectByStrPtr
func (xs UserExclude) FilterRejectByStrPtr(_StrPtr *string) (kept UserExclude, rejected UserExclude) {
	return lo.FilterReject(xs, func(entity User, index int) bool {
		return entity.StrPtr == _StrPtr
	})
}

// FilterRejectByStruct0
func (xs UserExclude) FilterRejectByStruct0(_Struct0 DefinedStruct0) (kept UserExclude, rejected UserExclude) {
	return lo.FilterReject(xs, func(entity User, index int) bool {
		return entity.Struct0 == _Struct0
	})
}

// FilterRejectByStructPtr0
func (xs UserExclude) FilterRejectByStructPtr0(_StructPtr0 *DefinedStruct0) (kept UserExclude, rejected UserExclude) {
	return lo.FilterReject(xs, func(entity User, index int) bool {
		return entity.StructPtr0 == _StructPtr0
	})
}

// FilterRejectByStruct1
func (xs UserExclude) FilterRejectByStruct1(_Struct1 DefinedStruct1) (kept UserExclude, rejected UserExclude) {
	return lo.FilterReject(xs, func(entity User, index int) bool {
		return entity.Struct1 == _Struct1
	})
}

// FilterRejectByStructPtr1
func (xs UserExclude) FilterRejectByStructPtr1(_StructPtr1 *DefinedStruct1) (kept UserExclude, rejected UserExclude) {
	return lo.FilterReject(xs, func(entity User, index int) bool {
		return entity.StructPtr1 == _StructPtr1
	})
}

// FilterRejectByChanSend0
func (xs UserExclude) FilterRejectByChanSend0(_ChanSend0 chan<- int) (kept UserExclude, rejected UserExclude) {
	return lo.FilterReject(xs, func(entity User, index int) bool {
		return entity.ChanSend0 == _ChanSend0
	})
}

// FilterRejectByChanSendPtr0
func (xs UserExclude) FilterRejectByChanSendPtr0(_ChanSendPtr0 *chan<- int) (kept UserExclude, rejected UserExclude) {
	return lo.FilterReject(xs, func(entity User, index int) bool {
		return entity.ChanSendPtr0 == _ChanSendPtr0
	})
}

// -- Find ------------------------------------
// FindByUserID
func (xs UserExclude) FindByUserID(_UserID string) (User, bool) {
	return lo.Find(xs, func(entity User) bool {
		return entity.UserID == _UserID
	})
}

// FindByInt
func (xs UserExclude) FindByInt(_Int int) (User, bool) {
	return lo.Find(xs, func(entity User) bool {
		return entity.Int == _Int
	})
}

// FindByIntPtr
func (xs UserExclude) FindByIntPtr(_IntPtr *int) (User, bool) {
	return lo.Find(xs, func(entity User) bool {
		return entity.IntPtr == _IntPtr
	})
}

// FindByBool
func (xs UserExclude) FindByBool(_Bool bool) (User, bool) {
	return lo.Find(xs, func(entity User) bool {
		return entity.Bool == _Bool
	})
}

// FindByBoolPtr
func (xs UserExclude) FindByBoolPtr(_BoolPtr *bool) (User, bool) {
	return lo.Find(xs, func(entity User) bool {
		return entity.BoolPtr == _BoolPtr
	})
}

// FindByStr
func (xs UserExclude) FindByStr(_Str string) (User, bool) {
	return lo.Find(xs, func(entity User) bool {
		return entity.Str == _Str
	})
}

// FindByStrPtr
func (xs UserExclude) FindByStrPtr(_StrPtr *string) (User, bool) {
	return lo.Find(xs, func(entity User) bool {
		return entity.StrPtr == _StrPtr
	})
}

// FindByStruct0
func (xs UserExclude) FindByStruct0(_Struct0 DefinedStruct0) (User, bool) {
	return lo.Find(xs, func(entity User) bool {
		return entity.Struct0 == _Struct0
	})
}

// FindByStructPtr0
func (xs UserExclude) FindByStructPtr0(_StructPtr0 *DefinedStruct0) (User, bool) {
	return lo.Find(xs, func(entity User) bool {
		return entity.StructPtr0 == _StructPtr0
	})
}

// FindByStruct1
func (xs UserExclude) FindByStruct1(_Struct1 DefinedStruct1) (User, bool) {
	return lo.Find(xs, func(entity User) bool {
		return entity.Struct1 == _Struct1
	})
}

// FindByStructPtr1
func (xs UserExclude) FindByStructPtr1(_StructPtr1 *DefinedStruct1) (User, bool) {
	return lo.Find(xs, func(entity User) bool {
		return entity.StructPtr1 == _StructPtr1
	})
}

// FindByChanSend0
func (xs UserExclude) FindByChanSend0(_ChanSend0 chan<- int) (User, bool) {
	return lo.Find(xs, func(entity User) bool {
		return entity.ChanSend0 == _ChanSend0
	})
}

// FindByChanSendPtr0
func (xs UserExclude) FindByChanSendPtr0(_ChanSendPtr0 *chan<- int) (User, bool) {
	return lo.Find(xs, func(entity User) bool {
		return entity.ChanSendPtr0 == _ChanSendPtr0
	})
}
