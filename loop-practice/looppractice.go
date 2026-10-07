package main

import "fmt"

func main() {
	for i := 1; i <= 20; i++ {
		if i%2 == 0 {
			fmt.Println(i, "is an even number")
		} else {
			fmt.Println(i, "is an odd number")
		}
		if i == 15 {
			break
		}
	}
	for j := 1; j <= 10; j++ {
		if j == 5 {
			continue
		}
		fmt.Println(j)
	}

	fruits := []string{"apple", "banana", "orange", "pineapple", "strawbwrry"}

	for position, name := range fruits {
		fmt.Println("In index ", position, ", The fruit name is: ", name)
	}
}
