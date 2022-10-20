package main

import "fmt"

func main() {
	// define map
	emails := make(map[string]string)

	// assign kv
	emails["Salam"] = "salam@example.com"
	emails["Anthony"] = "anthony@example.com"
	emails["Mike"] = "mike@example.com"

	fmt.Println(emails)
	fmt.Println(len(emails))     // 3
	fmt.Println(emails["Salam"]) // "salam@example.com"

	// delete from map
	delete(emails, "Mike")
	fmt.Println(emails)
	fmt.Println(len(emails)) // 2

	// declare map and assign kv
	otherEmails := map[string]string{"Nick": "nick@example.com"}

	fmt.Println(otherEmails)
}
