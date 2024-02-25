package datastructures

import "fmt"

type Animal interface {
	WhatTheySound() string
	NoOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name      string
	Color     string
	NoOfTeeth int
}

func Mammals() {
	dog := Dog{
		Name:  "Alan",
		Breed: "German Sheperd",
	}
	gorilla := Gorilla{
		Name:      "Coco",
		Color:     "Black",
		NoOfTeeth: 32,
	}

	PrintInfo(&dog)
	PrintInfo(&gorilla)
}

func PrintInfo(a Animal) {
	fmt.Println("This Animal Says", a.WhatTheySound(), "and has ", a.NoOfLegs(), "legs")
}

func (d *Dog) WhatTheySound() string {
	return "Bark"
}

func (d *Dog) NoOfLegs() int {
	return 4
}

func (g *Gorilla) WhatTheySound() string {
	return "Growl"
}

func (g *Gorilla) NoOfLegs() int {
	return 4
}
