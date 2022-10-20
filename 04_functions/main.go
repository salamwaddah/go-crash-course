package main

import "fmt"

func main() {
	fmt.Println(greeting("Salam"))
	fmt.Println(getSum(24, 10))
}

func greeting(name string) string {
	return "Hello " + name
}

func getSum(n1, n2 int) int {
	return n1 + n2
}
