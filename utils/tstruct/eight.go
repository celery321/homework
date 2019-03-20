package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Student8 struct {
	Id   int
	Name string
}

func main() {
	s := Student8{
		Id:   2,
		Name: "alice",
	}
	buf, err := json.Marshal(s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(buf))

}
