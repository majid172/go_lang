package main
import "fmt"
func main() {
	fmt.Println("Array Pointer Example in Go")

	// static sized array
	var n = [5]int{10, 20, 30, 40, 50}
	fmt.Println("Original array:", n)		

	// USE dynamic sized array is called slice in Go
	s := []int{100, 200, 300, 400, 500}
	fmt.Println("Original slice:", s)
	fmt.Println("Original slice length:", len(s))
	fmt.Println("Original slice capacity:", cap(s))

	// user input size of slice
	var size int
	fmt.Print("Enter size of slice: ")
	fmt.Scanln(&size)
	userSlice := make([]int,size)
	fmt.Println("Slice of size", size, "created.", userSlice)
	// user input elements of slice
	for i:=0; i<size; i++ {
		fmt.Printf("Enter element : ", i+1)
		fmt.Scanln(&userSlice[i])
	}
	fmt.Println("User defined slice:", userSlice)	
}