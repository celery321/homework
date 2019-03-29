package main

import (
	"fmt"
	"log"
	"net"
)

var mainContent = `HTTP/1.1 200 OK
Content-Type: text/html
Content-Length: %d

`

func handlecon(conn net.Conn) {
	var htmlBody = `<h1 style="color:red">hello golang</h1>`
	c := fmt.Sprintf(mainContent, len(htmlBody))
	fmt.Printf("%v", c)
	conn.Write([]byte(fmt.Sprintf(mainContent, len(htmlBody))))
	conn.Write([]byte(htmlBody))

}

func main() {
	addr := ":80"
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handlecon(conn)

	}
}
