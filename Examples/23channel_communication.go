package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//n1 := generateRandomNumber()
	//n2 := generateRandomNumber()
	pass := make(chan int)

	go generateRandomNumber(pass)
	go generateRandomNumber(pass)
	go generateRandomNumber(pass)
	go generateRandomNumber(pass)
	go generateRandomNumber(pass)
	go generateRandomNumber(pass)

	total := 0
	iteration := 0
	//use range to keep track of values in channel
	for n := range pass {
		total += n
		iteration++
		if iteration == 6 {
			close(pass)
		}
	}

	fmt.Println("Total:", total)
	// n1 := <-pass
	// n2 := <-pass

	//fmt.Println(n1 + n2)
	//fmt.Println(<-pass + <-pass)
}

func generateRandomNumber(receivedPass chan int) {
	rand.Seed(time.Now().UTC().UnixNano())

	sleep := rand.Intn(10)
	fmt.Println("Sleeping for seconds: ", sleep)
	time.Sleep(time.Duration(sleep) * time.Second)

	random := rand.Intn(100)
	fmt.Println(random)
	//return random
	receivedPass <- random

}
