package main

import (
	"fmt"
	"sync"
	"time"
)

var mutx sync.Mutex

func printer(s string) {
	mutx.Lock()
	for _, name := range s {

		fmt.Printf("%c", name)

		time.Sleep(time.Second)
	}
	mutx.Unlock()
}

func one(s string) {
	printer(s)
}

func two(s string) {
	printer(s)
}

func main() {

	go one("aaa")
	go two("bbb")
	for {

	}

}
