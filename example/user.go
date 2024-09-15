package main

type User struct {
	// Example
	UserID string

	// Primitive
	Int          int
	IntPtr       *int
	IntSlice     []int
	IntSlicePtr  []*int
	Bool         bool
	BoolPtr      *bool
	BoolSlice    []bool
	BoolSlicePtr []*bool
	Str          string
	StrPtr       *string
	StrSlice     []string
	StrSlicePtr  []*string

	// Function
	Fn1         func() string
	FnPtr1      *func() string
	FnSlice1    []func() string
	FnSlice1Ptr []*func() string
	Fn2         func()
	FnPtr2      *func()
	FnSlice2    []func()
	FnSlice2Ptr []*func()
	Fn3         func(func() string) func() int
	FnPtr3      *func(func() string) func() int
	FnSlice3    []func(func() string) func() int
	FnSlice3Ptr []*func(func() string) func() int
	Fn4         func(*func() string) *func() int
	FnPtr4      *func(*func() string) *func() int
	FnSlice4    []func(*func() string) *func() int
	FnSlice4Ptr []*func(*func() string) *func() int

	// Struct
	Struct0         DefinedStruct0
	StructPtr0      *DefinedStruct0
	StructSlice0    []DefinedStruct0
	StructSlice0Ptr []*DefinedStruct0
	Struct1         DefinedStruct1
	StructPtr1      *DefinedStruct1
	StructSlice1    []DefinedStruct1
	StructSlice1Ptr []*DefinedStruct1

	// Map
	Map1         map[string]int
	MapPtr1      *map[string]int
	MapSlice1    []map[string]int
	MapSlice1Ptr []*map[string]int
	Map2         map[string]func() string
	MapPtr2      *map[string]func() string
	MapSlice2    []map[string]func() string
	MapSlice2Ptr []*map[string]func() string

	// Channel
	Chan0             chan int
	ChanPtr0          *chan int
	ChanSlice0        []chan int
	ChanSlicePtr0     []*chan int
	Chan1             chan func() string
	ChanPtr1          *chan func() string
	ChanSlice1        []chan func() string
	ChanSlicePtr1     []*chan func() string
	ChanSend0         chan<- int
	ChanSendPtr0      *chan<- int
	ChanSendSlice0    []chan<- int
	ChanSendSlicePtr0 []*chan<- int
	ChanRecv0         <-chan int
	ChanRecvPtr0      *<-chan int
	ChanRecvSlice0    []<-chan int
	ChanRecvSlicePtr0 []*<-chan int

	// No Support
	// InlineStruct0 struct{ Name string } // NOTE: not supported
	// InlineStruct1 struct { Name string } // NOTE: not supported
}

type DefinedStruct0 struct{}

type DefinedStruct1 struct {
	Name string
}

//go:generate go run ../main.go --entity=*User --slice=UserPtrs --input=user.go --output=users_ptr_gen.go
type UserPtrs []*User

//go:generate go run ../main.go --entity=User --slice=UserList --input=user.go --output=users_list_gen.go
type UserList []User

//go:generate go run ../main.go --entity=User --slice=UserExclude --input=user.go --output=users_exclude_gen.go --exclude=Filter,KeyBy
type UserExclude []User

//go:generate go run ../main.go --entity=User --slice=UserRename --input=user.go --output=users_rename_gen.go --exclude=Filter,KeyBy --rename=Map:Loop
type UserRename []User
