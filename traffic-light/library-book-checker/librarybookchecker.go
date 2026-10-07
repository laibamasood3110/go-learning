package main

import "fmt"

func main() {
	booknames := [7]string{"english", "urdu", "maths", "computer", "physics", "chemistry", "islamiat"}
	availablebooks := booknames[2:6]

	fmt.Println("length of the slice is: ", len(availablebooks))
	fmt.Println("capacity of the slice is: ", cap(availablebooks))

	checkindex := 6

	switch {
	case checkindex == 0 || checkindex == 1 || checkindex == 6:
		fmt.Printf("the book no. %d is not available \n", checkindex)
		fallthrough
	case checkindex >= 2 && checkindex <= 5:
		fmt.Printf("the book no. %d is available", checkindex)
	default:
		fmt.Printf("you entered %d book number which is invalid", checkindex)
	}
}
