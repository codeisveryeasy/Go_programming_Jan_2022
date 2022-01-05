package main

import (
	"fmt"
	"reflect"
)

func main() {
	//declare slice
	var dynamicArray = []int{1, 2, 3, 4}
	var array = [4]int{1, 2, 3, 4}
	array[2] = 8
	//array[5] = 8  //not possible
	fmt.Printf("%T", dynamicArray)
	fmt.Printf("%T", array)
	fmt.Println()
	fmt.Println(reflect.TypeOf(dynamicArray).Kind())
	fmt.Println(reflect.TypeOf(array).Kind())

	var myslice = make([]int, 10, 20)

	fmt.Println("Capacity:", cap(myslice))
	fmt.Println("Length:", len(myslice))
	fmt.Println(myslice)
	myslice = append(myslice, 13)
	myslice = append(myslice, 14, 15)
	fmt.Println("Capacity:", cap(myslice))
	fmt.Println("Lenght:", len(myslice))
	fmt.Println(myslice)

	var martixSlice = [][]int{
		{0, 1, 2, 3},   //row with index 0
		{4, 5, 6, 7},   //row with index 1
		{8, 9, 10, 11}, //row with index 2
	}
	fmt.Println(reflect.TypeOf(martixSlice).Kind())

	var martixArray = [3][4]int{
		{0, 1, 2, 3},   //row with index 0
		{4, 5, 6, 7},   //row with index 1
		{8, 9, 10, 11}, //row with index 2
	}

	fmt.Println(martixArray)

	var myMulSlice = make([][]int, 10, 20)
	//fmt.Println(myMulSlice)
	//myMulSlice[0] = []int{1, 2, 3, 4}
	//myMulSlice[1] = []int{4, 5, 6}
	fmt.Println(myMulSlice)

	var myslice2 = make([][]int, 10, 20)
	fmt.Println(myslice2)
	myslice2[9] = []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	fmt.Println(myslice2)
	myslice3 := append(myslice2, []int{1, 1, 1, 1, 1, 1, 1, 1, 1})
	fmt.Println(myslice2)
	fmt.Println("Capacity:", cap(myslice2))
	fmt.Println("Lenght:", len(myslice2))
	fmt.Println(myslice3)

	var myDuoslice = make([][]int, 2)
	myDuoslice[0] = append(myDuoslice[0], 2, 4, 6, 8)
	myDuoslice[1] = append(myDuoslice[1], 1, 3, 5, 7)
	fmt.Println(myDuoslice)

	var friends = [8]string{"Oma", "Loki", "Kia", "Tik", "Piz", "Mia", "Uie", "Oie"}
	//filter operation using index positions
	fmt.Println(friends)
	//friends from 0 to n-1 i.e. 0 to 3 for 0, 1, 2
	fmt.Println(friends[0:4])
	//friends from n to last index
	fmt.Println(friends[4:])
	//all values in array
	fmt.Println(friends)
	fmt.Println(friends[0:])
	fmt.Println(friends[0:len(friends)])
	//last n values of array friends[len(friends)-n:]
	fmt.Println(friends[len(friends)-3:])

}
