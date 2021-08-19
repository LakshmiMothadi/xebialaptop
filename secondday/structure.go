package main

import "fmt"

type College struct {
	Name     string
	Location string
	Year     int
	Marks    int
}

var students []College

func Information() {

	student := College{"Maheswari", "chennai", 2017, 75} //type1

	student2 := College{ //type2
		Name:     "Vamsi",
		Location: "kadapa",
		Year:     2020,
		Marks:    90,
	}
	var student3 College //type3
	student3.Name = "lakshmi"
	student3.Location = "Tirupati"
	student3.Year = 2020
	student3.Marks = 81

	fmt.Println(student, "\n", student2, "\n", student3)
	fmt.Println(student3.Name)

}

func Percentage() {
	for _, candidate := range students {
		if candidate.Marks >= 75 {
			fmt.Printf(" %s", candidate.Name)
		}
	}
}
func insertStudent(user College) {
	students = append(students, user)
}
func main() {
	Information()

	insertStudent(College{
		Name:     "Sonu",
		Location: "kadapa",
		Year:     2020,
		Marks:    91,
	})
	insertStudent(College{
		Name:     "aishu",
		Location: "anatapur",
		Year:     2020,
		Marks:    65,
	})

	Percentage()
	shop()
	data()

	var details2 InMemory
	details2.Remove("remove")

}
