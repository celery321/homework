package main

import "net/http"

func hander(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))

}

func main() {

	http.HandleFunc("/itcast", hander)
	http.ListenAndServe("0.0.0.0:8000", nil)

}
