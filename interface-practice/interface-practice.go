package main

import "fmt"

type employeedata struct {
	name   string
	salary []int
}
type accountdetails struct {
	name       string
	balance    float64
	acc_number string
}

type productdetails struct {
	name  string
	price float64
}
type book struct {
	name string
}

func (p productdetails) describe() string {
	return fmt.Sprintf("%s has a price of Rs.%.02f", p.name, p.price)
}

func (a accountdetails) describe() string {
	return fmt.Sprintf("%s (%s) has balance of Rs.%.02f", a.acc_number, a.name, a.balance)
}
func (printdata employeedata) describe() string {
	return fmt.Sprintf("%s has %d", printdata.name, printdata.salary)
}
func (b book) describe() string {
	return fmt.Sprintf("the name of the book is %s", b.name)
}

type describer interface {
	describe() string
}

func printinfo(p describer) {
	fmt.Println(p.describe())
}

func main() {
	employee := []employeedata{
		{name: "kashif", salary: []int{80000, 80000, 80000, 80000}},
		{name: "ali", salary: []int{70000, 70000, 70000, 70000}},
		{name: "irum", salary: []int{60000, 60000, 60000, 60000}},
	}
	account := []accountdetails{
		{name: "Alaia", balance: 40000, acc_number: "A101"},
		{name: "Laiba", balance: 0.00, acc_number: "A102"},
		{name: "Aina", balance: 60000, acc_number: "A103"},
	}
	product := []productdetails{
		{name: "soap", price: 100},
		{name: "shampoo", price: 700},
	}
	bookdetails := []book{
		{name: "atomic habits"},
		{name: "the subtle art of not giving a f***"},
	}
	var details []describer
	for _, p := range product {
		details = append(details, p)
	}

	for _, e := range employee {
		details = append(details, e)
	}

	for _, a := range account {
		details = append(details, a)
	}
	for _, b := range bookdetails {
		details = append(details, b)
	}
	for _, value := range details {
		printinfo(value)

	}
}
