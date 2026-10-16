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

func (e employeedata) calculateavg() (float64, error) {
	if len(e.salary) == 0 {
		return 0, fmt.Errorf("employee %s has no salary list", e.name)
	}
	totalsum := 0
	for i := 0; i < len(e.salary); i++ {
		totalsum += e.salary[i]
		if e.salary[i] < 0 {
			return 0, fmt.Errorf("employee %s salary contain negative entry in month %d which is invalid", e.name, i+1)
		}
	}

	if totalsum == 0 {
		return 0, fmt.Errorf("employee %s has zero salary", e.name)
	}
	e.avgsalary = float64(totalsum) / float64(len(e.salary))
	return e.avgsalary, nil
}
func (printdata employeedata) describe() string {
	return fmt.Sprintf("%s has %.02f in %s", printdata.name, printdata.avgsalary, printdata.Address.city)
}

func main() {
	employee := map[string]employeedata{
		"rayyan":  {name: "kashif", salary: []int{80000, 80000, 80000, 0, 80000, 80000, 80000, 80000, 80000, 80000, 80000, 80000}, Address: address{street: 40, city: "Islamabad"}},
		"muazam":  {name: "ali", salary: []int{70000, 70000, 70000, 70000, 70000, 70000, 70000, 70000, 70000, 70000, 70000, 70000}, Address: address{street: 10, city: "Rawalpindi"}},
		"rabeet":  {name: "irum", salary: []int{60000, 60000, 60000, 60000, 60000, 60000, 60000, 60000, 60000, 60000, 60000, 60000}, Address: address{street: 16, city: "Lahore"}},
		"laiba":   {name: "laiba", salary: []int{}},
		"zunaira": {name: "zunaira", salary: []int{0, 0, 0, 0, 0, 0}},
		"Alaia":   {name: "alaia", salary: []int{55000, -55000, 55000, 55000}},
	}
	for key, value := range employee {
		esalary, err := value.calculateavg()
		if err != nil {
			fmt.Println("error: ", err)
			continue
		}
		value.avgsalary = esalary
		employee[key] = value
		fmt.Println(employee[key].describe())
		//fmt.Println(key, value.avgsalary)
	}
}
