package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to Group D's Week 4 Project!")
	asc()
	timenew()
	person := Person{firstName: "Savio", lastName: "Lovett", age: 27}
	fmt.Println("Full Name:", person.fullName())
	fmt.Println("Age:", person.age)
	var s int = factorial(person.age)
	fmt.Printf("Factorial of age of Savio- %d is - %d -- done by Supatra\n", person.age, s)
}
