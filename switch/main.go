package main

import "fmt"

func main() {
	a := 300

	switch a {
	case 1:
		fmt.Println("a is a 1")
	case 2, 3:
		fmt.Println("a is either 2 or 3")
	default:
		fmt.Println("a is neither 1 nor 2 or 3")
	}
}