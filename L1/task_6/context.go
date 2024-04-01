package main

import (
	"context"
	"fmt"
	"strconv"
	"time"
)

func main() {
	c := make(chan string)
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		i := 0
		for {

			select {
			case <-ctx.Done():
				c <- "Context has closed"
				return
			default:
				fmt.Println(strconv.Itoa(i))
			}

			i++
			time.Sleep(500 * time.Millisecond)
		}
	}(ctx)

	go func() {

		time.Sleep(3 * time.Second)
		cancel()
	}()

	fmt.Println(<-c)
}
