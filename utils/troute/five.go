package main

import "fmt"

func send(out chan<- int) {
	out <- 89
	close(out)
}

func recv(in <-chan int) {
	n := <-in
	fmt.Println(n)
}

func main() {
	ch := make(chan int)
	go func() {
		send(ch)
	}()

	recv(ch)
}
