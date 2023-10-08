package main

// использование глобальных переменных является плохой практикой.
var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	if len(v) >= 100 {
		justString = v[:100] // возможна паника out of range
	}
}

func createHugeString(size int) string {
	res := ""
	for i := 0; i < size; i++ {
		res += string(rune(i))
	}

	return res
}

func main() {
	someFunc() // не передаем никаких параметров а просто вызываем.
}
