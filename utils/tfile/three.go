package main

import (
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
	n, err := f.WriteString("hello world\r\n")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(n)
	off, _ := f.Seek(5, io.SeekStart)
	f.WriteAt([]byte("1111"), off)
	fmt.Println(off)
}
