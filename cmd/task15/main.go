package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	values := make([]any, 6)
	values[0] = "Hello"
	values[1] = int64(1)
	values[2] = float64(1.3)
	values[3] = []int{1, 2, 3}
	values[4] = true
	values[5] = ch

	for _, value := range values {
		valType := whichType(value)
		fmt.Println(valType)
	}
}

func whichType(val interface{}) string {
	switch val.(type) {
	case string:
		return "string"
	case int64:
		return "int64"
	case float64:
		return "float64"
	case []int:
		return "[]int"
	case bool:
		return "bool"
	case chan int:
		return "chan int"
	default:
		return "undefined type"
	}
}
