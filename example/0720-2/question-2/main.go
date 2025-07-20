package main

import "fmt"

type Speaker interface {
	Speak() string
}

type Dog struct {
	Name string
}

type Cat struct {
	Name string
}

type Robot struct {
	Model string
}

func (d Dog) Speak() string {
	return d.Name + "ワンワン！"
}

func (c Cat) Speak() string {
	return c.Name + "ニャーニャー！"
}

func (r Robot) Speak() string {
	return r.Model + "ビープ音です"
}

func main() {
	animals := []Speaker{
		Dog{Name: "ポチ"},
		Cat{Name: "たま"},
		Robot{Model: "バージョン1"},
	}

	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
}
