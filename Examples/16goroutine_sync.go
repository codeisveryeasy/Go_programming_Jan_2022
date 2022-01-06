package main

import (
	"fmt"
)

//without wait group
func main() {
	fmt.Println("Start the main thread")
	go counter(1, 10)
	go counter(2, 10)
	go counter(3, 10)
	go counter(4, 100)
	fmt.Println("End the main thread")
	//time.Sleep(time.Second)
}

func counter(id int, count int) {
	for i := 1; i <= count; i++ {
		//time.Sleep(time.Second)
		fmt.Printf("%v: %v \n", id, i)
	}
}
