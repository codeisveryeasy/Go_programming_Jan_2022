package main

import "fmt"

func main() {

	counter := 20

	if counter > 10 {
		fmt.Println("counter is greater than 10")
	} else {
		fmt.Println("counter is less than or equal to 10")
	}

	if counter > 10 && counter < 20 {
		fmt.Println("counter is between than 10 and 20")
	} else if counter > 20 {
		fmt.Println("counter is greater than20")
	} else {
		fmt.Println("I am free now! No more checks!!!!")
	}

	var check int

	if check = 71; check > 70 {
		fmt.Println("check is greater than 70")
	} else {
		fmt.Println("check is less than or equal to 70")
	}

}
