package main

import (
	"log"
	"net"
)

func main() {
	addr := ":8021"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	conn, err := listener.Accept()
	if err != nil {
		log.Println(err)
	}
	conn.Write([]byte("hwllo"))
	conn.Close()
}
