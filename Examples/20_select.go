package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Start the main thread")

	channel1 := make(chan string)
	channel2 := make(chan string)

	go printChannel1(channel1)
	go printChannel2(channel2)
	go printChannel3(channel1, channel2)

	fmt.Println("End the main thread")

	var inputSomething string
	fmt.Scanln(&inputSomething)

}

func printChannel1(channel1 chan string) {
	channel1 <- "one"
	time.Sleep(time.Second * 3)
}

func printChannel2(channel2 chan string) {
	channel2 <- "two"
	time.Sleep(time.Second * 2)
}

func printChannel3(channel1 chan string, channel2 chan string) {
	select {
	case received1 := <-channel1:
		fmt.Println(received1)
	case received2 := <-channel2:
		fmt.Println(received2)
	}
}
