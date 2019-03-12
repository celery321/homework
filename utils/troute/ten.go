package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	ch := make(chan int)
	qiut := make(chan bool)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
			time.Sleep(time.Second)
		}
		close(ch)
		qiut <- true
		runtime.Goexit()
	}()

	for {
		select {
		case num := <-ch:
			fmt.Println(num)
		case <-qiut:
			//break
			return

		}
		fmt.Println()
	}
}
