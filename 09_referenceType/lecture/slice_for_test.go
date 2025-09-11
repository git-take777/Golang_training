package main

import (
	"fmt"
	"testing"
)

func TestSliceFor(t *testing.T) {
	sliceFor()
}
func sliceFor() {
	s :=[]string{"A","B","C","D","E"}
	for _, v := range s {
		fmt.Println(v)
	}
	// for i, _ := range s {
	// 	fmt.Println(i)
	// }
	for i :=0; i < len(s); i++ {
		fmt.Println(s[i])
	}
}