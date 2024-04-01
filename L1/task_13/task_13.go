package main

import "fmt"

func main() {
	var a, b = 1, 2
	//Вывод значений до перестановки
	fmt.Println(a, b)
	a, b = b, a
	//Вывод значений после перестановки
	fmt.Println(a, b)
}
