package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(nums)
	fmt.Println(deleteElement(nums, 2))
}

func deleteElement(nums []int, idx int) []int {
	res := make([]int, 0)

	res = append(nums[:idx], nums[idx+1:]...)
	return res
}
