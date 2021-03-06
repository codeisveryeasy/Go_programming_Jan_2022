package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	var counter int64
	counter = 0
	// atomic.AddInt64(&counter, 1)
	// atomic.AddInt64(&counter, 1)
	// atomic.AddInt64(&counter, 1)
	// atomic.AddInt64(&counter, 5)

	var wg sync.WaitGroup
	fmt.Println("Start the main thread")

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			//fmt.Println("Welcome! to AAA!")
			//fmt.Printf("Go Routine:%v, Counter: %v \n", i, counter)
			atomic.AddInt64(&counter, 1)
			//counter++
			//wg.Done()
		}()

	}
	wg.Wait()
	fmt.Println("Counter with Atomic: ", counter)
	fmt.Println("End the main thread")
	//atomic.StoreInt64(&counter, 88)
	//fmt.Println("Counter: ", counter)
	//counter++
	//fmt.Println("Counter: ", counter)

}
