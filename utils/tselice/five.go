package main

import "fmt"

func noEmpty(data []string) []string {
	out := make([]string, 0)
	for _, str := range data {
		if str != "" {
			out = append(out, str)
		}
	}

	return out
}

func noEmpty2(data []string) []string {
	i := 0
	for _, str := range data {
		if str != "" {
			data[i] = str
			i++
		}
	}

	return data[:i]
}

func main() {
	data := []string{"aaa", "bbbb", "", "cccc", "", ""}

	afterData := noEmpty2(data)
	fmt.Println(afterData)
}
