package main

import (
	"fmt"
)

// binarySearch выполняет бинарный поиск в отсортированном массиве arr
// и возвращает индекс элемента x, если он найден, иначе -1.
func binarySearch(arr []int, x int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == x {
			return mid
		}

		if arr[mid] < x {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func main() {
	arr := []int{2, 3, 4, 10, 40}
	x := 10

	result := binarySearch(arr, x)
	if result != -1 {
		fmt.Printf("Элемент %d найден в позиции %d\n", x, result)
	} else {
		fmt.Printf("Элемент %d не найден в массиве\n", x)
	}
}
