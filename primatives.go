package main

import (
	"fmt"
)

func primatives() {
	bt := 1 == 1
	bf := 1 == 2
	fmt.Printf("%v, %T\n", bt, bt)
	fmt.Printf("%v, %T\n", bf, bf)

	var i int = 42
	fmt.Printf("%v, %T\n", i, i)

	var ui uint = 42
	fmt.Printf("%v, %T\n", ui, ui)

	a := 10             // 1010
	b := 3              // 0011
	fmt.Println(a & b)  // 0010
	fmt.Println(a | b)  // 1011
	fmt.Println(a ^ b)  // 1001
	fmt.Println(a &^ b) // 0100

	e := 8              // 2^3
	fmt.Println(e << 3) // 2^3 * 2^3 = 2^6
	fmt.Println(e >> 3) // 2^3 / 2^3 = 2^0

	var f float64 = 3.14
	f = 13.7e72
	f = 2.1E14
	fmt.Printf("%v, %T\n", f, f)

	fa := 10.2
	fb := 3.7
	fmt.Println(fa + fb)
	fmt.Println(fa - fb)
	fmt.Println(fa * fb)
	fmt.Println(fa / fb)

	var c complex64 = 1 + 2i
	fmt.Printf("%v, %T\n", c, c)
	fmt.Printf("%v, %T\n", real(c), real(c))
	fmt.Printf("%v, %T\n", imag(c), imag(c))

	s := "This is a string"
	fmt.Printf("%v, %T\n", s, s)

	r := 'r'
	fmt.Printf("%v, %T\n", r, r)
}
