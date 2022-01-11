package main

import "fmt"

/*
#include "./28usingcgo.c"
*/
import "C"

func main() {

	fmt.Println("I am in Go main!")
	C.inExternalCFile()

}
