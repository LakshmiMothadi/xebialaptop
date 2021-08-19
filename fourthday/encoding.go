package main

import (
	"encoding/json"
	"fmt"
)

type company map[string]interface{}

func convert() {
	employee := company{
		"name":     "Mothadi",
		"role":     "Trainee",
		"age":      23,
		"IsFemale": true,
	}
	print, err := json.Marshal(employee)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("\n%v", string(print))

}
