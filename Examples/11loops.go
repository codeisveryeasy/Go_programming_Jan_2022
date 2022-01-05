package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("------------")
	counter := 1
	for counter < 4 {
		fmt.Println(counter)
		counter++
	}
	fmt.Println("------------")
	for i := 0; ; i++ {
		if i == 8 {
			//break
			continue
		}
		if i == 12 {
			break
			//continue
		}
		fmt.Println(i)
	}
	fmt.Println("------------")
	//range
	friend := "Air Asia Academy"
	fmt.Println(len(friend))
	for range friend {
		fmt.Println("Hello....")
	}
	fmt.Println("------------")
	//range
	var scores = [4]int{81, 71, 26, 31}
	for i, v := range scores {
		fmt.Println(i, v)
	}
	//range with map
	keyvalue := map[string]string{
		"id":       "1",
		"name":     "AAA",
		"location": "KL",
		"country":  "Malaysia",
	}
	for key, value := range keyvalue {
		fmt.Printf("%v: %v \n", key, value)
	}

	for key, _ := range keyvalue {
		fmt.Printf("%v \n", key)
	}

	for _, value := range keyvalue {
		fmt.Printf("%v \n", value)
	}
}
