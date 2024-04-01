package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Необходимо передать аргумент командной строки")
		return
	}

	arg := os.Args[1]

	// Преобразуем аргумент к типу int
	numSeconds, err := strconv.Atoi(arg)
	if err != nil {
		fmt.Println("Ошибка при преобразовании в число:", err)
		return
	}

	// channel to transfer data
	data := make(chan string)
	reader := bufio.NewReader(os.Stdin)

	go func() {
		// loop through the channel and print data
		for d := range data {
			fmt.Printf("Received: %s\n", d)
		}
	}()

	// before ttl expires read user input and write to channel
	for now := time.Now(); time.Since(now) < time.Second*time.Duration(numSeconds); {
		input, _ := reader.ReadString('\n')
		data <- input
	}
	fmt.Println("Program has finished")
}
