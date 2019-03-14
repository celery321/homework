package main

import (
	"fmt"
	"net/http"
	"os"
)

func openSendfile(fname string, w http.ResponseWriter) {
	pathName := "c:/Go" + fname

	f, err := os.Open(pathName)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte("no file"))
		return
	}
	defer f.Close()

	buf := make([]byte, 4096)
	for {
		n, _ := f.Read(buf)
		if n == 0 {
			return
		}
		w.Write(buf[:n])
	}
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("url", r.URL)
	openSendfile(r.URL.String(), w)
}

func main() {
	http.HandleFunc("/", myHandler)
	http.ListenAndServe("127.0.0.1:8000", nil)
}
