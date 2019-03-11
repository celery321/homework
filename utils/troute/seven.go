package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())

	mytime := time.NewTimer(time.Second * 2)
	nowtime := <-mytime.C
	fmt.Println(nowtime)

	fmt.Println(time.Now())
	nowtim := <-time.After(time.Second * 2)
	fmt.Println(nowtim)
}
