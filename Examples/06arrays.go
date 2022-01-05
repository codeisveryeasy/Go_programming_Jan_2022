package main

import "fmt"

func main() {
	//declare array

	var scores = [4]int{81, 71, 26, 31}
	var scoresAuto = [...]int{81, 71, 26, 31, 27, 53, 44, 83}
	var emptyArray = [8]int{}

	fmt.Println(scores)
	fmt.Println(len(scores))
	fmt.Println(len(scoresAuto))
	fmt.Println(emptyArray)

	var specific = [9]int{0: 8, 8: 17}
	fmt.Println(specific)

	//[rows][cols]
	var martix = [3][4]int{
		{0, 1, 2, 3},   //row with index 0
		{4, 5, 6, 7},   //row with index 1
		{8, 9, 10, 11}, //row with index 2
	}

	fmt.Println(martix)

	passScores(scores, &scoresAuto)
}

func passScores(scores [4]int, scoresA *[8]int) {
	fmt.Println("In function")
	fmt.Println(scores)
	fmt.Println(*scoresA)
}
