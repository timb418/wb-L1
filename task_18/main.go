package main

import (
	"fmt"
	"sync"
)

type counter struct {
	iterations int
}

func (c *counter) increment() {
	c.iterations++
}

func main() {
	increment := counter{0}
	wg := sync.WaitGroup{}
	m := sync.Mutex{}

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			m.Lock()
			increment.increment()
			m.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println(increment)
}
