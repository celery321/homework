package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var rw sync.RWMutex

func readgo(in <-chan int, idx int) {
	for {
		rw.RLock()
		num := <-in
		fmt.Printf("%d %d\n", idx, num)
		rw.RUnlock()
	}
}

func writego(out chan<- int, index int) {
	for {
		num := rand.Intn(1000)
		//rw.Lock()
		out <- num
		fmt.Printf("%d %d\n", index, num)
		time.Sleep(time.Millisecond * 300)
		//rw.Unlock()
	}

}

func main() {
	rand.Seed(time.Now().UnixNano())

	//quit := make(chan bool)
	ch := make(chan int)

	for i := 0; i < 5; i++ {
		go readgo(ch, i+1)
	}

	for i := 0; i < 5; i++ {
		go writego(ch, i+1)
	}
	for {

	}
}
