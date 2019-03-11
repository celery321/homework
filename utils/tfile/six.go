package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func copyTodir(src, dst string) {
	f, err := os.Open(src)
	if err != nil {
		fmt.Println(err)

	}

	defer f.Close()

	fw, err := os.Create(dst)
	if err != nil {
		fmt.Println(err)
	}
	defer fw.Close()

	buf := make([]byte, 4096)

	for {
		n, err := f.Read(buf)
		if err != nil && err == io.EOF {
			fmt.Println("读完", n)
			return
		}
		fw.Write(buf[:n])

	}
}

func main() {
	fmt.Println(":")
	var path string
	fmt.Scan(&path)

	f, err := os.OpenFile(path, os.O_RDONLY, os.ModeDir)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	info, err := f.Readdir(-1)

	for _, filename := range info {
		if filename.IsDir() {
			// 是目录
			//fmt.Println(filename.Name())
		} else {
			if strings.HasSuffix(filename.Name(), ".mp4") {
				copyTodir(path+"/"+filename.Name(), "./"+filename.Name())
			}
		}
	}

}
