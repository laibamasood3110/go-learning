package main

import "fmt"

func main() {
	trafficLights := []string{"red", "green", "yellow"}
	fmt.Println("capacity of this slice is: ", cap(trafficLights))
	fmt.Println("length of this slice is: ", len(trafficLights))
	currentLight := "purple"
	switch currentLight {
	case "red":
		fmt.Println("stop now :')")
	case "yellow":
		fmt.Println("Be Ready ;)")
	case "green":
		fmt.Println("you can go now :)")
	default:
		fmt.Println("invalid input")
	}
}
