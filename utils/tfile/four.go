package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.OpenFile("C:/syncdata/1.txt", os.O_RDWR, 6)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	fmt.Println("ok")
	if err != nil {
		fmt.Println(err)
		return
	}
	reader := bufio.NewReader(f)
	for {
		buf, err := reader.ReadBytes('\n')
		if err != nil && err == io.EOF {
			fmt.Println("读取玩伴")
			return
		} else if err != nil {
			fmt.Println("读取玩伴")
		}
		fmt.Println(string(buf))
	}
}
