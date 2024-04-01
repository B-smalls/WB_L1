package main

import (
	"fmt"
	"sync"
)

type ConcurrentMap struct {
	Mutex sync.Mutex
	Entry map[int]int
}

func (m *ConcurrentMap) Add(key int, value int) {
	m.Mutex.Lock()
	m.Entry[key] = value
	m.Mutex.Unlock()
}

func main() {

	m := ConcurrentMap{
		Entry: make(map[int]int),
	}

	wg := sync.WaitGroup{}
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func(i int) {
			m.Add(i, i)
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Printf("%#v", m.Entry)
}
