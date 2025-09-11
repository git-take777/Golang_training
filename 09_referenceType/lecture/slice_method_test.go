package main

import (
	"fmt"
	"testing"
)
func sliceMethod() {
	// capは容量を気にする際に使用する。
	sl :=[]int{100,200,300}
	fmt.Println(sl)
	sl = append(sl,400)
	fmt.Println(sl)
	fmt.Println(len(sl))
	fmt.Println(cap(sl))
	sl2 := make([]int,5)
	fmt.Println(sl2)
	fmt.Println(len(sl2))
	fmt.Println(cap(sl2))

	sl2 = append(sl2,1,2,3,4,5,6,7,8,9,10)
	fmt.Println(len(sl2))
	fmt.Println(cap(sl2))
}

func TestSliceMethod(t *testing.T) {
	sliceMethod()
}