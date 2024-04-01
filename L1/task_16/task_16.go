package main

import (
	"fmt"
	"sort"
)

func main() {
	// Исходный массив
	arr := []int{3, 2, 5, 1, 4}

	// Функция для реализации интерфейса sort.Interface
	// необходимо для использования sort.Sort
	sortSlice := sort.IntSlice(arr)

	// Вызов быстрой сортировки
	sort.Sort(sortSlice)

	// Вывод отсортированного массива
	fmt.Println("Отсортированный массив:", arr)
}
