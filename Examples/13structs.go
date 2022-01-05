package main

import "fmt"

type student struct {
	name  string
	score int
	topic string
}

func (s *student) DisplayStudent(message string) {
	fmt.Printf("%v has scored %v in %v\n", s.name, s.score, s.topic)
	fmt.Println(message)
}

func (s *student) CreateStudent(name string, score int, topic string) {
	s.name = name
	s.score = score
	s.topic = topic
}

func main() {
	var s1 = student{}
	fmt.Println(s1)
	s1.name = "AAA"
	s1.score = 22
	s1.topic = "GO"
	fmt.Println(s1)

	var s2 = student{name: "OBB", score: 31, topic: "Network Programming"}
	fmt.Println(s2)

	var s3 = NewStudent("OMG", 44, "Spring Cloud")
	fmt.Println(*s3)
	//(*s3) not required
	//fmt.Printf("%v has scored %v in %v\n", (*s3).name, s3.score, s3.topic)
	s3.DisplayStudent("Hello receiver!")

	//s4.NewStudent("OMG", 44, "Spring Cloud")
	var s5 = student{}
	s5.CreateStudent("OK", 44, "Spring Cloud")
	fmt.Println(s5)

}

func NewStudent(name string, score int, topic string) *student {
	var s = student{
		name:  name,
		score: score,
		topic: topic,
	}
	return &s
}
