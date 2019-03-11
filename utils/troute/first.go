package main

import (
	"fmt"
	"time"
)

var ch = make(chan string)

func printer(s string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%s\n", s)
		time.Sleep(time.Microsecond)
	}
}

func one(s string) {

	printer(s)
	ch <- "1"

}

func two(s string) {
	<-ch
	printer(s)
}

func main() {
	go one("aaaaa")
	go two("bbbbb")
	for {

	}

}
