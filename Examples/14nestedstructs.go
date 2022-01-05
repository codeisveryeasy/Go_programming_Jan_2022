package main

import "fmt"

type student struct {
	name      string
	score     int
	topic     string
	institute school
}

type school struct {
	schoolname string
	city       string
}

func main() {
	school1 := school{schoolname: "Air Asia Academy", city: "KL"}
	student1 := student{
		name:      "Oni",
		score:     44,
		topic:     "Socket Programming",
		institute: school1}

	fmt.Println(student1)

	school2 := new(school)
	student2 := new(student)
	student2.name = "Uio"
	student2.topic = "Game Development"
	student2.score = 22
	school2.city = "KL"
	school2.schoolname = "MIT"
	student2.institute = *school2

	fmt.Println(*student2)
	fmt.Println(student2.name)

}
