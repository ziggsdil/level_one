package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isOnlyUnique("str"))
	fmt.Println(isOnlyUnique("abCdefAaf"))
}

func isOnlyUnique(str string) bool {
	newStr := strings.ToLower(str)
	bytes := []byte(newStr)
	m := make(map[byte]struct{})

	for i := 0; i < len(bytes); i++ {
		if _, ok := m[bytes[i]]; ok {
			return false
		}
		m[bytes[i]] = struct{}{}
	}
	return true
}
