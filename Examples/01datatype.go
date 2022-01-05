package main

import (
	"fmt"
	"reflect"
)

func main() {
	/*
		multi line
		comment

	*/
	//single line comment

	var name string
	name = "Air Asia Academy"

	var city string = "KL"
	pincode := 591211

	fmt.Println(name + "from" + city)
	fmt.Println(name, "from", city)
	//format specifier for variable value
	fmt.Printf("%v from %v  with pincode %v.\n", name, city, pincode)

	topic := "Go Programming"
	fmt.Printf("%v is very easy for programmers! \n", topic)
	fmt.Printf("Datatype of topic is %T. Value in topic is %v \n", topic, topic)
	fmt.Printf("Datatype of pincode is %T. Value in pincode is %v", pincode, pincode)

	var score float32 = 3.14
	fmt.Println()
	fmt.Printf("%T:  %v", score, score)

	pressure := 3.123456
	fmt.Println()
	fmt.Printf("%T:  %v", pressure, pressure)

	fmt.Println()
	fmt.Printf("With score of PI %-10v, I am in no pressure.", pressure)
	fmt.Printf("With score of PI %.2f \n", pressure)

	fmt.Println(reflect.TypeOf(pincode))

	fmt.Println(reflect.TypeOf(topic))

	// var (
	// 	friend = "Oki"
	// 	level = 3
	// )
	const PI = 3.14
	//PI = 3.1111

}
