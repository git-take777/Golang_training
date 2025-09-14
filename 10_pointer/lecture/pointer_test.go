package main

import (
	"fmt"
	"testing"
)

func Duble(x int) {
	x = x * 2
	fmt.Println(x)
}
func pointer1() {
	var x int = 100
	fmt.Println(x)
	Duble(x)
}

func pointer2() {
	// pointer型を計算
	var y int = 200

	var p *int =&y
	fmt.Println(p)
	fmt.Println(*p)

}

func Duble2(x int) {
	x = x * 4
	fmt.Println(x)
}

func TestPointer1(t *testing.T) {
	pointer1()
	pointer2()
}