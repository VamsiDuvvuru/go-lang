package main

import "fmt"

func main() {
	// Using Printf to format strings
	name := "Alice"
	age := 30
	phone := 1234567890
	fmt.Printf("Name: %s, Age: %d , Phone: %p ", name, age, phone)

	// Using Sprintf to create a formatted string
	formattedString := fmt.Sprintf("Hello, %s! You are %d years old.", name, age)
	fmt.Println(formattedString)

	// // Using Errorf to create an error message
	err := fmt.Errorf("An error occurred: %s is not a valid name", name)
	fmt.Println(err)
}
