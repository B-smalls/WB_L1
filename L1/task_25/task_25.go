package main

import (
	"fmt"
	"time"
)

// Функция sleep принимает количество секунд для ожидания и приостанавливает выполнение программы на указанное количество секунд.
func sleep(seconds int) {
	<-time.After(time.Duration(seconds) * time.Second)
}

func main() {
	fmt.Println("Начало программы")

	// Приостанавливаем выполнение программы на 3 секунды
	sleep(3)

	fmt.Println("Прошло 3 секунды")
}
