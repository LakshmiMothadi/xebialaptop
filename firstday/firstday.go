package main

import (
	"fmt"
	"io/ioutil"
	"reflect"
)

func sum(num int, size int) int {
	fmt.Println("enter number:")
	fmt.Scanln(&num)

	fmt.Println("enter size:")
	fmt.Scanln(&size)
	sum := 0 // another way for variable declaration ( declaration & intialization)

	for i := num; i <= size; i++ {
		sum = sum + i
	}
	fmt.Println("sum of ", size, "numbers:")
	return sum
}

func logical() {
	array := []int{2, 4, 5, 6}

	for i := 0; i < len(array); i++ { //for
		if array[i]%2 == 0 {
			fmt.Println("even")
		} else {
			fmt.Println("odd")
		}

	}
	for index, val := range array { //range
		if val == 2 {
			fmt.Println("index:", index)
		}
	}

	number := 3 // switch
	switch number {
	case 1:
		fmt.Println("it is string ")
	case 2:
		fmt.Println("it is float")
	case 3:
		fmt.Println("it is int")

	}

	fmt.Println(reflect.TypeOf(number)) // trying to print variable type using "reflect" package
	data := []string{"d", "a", "t", "a"}
	fmt.Println(reflect.TypeOf(data))

}

func reading() {

	adding := []byte("information")
	err := ioutil.WriteFile("text.txt", adding, 0644)

	if err != nil {
		fmt.Println(err)
	}

	data, err := ioutil.ReadFile("text.txt")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print(string(data))

}
