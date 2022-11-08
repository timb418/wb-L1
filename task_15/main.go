package main

import (
	"fmt"
	"unsafe"
)

var justString string

func main() {
	someFunc()
}

func someFunc() {
	v := createHugeString()
	justString = v[:100]
	// что плохо: глобальная переменная
	//2) отсечение строки как слайса, что разбивает строку на массив байт и в случае если там руны по несколько байт, то отсечёт неправильно и вообще, отсечет не 100 символов, а 100 байт, в русском языке это около 50 символов получается.
	// корректно:
	res := make([]rune, 0, 100)
	res = []rune("АААААААААААААААААААААААААААААААААААААААААААААААААААААААААААААААААААААААААААААААААААААААААААААААААААА")[:100]
	justString2 := string(res)

	fmt.Println("v = ", v)
	fmt.Println("justString = ", justString)
	fmt.Println("justString2 = ", justString2)

	fmt.Println("len(v) = ", len(v))
	fmt.Println("len(justString) = ", len(justString))
	fmt.Println("len(justString2) = ", len(justString2))

	fmt.Println("cap(v) = ", cap([]byte(v)))
	fmt.Println("cap(justString) = ", cap([]byte(justString)))
	fmt.Println("cap(justString2) = ", cap([]byte(justString2)))

	fmt.Println("sizeof(v) = ", unsafe.Sizeof(v))
	fmt.Println("sizeof(justString) = ", unsafe.Sizeof(justString))
	fmt.Println("sizeof(justString2) = ", unsafe.Sizeof(justString2))

	fmt.Printf("Pointer(v) = %p\n", &v)
	fmt.Printf("Pointer(justString) = %p\n", &justString)
	fmt.Printf("Pointer(justString2) = %p\n", &justString2)
}

func createHugeString() string {
	return "АААААААААААААААААААААААААААААААААААААААААААААААААААААААААААААААААААААААААААААААААААААААААААААААААААА"
}
