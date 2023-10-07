package main

import "fmt"

func main() {
	str := "aabcd"
	fmt.Println(isOnlyUnique(str))
}

func isOnlyUnique(str string) bool {
	bytes := []byte(str)
	m := make(map[byte]int)

	for i := 0; i < len(bytes); i++ {
		m[bytes[i]]++
		if val, ok := m[bytes[i]]; ok && val > 1 {
			return false
		}
	}
	return true
}
