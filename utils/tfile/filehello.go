package main

import (
	"io/ioutil"
	"os"
)

func FileHello() {
	body, _ := ioutil.ReadAll(os.Stdin)
	os.Stdout.Write(body)
}
