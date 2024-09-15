// Code generated by "go-gen-lo"; DO NOT EDIT.
// Based on information from https://github.com/snamiki1212/go-gen-lo

package main

import "github.com/samber/lo"

/************************************************
 ** lo basic methods
 ************************************************/

// LoFilter
func (xs UserRename) LoFilter(predicate func(User, int) bool) UserRename {
	return lo.Filter(xs, predicate)
}

// Loop
func (xs UserRename) Loop(iteratee func(item User, index int) User) UserRename {
	return lo.Map(xs, iteratee)
}

// FilterReject
func (xs UserRename) FilterReject(predicate func(User, int) bool) (UserRename, UserRename) {
	return lo.FilterReject(xs, predicate)
}

// Find
func (xs UserRename) Find(predicate func(User) bool) (User, bool) {
	return lo.Find(xs, predicate)
}

// ContainsBy
func (xs UserRename) ContainsBy(predicate func(User) bool) bool {
	return lo.ContainsBy(xs, predicate)
}

/************************************************
 ** lo extended methods
 ************************************************/

// -- Filter ------------------------------------

// LoFilterUserID
func (xs UserRename) LoFilterUserID(_UserID string) UserRename {
	return lo.Filter(xs, func(entity User, index int) bool {
		return entity.UserID == _UserID
	})
}

// LoFilterInt
func (xs UserRename) LoFilterInt(_Int int) UserRename {
	return lo.Filter(xs, func(entity User, index int) bool {
		return entity.Int == _Int
	})
}

// LoFilterIntPtr
func (xs UserRename) LoFilterIntPtr(_IntPtr *int) UserRename {
	return lo.Filter(xs, func(entity User, index int) bool {
		return entity.IntPtr == _IntPtr
	})
}

// LoFilterBool
func (xs UserRename) LoFilterBool(_Bool bool) UserRename {
	return lo.Filter(xs, func(entity User, index int) bool {
		return entity.Bool == _Bool
	})
}

// LoFilterBoolPtr
func (xs UserRename) LoFilterBoolPtr(_BoolPtr *bool) UserRename {
	return lo.Filter(xs, func(entity User, index int) bool {
		return entity.BoolPtr == _BoolPtr
	})
}

// LoFilterStr
func (xs UserRename) LoFilterStr(_Str string) UserRename {
	return lo.Filter(xs, func(entity User, index int) bool {
		return entity.Str == _Str
	})
}

// LoFilterStrPtr
func (xs UserRename) LoFilterStrPtr(_StrPtr *string) UserRename {
	return lo.Filter(xs, func(entity User, index int) bool {
		return entity.StrPtr == _StrPtr
	})
}

// LoFilterStruct0
func (xs UserRename) LoFilterStruct0(_Struct0 DefinedStruct0) UserRename {
	return lo.Filter(xs, func(entity User, index int) bool {
		return entity.Struct0 == _Struct0
	})
}

// LoFilterStructPtr0
func (xs UserRename) LoFilterStructPtr0(_StructPtr0 *DefinedStruct0) UserRename {
	return lo.Filter(xs, func(entity User, index int) bool {
		return entity.StructPtr0 == _StructPtr0
	})
}

// LoFilterStruct1
func (xs UserRename) LoFilterStruct1(_Struct1 DefinedStruct1) UserRename {
	return lo.Filter(xs, func(entity User, index int) bool {
		return entity.Struct1 == _Struct1
	})
}

// LoFilterStructPtr1
func (xs UserRename) LoFilterStructPtr1(_StructPtr1 *DefinedStruct1) UserRename {
	return lo.Filter(xs, func(entity User, index int) bool {
		return entity.StructPtr1 == _StructPtr1
	})
}

// LoFilterChanSend0
func (xs UserRename) LoFilterChanSend0(_ChanSend0 chan<- int) UserRename {
	return lo.Filter(xs, func(entity User, index int) bool {
		return entity.ChanSend0 == _ChanSend0
	})
}

