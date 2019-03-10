package main

import "fmt"

type person3 struct {
	name     string
	age      int
	flg      bool
	interset []string
}

func initFunc(p *person3) {
	p.name = "nami"
	p.age = 18
	p.flg = true
	p.interset = append(p.interset, "shoping")
	p.interset = append(p.interset, "sleeping")
	//p.interset = []string{"aaa", "bbb"}
}

func initFunc2() *person3 {
	p := new(person3)
	p.name = "nami"
	p.age = 18
	p.flg = true
	p.interset = append(p.interset, "shoping")
	p.interset = append(p.interset, "sleeping")
	//p.interset = []string{"aaa", "bbb"}
	return p
}
func main() {
	//var p person3
	//initFunc(&p)
	//fmt.Println(p)
	p2 := initFunc2()
	fmt.Println(p2)

}
