package main

import (
	"fmt"
	"reflect"
)

func main() {
	// реализовать пересечение двух
	// неупорядоченных множеств
	nums1 := []int{2, 2, 5, 1, 23, 4, 6}
	nums2 := []int{3, 1, 4, 2, 6}
	expectedResult := []int{2, 1, 4, 6}

	fmt.Println(intersection(nums1, nums2))
	fmt.Println(reflect.DeepEqual(intersection(nums1, nums2), expectedResult))
}

func intersection(nums1, nums2 []int) []int {
	small, big := nums1, nums2
	// находим наименьший массив, чтобы в дальнешим заполнить им мапу.
	if len(small) > len(big) {
		big, small = small, big
	}

	// создаем мапу в которую будем заносить данные из малого массива
	m := make(map[int]struct{}, len(small))
	for _, key := range small {
		m[key] = struct{}{}
	}

	res := make([]int, 0)
	// пробегаемаеся по значениям наибольшего массива и проверяем есть ли такое значение в мапе
	for _, key := range big {
		if _, ok := m[key]; ok { // если значение есть то добавляем в результирующий массив
			res = append(res, key)
			delete(m, key) // удалем из мапы если уже добавили чтобы не было дубликатов
		}
	}
	return res
}
