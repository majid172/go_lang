package main
import "fmt"
var GlobalVar string = "I am a global variable"

func main() {
	// localVar := "I am a local variable"
	a := 10
	b := 20
	sumResult := sum(a, b)
	fmt.Println("Sum:", sumResult)
	arrayEx()
}

func sum(x int, y int) int {
	return x + y
}

func arrayEx() {
	numbers := [...]int{1, 2, 3, 4, 5}
	fmt.Println("Array elements:", numbers[:])
}