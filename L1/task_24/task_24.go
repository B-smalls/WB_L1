package main

import (
	"fmt"
	"math"
)

// Point представляет координаты точки.
type Point struct {
	x float64
	y float64
}

// NewPoint создает новую точку с заданными координатами.
func NewPoint(x, y float64) Point {
	return Point{x, y}
}

// Distance вычисляет расстояние между двумя точками.
func Distance(p1, p2 Point) float64 {
	dx := p2.x - p1.x
	dy := p2.y - p1.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	// Создаем две точки
	p1 := NewPoint(1, 2)
	p2 := NewPoint(4, 6)

	// Вычисляем расстояние между ними
	distance := Distance(p1, p2)

	// Выводим результат
	fmt.Printf("Расстояние между точками: %f\n", distance)
}
