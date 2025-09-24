package main

import "fmt"

func main() {
	// 数学的な定数
	fmt.Println(math.Pi)
	fmt.Println(math.Sqrt2)
	// 2の平方根
	fmt.Println(math.MaxFloat32)
	// float32で0以外の最小値
	fmt.Println(math.SmallesNonzeroFloat32)
	// float64で表現可能な最大値
	fmt.Println(math.MaxFloat64)
	// float64で0以外の最小値
	fmt.Println(math.SmallesNonzeroFloat64)

	// int ver
	fmt.Println(math.MaxInt64)
	fmt.Println(math.MinInt64)

	// 絶対値
	fmt.Println(math.Abs(100))
	fmt.Println(math.Abs(-100))

	// 累乗を求める
	fmt.Println(math.Pow(0, 3))
	// 累乗を求める
	fmt.Println(math.Pow(3, 2))

	// 平方根、立方根
	fmt.Println(math.Sqrt(9))
	fmt.Println(math.Cbrt(27))

	// 最大値、最小値
	fmt.Println(math.Max(10, 20))
	fmt.Println(math.Min(10, 20))
}