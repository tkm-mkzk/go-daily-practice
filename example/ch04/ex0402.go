package main

import "fmt"

func ex0402() {
	x := 10
	if x > 5 {
		a, x := 20, 5
		fmt.Println(a, x)
	}
	fmt.Println(x)
}
