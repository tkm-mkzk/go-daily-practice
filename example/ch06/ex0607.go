package main

import "fmt"

func failedUpdate2(px *int) {
	x2 := 20
	px = &x2
}

func update(px *int) {
	*px = 20
}

func ex0607() {
	x := 10
	failedUpdate2(&x)
	fmt.Println(x) // 10
	update(&x)
	fmt.Println(x) // 20
}
