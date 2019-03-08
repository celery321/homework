package tinterfacer

import "fmt"

type person struct {
	name string
	sex  string
	age  int
}

func (p *person) SayHello() {
	fmt.Printf("名字是:%s 性别是:%s 年龄是:%d", p.name, p.sex, p.age)
}

func (p *person) editHello(name string, sex string, age int) {
	p.name = name
	p.sex = sex
	p.name = name
}

func main() {
	p1 := person{"王宝强", "男", 36}
	p1.SayHello()

	f := person{"马蓉", "女", 29}
	f.editHello("李薇薇", "女", 28)
	f.SayHello()
}
