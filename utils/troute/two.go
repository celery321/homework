package main

import (
	"fmt"
	"strconv"
)

var ch = make(chan string, 10)

func one() {
	for i := 0; i < 10; i++ {
		ch <- strconv.Itoa(i)
		fmt.Println(i)
	}
}

func two() {
	for {
		fmt.Println(<-ch)

	}

}

func main() {
	go one()
	go two()
	for {

	}
}
