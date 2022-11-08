package main

import "fmt"

func main() {
	firstSet := []int{1, 2, 3, 4, 5}
	secondSet := []int{3, 4, 5, 6, 7, 8}
	fmt.Println(getTwoSetsIntersection(firstSet, secondSet))

	firstSetOfStrings := []string{"cat", "tree", "dog", "animal", "42"}
	secondSetOfStrings := []string{"554", "qwerty", "cat", "dog"}
	fmt.Println(getTwoSetsIntersection(firstSetOfStrings, secondSetOfStrings))
	// Output:
	// [3 4 5]
	// [cat dog]
}

func getTwoSetsIntersection[T comparable](first []T, second []T) []T {
	set := make([]T, 0)
	hash := make(map[T]struct{})

	for _, v := range first {
		hash[v] = struct{}{}
	}

	for _, v := range second {
		if _, ok := hash[v]; ok {
			set = append(set, v)
		}
	}

	return set
}
