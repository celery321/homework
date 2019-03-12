package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i * i
			time.Sleep(time.Second * 2)
		}
		close(ch)
	}()
	for {

		select {
		case num := <-ch:
			fmt.Println(num)
		case <-time.After(time.Second * 3):
			fmt.Println("超时")
			return
		}
	}

}