// LoFilterChanSendPtr0
func (xs UserRename) LoFilterChanSendPtr0(_ChanSendPtr0 *chan<- int) UserRename {
	return lo.Filter(xs, func(entity User, index int) bool {
		return entity.ChanSendPtr0 == _ChanSendPtr0
	})
}

// -- KeyBy ------------------------------------

// LoKeyByUserID
func (xs UserRename) LoKeyByUserID() map[string]User {
	return lo.KeyBy(xs, func(entity User) string {
		return entity.UserID
	})
}

// LoKeyByInt
func (xs UserRename) LoKeyByInt() map[int]User {
	return lo.KeyBy(xs, func(entity User) int {
		return entity.Int
	})
}

// LoKeyByIntPtr
func (xs UserRename) LoKeyByIntPtr() map[*int]User {
	return lo.KeyBy(xs, func(entity User) *int {
		return entity.IntPtr
	})
}

// LoKeyByBool
func (xs UserRename) LoKeyByBool() map[bool]User {
	return lo.KeyBy(xs, func(entity User) bool {
		return entity.Bool
	})
}

// LoKeyByBoolPtr
func (xs UserRename) LoKeyByBoolPtr() map[*bool]User {
	return lo.KeyBy(xs, func(entity User) *bool {
		return entity.BoolPtr
	})
}

// LoKeyByStr
func (xs UserRename) LoKeyByStr() map[string]User {
	return lo.KeyBy(xs, func(entity User) string {
		return entity.Str
	})
}

// LoKeyByStrPtr
func (xs UserRename) LoKeyByStrPtr() map[*string]User {
	return lo.KeyBy(xs, func(entity User) *string {
		return entity.StrPtr
	})
}

// LoKeyByStruct0
func (xs UserRename) LoKeyByStruct0() map[DefinedStruct0]User {
	return lo.KeyBy(xs, func(entity User) DefinedStruct0 {
		return entity.Struct0
	})
}

// LoKeyByStructPtr0
func (xs UserRename) LoKeyByStructPtr0() map[*DefinedStruct0]User {
	return lo.KeyBy(xs, func(entity User) *DefinedStruct0 {
		return entity.StructPtr0
	})
}

// LoKeyByStruct1
func (xs UserRename) LoKeyByStruct1() map[DefinedStruct1]User {
	return lo.KeyBy(xs, func(entity User) DefinedStruct1 {
		return entity.Struct1
	})
}

// LoKeyByStructPtr1
func (xs UserRename) LoKeyByStructPtr1() map[*DefinedStruct1]User {
	return lo.KeyBy(xs, func(entity User) *DefinedStruct1 {
		return entity.StructPtr1
	})
}

// LoKeyByChanSend0
func (xs UserRename) LoKeyByChanSend0() map[chan<- int]User {
	return lo.KeyBy(xs, func(entity User) chan<- int {
		return entity.ChanSend0
	})
}

// LoKeyByChanSendPtr0
func (xs UserRename) LoKeyByChanSendPtr0() map[*chan<- int]User {
	return lo.KeyBy(xs, func(entity User) *chan<- int {
		return entity.ChanSendPtr0
	})
}

// -- GroupBy ------------------------------------

// GroupByUserID
func (xs UserRename) GroupByUserID() map[string]UserRename {
	return lo.GroupBy(xs, func(entity User) string {
		return entity.UserID
	})
}

// GroupByInt
func (xs UserRename) GroupByInt() map[int]UserRename {
	return lo.GroupBy(xs, func(entity User) int {
		return entity.Int
	})
}

// GroupByIntPtr
func (xs UserRename) GroupByIntPtr() map[*int]UserRename {
	return lo.GroupBy(xs, func(entity User) *int {
		return entity.IntPtr
	})
}

// GroupByBool
func (xs UserRename) GroupByBool() map[bool]UserRename {
	return lo.GroupBy(xs, func(entity User) bool {
		return entity.Bool
	})
}

