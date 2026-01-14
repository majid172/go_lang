package main
import "fmt"

type User struct {
	name  string
	email string
	age   int
}

func main() {
	fmt.Println("Enter Variable")
	var name string
	var email string
	var age int
	fmt.Scanln(&name, &email, &age)

	user1 := User{name: name, email: email, age: age}
	fmt.Println("User:", user1.age ,user1.name, user1.email)
}
	