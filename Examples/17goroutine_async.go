package main

import (
	"fmt"
	"sync"
	"time"
)

//without wait group
func main() {
	fmt.Println("Start the main thread")
	var wg sync.WaitGroup
	wg.Add(5)

	go counter(1, 5, &wg)
	go counter(2, 5, &wg)
	go counter(3, 5, &wg)
	go counter(4, 5, &wg)
	go counter(5, 5, &wg)
	fmt.Println("End the main thread")
	wg.Wait()
	//time.Sleep(time.Second)
}

func counter(id int, count int, wg *sync.WaitGroup) {
	for i := 1; i <= count; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%v: %v \n", id, i)
	}
	wg.Done()
}
