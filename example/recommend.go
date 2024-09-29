package main

type Recommend struct {
	ID  string
	Num int32
}

//go:generate go run ../main.go --entity=*Recommend --slice=Recommends --input=recommend.go --output=recommend_gen.go --rename="(.*):Lo\{0}" --include=Filter$,FilterBy,FilterReject$,UniqBy,GroupBy,KeyBy,ContainsBy$,EveryBy,SomeBy,Find,ToSlicePtr,FromSlicePtr
type Recommends []*Recommend
