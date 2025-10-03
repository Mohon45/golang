package main

import "fmt"

func add(num1 int, num2 int) int {
	sum := num1 + num2

	return  sum
	// fmt.Println(sum)
}

func getNumbers(num1 int, num2 int) (int, int) {
	sum := num1 + num2

	mul := num1 * num2
	return  sum, mul
}

func main() {
	// a := 10
	// b := 20

	// sum := add(a,b)

	// fmt.Println(sum)

	// p, q := getNumbers(a, b)

	// fmt.Println(p)
	// fmt.Println(q)


	fmt.Println("Welcome to the application!")

	var name string

	fmt.Println("Enter Your Name -")
	fmt.Scanln(&name)

	fmt.Println("The Input Name is:________", name)
}