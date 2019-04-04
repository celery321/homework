package models

type User1 struct {
	Id   int
	Name string `orm:"size(100)"`
}

type User2 struct {
	Id   int
	Name string `orm:"size(100)"`
}
