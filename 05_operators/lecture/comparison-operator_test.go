package main

import (
	"fmt"
	"testing"
)

func comparisonOperator() {
	fmt.Println(1 == 1)
	fmt.Println(1 == 2)
	fmt.Println(1 > 4)
	fmt.Println(1 < 4)
	fmt.Println(1 >= 1)
	fmt.Println(1 <= 1)
	fmt.Println(1 != 1)
}

func TestComparisonOperator(t *testing.T) {
	comparisonOperator()
}