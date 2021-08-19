package main

import (
	"fmt"
)

func sum(num int, size int) int {
	fmt.Println("enter number:")
	fmt.Scanln(&num)

	fmt.Println("enter size:")
	fmt.Scanln(&size)

	sum := num + size
	return sum
}
func main() {
	sum()
}
