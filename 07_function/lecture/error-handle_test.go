package main

import (
	"fmt"
	"strconv"
	"testing"
)

func errorHandle() {
	var s string= "Hello"
	i,_ := strconv.Atoi(s)
	fmt.Println(i) //0なのはなぜ?
	fmt.Printf("%T\n", i)

}

func TestErrorHandle(t *testing.T) {
	errorHandle()
}