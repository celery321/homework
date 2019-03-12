package main

import (
	"fmt"
	"math/rand"
	"time"
)

func producer(out chan<- int, idx int) {
	for i := 0; i < 5; i++ {
		num := rand.Intn(800)
		fmt.Printf("生产者%d %d\n", num, idx)
		out <- num
	}
	close(out)
}

func consumer(in <-chan int, idx int) {
	for num := range in {
		fmt.Printf("---消费者:%d %d\n", idx, num)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	p := make(chan int)

	for i := 0; i < 5; i++ {

		go producer(p, i+1)
	}

	for i := 0; i < 5; i++ {

		go consumer(p, i+1)
	}

	for {

	}
}
