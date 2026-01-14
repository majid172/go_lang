package main
import "fmt"
type Person struct {
	firstName string
	age       int
}

// receiver function to print details of Person
func (p Person) printDetails() {
	fmt.Println("First Name:", p.firstName)
	fmt.Println("Age:", p.age)
}

func print(p Person) {
	fmt.Println("First Name:", p.firstName)
	fmt.Println("Age:", p.age)
}

func main(){
	p1 := Person{firstName: "John", age: 25}
	p2 := Person{"Doe", 30}

	p1.printDetails()
	print(p2)

}