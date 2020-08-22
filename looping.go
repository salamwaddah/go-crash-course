package main

import "fmt"

func looping() {
	for i, j := 0, 0; i < 5; i, j = i+1, j+1 {
		fmt.Println(i, j)
	}

	x := 0
	for x < 5 {
		fmt.Println(x)
		x++
	}

	n := 0
	for {
		fmt.Println(n)
		if n == 10 {
			break
		}
		n++
	}

Loop:
	for i := 0; i <= 3; i++ {
		for j := 0; j <= 3; j++ {
			fmt.Println(i * j)
			if i*j >= 3 {
				break Loop
			}
		}
	}

	s := []int{1, 2, 3}
	fmt.Println(s)
	for k, v := range s {
		fmt.Println(k, v)
	}
}
