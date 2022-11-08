package main

import "fmt"

// Дана переменная int64. Разработать программу которая устанавливает i-й бит в
// 1 или 0

func main() {
	var number int64 = 100 // 1100100
	var pos uint = 1

	fmt.Println(setBit(number, pos))   // 102 == 1100110
	fmt.Println(clearBit(number, pos)) // 100 == 1100100
}

func setBit(n int64, pos uint) int64 {
	return n | 1<<pos
}

func clearBit(n int64, pos uint) int64 {
	return n & ^(1 << pos)
}
