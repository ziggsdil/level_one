package main

import (
	"fmt"
	"sort"
)

/*
Реализовать быструю сотртировку встренными методами языка.
*/

func main() {
	nums := []int{3, 4, 5, 2, 1, 2}
	fmt.Printf("nums: %v, is sorted = %v\n", nums, sort.IntsAreSorted(nums))
	sortedArr := quickSort(nums)
	fmt.Printf("nums: %v, is sorted = %v\n", sortedArr, sort.IntsAreSorted(sortedArr))
}

func quickSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	pivot := nums[0]
	var less, greater []int
	for _, num := range nums[1:] {
		if num <= pivot {
			less = append(less, num)
		} else {
			greater = append(greater, num)
		}
	}
	res := append(quickSort(less), pivot)
	res = append(res, quickSort(greater)...)
	return res
}
