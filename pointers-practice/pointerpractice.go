package main

import "fmt"

type employeedata struct {
	name      string
	salary    []int
	avgsalary float64
	Address   address
}
type address struct {
	street int
	city   string
}

func giveRaisee(e employeedata, amount float64) {
	e.avgsalary = e.avgsalary + amount
}

func (e *employeedata) giveRaise(amount float64) {
	e.avgsalary = e.avgsalary + amount
}

func main() {
	var emp employeedata
	emp.name = "Haroon"
	emp.salary = []int{70000, 70000, 70000, 70000, 70000, 70000, 70000, 70000, 70000, 70000, 70000, 70000}
	totalsum := 0
	for i := 0; i < len(emp.salary); i++ {
		totalsum += emp.salary[i]
	}
	emp.avgsalary = float64(totalsum) / float64(len(emp.salary))
	fmt.Println(emp.avgsalary)

	giveRaisee(emp, 5000)
	fmt.Println(emp.name, emp.avgsalary)

	emp.giveRaise(10000)
	fmt.Println(emp.name, emp.avgsalary)
}
