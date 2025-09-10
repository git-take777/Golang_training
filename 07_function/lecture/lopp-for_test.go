package main

import (
	"fmt"
	"testing"
)

func loopFor() {
	for i := 0; i < 10; i++ {
		i++
		if (i < 4) {
			break
		}
		fmt.Println(i)
	}
	for i := 0; i < 10; i++ {
		i++
		if (i == 4) {
			continue
		}
		fmt.Println(i)
	}
	arr:=[3]int{1,2,3}
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func TestLoopFor(t *testing.T) {
	loopFor()
}