package main

import "fmt"

func main() {
	// MAIN TYPES
	// string
	// bool
	// int
	// int  int8  int16  int32  int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	// Using var
	var name = "Salam"
	var age = 27
	const isCool = true
	var size float32 = 420.69

	// Shorthand
	// name := "Salam"
	// email := "go@example.com"

	fmt.Println(name, age, isCool, size)
	fmt.Printf("%T\n", size)
}
