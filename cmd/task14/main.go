package main

import "fmt"

// Поменять местами два числа без создания временной переменной
func main() {
	x, y := 1, 2
	fmt.Printf("x = %d y = %d\n", x, y)

	x, y = y, x
	fmt.Printf("x = %d y = %d\n", x, y)
}
