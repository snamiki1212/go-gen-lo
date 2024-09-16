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

// ContainsBy
func (xs UserExclude) ContainsBy(predicate func(User) bool) bool {
	return lo.ContainsBy(xs, predicate)
}

// EveryBy
func (xs UserExclude) EveryBy(predicate func(item User) bool) bool {
	return lo.EveryBy(xs, predicate)
}

// SomeBy
func (xs UserExclude) SomeBy(predicate func(item User) bool) bool {
	return lo.SomeBy(xs, predicate)
}

/************************************************
 ** lo extended methods
 ************************************************/

// -- UniqBy ------------------------------------

// UniqByUserID
func (xs UserExclude) UniqByUserID() UserExclude {
	return lo.UniqBy(xs, func(entity User) string {
		return entity.UserID
	})
}

// UniqByInt
func (xs UserExclude) UniqByInt() UserExclude {
	return lo.UniqBy(xs, func(entity User) int {
		return entity.Int
	})
}

// UniqByIntPtr
func (xs UserExclude) UniqByIntPtr() UserExclude {
	return lo.UniqBy(xs, func(entity User) *int {
		return entity.IntPtr
	})
}

// UniqByBool
func (xs UserExclude) UniqByBool() UserExclude {
	return lo.UniqBy(xs, func(entity User) bool {
		return entity.Bool
	})
}

// UniqByBoolPtr
func (xs UserExclude) UniqByBoolPtr() UserExclude {
	return lo.UniqBy(xs, func(entity User) *bool {
		return entity.BoolPtr
	})
}

// UniqByStr
func (xs UserExclude) UniqByStr() UserExclude {
	return lo.UniqBy(xs, func(entity User) string {
		return entity.Str
	})
}

// UniqByStrPtr
func (xs UserExclude) UniqByStrPtr() UserExclude {
	return lo.UniqBy(xs, func(entity User) *string {
		return entity.StrPtr
	})
}

// UniqByStruct0
func (xs UserExclude) UniqByStruct0() UserExclude {
	return lo.UniqBy(xs, func(entity User) DefinedStruct0 {
		return entity.Struct0
	})
}

// UniqByStructPtr0
func (xs UserExclude) UniqByStructPtr0() UserExclude {
	return lo.UniqBy(xs, func(entity User) *DefinedStruct0 {
		return entity.StructPtr0
	})
}

// UniqByStruct1
func (xs UserExclude) UniqByStruct1() UserExclude {
	return lo.UniqBy(xs, func(entity User) DefinedStruct1 {
		return entity.Struct1
	})
}

// UniqByStructPtr1
func (xs UserExclude) UniqByStructPtr1() UserExclude {
	return lo.UniqBy(xs, func(entity User) *DefinedStruct1 {
		return entity.StructPtr1
	})
}

// UniqByChanSend0
func (xs UserExclude) UniqByChanSend0() UserExclude {
	return lo.UniqBy(xs, func(entity User) chan<- int {
		return entity.ChanSend0
	})
}

// UniqByChanSendPtr0
func (xs UserExclude) UniqByChanSendPtr0() UserExclude {
	return lo.UniqBy(xs, func(entity User) *chan<- int {
		return entity.ChanSendPtr0
	})
}

// -- ContainsBy ------------------------------------

// ContainsByUserID
func (xs UserExclude) ContainsByUserID(_UserID string) bool {
	return lo.ContainsBy(xs, func(entity User) bool {
		return entity.UserID == _UserID
	})
}

// ContainsByInt
func (xs UserExclude) ContainsByInt(_Int int) bool {
	return lo.ContainsBy(xs, func(entity User) bool {
		return entity.Int == _Int
	})
}

// ContainsByIntPtr
func (xs UserExclude) ContainsByIntPtr(_IntPtr *int) bool {
	return lo.ContainsBy(xs, func(entity User) bool {
		return entity.IntPtr == _IntPtr
	})
}

// ContainsByBool
func (xs UserExclude) ContainsByBool(_Bool bool) bool {
	return lo.ContainsBy(xs, func(entity User) bool {
		return entity.Bool == _Bool
	})
}

// ContainsByBoolPtr
func (xs UserExclude) ContainsByBoolPtr(_BoolPtr *bool) bool {
	return lo.ContainsBy(xs, func(entity User) bool {
		return entity.BoolPtr == _BoolPtr
	})
}

// ContainsByStr
func (xs UserExclude) ContainsByStr(_Str string) bool {
	return lo.ContainsBy(xs, func(entity User) bool {
		return entity.Str == _Str
	})
}

// ContainsByStrPtr
func (xs UserExclude) ContainsByStrPtr(_StrPtr *string) bool {
	return lo.ContainsBy(xs, func(entity User) bool {
		return entity.StrPtr == _StrPtr
	})
}

// ContainsByStruct0
func (xs UserExclude) ContainsByStruct0(_Struct0 DefinedStruct0) bool {
	return lo.ContainsBy(xs, func(entity User) bool {
		return entity.Struct0 == _Struct0
	})
}

// ContainsByStructPtr0
func (xs UserExclude) ContainsByStructPtr0(_StructPtr0 *DefinedStruct0) bool {
	return lo.ContainsBy(xs, func(entity User) bool {
		return entity.StructPtr0 == _StructPtr0
	})
}

// ContainsByStruct1
func (xs UserExclude) ContainsByStruct1(_Struct1 DefinedStruct1) bool {
	return lo.ContainsBy(xs, func(entity User) bool {
		return entity.Struct1 == _Struct1
	})
}

// ContainsByStructPtr1
func (xs UserExclude) ContainsByStructPtr1(_StructPtr1 *DefinedStruct1) bool {
	return lo.ContainsBy(xs, func(entity User) bool {
		return entity.StructPtr1 == _StructPtr1
	})
}

// ContainsByChanSend0
func (xs UserExclude) ContainsByChanSend0(_ChanSend0 chan<- int) bool {
	return lo.ContainsBy(xs, func(entity User) bool {
		return entity.ChanSend0 == _ChanSend0
	})
}

// ContainsByChanSendPtr0
func (xs UserExclude) ContainsByChanSendPtr0(_ChanSendPtr0 *chan<- int) bool {
	return lo.ContainsBy(xs, func(entity User) bool {
		return entity.ChanSendPtr0 == _ChanSendPtr0
	})
}
