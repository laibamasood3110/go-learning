package main

import "fmt"

type studentdata struct {
	name       string
	marks      []int
	percentage float64
}

func main() {
	student := []studentdata{
		{name: "Alaya", marks: []int{54, 87, 42, 98, 56}},
		{name: "Aina", marks: []int{76, 83, 72, 90, 40}},
		{name: "Aarish", marks: []int{83, 96, 55, 76, 90}},
	}
	totalmarks := 500
	for i := 0; i < len(student); i++ {
		allsubmarks := 0
		for j := 0; j < len(student[i].marks); j++ {
			allsubmarks += student[i].marks[j]

		}
		student[i].percentage = (float64(allsubmarks) / float64(totalmarks)) * 100
		fmt.Printf("%s scored %.02f%%\n", student[i].name, student[i].percentage)
	}
}

/*func main() {
	var student1 studentdata

	student1.name = "Alaya"
	student1.marks = []int{86, 76, 72, 83, 98}

	getmarks := 0
	for i := 0; i < len(student1.marks); i++ {
		getmarks += student1.marks[i]
	}
	fmt.Println(getmarks)
	student1.percentage = (float64(getmarks) / 500) * 100
	fmt.Printf("%s scored   : %.2f%%\n", student1.name, student1.percentage)

}*/
