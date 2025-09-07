package main

import (
	"fmt"
	"testing"
)

func arithmeticOperators() {
	fmt.Println(1 + 4)
	fmt.Println(1 - 4)
	fmt.Println(1 * 4)
	fmt.Println(60 / 4)
	fmt.Println(1 % 4)

	n := 0
	n += 4 // n = n + 4
	fmt.Println(n)

	s := "Hello "
	s += "World"
	fmt.Println(s)
}

func TestArithmeticOperators(t *testing.T) {
	arithmeticOperators()
}
