package main

import (
	"fmt"
	"testing"
)

func TestInterface1(t * testing.T) {
	interface1()
}

func interface1() {
	var x interface{}
	fmt.Printf("%T\n", x)

	x = 100
	fmt.Println(x)
}