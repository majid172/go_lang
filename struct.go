package main

import "fmt"

type User struct {
	name  string
	email string
	age   int
}

func printUser(u User) {
	fmt.Println("User Details:")
	fmt.Println("Name:", u.name)
	fmt.Println("Email:", u.email)
	fmt.Println("Age:", u.age)
}

func main() {
	var name string
	var email string
	var age int

	// Ask user for input clearly
	fmt.Print("Enter name: ")
	fmt.Scanln(&name)

	fmt.Print("Enter email: ")
	fmt.Scanln(&email)

	fmt.Print("Enter age: ")
	fmt.Scanln(&age)

	// Create User struct
	user1 := User{
		name:  name,
		email: email,
		age:   age,
	}

	user2 := User{"Alice", "alice@example.com", 30}

	// Print details
	printUser(user1)
	printUser(user2)
}
