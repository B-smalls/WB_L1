package main

import (
	"fmt"
)

func setBit(bitValue int64, position, value int) int64 {
	if value == 1 {
		return bitValue | 1<<position // Устанавливаем i-й бит в 1
	} else {
		return bitValue &^ (1 << position) // Устанавливаем i-й бит в 0
	}

}

func main() {

	var number int64 = 5

	fmt.Println(setBit(number, 1, 1))
}
