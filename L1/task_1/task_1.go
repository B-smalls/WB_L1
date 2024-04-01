package main

//Дана структура Human (с произвольным набором полей и методов).
//Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

import (
	"fmt"
)

// Базовая структура Human
type Human struct {
	Name    string
	Message string
}

// Метод структуры Human.
func (student Human) Say() {
	fmt.Println("Hello! My name is ", student.Name)
}

// Метод структуры Human.
func (student Human) Think(message string) {
	student.Message = message
	fmt.Println("I think aboout...", student.Message)
}

type Action struct {
	Human //Анонимное поле

	//Могут быть свои поля
}

func main() {
	var Action Action

	//Вызов методов структуры Human через объект типа Action
	Action.Name = "Vasua"
	Action.Say()
	Action.Think("ЯЯЯ")
}
