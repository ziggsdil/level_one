package main

import "fmt"

func main() {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	target := 9

	idx := binarySearch(nums, target)
	fmt.Println("idx of target is: ", idx)
}

// возвращает индекс найденного элемента.
func binarySearch(nums []int, target int) int {
	l, r, mid := 0, len(nums)-1, 0

	for l <= r {
		mid = (l + r) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1 // если такого числа нет
}
