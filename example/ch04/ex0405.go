package main

import (
	"fmt"
	"math/rand"
)

func ex0405() {
	// rand.Seed(time.Now().UnixNano()) // シードの指定
	n := rand.Intn(10) // 0から9の整数をランダムに生成
	if n == 0 {
		fmt.Println("n is zero", n)
	} else if n > 5 {
		fmt.Println("n is greater than 5", n)
	} else {
		fmt.Println("n is less than or equal to 5", n)
	}
}
