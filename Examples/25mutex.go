package main

import (
	"fmt"
	"sync"
)

func main() {

	var counter int64
	counter = 0
	// atomic.AddInt64(&counter, 1)
	// atomic.AddInt64(&counter, 1)
	// atomic.AddInt64(&counter, 1)
	// atomic.AddInt64(&counter, 5)

	var wg sync.WaitGroup
	var mu sync.Mutex
	fmt.Println("Start the main thread")

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			//fmt.Println("Welcome! to AAA!")
			//fmt.Printf("Go Routine:%v, Counter: %v \n", i, counter)
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()

		}()

	}
	wg.Wait()
	fmt.Println("Counter with Mutex: ", counter)
	fmt.Println("End the main thread")
	//atomic.StoreInt64(&counter, 88)
	//fmt.Println("Counter: ", counter)
	//counter++
	//fmt.Println("Counter: ", counter)

}
