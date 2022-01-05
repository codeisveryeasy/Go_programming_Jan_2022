package main

import "fmt"

func main() {
	//map is made with key: value pairs
	var students = map[string]int{"okie": 21, "mori": 22, "yuko": 26, "sia": 20}
	fmt.Println(students)
	fmt.Println(students["sia"])
	fmt.Println(students["pia"])
	students["pia"] = 30
	fmt.Println(students)

	//declare using make
	studentsScore := make(map[string]int)
	studentsScore["sia"] = 33
	studentsScore["ohi"] = 35

	studentsScore["hia"] = 31
	studentsScore["kia"] = 36
	fmt.Println(studentsScore)

	//delete
	delete(studentsScore, "hia")
	fmt.Println(studentsScore)

}
