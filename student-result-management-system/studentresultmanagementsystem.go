package main

import "fmt"

func main() {
	studentnames := []string{"Alaya", "Aina", "Zimal", "Aarish", "Rabeet"}
	subjects := []string{"maths", "computer", "english", "chemistry"}

	allMarks := [][]int{
		{47, 65, 73, 80},
		{74, 87, 95, 63},
		{20, 30, 24, 40},
		{92, 87, 64, 74},
		{74, 93, 85, 100},
	}
	classtotal := 0
	var classpercentage float64
	var percentage float64
	var newhighmark string
	max_marks := 400
	passedstudents := 0
	failedstudents := 0
	highestmax_marks := max_marks * len(studentnames)

	for i := 0; i < len(studentnames); i++ {
		highestmarks := 0
		total := 0

		for j := 0; j < len(subjects); j++ {
			total += allMarks[i][j]
			if allMarks[i][j] >= highestmarks {
				highestmarks = allMarks[i][j]
				newhighmark = subjects[j]
			}
		}
		classtotal += total
		percentage = (float64(total) / float64(max_marks)) * 100
		if percentage >= 90 {
			fmt.Println("Excellent! You got an A+ Grade")
		} else if percentage >= 80 {
			fmt.Println("Congratulations!! You got an A grade ")
		} else if percentage >= 70 {
			fmt.Println("GOOD!! You have got a B grade")
		} else if percentage >= 50 {
			fmt.Println("you got a C grade")
		} else {
			fmt.Println("YOU Failed!!")
		}

		if percentage >= 50 {
			passedstudents++
		} else {
			failedstudents++
		}
		fmt.Printf("Student      : %s\n", studentnames[i])
		fmt.Printf("Total marks  : %d / %d\n", total, max_marks)
		fmt.Printf("Percentage   : %.2f%%\n", percentage)
		//fmt.Printf("Grade        : %s\n", grade)
		fmt.Printf("Best subject : %s (%d marks)\n", newhighmark, highestmarks)
		fmt.Println("------------------------------")

	}

	classpercentage = (float64(classtotal) / float64(highestmax_marks)) * 100
	fmt.Println("========= CLASS SUMMARY =========")
	fmt.Printf("Class average : %.2f%%\n", classpercentage)
	fmt.Printf("Passed        : %d\n", passedstudents)
	fmt.Printf("Failed        : %d\n", failedstudents)
}
