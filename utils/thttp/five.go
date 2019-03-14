package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	resp, err := http.Get("http://www.itcast.cn")
	if err != nil {
		fmt.Printf("", err)
		return
	}

	defer resp.Body.Close()

	fmt.Println("Header", resp.Header)
	fmt.Println("Header", resp.Status)
	fmt.Println("Header", resp.StatusCode)
	fmt.Println("Header", resp.Proto)
	fmt.Println("Header", resp.Body)

	buf := make([]byte, 4096)
	var result string
	for {
		n, _ := resp.Body.Read(buf)
		if n == 0 {
			fmt.Println("完成")
			break
		}
		if err != nil && err != io.EOF {
			fmt.Println(err)
			break
		}
		result += string(buf[:n])
	}
	fmt.Printf("|%v|\n", result)

}
