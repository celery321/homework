package main

import (
	"fmt"
	"net/http"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this is a web server"))
	fmt.Println("header ", r.Header["11"])
	for key, value := range r.Header {
		fmt.Println(key, value)
	}
	fmt.Println("URL", r.URL)
	fmt.Println("Method", r.Method)
	fmt.Println("Host", r.Host)
	fmt.Println("RemoteAddr", r.RemoteAddr)
	fmt.Println("Body", r.Body)
}

func main() {
	http.HandleFunc("/itcast", myHandler)
	http.ListenAndServe("127.0.0.1:8000", nil)
}
