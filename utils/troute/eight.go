package main

import (
	"fmt"
	"time"
)

func main() {
	mytime := time.NewTimer(time.Second * 3)
	mytime.Reset(1 * time.Second) //停止
	go func() {
		<-mytime.C
		fmt.Println("完毕")
	}()
	//mytime.Stop() // 重置
	for {

	}
}
