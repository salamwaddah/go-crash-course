package main

import "fmt"

func conditionals() {
	ifs()
	switches()
}

func ifs() {
	if true {
		fmt.Println("The test is true")
	}

	number := 50
	guess := -30
	if guess < 1 || guess > 100 {
		fmt.Println("Guess must be between 1 and 100")
	} else {
		if guess < number {
			fmt.Println("Too low")
		}
		if guess > number {
			fmt.Println("Too high")
		}
		if guess == number {
			fmt.Println("You got it")
		}
		fmt.Println(guess >= number, guess <= number, guess != number)
	}
}

func switches() {
	switch i := 1 + 3; i {
	case 1, 3, 5:
		fmt.Println("Odd")
	case 2, 4, 6:
		fmt.Println("Even")
	default:
		fmt.Println("Neiter")
	}

	var v interface{} = 1
	switch v.(type) {
	case int:
		fmt.Println("Integer")
		if v == 1 {
			break
		}
		fmt.Println("Number")
	case string:
		fmt.Println("String")
	default:
		fmt.Println("Another type")
	}
}
