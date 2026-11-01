package main

import "fmt"

type dog struct{}

func (d dog) speak() string {
	return "woof"
}

type cat struct{}

func (c cat) speak() string {
	return "meoww"
}

type cow struct{}

func (c cow) speak() string {
	return "moo"
}

type sound interface {
	speak() string
}

func printsound(s sound) {
	fmt.Println(s.speak())
}

func main() {
	c := cat{}
	d := dog{}
	animals := []sound{c, d, cow{}}

	for _, voices := range animals {
		printsound(voices)
	}
}
