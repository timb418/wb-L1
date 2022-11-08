package main

import "fmt"

//Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из
//массива, во второй — результат операции x*2, после чего данные из второго
//канала должны выводиться в stdout

func doubleGivenValue(firstChan chan int, secondChan chan int) {
	for v := range firstChan {
		secondChan <- v * 2
	}
}

func main() {
	numbers := [...]int{1, 2, 3, 4, 5, 6, 7, 8}

	firstChan := make(chan int)
	secondChan := make(chan int)
	defer close(firstChan)
	defer close(secondChan)

	go doubleGivenValue(firstChan, secondChan)

	for _, val := range numbers {
		firstChan <- val
		fmt.Println(<-secondChan)
	}

}
