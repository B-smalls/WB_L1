package main

import (
	"fmt"
	"sync"
)

// Counter представляет структуру счетчика.
type Counter struct {
	mu    sync.Mutex
	value int
}

// Inc инкрементирует значение счетчика.
func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// Value возвращает текущее значение счетчика.
func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func main() {
	// Создаем новый счетчик.
	counter := Counter{}

	// Создаем WaitGroup для ожидания завершения всех горутин.
	var wg sync.WaitGroup

	// Запускаем горутины для инкрементации счетчика.
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Inc()
		}()
	}

	// Ждем завершения всех горутин.
	wg.Wait()

	// Выводим итоговое значение счетчика.
	fmt.Println("Итоговое значение счетчика:", counter.Value())
}
