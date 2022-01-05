package main

import "fmt"

func main() {
	name := "AAA"
	//normal function call - pass parameter by value
	printMessage(name)
	fmt.Println("In main:", name)
	//pass parameter by reference
	printMessageNow(&name)

}

func printMessageNow(message *string) {
	*message = "I am from Mars!"
	fmt.Println("In function:", message)
}

func printMessage(message string) {
	message = "I am alien!"
	fmt.Println("In function:", message)
}
