package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	c := make(chan string, 6)
	done := make(chan struct{})

	go func() {
		i := 0
		for {
			select {
			case c <- strconv.Itoa(i):
			case <-done:
				close(c)
				return
			}

			i++
			time.Sleep(100 * time.Millisecond)
		}
	}()

	go func() {
		time.Sleep(2 * time.Second)
		done <- struct{}{}
	}()

	for i := range c {
		fmt.Printf("Received: %s\n", i)
	}
}
