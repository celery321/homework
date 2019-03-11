package main

import "fmt"

func main() {
	ch := make(chan int)

	var send chan<- int = ch
	send <- 789
	num := <-send

	fmt.Println(num)
}
