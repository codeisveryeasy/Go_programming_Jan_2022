package main

import "fmt"

func main() {
	var scores = [4]int{1, 2, 3, 4}
	var scoresAuto = [...]int{1, 2, 3, 4}
	name := "Hello"
	printMessage(&name, scores, scoresAuto)
	var martix = [3][4]int{
		{0, 1, 2, 3},   //row with index 0
		{4, 5, 6, 7},   //row with index 1
		{8, 9, 10, 11}, //row with index 2
	}
	fmt.Println(martix)
}

func printMessage(message *string, scores [4]int, scoresAuto [4]int) {
	*message = "World"
	fmt.Println("Message: ", *message, scores, len(scoresAuto))
}
