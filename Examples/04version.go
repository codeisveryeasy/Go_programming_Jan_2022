package main

import (
	"fmt"
	"os"
	"reflect"
	"runtime"
	"time"
)

func main() {
	fmt.Println(runtime.Version())
	fmt.Println(runtime.GOROOT())

	//fmt.Println(os.Getenv("PATH"))
	fmt.Println(os.Getenv("JAVA_HOME"))
	fmt.Println(os.Getenv("OS"))
	fmt.Println(os.Getenv("NUMBER_OF_PROCESSORS"))
	//environmentVariable := "JAVA_HOME"
	environmentVariable := "TEMP"
	value, present := os.LookupEnv(environmentVariable)
	if present {
		fmt.Printf("%v:%v", environmentVariable, value)
	} else {
		fmt.Printf("%v is not defined!", environmentVariable)
	}
	fmt.Println()
	currentTime := time.Now()
	fmt.Println(currentTime.Hour())
	fmt.Println(currentTime.Minute())
	fmt.Println(currentTime.Second())
	fmt.Println(currentTime.Weekday())

	fmt.Println(reflect.TypeOf(currentTime).Kind())

}
