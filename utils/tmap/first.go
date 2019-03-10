package main

import "fmt"

func main() {
	m1 := make(map[int]string, 1)
	m1[100] = "1"
	m1[200] = "bbb"
	fmt.Println(m1)

	m1[200] = "yello"
	fmt.Println(m1)

	m8 := map[int]string{1: "aaa",
		2: "bbb",
		3: "cccc",
	}
	for k, v := range m8 {
		fmt.Println(k, v)
	}

	if value, ok := m8[1]; ok {
		fmt.Println(value, ok)
	} else {
		fmt.Println("不存在")
	}

}
