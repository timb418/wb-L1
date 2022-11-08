//package main
//
//import (
//	"fmt"
//	"sync"
//)
//
//func sumOfSquares(num int) {
//	m := sync.Mutex{}
//	m.Lock()
//	sum += num * num
//	m.Unlock()
//	fmt.Println(num)
//}
//
//var sum int
//
//func main() {
//	arr := [5]int{2, 4, 6, 8, 10}
//	wg := sync.WaitGroup{}
//	//m := sync.Mutex{}
//	for _, num := range arr {
//		wg.Add(1)
//
//		go func(num int) {
//			sumOfSquares(num)
//			wg.Done()
//		}(num)
//	}
//	wg.Wait()
//
//	fmt.Println(sum)
//}

package main

import (
	"fmt"
	"sync"
)

//func sumOfSquares(num int) {
//	m := sync.Mutex{}
//	m.Lock()
//	sum += num * num
//	m.Unlock()
//	fmt.Println(num)
//}

func main() {
	arr := [5]int{2, 4, 6, 8, 10}
	var sum int

	wg := sync.WaitGroup{}
	m := sync.Mutex{}

	for _, num := range arr {
		wg.Add(1)

		go func(num int) {
			m.Lock()
			sum += num * num
			m.Unlock()
			wg.Done()
		}(num)
	}
	wg.Wait()

	fmt.Printf("Sum is %d", sum)
}
