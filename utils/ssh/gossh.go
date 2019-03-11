package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	//host := os.Args[1]
	//password := os.Args[2]
	cmd := exec.Command("", "./")
	ppReader, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(ppReader)
}
