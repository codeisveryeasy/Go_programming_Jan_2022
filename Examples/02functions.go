package main

import "fmt"

func main() {
	moon := "Europa"
	fmt.Printf("Moon: %v \n", moon)
	welcome()
	printMessage("Function with friends!", 8)
	printMessage("Function with foes!", 2)

	total := calculator(9, 3)
	fmt.Println(total)
	n1, n2, n3 := multipleValues(13, 2)
	fmt.Printf("%v * %v = %v \n", n1, n2, n3)

	a1, a2, a3 := multiplePlanets("Mars")
	fmt.Printf("%v supports life is %v. It is %v miles from sun\n", a1, a2, a3)

	b1, b2, b3 := multiplePlanetsNamed("Earth")
	fmt.Printf("%v supports life is %v. It is %v miles from sun\n", b1, b2, b3)

}

func multiplePlanetsNamed(s1 string) (planetName string, supportLife bool, distanceFromSun int) {
	planetName = s1
	supportLife = true
	distanceFromSun = 965485415454
	return
}

func multiplePlanets(s1 string) (string, bool, int) {
	supportLife := true
	distanceFromSun := 965485415454
	return s1, supportLife, distanceFromSun
}

func multipleValues(i1, i2 int) (int, int, int) {
	multiple := i1 * i2
	return i1, i2, multiple
}

func calculator(n1 int, n2 int) int {
	sum := n1 + n2
	return sum
}

func welcome() {
	moon := "Ganymede"
	fmt.Println("Welcome to another Moon:", moon)
}

func printMessage(message string, count int) {
	fmt.Println("Message: ", message, count)
}
