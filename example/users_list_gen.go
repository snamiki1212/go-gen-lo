// Code generated by "go-gen-lo"; DO NOT EDIT.
// Based on information from https://github.com/snamiki1212/go-gen-lo

package main

import "github.com/samber/lo"

// Filter
func (xs UserList) Filter(predicate func(User, int) bool) UserList {
	return lo.Filter(xs, predicate)
}

// Map
func (xs UserList) Map(iteratee func(item User, index int) User) UserList {
	return lo.Map(xs, iteratee)
}

// FilterReject
func (xs UserList) FilterReject(predicate func(User, int) bool) (UserList, UserList) {
	return lo.FilterReject(xs, predicate)
}

// Find
func (xs UserList) Find(predicate func(User) bool) (User, bool) {
	return lo.Find(xs, predicate)
}

// FilterByUserID
func (xs UserList) FilterByUserID(_UserID string) UserList {
	return lo.Filter(xs, func(entity User, index int) bool {
		return entity.UserID == _UserID
	})
}

// FilterByInt
func (xs UserList) FilterByInt(_Int int) UserList {
	return lo.Filter(xs, func(entity User, index int) bool {
		return entity.Int == _Int
	})
}

// FilterByIntPtr
func (xs UserList) FilterByIntPtr(_IntPtr *int) UserList {
	return lo.Filter(xs, func(entity User, index int) bool {
		return entity.IntPtr == _IntPtr
	})
}

// FilterByBool
func (xs UserList) FilterByBool(_Bool bool) UserList {
	return lo.Filter(xs, func(entity User, index int) bool {
		return entity.Bool == _Bool
	})
}

// FilterByBoolPtr
func (xs UserList) FilterByBoolPtr(_BoolPtr *bool) UserList {
	return lo.Filter(xs, func(entity User, index int) bool {
		return entity.BoolPtr == _BoolPtr
	})
}

// FilterByStr
func (xs UserList) FilterByStr(_Str string) UserList {
	return lo.Filter(xs, func(entity User, index int) bool {
		return entity.Str == _Str
	})
}

// FilterByStrPtr
func (xs UserList) FilterByStrPtr(_StrPtr *string) UserList {
	return lo.Filter(xs, func(entity User, index int) bool {
		return entity.StrPtr == _StrPtr
	})
}

// FilterByStruct0
func (xs UserList) FilterByStruct0(_Struct0 DefinedStruct0) UserList {
	return lo.Filter(xs, func(entity User, index int) bool {
		return entity.Struct0 == _Struct0
	})
}

// FilterByStructPtr0
func (xs UserList) FilterByStructPtr0(_StructPtr0 *DefinedStruct0) UserList {
	return lo.Filter(xs, func(entity User, index int) bool {
		return entity.StructPtr0 == _StructPtr0
	})
}

// FilterByStruct1
func (xs UserList) FilterByStruct1(_Struct1 DefinedStruct1) UserList {
	return lo.Filter(xs, func(entity User, index int) bool {
		return entity.Struct1 == _Struct1
	})
}

// FilterByStructPtr1
func (xs UserList) FilterByStructPtr1(_StructPtr1 *DefinedStruct1) UserList {
	return lo.Filter(xs, func(entity User, index int) bool {
		return entity.StructPtr1 == _StructPtr1
	})
}

// FilterByChanSend0
func (xs UserList) FilterByChanSend0(_ChanSend0 chan<- int) UserList {
	return lo.Filter(xs, func(entity User, index int) bool {
		return entity.ChanSend0 == _ChanSend0
	})
}

// FilterByChanSendPtr0
func (xs UserList) FilterByChanSendPtr0(_ChanSendPtr0 *chan<- int) UserList {
	return lo.Filter(xs, func(entity User, index int) bool {
		return entity.ChanSendPtr0 == _ChanSendPtr0
	})
}

// KeyByUserID
func (xs UserList) KeyByUserID() map[string]User {
	return lo.KeyBy(xs, func(entity User) string {
		return entity.UserID
	})
}

