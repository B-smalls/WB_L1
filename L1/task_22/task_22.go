package main

import (
	"fmt"
	"math/big"
)

func main() {

	first := big.NewFloat(19980435006.234)
	second := big.NewFloat(20400761021.3214)

	// addition
	add := new(big.Float)
	fmt.Printf("%f + %f = ", first, second)
	fmt.Println(add.Add(first, second))

	// subtraction
	sub := new(big.Float)
	fmt.Printf("%f - %f = ", first, second)
	fmt.Println(sub.Sub(first, second))

	// division
	div := new(big.Float)
	fmt.Printf("%f / %f = ", first, second)
	fmt.Println(div.Quo(second, first))
}
