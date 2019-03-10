package main

import (
	"fmt"
	"strings"
)

func wordCount(str string) map[string]int {

	s := strings.Fields(str)
	m := make(map[string]int)

	//for i:=0; i < len(s); i++{
	//	if _, ok := m[s[i]]; ok {
	//		m[s[i]] = m[s[i]]+1
	//	} else {
	//		m[s[i]] = 1
	//	}
	//}

	for _, v := range s {
		if _, ok := m[v]; ok {
			m[v] = m[v] + 1
		} else {
			m[v] = 1
		}
	}

	return m

}

func ttt(p *map[string]int) {

	fmt.Printf("%p %p", &p, *p)
}

func main() {
	str := "i love my work and i love my fanily too"

	mret := wordCount(str)

	fmt.Println(mret)

	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2
	fmt.Printf("%p %p\n", m, &m)
	ttt(&m)
	fmt.Println(m)
}
