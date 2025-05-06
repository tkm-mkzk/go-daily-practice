package main

import "fmt"

func ex0712() {
	var s *string
	fmt.Println(s == nil)
	var i interface{}
	fmt.Println(i == nil)
	i = s
	fmt.Println(i == nil)
}
