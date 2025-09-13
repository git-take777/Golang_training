package main

import (
	"fmt"
	"testing"
)
func map1() {
 var m = map[string]int{
	"apple":  100,
	"banana": 200,
	"orange": 300,
}
fmt.Println(m)
m2 := map[int]string{
	1: "apple",
	2: "banana",
	3: "orange",
}
fmt.Println(m2)
m4 := map[string]int{
	"apple":  100,
	"banana": 200,
	"orange": 300,
}
fmt.Println(m4)
m4["apple"] = 5000
fmt.Println(m4)
fmt.Println(m2[4])
i, v := m2[4]
fmt.Println(i)
fmt.Println(v) //false
}

func TestMap(t *testing.T) {
	map1()
}