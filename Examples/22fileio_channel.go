package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println("Start the main thread")
	//var wg sync.WaitGroup
	//wg.Add(4)
	// channel1 := make(chan int)
	// channel2 := make(chan int)
	// channel3 := make(chan int)
	// channel4 := make(chan int)
	channel := make(chan int)

	// go writeToFile(80000, "file1.txt", channel1)
	// go writeToFile(80000, "file2.txt", channel2)
	// go writeToFile(80000, "file3.txt", channel3)
	// go writeToFile(80000, "file4.txt", channel4)
	go writeToFile(80000, "file1.txt", channel)
	go writeToFile(80000, "file2.txt", channel)
	go writeToFile(80000, "file3.txt", channel)
	go writeToFile(80000, "file4.txt", channel)
	go writeToFile(80000, "file5.txt", channel)
	go writeToFile(80000, "file6.txt", channel)
	go writeToFile(80000, "file7.txt", channel)
	go writeToFile(80000, "file8.txt", channel)

	//check1 := <-channel1
	// <-channel1
	// <-channel2
	// <-channel3
	// <-channel4
	<-channel
	<-channel
	<-channel
	<-channel
	<-channel
	<-channel
	<-channel
	<-channel

	//fmt.Println(check1)

}

func writeToFile(countOfLines int, fileName string, channel chan int) {
	fileHandle, error := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)

	if error != nil {
		fmt.Println("Error: ", error)
		return
	}

	defer fileHandle.Close()

	for i := 0; i <= countOfLines; i++ {
		content := fmt.Sprintf("Line number: %v \n", i)
		_, error := fileHandle.WriteString(content)
		if error != nil {
			fmt.Println("Not able to write to file", error)
		}
	}

	fmt.Printf("Just written %v lines to the file: %v \n", countOfLines, fileName)
	channel <- 1

}
