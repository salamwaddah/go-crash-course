package main

import (
	"fmt"
)

func arrays() {
	grades := [...]int{99, 74, 96}
	fmt.Printf("Grades: %v\n", grades)
	fmt.Println(len(grades))

	var students [1]string
	students[0] = "Salam"
	fmt.Printf("Students: %v\n", students)

	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	b := a[:]
	c := a[3:]
	d := a[:6]
	e := a[3:6]
	a[2] = 99
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
}
