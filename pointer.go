package main 
import "fmt"

func changeValue(p *int) {
	*p = 20
}

func main() {
	fmt.Println("Pointer Example in Go")
	var x int = 10
	var p  = &x
	changeValue(p)

	fmt.Println("========")
	fmt.Println("Address of x:", p,"Value of x using pointer p:", *p)
	fmt.Println("======pointer example======")

}