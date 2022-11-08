package main

import "fmt"

type Human struct {
	name string
	age  int
}

type Action struct {
	Human
}

func (h *Human) getName() string {
	return h.name
}

func main() {
	h := Human{name: "Ivan", age: 30}
	action := Action{h}

	fmt.Println(h.getName())
	fmt.Println(action.Human.getName())
}
