package main

import "fmt"

/*func power(b, e int) int {
	if e == 0 {
		return 1
	}
	return b * power(b, e-1)
}

func main() {
	result := power(5, 3)
	fmt.Println(result)
}*/

func fibonacci(n int) int {
	if n == 1 {
		return 1
	}
	if n == 0 {
		return 0
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	result := fibonacci(7)
	fmt.Println(result)
}
