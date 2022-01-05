package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var myInput string
	fmt.Printf("Input integer:")
	fmt.Scanf("%s\n", &myInput)
	fmt.Printf("Integer: %v (Type: %T)\n", myInput, myInput)
	myInput = strings.TrimSpace(strings.Replace(myInput, "\n", "", -1))
	n1, err := strconv.ParseFloat(myInput, 64)
	if err != nil {
		fmt.Println(err)
	}

	var reader = bufio.NewReader(os.Stdin)
	var other string
	fmt.Printf("Input other integer:")
	other, _ = reader.ReadString('\n')
	other = strings.TrimSpace(strings.Replace(other, "\n", "", -1))
	fmt.Printf("Other: %v (Type: %T) \n", other, other)

	n2, err := strconv.ParseFloat(other, 64)
	if err != nil {
		fmt.Println(err)
	}

	sum := n1 + n2
	fmt.Printf("Sum: %v (Type: %T)", sum, sum)

	fmt.Print("Please enter first number: ")
	var n3 int
	fmt.Scanln(&n3) // take input from user
	fmt.Print("Please enter Second number: ")
	var n4 int
	fmt.Scanln(&n4) // take input from user
	result := n3 + n4
	fmt.Println("Sum of two numbers: ", result)

}
