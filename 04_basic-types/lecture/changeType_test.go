package main

import (
	"fmt"
	"strconv"
)

func changeType() {
	var i int = 100
	fl64 := float64(i)
	fmt.Printf("%T\n", fl64)

	// 文字列型(string)→数値型への変換
	var str1 string = "100"
	fmt.Printf("%T\n", str1)
	i, err :=strconv.Atoi(str1)
	if err !=nil {
		fmt.Println(err)
	}
	fmt.Println(i)
	fmt.Println("i =%T\n", i)
}