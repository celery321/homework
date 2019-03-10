package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("C:/syncdata/2.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	fmt.Println()
}
