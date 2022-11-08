package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func worker(id int, c chan int) {
	for data := range c {
		time.Sleep(time.Second)
		fmt.Printf("Worker %d with data %d\n", id, data)
	}
}

func main() {
	c := make(chan int)
	defer close(c)

	workersAmount := 3

	for i := 1; i <= workersAmount; i++ {
		go worker(i, c)
	}

	signalHandle := make(chan os.Signal)
	signal.Notify(signalHandle, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-signalHandle
		close(c)
		fmt.Println("Exiting ...")
		os.Exit(1)
	}()

	for {
		data := rand.Intn(50)
		c <- data
	}

}
