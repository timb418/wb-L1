package main

import (
	"fmt"
	"time"
)

func sleepAfter(s time.Duration) {
	<-time.After(s)
}

// эти два механизма эквивалентны
func sleepNewTimer(s time.Duration) {
	<-time.NewTimer(s).C
}

func main() {
	fmt.Println("sleepAfter execution")
	sleepAfter(time.Second * 2)
	fmt.Println("sleepNewTimer execution")
	sleepNewTimer(time.Second * 2)

}
