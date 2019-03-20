package main

import "fmt"

func main() {
	fmt.Println()
	fmt.Printf("%v\n", "a")
	s := fmt.Sprintf("http://%s.%s", "www.baidu", "com")
	fmt.Println(s)

	var s1 string
	var n int
	//var line string
	for {
		fmt.Print("> ")
		fmt.Scan(&s1, &n)
		fmt.Println(s1, n)
		if s1 == "stop" {
			break
		}
	}
}
