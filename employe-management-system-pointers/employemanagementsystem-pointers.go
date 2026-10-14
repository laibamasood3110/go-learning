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

func (e *employeedata) give_raise(amount int) {
	for i := 0; i < len(e.salary); i++ {
		e.salary[i] = e.salary[i] + amount
	}

}

func (e *employeedata) calculateavg() {
	totalsum := 0
	for i := 0; i < len(e.salary); i++ {
		totalsum += e.salary[i]
	}
	e.avgsalary = float64(totalsum) / float64(len(e.salary))

}

func (e *employeedata) giveRaise(amount float64) {
	e.avgsalary = e.avgsalary + amount
}
func (printdata employeedata) describe() string {
	return fmt.Sprintf("%s has %.02f in %s", printdata.name, printdata.avgsalary, printdata.Address.city)
}

func main() {
	employee := map[string]employeedata{
		"rayyan": {name: "kashif", salary: []int{80000, 80000, 80000, 80000, 80000, 80000, 80000, 80000, 80000, 80000, 80000, 80000}, Address: address{street: 40, city: "Islamabad"}},
		"muazam": {name: "ali", salary: []int{70000, 70000, 70000, 70000, 70000, 70000, 70000, 70000, 70000, 70000, 70000, 70000}, Address: address{street: 10, city: "Rawalpindi"}},
		"rabeet": {name: "irum", salary: []int{60000, 60000, 60000, 60000, 60000, 60000, 60000, 60000, 60000, 60000, 60000, 60000}, Address: address{street: 16, city: "Lahore"}},
	}
	for key, value := range employee {
		value.give_raise(5000)
		value.calculateavg()
		//value.giveRaise(5000)
		employee[key] = value
		fmt.Println(employee[key].describe())
		//fmt.Println(key, value.avgsalary)
	}
	fmt.Println(employee["rabeet"].salary)
}
