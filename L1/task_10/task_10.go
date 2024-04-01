package main

import (
	"fmt"
	"sort"
)

func main() {
	// Исходные данные
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	// Сортировка исходных данных
	sort.Float64s(temperatures)

	// Создание словаря для хранения групп температур
	temperatureGroups := make(map[int][]float64)

	// Разделение данных на интервалы
	for _, temp := range temperatures {
		group := int(temp/10) * 10 // Находим группу для текущей температуры
		temperatureGroups[group] = append(temperatureGroups[group], temp)
	}

	// Вывод результатов
	for group, temps := range temperatureGroups {
		fmt.Printf("Группа %d: %v\n", group, temps)
	}
}
