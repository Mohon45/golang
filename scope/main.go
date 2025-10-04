package main

import (
	"fmt"

	"example.com/mathlib"
)

func main() {
	a := 10
	b := 20

	// add(a,b)

	fmt.Println("Showing Custom Packages:")

	mathlib.Add(a, b)
}