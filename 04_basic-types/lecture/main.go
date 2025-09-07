package main

import "fmt"

func main() {
	var i int = 100

	var i2 int64 = 200

	// fmt	.Println(i + i2) //これはデータ型が異なるのでエラーになる
	fmt.Printf("%T\n", i2)
	fmt	.Println(i + int(i2))
}