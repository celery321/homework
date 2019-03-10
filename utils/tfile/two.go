package main

import (
	"fmt"
	"os"
)

func main() {

	f, err := os.OpenFile("C:/syncdata/1.txt", os.O_RDWR, 6)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = f.WriteString("!!!!!")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("ok")
}