// KeyByInt
func (xs UserList) KeyByInt() map[int]User {
	return lo.KeyBy(xs, func(entity User) int {
		return entity.Int
	})
}

// KeyByIntPtr
func (xs UserList) KeyByIntPtr() map[*int]User {
	return lo.KeyBy(xs, func(entity User) *int {
		return entity.IntPtr
	})
}

// KeyByBool
func (xs UserList) KeyByBool() map[bool]User {
	return lo.KeyBy(xs, func(entity User) bool {
		return entity.Bool
	})
}

// KeyByBoolPtr
func (xs UserList) KeyByBoolPtr() map[*bool]User {
	return lo.KeyBy(xs, func(entity User) *bool {
		return entity.BoolPtr
	})
}

// KeyByStr
func (xs UserList) KeyByStr() map[string]User {
	return lo.KeyBy(xs, func(entity User) string {
		return entity.Str
	})
}

// KeyByStrPtr
func (xs UserList) KeyByStrPtr() map[*string]User {
	return lo.KeyBy(xs, func(entity User) *string {
		return entity.StrPtr
	})
}

// KeyByStruct0
func (xs UserList) KeyByStruct0() map[DefinedStruct0]User {
	return lo.KeyBy(xs, func(entity User) DefinedStruct0 {
		return entity.Struct0
	})
}

// KeyByStructPtr0
func (xs UserList) KeyByStructPtr0() map[*DefinedStruct0]User {
	return lo.KeyBy(xs, func(entity User) *DefinedStruct0 {
		return entity.StructPtr0
	})
}

// KeyByStruct1
func (xs UserList) KeyByStruct1() map[DefinedStruct1]User {
	return lo.KeyBy(xs, func(entity User) DefinedStruct1 {
		return entity.Struct1
	})
}

// KeyByStructPtr1
func (xs UserList) KeyByStructPtr1() map[*DefinedStruct1]User {
	return lo.KeyBy(xs, func(entity User) *DefinedStruct1 {
		return entity.StructPtr1
	})
}

// KeyByChanSend0
func (xs UserList) KeyByChanSend0() map[chan<- int]User {
	return lo.KeyBy(xs, func(entity User) chan<- int {
		return entity.ChanSend0
	})
}

// KeyByChanSendPtr0
func (xs UserList) KeyByChanSendPtr0() map[*chan<- int]User {
	return lo.KeyBy(xs, func(entity User) *chan<- int {
		return entity.ChanSendPtr0
	})
}

// GroupByUserID
func (xs UserList) GroupByUserID() map[string]UserList {
	return lo.GroupBy(xs, func(entity User) string {
		return entity.UserID
	})
}

// GroupByInt
func (xs UserList) GroupByInt() map[int]UserList {
	return lo.GroupBy(xs, func(entity User) int {
		return entity.Int
	})
}

// GroupByIntPtr
func (xs UserList) GroupByIntPtr() map[*int]UserList {
	return lo.GroupBy(xs, func(entity User) *int {
		return entity.IntPtr
	})
}

// GroupByBool
func (xs UserList) GroupByBool() map[bool]UserList {
	return lo.GroupBy(xs, func(entity User) bool {
		return entity.Bool
	})
}

// GroupByBoolPtr
func (xs UserList) GroupByBoolPtr() map[*bool]UserList {
	return lo.GroupBy(xs, func(entity User) *bool {
		return entity.BoolPtr
	})
}

// GroupByStr
func (xs UserList) GroupByStr() map[string]UserList {
	return lo.GroupBy(xs, func(entity User) string {
		return entity.Str
	})
}

// GroupByStrPtr
func (xs UserList) GroupByStrPtr() map[*string]UserList {
	return lo.GroupBy(xs, func(entity User) *string {
		return entity.StrPtr
	})
}

