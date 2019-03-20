package main

import "fmt"

type Student1 struct {
	Id   int
	Name string
}

func main() {
	s1 := Student1{
		Id:   2,
		Name: "alice",
	}

	var p *Student1
	p = &s1
	p.Id = 2
	fmt.Println(s1)

	var p1 *int
	p1 = &s1.Id
	*p1 = 3
	fmt.Println(s1)

}
