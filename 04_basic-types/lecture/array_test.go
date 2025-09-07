package main

import (
	"fmt"
	"testing"
)
func array1(){
	var arr1 [3]int
	fmt.Printf("%T\n", arr1)
}

func TestArray1(t * testing.T) {
	array1()
}