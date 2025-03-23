package main

import "fmt"

func main() {
	x := 20
	if x > 10 {
		fmt.Println("1:", x)
		x := 5
		fmt.Println("2:", x)
		if x < 10 {
			x := 100
			fmt.Println("3:", x)
		}
		fmt.Println("4:", x)
	}
	fmt.Println("5:", x)

	text := "Go言語"
	for i, r := range text {
		fmt.Println(i, r, string(r))
	}

	score := 85
	if score >= 90 {
		fmt.Println("優秀")
	} else if score >= 80 {
		fmt.Println("良好")
	} else if score >= 70 {
		fmt.Println("普通")
	} else {
		fmt.Println("要努力")
	}

	// 上記をswitch文で書き換え
	switch {
	case score >= 90:
		fmt.Println("優秀")
	case score >= 80:
		fmt.Println("良好")
	case score >= 70:
		fmt.Println("普通")
	default:
		fmt.Println("要努力")
	}

	values := []int{10, 20, 30}
	for i, v := range values {
		values[i] = v * 2
	}
	fmt.Println()
	fmt.Println(values)
}
