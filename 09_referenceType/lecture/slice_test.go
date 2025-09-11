package main

import (
	"fmt"
	"testing"
)

func slice1() {
	slice1 := []int{1, 2, 3, 4, 5}
	fmt.Printf("%T\n", slice1)
	fmt.Println(slice1)
	sl3 := []string{"A", "B", "C"}
	fmt.Println(sl3)
	sl4 := make([]int, 5)
	sl5 := make([]int, 5, 10)

	fmt.Println(sl4)
	fmt.Println(sl5)

	slice1[0] = 100
	fmt.Println(slice1)

	fmt.Println(slice1[2:4])
	fmt.Println(slice1[:4])
	fmt.Println(slice1[2:])
	fmt.Println(slice1[:])
}

func TestSlice1(t*testing.T) {
	slice1()
}