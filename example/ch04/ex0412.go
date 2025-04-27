package main

import "fmt"

func ex0412() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println(i, "FizzBuzz")
			continue
		}
		if i%3 == 0 {
			fmt.Println(i, "Fizz")
			continue
		}
		if i%5 == 0 {
			fmt.Println(i, "Buzz")
			continue
		}
		fmt.Println(i)
	}
}
