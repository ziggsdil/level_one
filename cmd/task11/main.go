package main

import "fmt"

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	for key, val := range groupTemperatures(temperatures) {
		fmt.Printf("%d: %v\n", key, val)
	}
}

func groupTemperatures(temps []float64) map[int][]float64 {
	res := make(map[int][]float64)

	for _, temp := range temps {
		group := int(temp/10) * 10
		res[group] = append(res[group], temp)
	}
	return res
}
