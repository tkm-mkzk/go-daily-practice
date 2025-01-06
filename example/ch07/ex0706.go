package main

import "fmt"

func ex0706() {
	type Score int
	type HighScore Score

	type Person struct {
		LastName  string
		FirstName string
		Age       int
	}

	type Employee Person

	var i int = 300
	var s Score = 100
	var hs HighScore = 200
	// hs = s
	// s = i
	s = Score(i)
	hs = HighScore(s)
	fmt.Println(s, hs)
	hhs := hs + 20
	fmt.Println(hhs)
}
