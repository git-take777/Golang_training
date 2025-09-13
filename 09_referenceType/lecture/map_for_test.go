package main

import (
	"fmt"
	"testing"
)

func MapFor() {
	m := map[string]int{"Apple":100, "Banana":200, "Orange":300}
	for k, v := range m {
		fmt.Println(k, v)
	}
}
func TestMapFor(t *testing.T) {
	MapFor()
}