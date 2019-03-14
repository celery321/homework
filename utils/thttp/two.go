package main

import (
	"fmt"
	"net"
)

func main() {
	con, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println(err)
	}
	defer con.Close()

	httpRequest := "GET /itcast HTTP/1.1\r\nHost:127.0.0.1:8000\r\n\r\n"

	con.Write([]byte(httpRequest))
	buf := make([]byte, 4096)
	n, _ := con.Read(buf)
	if n == 0 {
		return
	}
	fmt.Printf("|%s|", string(buf[:n]))
}
