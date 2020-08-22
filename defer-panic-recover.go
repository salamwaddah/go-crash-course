package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func deferring() {
	fmt.Println("Start")
	defer fmt.Println("Middle")
	fmt.Println("End")

	// result is "start"
	a := "start"
	defer fmt.Println(a)
	a = "end"

	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
}

func panicker() {
	fmt.Println("about to panic")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error: ", err)
		}
	}()
	panic("something went wrong")
	fmt.Println("done panicking")
}
