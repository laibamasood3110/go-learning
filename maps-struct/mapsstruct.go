package main

import "fmt"

type address struct {
	street int
	city   string
}

type employeedata struct {
	name      string
	salaries  []int
	avgsalary float64
	address   address
}

func printemployee(e employeedata) {
	fmt.Println(e.name, e.avgsalary)
}

func (e employeedata) calculateavg() float64 {
	totalsum := 0
	for i := 0; i < len(e.salaries); i++ {
		totalsum += e.salaries[i]
	}
	e.avgsalary = float64(totalsum) / float64(len(e.salaries))
	return e.avgsalary
}

func main() {
	employee := map[string]employeedata{
		"kashif": {name: "kashif", salaries: []int{80000, 80000, 80000, 80000, 80000, 80000, 80000, 80000, 80000, 80000, 80000, 80000}, address: address{street: 40, city: "Islamabad"}},
		"ali":    {name: "ali", salaries: []int{70000, 70000, 70000, 70000, 70000, 70000, 70000, 70000, 70000, 70000, 70000, 70000}},
		"irum":   {name: "irum", salaries: []int{60000, 60000, 60000, 60000, 60000, 60000, 60000, 60000, 60000, 60000, 60000, 60000}},
	}

	for key, value := range employee {
		value.avgsalary = value.calculateavg()
		fmt.Println(key, value.avgsalary)
		employee[key] = value
	}
	fmt.Println(employee["ali"].avgsalary)

	printemployee(employee["kashif"])
	printemployee(employee["irum"])
	printemployee(employee["ali"])

	fmt.Println(employee["kashif"].address.street, employee["kashif"].address.city)
	fmt.Println(employee["ali"].address.street, employee["ali"].address.city)
}
