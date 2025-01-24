package main

import "fmt"

func failedUpdate(g *int) {
	x := 10
	g = &x
}

func ex0606() {
	var f *int
	failedUpdate(f)
	fmt.Println(f)
}
