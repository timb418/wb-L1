package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(checkAllLettersUnique("abcd"))      // expected true
	fmt.Println(checkAllLettersUnique("abCdefAaf")) // expected false
	fmt.Println(checkAllLettersUnique("aabcd"))     // expected false
}

func checkAllLettersUnique(s string) bool {
	s = strings.ToLower(s)
	var onceSeen []rune
	for _, symbol := range s {
		if !contains(onceSeen, symbol) {
			onceSeen = append(onceSeen, symbol)
		}
	}
	if len(onceSeen) != len(s) {
		return false
	}
	return true
}

func contains(s []rune, compareTo rune) bool {
	for _, a := range s {
		if a == compareTo {
			return true
		}
	}
	return false
}
