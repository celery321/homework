package utils

import (
	"fmt"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")

}

func httphello() {
	http.HandleFunc("/", IndexHandler)
	http.ListenAndServe("127.0.0.1:80", nil)

}
