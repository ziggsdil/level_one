package main

import "fmt"

func main() {
	str := "главрыба"
	fmt.Println(string(reverse(str)))
}

func reverse(str string) []rune {
	runes := []rune(str)
	res := make([]rune, len(runes))
	l := len(runes)

	for i, run := range runes {
		res[l-1-i] = run
	}

	return res
}
