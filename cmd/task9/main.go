package main

import "fmt"

// дана переменная int64, разработать прогармму
// которая устнавливает i-й бит в 1 или 0

func main() {
	var x int64 = 123
	fmt.Println(int64ToBinary(x)) //   1 1 1 1 0 1 1
	num := changeBit(x, 5, false)
	fmt.Println(int64ToBinary(num)) // 1 0 1 1 0 1 1
}

func int64ToBinary(num int64) string {
	res := ""
	for i := 63; i >= 0; i-- {
		bit := (num >> uint(i)) & 1
		res += fmt.Sprintf("%d", bit)
	}
	return res
}

func changeBit(num int64, idx int, val bool) int64 {
	mask := int64(1) << idx // маска на позицию которую нужно поменять
	if !val {
		num = num &^ mask // и исключащее или
	} else {
		num = num | mask // или
	}

	return num
}
