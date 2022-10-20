package main

import "fmt"

func main() {
	x := 10
	y := 12

	// if else
	if x <= y {
		fmt.Printf("%d is less or equal to %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}

	// else if
	color := "red"

	if color == "red" {
		fmt.Println("color is red")
	} else if color == "green" {
		fmt.Println("color is green")
	} else {
		fmt.Println("another color")
	}

	switch color {
	case "red":
		fmt.Println("color is red")
	case "green":
		fmt.Println("color is green")
	default:
		fmt.Println("another color")
	}
}