// GroupByStruct0
func (xs UserList) GroupByStruct0() map[DefinedStruct0]UserList {
	return lo.GroupBy(xs, func(entity User) DefinedStruct0 {
		return entity.Struct0
	})
}

// GroupByStructPtr0
func (xs UserList) GroupByStructPtr0() map[*DefinedStruct0]UserList {
	return lo.GroupBy(xs, func(entity User) *DefinedStruct0 {
		return entity.StructPtr0
	})
}

// GroupByStruct1
func (xs UserList) GroupByStruct1() map[DefinedStruct1]UserList {
	return lo.GroupBy(xs, func(entity User) DefinedStruct1 {
		return entity.Struct1
	})
}

// GroupByStructPtr1
func (xs UserList) GroupByStructPtr1() map[*DefinedStruct1]UserList {
	return lo.GroupBy(xs, func(entity User) *DefinedStruct1 {
		return entity.StructPtr1
	})
}

// GroupByChanSend0
func (xs UserList) GroupByChanSend0() map[chan<- int]UserList {
	return lo.GroupBy(xs, func(entity User) chan<- int {
		return entity.ChanSend0
	})
}

// GroupByChanSendPtr0
func (xs UserList) GroupByChanSendPtr0() map[*chan<- int]UserList {
	return lo.GroupBy(xs, func(entity User) *chan<- int {
		return entity.ChanSendPtr0
	})
}

// FilterRejectByUserID
func (xs UserList) FilterRejectByUserID(_UserID string) (kept UserList, rejected UserList) {
	return lo.FilterReject(xs, func(entity User, index int) bool {
		return entity.UserID == _UserID
	})
}

// FilterRejectByInt
func (xs UserList) FilterRejectByInt(_Int int) (kept UserList, rejected UserList) {
	return lo.FilterReject(xs, func(entity User, index int) bool {
		return entity.Int == _Int
	})
}

// FilterRejectByIntPtr
func (xs UserList) FilterRejectByIntPtr(_IntPtr *int) (kept UserList, rejected UserList) {
	return lo.FilterReject(xs, func(entity User, index int) bool {
		return entity.IntPtr == _IntPtr
	})
}

// FilterRejectByBool
func (xs UserList) FilterRejectByBool(_Bool bool) (kept UserList, rejected UserList) {
	return lo.FilterReject(xs, func(entity User, index int) bool {
		return entity.Bool == _Bool
	})
}

// FilterRejectByBoolPtr
func (xs UserList) FilterRejectByBoolPtr(_BoolPtr *bool) (kept UserList, rejected UserList) {
	return lo.FilterReject(xs, func(entity User, index int) bool {
		return entity.BoolPtr == _BoolPtr
	})
}

// FilterRejectByStr
func (xs UserList) FilterRejectByStr(_Str string) (kept UserList, rejected UserList) {
	return lo.FilterReject(xs, func(entity User, index int) bool {
		return entity.Str == _Str
	})
}

// FilterRejectByStrPtr
func (xs UserList) FilterRejectByStrPtr(_StrPtr *string) (kept UserList, rejected UserList) {
	return lo.FilterReject(xs, func(entity User, index int) bool {
		return entity.StrPtr == _StrPtr
	})
}

// FilterRejectByStruct0
func (xs UserList) FilterRejectByStruct0(_Struct0 DefinedStruct0) (kept UserList, rejected UserList) {
	return lo.FilterReject(xs, func(entity User, index int) bool {
		return entity.Struct0 == _Struct0
	})
}

// FilterRejectByStructPtr0
func (xs UserList) FilterRejectByStructPtr0(_StructPtr0 *DefinedStruct0) (kept UserList, rejected UserList) {
	return lo.FilterReject(xs, func(entity User, index int) bool {
		return entity.StructPtr0 == _StructPtr0
	})
}

// FilterRejectByStruct1
func (xs UserList) FilterRejectByStruct1(_Struct1 DefinedStruct1) (kept UserList, rejected UserList) {
	return lo.FilterReject(xs, func(entity User, index int) bool {
		return entity.Struct1 == _Struct1
	})
}

