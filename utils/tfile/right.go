package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.OpenFile("a.txt", os.O_CREATE|os.O_APPEND|os.O_APPEND, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	buf := make([]byte, 4096)

	//r := bufio.NewReader(f)
	//n,err := r.Read(buf)
	//
	io.Copy(os.Stdout, f)
	r := bufio.NewReader(f)
	for {

		n, err := r.Read(buf)
		if err == nil {
			fmt.Println(buf[:n])
		} else if err == io.EOF {
			break
		} else {
			log.Fatal(err)
		}

	}

	//for {
	//	n, err := os.Stdin.Read(buf)
	//	if err != nil{
	//		return
	//	}
	//	os.Stdout.Write(buf[:n])
	//}
	//
	//r1 := bufio.NewScanner(f)
	//for r1.Scan() {
	//	fmt.Println(r1.Text())
	//}

}
