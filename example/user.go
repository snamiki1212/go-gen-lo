package main

//go:generate go run ../main.go --entity=*User --slice=Users --input=user.go --output=user_gen.go
type User struct {
	UserID string
	Age    int64
	// Callback func() string
	Map map[string]string
	// Chann    chan string
	// IF       interface{}
	// Any      any
	item Item
}

type Users []*User

type Item struct{}
