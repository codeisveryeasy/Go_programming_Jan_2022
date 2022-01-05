package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	filedata, error := ioutil.ReadFile("filesample.txt")
	if error != nil {
		fmt.Println(error)
		return
	}
	fmt.Println("File read success")
	fmt.Println(filedata)
	fmt.Println(string(filedata))

}
