package main

import (
	"fmt"
	"strconv"
)

type Stringer interface {
	String() string
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return p.Name + "-----" + strconv.Itoa(p.Age)
}

func main() {
	p := Person{Name: "太郎", Age: 25}

	var stringer Stringer = p

	fmt.Println(stringer.String())
}
