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

func (e employeedata) calculateavg() float64 {
	totalsum := 0
	for i := 0; i < len(e.salary); i++ {
		totalsum += e.salary[i]
	}
	e.avgsalary = float64(totalsum) / float64(len(e.salary))
	return e.avgsalary
}
func (printdata employeedata) describe() string {
	return fmt.Sprintf("%s has %.02f in %s\n", printdata.name, printdata.avgsalary, printdata.Address.city)
}

func main() {
	employee := map[string]employeedata{
		"rayyan": {name: "kashif", salary: []int{80000, 80000, 80000, 80000, 80000, 80000, 80000, 80000, 80000, 80000, 80000, 80000}, Address: address{street: 40, city: "Islamabad"}},
		"muazam": {name: "ali", salary: []int{70000, 70000, 70000, 70000, 70000, 70000, 70000, 70000, 70000, 70000, 70000, 70000}, Address: address{street: 10, city: "Rawalpindi"}},
		"rabeet": {name: "irum", salary: []int{60000, 60000, 60000, 60000, 60000, 60000, 60000, 60000, 60000, 60000, 60000, 60000}, Address: address{street: 16, city: "Lahore"}},
	}
	for key, value := range employee {
		value.avgsalary = value.calculateavg()
		employee[key] = value
		fmt.Println(employee[key].describe())
		//fmt.Println(key, value.avgsalary)
	}
}