// GroupByBoolPtr
func (xs UserRename) GroupByBoolPtr() map[*bool]UserRename {
	return lo.GroupBy(xs, func(entity User) *bool {
		return entity.BoolPtr
	})
}

// GroupByStr
func (xs UserRename) GroupByStr() map[string]UserRename {
	return lo.GroupBy(xs, func(entity User) string {
		return entity.Str
	})
}

// GroupByStrPtr
func (xs UserRename) GroupByStrPtr() map[*string]UserRename {
	return lo.GroupBy(xs, func(entity User) *string {
		return entity.StrPtr
	})
}

// GroupByStruct0
func (xs UserRename) GroupByStruct0() map[DefinedStruct0]UserRename {
	return lo.GroupBy(xs, func(entity User) DefinedStruct0 {
		return entity.Struct0
	})
}

// GroupByStructPtr0
func (xs UserRename) GroupByStructPtr0() map[*DefinedStruct0]UserRename {
	return lo.GroupBy(xs, func(entity User) *DefinedStruct0 {
		return entity.StructPtr0
	})
}

// GroupByStruct1
func (xs UserRename) GroupByStruct1() map[DefinedStruct1]UserRename {
	return lo.GroupBy(xs, func(entity User) DefinedStruct1 {
		return entity.Struct1
	})
}

// GroupByStructPtr1
func (xs UserRename) GroupByStructPtr1() map[*DefinedStruct1]UserRename {
	return lo.GroupBy(xs, func(entity User) *DefinedStruct1 {
		return entity.StructPtr1
	})
}

// GroupByChanSend0
func (xs UserRename) GroupByChanSend0() map[chan<- int]UserRename {
	return lo.GroupBy(xs, func(entity User) chan<- int {
		return entity.ChanSend0
	})
}

// GroupByChanSendPtr0
func (xs UserRename) GroupByChanSendPtr0() map[*chan<- int]UserRename {
	return lo.GroupBy(xs, func(entity User) *chan<- int {
		return entity.ChanSendPtr0
	})
}

// -- FilterReject ------------------------------------

// FilterRejectByUserID
func (xs UserRename) FilterRejectByUserID(_UserID string) (kept UserRename, rejected UserRename) {
	return lo.FilterReject(xs, func(entity User, index int) bool {
		return entity.UserID == _UserID
	})
}

// FilterRejectByInt
func (xs UserRename) FilterRejectByInt(_Int int) (kept UserRename, rejected UserRename) {
	return lo.FilterReject(xs, func(entity User, index int) bool {
		return entity.Int == _Int
	})
}

// FilterRejectByIntPtr
func (xs UserRename) FilterRejectByIntPtr(_IntPtr *int) (kept UserRename, rejected UserRename) {
	return lo.FilterReject(xs, func(entity User, index int) bool {
		return entity.IntPtr == _IntPtr
	})
}

// FilterRejectByBool
func (xs UserRename) FilterRejectByBool(_Bool bool) (kept UserRename, rejected UserRename) {
	return lo.FilterReject(xs, func(entity User, index int) bool {
		return entity.Bool == _Bool
	})
}

// FilterRejectByBoolPtr
func (xs UserRename) FilterRejectByBoolPtr(_BoolPtr *bool) (kept UserRename, rejected UserRename) {
	return lo.FilterReject(xs, func(entity User, index int) bool {
		return entity.BoolPtr == _BoolPtr
	})
}

// FilterRejectByStr
func (xs UserRename) FilterRejectByStr(_Str string) (kept UserRename, rejected UserRename) {
	return lo.FilterReject(xs, func(entity User, index int) bool {
		return entity.Str == _Str
	})
}

// FilterRejectByStrPtr
func (xs UserRename) FilterRejectByStrPtr(_StrPtr *string) (kept UserRename, rejected UserRename) {
	return lo.FilterReject(xs, func(entity User, index int) bool {
		return entity.StrPtr == _StrPtr
	})
}

