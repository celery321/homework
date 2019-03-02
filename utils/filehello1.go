package utils

import (
	"log"
	"os"
)

func FileHello1() {
	f, err := os.Open("1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
}
