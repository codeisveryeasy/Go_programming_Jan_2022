package main

import (
	"fmt"

	"mymodule/celebrate"
)

//var message = "Hello!!!!Go"

func main() {
	fmt.Println(celebrate.GetMessage(), "", celebrate.GetYear())
	printMessage()
}

func printMessage() {
	fmt.Println(celebrate.GetMessage(), "", celebrate.GetYear())
}
