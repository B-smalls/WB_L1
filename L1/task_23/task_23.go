package main

import "fmt"

func main() {
	// Создаем срез
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Исходный срез:", slice)

	// Указываем индекс элемента, который нужно удалить
	index := 2

	// Удаляем элемент из среза
	slice = append(slice[:index], slice[index+1:]...)
	fmt.Println("Срез после удаления элемента:", slice)
}
