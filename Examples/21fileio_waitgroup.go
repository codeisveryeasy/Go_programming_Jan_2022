package main

import (
	"fmt"
	"os"
	"sync"
)

func main() {

	fmt.Println("Start the main thread")
	var wg sync.WaitGroup
	wg.Add(4)

	go writeToFile(80000, "file1.txt", &wg)
	go writeToFile(80000, "file2.txt", &wg)
	go writeToFile(80000, "file3.txt", &wg)
	go writeToFile(80000, "file4.txt", &wg)

	wg.Wait()
}

func writeToFile(countOfLines int, fileName string, wg *sync.WaitGroup) {
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
	wg.Done()

}
