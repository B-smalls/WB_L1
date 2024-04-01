package main

import "fmt"

/*Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(2^2+3^2+4^2….) с использованием конкурентных вычислений.
 */

// Функция для рассчета квадрата числа. Результат передается дальше в канал result
func square(x int, result chan int) {
	result <- x * x
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	resultChan := make(chan int)

	for _, num := range numbers {
		go square(num, resultChan)
	}

	sum := 0
	for range numbers {
		sum += <-resultChan
	}

	fmt.Println(sum)
}
