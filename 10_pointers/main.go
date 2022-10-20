package main

import "fmt"

func main() {
	a := 5
	b := &a

	fmt.Println(a, b)     // 5, 0x
	fmt.Printf("%T\n", a) // int
	fmt.Printf("%T\n", b) // *int

	// reading memory address values
	fmt.Println(*b)  // 5
	fmt.Println(*&a) //5

	// change value with pointer
	*b = 10
	fmt.Println(a) // 10
}
