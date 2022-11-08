package main

import "fmt"

func main() {
	fmt.Println(binarySearch([]int{1, 5, 6, 7, 10, 12, 15, 18}, 15))
}

func binarySearch(sortedList []int, target int) int {
	var lo = 0
	var hi = len(sortedList) - 1

	for lo <= hi {
		var mid = lo + (hi-lo)/2
		var midValue = sortedList[mid]

		if midValue == target {
			return mid
		} else if midValue > target {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}

	return -1
}
