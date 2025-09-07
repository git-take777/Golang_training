package main

import (
	"fmt"
	"testing"
)
func comparisonOperators() {
	// Practiceなので冗長になっている。GoはPHPと違って 「===」はない。
	fmt.Println(true && false == true)
	fmt.Println(true && true == true)

	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(false || true)
	fmt.Println(false || false)

	fmt.Println(!true)
	fmt.Println(!false)
}

func TestLogicalOperators(t *testing.T) {
	comparisonOperators()
}