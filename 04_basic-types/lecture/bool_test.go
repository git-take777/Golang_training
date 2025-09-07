package main

import (
	"fmt"
	"testing"
)

func Testbool1(t *testing.T) {
 bool1()
}

func bool1() {
	var b1 bool = true
	fmt.Println(b1)
	b2 := false
	fmt.Println(b2)
	fmt.Printf("%T, %T\n", b1, b2)
}