package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func writeToChannel(ch chan<- string) {
	for {
		data := "Some data" // Данные передающиеся в канал
		ch <- data
		time.Sleep(time.Second) // Пауза в одну секунду между записями
	}
}

func worker(id int, ch <-chan string, done <-chan struct{}) {
	for {
		select {
		case data := <-ch:
			fmt.Printf("Worker %d received data: %s\n", id, data)
		case <-done:
			fmt.Printf("Worker %d stopped\n", id)
			return
		}
	}
}

func main() {
	// Проверяем, что передан хотя бы один аргумент командной строки
	if len(os.Args) < 2 {
		fmt.Println("Необходимо передать аргумент командной строки")
		return
	}

	arg := os.Args[1]

	// Преобразуем аргумент к типу int
	numWorkers, err := strconv.Atoi(arg)
	if err != nil {
		fmt.Println("Ошибка при преобразовании в число:", err)
		return
	}

	channel := make(chan string)
	done := make(chan struct{})

	// Запускаем функцию для записи данных в канал
	go writeToChannel(channel)

	// Запускаем воркеров
	for i := 1; i <= numWorkers; i++ {
		go worker(i, channel, done)
	}

	// Ждем сигнала о завершении работы программы
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan

	// Отправляем сигнал завершения всем воркерам
	close(done)

	// Ожидаем завершения всех воркеров
	for i := 1; i <= numWorkers; i++ {
		<-done
	}

	fmt.Println("Программа завершена")
}
