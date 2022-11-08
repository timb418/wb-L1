package main

import (
	"fmt"
	"math/big"
)

// Разработать программу, которая перемножает, делит, складывает, вычитает
// две числовых переменных a,b, значение которых > 2^20
func main() {
	a := new(big.Int)
	a.SetString("456789000000000000000000", 10)
	b := new(big.Int)
	b.SetString("387654000000001230000000", 10)

	division := new(big.Int)
	division.Div(a, b)
	fmt.Printf("division       result is %v\n", division)

	multiplication := new(big.Int)
	multiplication.Mul(a, b)
	fmt.Printf("multiplication result is %v\n", multiplication)

	sum := new(big.Int)
	sum.Add(a, b)
	fmt.Printf("sum            result is %v\n", sum)

	subtraction := new(big.Int)
	subtraction.Sub(a, b)
	fmt.Printf("subtraction    result is %v\n", subtraction)
}