// FilterRejectByStruct0
func (xs UserRename) FilterRejectByStruct0(_Struct0 DefinedStruct0) (kept UserRename, rejected UserRename) {
	return lo.FilterReject(xs, func(entity User, index int) bool {
		return entity.Struct0 == _Struct0
	})
}

// FilterRejectByStructPtr0
func (xs UserRename) FilterRejectByStructPtr0(_StructPtr0 *DefinedStruct0) (kept UserRename, rejected UserRename) {
	return lo.FilterReject(xs, func(entity User, index int) bool {
		return entity.StructPtr0 == _StructPtr0
	})
}

// FilterRejectByStruct1
func (xs UserRename) FilterRejectByStruct1(_Struct1 DefinedStruct1) (kept UserRename, rejected UserRename) {
	return lo.FilterReject(xs, func(entity User, index int) bool {
		return entity.Struct1 == _Struct1
	})
}

// FilterRejectByStructPtr1
func (xs UserRename) FilterRejectByStructPtr1(_StructPtr1 *DefinedStruct1) (kept UserRename, rejected UserRename) {
	return lo.FilterReject(xs, func(entity User, index int) bool {
		return entity.StructPtr1 == _StructPtr1
	})
}

// FilterRejectByChanSend0
func (xs UserRename) FilterRejectByChanSend0(_ChanSend0 chan<- int) (kept UserRename, rejected UserRename) {
	return lo.FilterReject(xs, func(entity User, index int) bool {
		return entity.ChanSend0 == _ChanSend0
	})
}

// FilterRejectByChanSendPtr0
func (xs UserRename) FilterRejectByChanSendPtr0(_ChanSendPtr0 *chan<- int) (kept UserRename, rejected UserRename) {
	return lo.FilterReject(xs, func(entity User, index int) bool {
		return entity.ChanSendPtr0 == _ChanSendPtr0
	})
}

// -- Find ------------------------------------

// FindByUserID
func (xs UserRename) FindByUserID(_UserID string) (User, bool) {
	return lo.Find(xs, func(entity User) bool {
		return entity.UserID == _UserID
	})
}

// FindByInt
func (xs UserRename) FindByInt(_Int int) (User, bool) {
	return lo.Find(xs, func(entity User) bool {
		return entity.Int == _Int
	})
}

// FindByIntPtr
func (xs UserRename) FindByIntPtr(_IntPtr *int) (User, bool) {
	return lo.Find(xs, func(entity User) bool {
		return entity.IntPtr == _IntPtr
	})
}

// FindByBool
func (xs UserRename) FindByBool(_Bool bool) (User, bool) {
	return lo.Find(xs, func(entity User) bool {
		return entity.Bool == _Bool
	})
}

// FindByBoolPtr
func (xs UserRename) FindByBoolPtr(_BoolPtr *bool) (User, bool) {
	return lo.Find(xs, func(entity User) bool {
		return entity.BoolPtr == _BoolPtr
	})
}

// FindByStr
func (xs UserRename) FindByStr(_Str string) (User, bool) {
	return lo.Find(xs, func(entity User) bool {
		return entity.Str == _Str
	})
}

// FindByStrPtr
func (xs UserRename) FindByStrPtr(_StrPtr *string) (User, bool) {
	return lo.Find(xs, func(entity User) bool {
		return entity.StrPtr == _StrPtr
	})
}

// FindByStruct0
func (xs UserRename) FindByStruct0(_Struct0 DefinedStruct0) (User, bool) {
	return lo.Find(xs, func(entity User) bool {
		return entity.Struct0 == _Struct0
	})
}

// FindByStructPtr0
func (xs UserRename) FindByStructPtr0(_StructPtr0 *DefinedStruct0) (User, bool) {
	return lo.Find(xs, func(entity User) bool {
		return entity.StructPtr0 == _StructPtr0
	})
}

// FindByStruct1
func (xs UserRename) FindByStruct1(_Struct1 DefinedStruct1) (User, bool) {
	return lo.Find(xs, func(entity User) bool {
		return entity.Struct1 == _Struct1
	})
}

