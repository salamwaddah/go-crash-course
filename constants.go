package main

import (
	"fmt"
)

const ci int16 = 27

const (
	_ = iota
	cat
	dog
	snake
)

func constants() {
	const (
		ci int     = 42
		cs string  = "foo"
		cf float32 = 3.14
		cb bool    = true
	)

	fmt.Printf("%v, %T\n", ci, ci)
	fmt.Printf("%v, %T\n", cs, cs)
	fmt.Printf("%v, %T\n", cf, cf)
	fmt.Printf("%v, %T\n", cb, cb)

	var animal int = cat
	fmt.Printf("%v", animal == cat)
}
