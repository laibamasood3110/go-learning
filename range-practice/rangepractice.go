package main

import "fmt"

func main() {
	favbooks := []string{"maths", "computer", "chemistry", "novels"}

	for position, name := range favbooks {
		fmt.Println("the index of the book is: ", position, "And the name of the book is: ", name)
	}
}
