package main

import (
	"encoding/json"
	"fmt"
)

type Brand struct {
	Name     string
	Role     string
	Age      int
	IsFemale bool
}

func data(mobile Brand) {
	//var mobile Brand
	// data in JSON format
	Data := []byte(`{      
        "Name": "Mothadi",  
        "Role": "Trainee",
        "Age": 23,
		"IsFemale": true
    }`)

	err := json.Unmarshal(Data, &mobile) // decoding
	if err != nil {
		fmt.Println(err)
		fmt.Println("\nDecoded data:", mobile)
	}
}
