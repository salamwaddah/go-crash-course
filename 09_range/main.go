package main

import "fmt"

func main() {
	ids := []int{24, 10, 1994, 13, 05, 1996}

	// loop through ids
	for i, id := range ids {
		fmt.Printf("%d - ID %d\n", i, id)
	}

	// not index loop through ids
	for _, id := range ids {
		fmt.Printf("ID %d\n", id)
	}

	// add all ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum", sum)

	// range with maps
	emails := map[string]string{"Nick": "nick@example.com", "Mike": "mike@example.com"}

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}

	for k := range emails {
		fmt.Printf("Name: %s\n", k)
	}

}
