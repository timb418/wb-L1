package main

import "fmt"

func main() {
	fmt.Println(getActualDataType(nil))
	fmt.Println(getActualDataType(42))
	fmt.Println(getActualDataType("str"))
	fmt.Println(getActualDataType(true))
	fmt.Println(getActualDataType(make(chan int)))
	fmt.Println(getActualDataType(42.5))
}

func getActualDataType(value interface{}) string {
	switch value.(type) {
	case nil:
		return "Passed interface is nil"
	case int:
		return "Passed interface is an int"
	case string:
		return "Passed interface is a string"
	case bool:
		return "Passed interface is a boolean"
	case chan int:
		return "Passed interface is a chan int"
	default:
		return "Passed interface type is unknown"
	}
}
