package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("C:/syncdata/1.mp4")
	if err != nil {
		fmt.Println(err)

	}

	defer f.Close()

	fw, err := os.Create("C:/syncdata/2.mp4")
	if err != nil {
		fmt.Println(err)
	}
	defer fw.Close()

	buf := make([]byte, 4096)

	for {

		n, err := f.Read(buf)
		if err != nil && err == io.EOF {
			fmt.Println("读完", n)
			return
		}
		fw.Write(buf[:n])

	}
}
