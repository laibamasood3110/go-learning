package main

import "fmt"

type accountdetails struct {
	name       string
	balance    float64
	acc_number string
}

func (a accountdetails) balancecheck() error {
	if a.balance < 0 {
		return fmt.Errorf("the balance in your account %s is %.02f which is not valid", a.acc_number, a.balance)
	}
	/*if a.balance == 0 {
		return fmt.Errorf("your account:%s has balance of Rs.%.02f.Please recharge your account %s", a.acc_number, a.balance, a.name)
	}*/
	return nil
}

func (a *accountdetails) deposit(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("deposit of amount less that or equal to 0 is invalid for account:%s", a.acc_number)
	}
	a.balance = a.balance + amount
	fmt.Printf("deposit of Rs.%.02f is done to your account:%s\n", amount, a.acc_number)
	return nil
}
func (a *accountdetails) withdraw(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("withdrawl of amount less than or equal to 0 is invalid from account:%s", a.acc_number)
	}
	if amount > a.balance {
		return fmt.Errorf("the amount you entered is bigger than the balance of the account:%s which is %.02f", a.acc_number, a.balance)
	}
	a.balance = a.balance - amount
	fmt.Printf("withdraw of Rs.%.02f is withdrwed from your account:%s\n", amount, a.acc_number)
	return nil
}

func transfer(holder map[string]accountdetails, senderNo, recieverNo string, amount float64) error {
	sender, ok := holder[senderNo]
	if !ok {
		return fmt.Errorf("the account no:%s doesnot exist.", senderNo)
	}
	reciever, ok := holder[recieverNo]
	if !ok {
		return fmt.Errorf("reciever account %s doesnot exist", recieverNo)
	}
	err := sender.withdraw(amount)
	if err != nil {
		return err
	}
	err2 := reciever.deposit(amount)
	if err2 != nil {
		return err2
	}
	holder[senderNo] = sender
	holder[recieverNo] = reciever
	return nil
}

func (a accountdetails) describe() string {
	return fmt.Sprintf("%s (%s) has balance of Rs.%.02f", a.acc_number, a.name, a.balance)
}

func main() {
	holder := map[string]accountdetails{
		"A101": {name: "Alaia", balance: 40000, acc_number: "A101"},
		"A102": {name: "Laiba", balance: 0.00, acc_number: "A102"},
		"A103": {name: "Aina", balance: 60000, acc_number: "A103"},
		"A104": {name: "Aimen", balance: -4000, acc_number: "A104"},
		"A105": {name: "Hala", balance: 80000, acc_number: "A105"},
		"A106": {name: "Alia", balance: 20000, acc_number: "A106"},
	}
	for key, value := range holder {
		err1 := value.balancecheck()
		if err1 != nil {
			fmt.Println("Error:", err1)
			continue //for withdraw
		}
		fmt.Println(value.describe())
		err2 := value.deposit(5000)
		if err2 != nil {
			fmt.Println("Error:", err2)
			continue
		}
		holder[key] = value
		fmt.Println(value.describe())
		err3 := value.withdraw(1000)
		if err3 != nil {
			fmt.Println("Error:", err3)
			continue
		}
		holder[key] = value
		fmt.Println(value.describe())

	}
	err := transfer(holder, "A101", "A102", 5000)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(holder["A101"].describe())
	fmt.Println(holder["A102"].describe())

	err1 := transfer(holder, "A100", "A102", 5000)
	if err1 != nil {
		fmt.Println("Error:", err1)
	}
	fmt.Println(holder["A100"].describe())
	fmt.Println(holder["A102"].describe())

	err3 := transfer(holder, "A106", "A105", -5000)
	if err3 != nil {
		fmt.Println("Error:", err3)
	}
	fmt.Println(holder["A106"].describe())
	fmt.Println(holder["A105"].describe())

	err4 := transfer(holder, "A102", "A104", 10000)
	if err4 != nil {
		fmt.Println("Error:", err4)
	}
	fmt.Println(holder["A102"].describe())
	fmt.Println(holder["A104"].describe())

	for _, value := range holder {
		fmt.Println(value.describe())
	}
}
