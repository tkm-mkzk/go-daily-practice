package main

import "fmt"

func ex0401() {
	x := 10
	if x > 5 {
		fmt.Println(x)
		x := 5
		fmt.Println(x)
	}
	fmt.Println(x)
}
