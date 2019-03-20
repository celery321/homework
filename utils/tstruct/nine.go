package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Student9 struct {
	Id   int
	Name string
}

func main() {
	str := `{"Id":2, "Name":"bob"}`
	var s Student9
	err := json.Unmarshal([]byte(str), &s)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(s)
}
