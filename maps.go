package main

import (
	"fmt"
	"reflect"
)

func maps() {
	statePopulation := map[string]int{
		"California": 39000,
		"Texas":      27000,
		"Florida":    20000,
	}

	statePopulation["Georgia"] = 10000
	delete(statePopulation, "Texas")
	invalidState, ok := statePopulation["NON"]

	manipulatedMap := statePopulation // References, wont copy
	delete(manipulatedMap, "Florida")

	fmt.Println(statePopulation, len(statePopulation), invalidState, ok)
}

type Doctor struct {
	number     int
	ActorName  string // Pascal case exports globally
	companions []string
}

func structs() { // Basically like JavaScript Objects :D
	aDoctor := Doctor{
		number:    3,
		ActorName: "Salam",
		companions: []string{
			"Chessy",
		},
	}
	fmt.Println(aDoctor)
	fmt.Println(aDoctor.ActorName)

	theDoctor := struct{ name string }{name: "Salam"}
	fmt.Println(theDoctor)
	fmt.Println(theDoctor.name)

	compositions()
}

type Animal struct {
	name   string `required max:"100"`
	origin string
}

type Bird struct {
	Animal
	speedKPH float64
	canFly   bool
}

func compositions() {
	b := Bird{}
	b.name = "Emu"
	b.origin = "Australia"
	b.speedKPH = 48.1
	b.canFly = false

	fmt.Println(b)
	fmt.Println(b.name)

	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("name")
	fmt.Println(field.Tag)
}
