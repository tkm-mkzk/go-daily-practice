package main

import "fmt"

func div1(numerator int, denominator int) int { // 分子, 分母
	if denominator == 0 {
		return 0
	}
	return numerator / denominator
}

func ex0501() {
	result := div1(5, 2)
	fmt.Println(result)
}
