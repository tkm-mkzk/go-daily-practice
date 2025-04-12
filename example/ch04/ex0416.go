package main

import "fmt"

func ex0416() {
	m := map[string]int{
		"a": 1,
		"c": 3,
		"b": 2,
	}

	for i := 0; i < 3; i++ {
		fmt.Println("ループ", i)
		for k, v := range m {
			fmt.Println(k, v)
		}
	}
}