// FilterRejectByStructPtr1
func (xs UserList) FilterRejectByStructPtr1(_StructPtr1 *DefinedStruct1) (kept UserList, rejected UserList) {
	return lo.FilterReject(xs, func(entity User, index int) bool {
		return entity.StructPtr1 == _StructPtr1
	})
}

// FilterRejectByChanSend0
func (xs UserList) FilterRejectByChanSend0(_ChanSend0 chan<- int) (kept UserList, rejected UserList) {
	return lo.FilterReject(xs, func(entity User, index int) bool {
		return entity.ChanSend0 == _ChanSend0
	})
}

// FilterRejectByChanSendPtr0
func (xs UserList) FilterRejectByChanSendPtr0(_ChanSendPtr0 *chan<- int) (kept UserList, rejected UserList) {
	return lo.FilterReject(xs, func(entity User, index int) bool {
		return entity.ChanSendPtr0 == _ChanSendPtr0
	})
}

// FindByUserID
func (xs UserList) FindByUserID(_UserID string) (User, bool) {
	return lo.Find(xs, func(entity User) bool {
		return entity.UserID == _UserID
	})
}

// FindByInt
func (xs UserList) FindByInt(_Int int) (User, bool) {
	return lo.Find(xs, func(entity User) bool {
		return entity.Int == _Int
	})
}

// FindByIntPtr
func (xs UserList) FindByIntPtr(_IntPtr *int) (User, bool) {
	return lo.Find(xs, func(entity User) bool {
		return entity.IntPtr == _IntPtr
	})
}

// FindByBool
func (xs UserList) FindByBool(_Bool bool) (User, bool) {
	return lo.Find(xs, func(entity User) bool {
		return entity.Bool == _Bool
	})
}

// FindByBoolPtr
func (xs UserList) FindByBoolPtr(_BoolPtr *bool) (User, bool) {
	return lo.Find(xs, func(entity User) bool {
		return entity.BoolPtr == _BoolPtr
	})
}

// FindByStr
func (xs UserList) FindByStr(_Str string) (User, bool) {
	return lo.Find(xs, func(entity User) bool {
		return entity.Str == _Str
	})
}

// FindByStrPtr
func (xs UserList) FindByStrPtr(_StrPtr *string) (User, bool) {
	return lo.Find(xs, func(entity User) bool {
		return entity.StrPtr == _StrPtr
	})
}

// FindByStruct0
func (xs UserList) FindByStruct0(_Struct0 DefinedStruct0) (User, bool) {
	return lo.Find(xs, func(entity User) bool {
		return entity.Struct0 == _Struct0
	})
}

// FindByStructPtr0
func (xs UserList) FindByStructPtr0(_StructPtr0 *DefinedStruct0) (User, bool) {
	return lo.Find(xs, func(entity User) bool {
		return entity.StructPtr0 == _StructPtr0
	})
}

// FindByStruct1
func (xs UserList) FindByStruct1(_Struct1 DefinedStruct1) (User, bool) {
	return lo.Find(xs, func(entity User) bool {
		return entity.Struct1 == _Struct1
	})
}

// FindByStructPtr1
func (xs UserList) FindByStructPtr1(_StructPtr1 *DefinedStruct1) (User, bool) {
	return lo.Find(xs, func(entity User) bool {
		return entity.StructPtr1 == _StructPtr1
	})
}

// FindByChanSend0
func (xs UserList) FindByChanSend0(_ChanSend0 chan<- int) (User, bool) {
	return lo.Find(xs, func(entity User) bool {
		return entity.ChanSend0 == _ChanSend0
	})
}

// FindByChanSendPtr0
func (xs UserList) FindByChanSendPtr0(_ChanSendPtr0 *chan<- int) (User, bool) {
	return lo.Find(xs, func(entity User) bool {
		return entity.ChanSendPtr0 == _ChanSendPtr0
	})
}
