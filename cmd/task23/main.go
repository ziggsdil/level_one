package main

import (
	"fmt"
	"math/big"
)

func main() {
	q := big.NewInt(1 << 20)
	w := big.NewInt(1 << 20)

	mul := new(big.Int).Mul(q, w)
	sub := new(big.Int).Sub(q, w)
	sum := new(big.Int).Add(q, w)
	div := new(big.Int).Div(q, w)

	fmt.Println(mul)
	fmt.Println(sub)
	fmt.Println(sum)
	fmt.Println(div)
}
