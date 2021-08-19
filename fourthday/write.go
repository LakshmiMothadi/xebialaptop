package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func writing() {
	adding := []byte(`{
	        "Name": "Mothadi",
	        "Role": "Trainee",
	        "Age": 23,
			"IsFemale": true
	    }`)

	err := ioutil.WriteFile("write.txt", adding, 0644)
	if err != nil {
		fmt.Println(err)
	}

	data, err := ioutil.ReadFile("write.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
	print, _ := json.Marshal(data)
	fmt.Println(string(print))

}

type code struct {
}

func anotherWrite() {
	var details2 code
	convert := []byte(`{
			"json":"GoLang"
		    "json":2009
	}`)

	err := ioutil.WriteFile("write2.txt", convert, 0644)

	if err != nil {
		fmt.Println(err)
	}

	convertData, err := ioutil.ReadFile("write2.txt")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print(string(convertData))
	errErr := json.Unmarshal(convertData, &details2)
	if errErr != nil {
		fmt.Println(err)
	}
	jsonFile, _ := os.Open("write2.txt")

	fmt.Printf("\n%v", &jsonFile)

}
