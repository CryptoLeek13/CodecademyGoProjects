package main

import "fmt"

func main() {
	// ask for user name and return input
	var name string
	fmt.Println("Please enter your name!")
	fmt.Scan(&name)
	fmt.Println("Hello ", name)
	// create age attribute
	var age int
	fmt.Println("Please enter your age!")
	fmt.Scan(&age)
	fmt.Printf("Your name is %s and you are %d years old", name, age)
	// store new user message within newMessage var
	newMessage := fmt.Sprintf("Hello %s and you are finally %d years old", name, age)
	fmt.Println(newMessage)
}
