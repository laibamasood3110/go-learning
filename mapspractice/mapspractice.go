package main

import "fmt"

func main() {

	salaries := map[string]float64{
		"kamran": 74000,
		"bilal":  89000,
		"umer":   65500,
	}
	for key, value := range salaries {
		fmt.Println(key, value)
	}
	//fmt.Println("average salary of kamran is:", salaries["kamran"])
}