// FindByStructPtr1
func (xs UserRename) FindByStructPtr1(_StructPtr1 *DefinedStruct1) (User, bool) {
	return lo.Find(xs, func(entity User) bool {
		return entity.StructPtr1 == _StructPtr1
	})
}

// FindByChanSend0
func (xs UserRename) FindByChanSend0(_ChanSend0 chan<- int) (User, bool) {
	return lo.Find(xs, func(entity User) bool {
		return entity.ChanSend0 == _ChanSend0
	})
}

// FindByChanSendPtr0
func (xs UserRename) FindByChanSendPtr0(_ChanSendPtr0 *chan<- int) (User, bool) {
	return lo.Find(xs, func(entity User) bool {
		return entity.ChanSendPtr0 == _ChanSendPtr0
	})
}

// -- ContainsBy ------------------------------------

// ContainsByUserID
func (xs UserRename) ContainsByUserID(_UserID string) bool {
	return lo.ContainsBy(xs, func(entity User) bool {
		return entity.UserID == _UserID
	})
}

// ContainsByInt
func (xs UserRename) ContainsByInt(_Int int) bool {
	return lo.ContainsBy(xs, func(entity User) bool {
		return entity.Int == _Int
	})
}

// ContainsByIntPtr
func (xs UserRename) ContainsByIntPtr(_IntPtr *int) bool {
	return lo.ContainsBy(xs, func(entity User) bool {
		return entity.IntPtr == _IntPtr
	})
}

// ContainsByBool
func (xs UserRename) ContainsByBool(_Bool bool) bool {
	return lo.ContainsBy(xs, func(entity User) bool {
		return entity.Bool == _Bool
	})
}

// ContainsByBoolPtr
func (xs UserRename) ContainsByBoolPtr(_BoolPtr *bool) bool {
	return lo.ContainsBy(xs, func(entity User) bool {
		return entity.BoolPtr == _BoolPtr
	})
}

// ContainsByStr
func (xs UserRename) ContainsByStr(_Str string) bool {
	return lo.ContainsBy(xs, func(entity User) bool {
		return entity.Str == _Str
	})
}

// ContainsByStrPtr
func (xs UserRename) ContainsByStrPtr(_StrPtr *string) bool {
	return lo.ContainsBy(xs, func(entity User) bool {
		return entity.StrPtr == _StrPtr
	})
}

// ContainsByStruct0
func (xs UserRename) ContainsByStruct0(_Struct0 DefinedStruct0) bool {
	return lo.ContainsBy(xs, func(entity User) bool {
		return entity.Struct0 == _Struct0
	})
}

// ContainsByStructPtr0
func (xs UserRename) ContainsByStructPtr0(_StructPtr0 *DefinedStruct0) bool {
	return lo.ContainsBy(xs, func(entity User) bool {
		return entity.StructPtr0 == _StructPtr0
	})
}

// ContainsByStruct1
func (xs UserRename) ContainsByStruct1(_Struct1 DefinedStruct1) bool {
	return lo.ContainsBy(xs, func(entity User) bool {
		return entity.Struct1 == _Struct1
	})
}

// ContainsByStructPtr1
func (xs UserRename) ContainsByStructPtr1(_StructPtr1 *DefinedStruct1) bool {
	return lo.ContainsBy(xs, func(entity User) bool {
		return entity.StructPtr1 == _StructPtr1
	})
}

// ContainsByChanSend0
func (xs UserRename) ContainsByChanSend0(_ChanSend0 chan<- int) bool {
	return lo.ContainsBy(xs, func(entity User) bool {
		return entity.ChanSend0 == _ChanSend0
	})
}

// ContainsByChanSendPtr0
func (xs UserRename) ContainsByChanSendPtr0(_ChanSendPtr0 *chan<- int) bool {
	return lo.ContainsBy(xs, func(entity User) bool {
		return entity.ChanSendPtr0 == _ChanSendPtr0
	})
}
