package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args)
	fmt.Println("Length of os.Args: ", len(os.Args))
	fmt.Println("Count of arguments: ", len(os.Args)-1)
	//os.Args[1] to .... length of Array
}
