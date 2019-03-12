package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

func main() {
	//host := os.Args[1]
	//password := os.Args[2]
	cmd := exec.Command("ping ", "-w 1 8.8.8.8")

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := cmd.Start(); err != nil {
		fmt.Println(err)
		return
	}

	bytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		fmt.Println(err)
		return
	}

	err1 := cmd.Wait()
	if err1 != nil {
		fmt.Println(err1)
	}

	fmt.Println(bytes)
}
