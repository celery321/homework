package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	addr := "www.baidu.com:80"
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(conn.RemoteAddr().String())
	fmt.Println(conn.LocalAddr().String())
	n, err := conn.Write([]byte("GET / HTTP/1.1\r\n\r\n"))
	if err != nil {
		log.Println(err)
	}
	buf := make([]byte, 4096)
	n, err = conn.Read(buf)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(buf[:n]))
	conn.Close()

}
