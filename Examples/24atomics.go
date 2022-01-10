package main

import (
	"fmt"
	"sync"
)

func main() {

	counter := 0
	var wg sync.WaitGroup
	fmt.Println("Start the main thread")

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			//fmt.Println("Welcome! to AAA!")
			//fmt.Printf("Go Routine:%v, Counter: %v \n", i, counter)
			counter++
			wg.Done()
		}()

	}
	wg.Wait()
	fmt.Println("Counter: ", counter)
	fmt.Println("End the main thread")

}
