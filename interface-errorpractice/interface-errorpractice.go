package main

import "fmt"

type negativebalancecheck struct {
	balance float64
}

func (n negativebalancecheck) Error() string {
	fmt.Println("my Error method ran")
	return fmt.Sprintf("you have entered %.02f, a negative value which is not acceptable", n.balance)
}

func check(b float64) error {
	if b < 0 {
		return negativebalancecheck{balance: b}
	}
	fmt.Println(b)
	return nil

}
func main() {
	checkbalance := negativebalancecheck{balance: -1000}
	_ = checkbalance
	fmt.Println(check(-500))
}
