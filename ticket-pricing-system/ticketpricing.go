package main

import "fmt"

func main() {
	age := 78
	//fmt.Printf("the datatype of age is %T \n", age)
	isWeekend := false
	if isWeekend {
		fmt.Println("today is weekend day")
	} else {
		fmt.Println("Today is a weekday!!")
	}
	if age <= 12 {
		fmt.Println("price of the ticket for children is Rs.2000")
	}
	if age > 60 {
		fmt.Println("Price of tickets for senior citizens is Rs.2500")
	}
	if age > 12 && age <= 60 {
		if isWeekend {
			fmt.Println("Price of tickets for adults on weekend is Rs.3500")
		} else {
			fmt.Println("Price of tickets for adults on weekdays is Rs.3000")
		}

	}

}
