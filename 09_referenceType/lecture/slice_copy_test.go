package main

import (
	"fmt"
	"testing"
)

func sliceCopy() {
	sl :=[]int{100,200,300,400,500}
	sl2 := make([]int,5, 10)
	// sl2 = copy(sl2, sl)→なぜこれはダメ??
	n := copy(sl2, sl)

	fmt.Println(n, sl2)
	fmt.Println(len(sl2))
	fmt.Println(cap(sl2))
}

func TestSliceCopy(t *testing.T) {
	sliceCopy()
}