package main

import (
	"fmt"
)

func data() {
	var details = make(map[string][]string)
	details["Julia"] = []string{" cricket", "drawing"}
	details["Sophie"] = []string{" drawing"}
	details["mila"] = []string{" drawing"}
	details["Emma"] = []string{"tennis", "kabaddi"}
	details["neha"] = []string{"running"}
	details["Abhi"] = []string{" photography", "cricket"}
	details["Noor"] = []string{"cricket"}
	details["Elin"] = []string{" hockey"}
	details["Sara"] = []string{" cricket", "kabaddi"}
	details["Yara"] = []string{"tennis"}

	for key, Hobbies := range details {
		fmt.Println("Key:", key, "=>", Hobbies)
		//fmt.Println(hobbies(

	}

	//fmt.Println(details)

}
