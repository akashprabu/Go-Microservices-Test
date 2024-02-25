package datastructures

import "log"

func Loops() {
	var i int
	for i = 0; i < 5; i++ {
		log.Println(i)
	}
	for j := 0; j <= 5; j++ {
		log.Println(j)
	}

	arr := []int{8, 9, 4, 5, 6, 2, 1, 0, 3, 7}
	for index, value := range arr {
		log.Println("Index:", index, "Value:", value)
	}
	for _, value := range arr {
		log.Println("Value:", value)
	}

	myMap := make(map[string]string)
	myMap["Dog"] = "Bark"
	myMap["Cat"] = "Meow"
	myMap["Cow"] = "Moo"
	for key, value := range myMap {
		log.Println("Key:", key, "Value:", value)
	}

	string1 := "Hello World"
	for index, value := range string1 {
		log.Println("Index:", index, "Value:", value)
	}
	for index, value := range string1 {
		log.Println("Index:", index, "Value:", string(value))
	}

}
