package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())
	myticker := time.NewTicker(time.Second)

	go func() {
		for {

			nowtime := <-myticker.C
			fmt.Println(nowtime)
		}

	}()

	for {

	}
}
