package main

import "fmt"

func main() {
	a := []string{"A", "B", "C", "D", "E"}
	i := 3
	a = remove(a, i)
	fmt.Println(a)
}

func remove(a []string, i int) []string {
	copy(a[i:], a[i+1:])

	return a[:len(a)-1]
}
