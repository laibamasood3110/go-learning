package main

import "fmt"

func main() {
	var name string = "Laiba"
	var maths int = 76
	english := 81
	computer := 69
	total_marks := maths + english + computer
	max_marks := 300
	var parecentage float64
	parecentage = (float64(total_marks) / float64(max_marks)) * 100
	fmt.Printf("Name of the Student is: %s\n", name)
	fmt.Printf("total marks= %d\n", total_marks)
	fmt.Printf("Percentage= %.2f\n", parecentage)

	if parecentage >= 90 {
		fmt.Println("Excellent! You got an A+ Grade")
	} else if parecentage >= 80 {
		fmt.Println("Congratulations!! You got an A grade ")
	} else if parecentage >= 70 {
		fmt.Println("GOOD!! You have got a B grade")
	} else if parecentage >= 50 {
		fmt.Println("you got a C grade")
	} else {
		fmt.Println("YOU Failed!!")
	}

}
