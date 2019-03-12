package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var value int
var rw1 sync.RWMutex

func readgo5(idx int) {
	for {
		rw1.RLock()
		num := value
		fmt.Printf("读%d %d\n", idx, num)
		rw1.RUnlock()
		time.Sleep(time.Second)
	}
}

func writego5(index int) {
	for {
		num := rand.Intn(1000)
		rw1.Lock()
		fmt.Printf("====%d %d\n", index, num)
		time.Sleep(time.Millisecond * 300)
		rw1.Unlock()
	}

}

func main() {
	rand.Seed(time.Now().UnixNano())

	//quit := make(chan bool)

	for i := 0; i < 5; i++ {
		go readgo5(i + 1)
	}

	for i := 0; i < 5; i++ {
		go writego5(i + 1)
	}
	for {

	}
}
