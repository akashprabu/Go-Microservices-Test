package main

import (
	"fmt"
	datastructures "gobasics/data-structures"
	"log"
	"time"
)

var temp3 string

var (
	temp  string
	temp2 string
)

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Date        time.Time
}

func (m *User) printName() string {
	return m.FirstName
}

func main() {
	fmt.Println("Hello world")

	var whatToSay string
	whatToSay = "Goodbye, cruel world"
	fmt.Println(whatToSay)
	var i int
	i = 7
	fmt.Println("i is set to", i)
	whatWasSaid := saySomething()
	fmt.Println("The function returned", whatWasSaid)
	whatWasSaid2, otherThingThatWasSaid := saySomething2()
	fmt.Println("The function returned", whatWasSaid2, "and", otherThingThatWasSaid)
	fmt.Printf("The function returned %s and %s\n", whatWasSaid2, otherThingThatWasSaid)

	color := "Green"
	log.Println("Color is ", color)
	changeColour(&color)
	log.Println("Color is ", color)

	log.Println("User Struct", whatEver())
	myVar := whatEver()
	log.Println("User's FirstName: ", myVar.printName())

	// datastructures.TestMap()
	// datastructures.Loops()
	datastructures.Mammals()
	intChan := make(chan int)
	defer close(intChan)
	go datastructures.Channel(intChan)
	answer := <-intChan
	log.Println("The RandomNumber is :", answer)
	datastructures.MarshallAndUnMarshall()
}

func saySomething() string {
	return "something"
}

func saySomething2() (string, string) {
	return "something2", "else"
}

func changeColour(s *string) {
	newValue := "Red"
	*s = newValue
}

func whatEver() User {
	var person1 User
	person1 = User{
		FirstName:   "AKASH",
		LastName:    "A C",
		PhoneNumber: "99999999",
		Date:        time.Now(),
	}
	return person1
}
