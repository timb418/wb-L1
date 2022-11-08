package main

import "fmt"

// Разработать программу, которая переворачивает подаваемую на ход строку
// например: «главрыба — абырвалг». Символы могут быть unicode.

func main() {
	fmt.Println(reverseString("главрыба"))

}

func reverseString(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
