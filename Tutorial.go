package main

import "fmt"

func main() {
	fmt.Println("Welcome to my quiz game!")

	fmt.Printf("Enter your name: ")

	var name string

	fmt.Scan(&name)

	fmt.Printf("Hello %v, Welcome to the game!\n", name)

	fmt.Println ("Enter your age: ")
	var age uint
	fmt.Scan(&age)

	if age >= 10 {
		fmt.Println("Yay you can play!")
	} else {
	fmt.Println("You cannot play!")
	return
	}
fmt.Printf("What is ")
}
