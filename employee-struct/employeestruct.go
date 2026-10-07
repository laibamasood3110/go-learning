package main

import "fmt"

type employeedata struct {
	name      string
	salaries  []int
	avgsalary float64
}

func main() {
	employee := []employeedata{
		{name: "kashif", salaries: []int{80000, 80000, 80000, 80000, 80000, 80000, 80000, 80000, 80000, 80000, 80000, 80000}},
		{name: "ali", salaries: []int{70000, 70000, 70000, 70000, 70000, 70000, 70000, 70000, 70000, 70000, 70000, 70000}},
		{name: "irum", salaries: []int{60000, 60000, 60000, 60000, 60000, 60000, 60000, 60000, 60000, 60000, 60000, 60000}},
	}
	for i := 0; i < len(employee); i++ {
		totalsum := 0
		var totalnumofsalaries int
		for j := 0; j < len(employee[i].salaries); j++ {
			totalsum += employee[i].salaries[j]
		}
		totalnumofsalaries += len(employee[i].salaries)
		fmt.Println("total number of salaries are: ", totalnumofsalaries)
		employee[i].avgsalary = float64(totalsum) / float64(totalnumofsalaries)
		fmt.Printf("%s's average salary for a month is: %.02f\n", employee[i].name, employee[i].avgsalary)
	}
}
