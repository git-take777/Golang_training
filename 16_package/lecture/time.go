package main

import (
	"fmt"
	"time"
)
func main() {
	// 時間の生成
	t := time.Now()
	fmt.Println(t)

	// フォーマット
	// fmt.Println(t.Format("2006-01-02 15:04:05"))
	// fmt.Println(t.Format("2006/01/02 15:04:05"))
	// fmt.Println(t.Format("2006年01月02日 15時04分05秒"))

	// 指定した時間
	t2 := time.Date(2025, 6,10,20,30,40,50, time.Local)
	// fmt.Println(t2)
	// fmt.Println(t2.Year())
	// fmt.Println(t2.Month())
	// fmt.Println(t2.Weekday())
	// fmt.Println(t2.Day())
	// fmt.Println(t2.Hour())
	// fmt.Println(t2.Minute())
	// fmt.Println(t2.Second())
	// fmt.Println(t2.Nanosecond())
	// fmt.Println(t2.Zone())

	// fmt.Println(time.Hour)
	// fmt.Println("%T\n", time.Hour)
	// fmt.Println(time.Minute)
	// fmt.Println(time.Second)
	// fmt.Println(time.Millisecond)
	// fmt.Println(time.Microsecond)
	// fmt.Println(time.Nanosecond)




	// 時間の差
	diff := t2.Sub(t)
	fmt.Println(diff)

	// 秒、分、時間
	fmt.Println(diff.Seconds())
	fmt.Println(diff.Minutes())
	fmt.Println(diff.Hours())

	// 加算、減算
	// t3 := t.Add(diff)
	// fmt.Println(t3)

	// t4 := t.Add(-diff)
	// fmt.Println(t4)

	t5 := time.Date(2025,7,29,0,0,0,0,time.Local)
	t6 := time.Now()
	fmt.Println(t6)

	// t5 - t6
	diff2 := t5.Sub(t6)
	fmt.Println(diff2)
	// タイマー
	// timer := time.NewTimer(3 * time.Second)
	// <-timer.C
	// fmt.Println("Timer expired")

	//. 指定時間のスリープ
	time.Sleep(5 * time.Second)
	fmt.Println("5秒経過")
}