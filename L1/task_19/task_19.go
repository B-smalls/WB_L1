package main

import (
	"fmt"
)

// Функция, которая переворачивает строку
func reverseString(str string) string {
	// Переводим строку в массив рун (unicode символов)
	runes := []rune(str)

	// Получаем длину строки
	length := len(runes)

	// Переворачиваем массив рун
	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	// Возвращаем перевернутую строку
	return string(runes)
}

func main() {
	// Строка для переворота
	str := "главрыба — абырвалг"

	// Переворачиваем строку
	reversed := reverseString(str)

	// Выводим результат
	fmt.Println("Исходная строка:", str)
	fmt.Println("Перевернутая строка:", reversed)
}
