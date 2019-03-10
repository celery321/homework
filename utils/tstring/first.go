package main

import (
	"fmt"
	"strings"
)

func main() {
	str11 := "i love my work and i love my family too"
	ret := strings.Split(str11, "")

	for _, s := range ret {
		fmt.Println(s)
	}

	str11 = "xxx.exe"
	flg := strings.HasSuffix(str11, ".exe")
	fmt.Println(flg)

}
