package main

import (
	"fmt"
	"strconv"
)

func main() {
	var number int = 42
	fmt.Printf("%v, %T\n", number, number)

	var word string
	word = strconv.Itoa(number)
	fmt.Printf("%v, %T", word, word)
}
