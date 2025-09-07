package main

import "fmt"

func Testfloat1() {
	float1()
}

func float1() {
	var f float64 = 3.14
	fmt.Println(f)
	fl  :=3.2
	fmt.Println(fl)
	fmt.Println(f + fl)
	fmt.Printf("%T, %T\n", f, fl)

	var fl32 float32 = 1.23
	fmt.Printf("%T\n", fl32)
	// 基本使わない

	// 型変換
	fmt.Printf("%T\n", float64(fl32))

	zero :=0.0
	pinf :=1.0 / zero
	fmt.Println(pinf) // 無限大:  Inf
	ninf := -1.0 / zero //- 無限大: - Inf
	fmt.Println(ninf)
	nan := zero / zero
	fmt.Println(nan)
}