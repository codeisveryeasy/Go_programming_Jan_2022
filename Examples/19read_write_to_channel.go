package main

import "fmt"

func main() {
	fmt.Println("Start the main thread")
	var common = make(chan int, 2)
	//call go routine

	go printChannel2(common)
	go printChannel1(common)
	// go printChannel3(common)
	// go printChannel4(common)

	fmt.Println("End the main thread")

	var inputSomething string
	fmt.Scanln(&inputSomething)

}

func printChannel1(common chan int) {
	fmt.Println("Print Channel 1 Writing")
	common <- 200

}

//common <-chan int will make it read-only
func printChannel3(common chan int) {
	fmt.Println("Print Channel 3 Writing")
	common <- 100
	//common <- 200

}

func printChannel2(common chan int) {
	received := <-common
	fmt.Println("Print Channel 2 Reading", received)
}

func printChannel4(common chan int) {
	received := <-common
	fmt.Println("Print Channel 4 Reading", received)
}
