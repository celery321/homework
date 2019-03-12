package main

import (
	"fmt"
	"math/rand"
	"time"
)

func readgo6(in <-chan int, idx int) {
	for {
		num := <-in
		fmt.Printf("读%d %d\n", idx, num)
		time.Sleep(time.Second)
	}
}

func writego6(out chan<- int, index int) {
	for {
		num := rand.Intn(1000)
		out <- num
		fmt.Printf("====%d %d\n", index, num)
		time.Sleep(time.Millisecond * 300)
	}

}

func main() {
	rand.Seed(time.Now().UnixNano())

	ch6 := make(chan int)

	for i := 0; i < 5; i++ {
		go readgo5(ch6, i+1)
	}

	for i := 0; i < 5; i++ {
		go writego5(ch6, i+1)
	}
	for {

	}
}
