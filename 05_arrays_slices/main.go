package main

import "fmt"

func main() {
	// Arrays
	var fruitArr [2]string
	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"

	fmt.Println(fruitArr)
	fmt.Println(fruitArr[0])
	fmt.Println(fruitArr[1])

	// Declare and assign
	anotherFruitArr := [2]string{"Banana", "Grape"}

	fmt.Println(anotherFruitArr)

	// Slices
	fruitSlice := []string{"Watermelon", "Tomato", "Lemon"}

	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:3]) // start at index 1 but stop before 3

}
