package main

import "fmt"

// Функция для нахождения пересечения двух множеств, представленных отображениями
func intersection(set1, set2 map[int]bool) map[int]bool {
	result := make(map[int]bool)

	// Перебираем элементы первого множества
	for k := range set1 {
		// Если элемент присутствует во втором множестве, добавляем его в результат
		if set2[k] {
			result[k] = true
		}
	}

	return result
}

// Функция для нахождения пересечения двух неупорядоченных множеств, представленных срезами
/*func intersection(set1, set2 []int) []int {
	setMap := make(map[int]bool)
	var result []int

	// Создаем мапу для первого множества для быстрого поиска
	for _, val := range set1 {
		setMap[val] = true
	}

	// Проверяем элементы второго множества на наличие в первом
	for _, val := range set2 {
		if setMap[val] {
			result = append(result, val)
		}
	}

	return result
}*/

func main() {
	// Примеры двух неупорядоченных множеств
	set1 := map[int]bool{1: true, 2: true, 3: true, 4: true}
	set2 := map[int]bool{3: true, 4: true, 5: true, 6: true}

	// Находим пересечение множеств
	result := intersection(set1, set2)

	// Выводим результат
	fmt.Println("Пересечение множеств:", result)
}
