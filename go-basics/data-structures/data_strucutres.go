package datastructures

import (
	"log"
	"sort"
)

var s *string

func TestMap() {
	var myMap map[string]string
	myMap = make(map[string]string)
	myMap["Dog"] = "Bark"
	myMap["Cat"] = "Meow"
	myMap["Cow"] = "Moo"
	log.Println(myMap)
	log.Println(myMap["Dog"])

	animal := "Lion"
	s = &animal
	log.Println("Address of Animal:", &animal)
	log.Println("Address of Animal from s:", s)
	log.Println("Address of s:", &s)
	log.Println("Value of Animal:", animal)
	log.Println("Value of Animal from s:", *s)
	*s = "Tiger"
	log.Println("After changing Value through s")
	log.Println("Value of Animal:", animal)
	log.Println("Value of Animal from s:", *s)
	changeMapValue(myMap)
	log.Println(myMap)
	TestSlices()
}

func changeMapValue(temp map[string]string) {
	temp["Dog"] = "Neigh"
}

func TestSlices() {
	var arr []int
	arr = append(arr, 3)
	arr = append(arr, 2)
	arr = append(arr, 1)
	log.Println(arr)
	sort.Ints(arr)
	log.Println(arr)
	log.Println(arr[0:1])

	arr2 := []string{"dog", "cat", "cow"}
	log.Println(arr2)
	sort.Strings(arr2)
	log.Println(arr2)

	defer switchCases(arr2)
	log.Println(arr2)
}

func switchCases(arr []string) {
	a, b := 1, 2
	c := a + b
	if c > 3 {
		log.Println("Greater than 3")
	} else if c < 3 {
		log.Println("Less than 3")
	} else {
		log.Println("It is equal to 3")
	}

	n := arr[1]
	switch n {
	case "dog":
		log.Println("Noting to Do")
		log.Println("Bye Then")
	case "cow":
		log.Println("Changing cow to ant")
		arr[1] = "ant"
		fallthrough
	case "cat":
		log.Println("Noting to Do")
		log.Println("Bye Then from cat")
	default:
		log.Println("Noting to Do")
		log.Println("Bye Then")
	}
	log.Println(arr)
}
