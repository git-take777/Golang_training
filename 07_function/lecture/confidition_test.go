package main

import (
	"fmt"
	"testing"
)
func confidition() {
	a := 10
	if a > 5 {
		fmt.Println("a is greater than 5")
	} else if a == 6 {
		fmt.Println("a is 6")
	} else {
		fmt.Println("a is 5 or less")
	}
}

func TestCondition(t *testing.T) {
	confidition()
}