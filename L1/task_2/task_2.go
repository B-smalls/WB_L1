package main

import (
	"fmt"
	"sync"
)

//Написать программу, которая конкурентно рассчитает значение
//квадратов чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

func main() {
	num_list := [...]int{2, 4, 6, 8, 10}

	var waitGroup sync.WaitGroup

	waitGroup.Add(len(num_list))

	for _, num := range num_list {
		go func(num int) {
			fmt.Println(num * num)

			waitGroup.Done()
		}(num)
	}

	waitGroup.Wait()
}
