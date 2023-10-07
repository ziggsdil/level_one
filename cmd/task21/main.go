package main

import (
	"fmt"
	"strings"
)

/*
Разработать программу, которая переворачивает
слова в строке. Пример: «snow dog sun — sun dog snow».
*/

func main() {
	str := "snow dog sun"

	fields := strings.Fields(str)
	fmt.Println(reverse(fields))
}

func reverse(arr []string) string {
	res := make([]string, len(arr))
	j := 0
	for i := len(arr) - 1; i >= 0; i-- {
		res[j] = arr[i]
		j++
	}

	return strings.Join(res, " ")
}
