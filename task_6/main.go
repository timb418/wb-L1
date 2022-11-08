package main

import (
	"fmt"
	"sync"
	"time"
)

func boolRoutine(exit chan bool) {
	for {
		select {
		case <-exit:
			fmt.Printf("exit bool goroutine\n")
			return
		default:
			fmt.Printf("perform some work in bool goroutine\n")
			time.Sleep(time.Millisecond * 250)
		}
	}
}

func newChan() chan int {
	ch := make(chan int)
	go func() {
		n := 1
		for {
			select {
			case ch <- n:
				fmt.Println("newChan() ", n)
				n++
			case <-ch:
				return
			}
		}
	}()
	return ch
}

func main() {
	exit := make(chan bool)
	defer close(exit)
	go boolRoutine(exit)

	time.Sleep(time.Second * 3)
	exit <- true

	number := newChan()
	fmt.Println(<-number)
	fmt.Println(<-number)
	number <- 333 // отправка любого номера в этот канал завершает его

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		fmt.Println("anon func")
		time.Sleep(time.Second * 2)
		defer wg.Done()
	}()
	wg.Wait() // останавливает основную рутину

}
