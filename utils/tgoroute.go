package main

import (
	"fmt"
	"time"
)

func sing() {
	for i := 0; i < 10; i++ {
		fmt.Println("sing")
		time.Sleep(time.Second)
	}

}

func dance() {
	for i := 0; i < 10; i++ {
		fmt.Println("dance")
		time.Sleep(time.Second)
	}

}

func main() {

	go sing()
	go dance()
	for {

	}
}
