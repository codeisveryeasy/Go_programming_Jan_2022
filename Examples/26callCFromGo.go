package main

// #include<stdio.h>
// void callC(){
// 	printf("I am the function in C!");
// }
import "C"
import "fmt"

func main() {
	fmt.Println("I am in Go main!")
	C.callC()
}
