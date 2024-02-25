package datastructures

import (
	"encoding/json"
	"log"
)

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func MarshallAndUnMarshall() {
	myJson := `
	[
		{
			"first_name":"Akash",
			"last_name":"A C",
			"hair_color":"Black",
			"has_dog": true
		},
		{
			"first_name":"Bruce",
			"last_name":"Wayne",
			"hair_color":"Black",
			"has_dog":false
		}
	]
	`
	var person []User
	err := json.Unmarshal([]byte(myJson), &person)
	if err != nil {
		log.Fatalln("Error unmarshalling json:", err)
	}
	log.Println("UnMarshalled:", person)

	// write json from a struct

	var mySlice = []User{
		{
			FirstName: "Clark",
			LastName:  "Kent",
			HairColor: "Black",
			HasDog:    true,
		},
		{
			FirstName: "Bruce",
			LastName:  "Wayne",
			HairColor: "Black",
			HasDog:    false,
		},
	}
	m1, err := json.Marshal(mySlice)
	if err != nil {
		log.Fatalln("Error marshalling User", err)
	}
	log.Println(string(m1))
}
