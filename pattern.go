package main
import "fmt"
func main() {
	fmt.Println("Pattern Matching in Go")
	var n int
	fmt.Scanln(&n)
	fmt.Println("Triangle Pattern:")
	trianglePattern(n)
	// fmt.Println("Right Triangle Pattern:")
	rightTrianglePattern(n)

}

// left triangle pattern
func trianglePattern(n int) {
	for i :=1; i<=n; i++ {
		for j:=1; j<=i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
	for i :=n-1; i>=1; i-- {
		for j:=1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

// right triangle pattern
func rightTrianglePattern(n int){
	for i :=1; i<=n; i++ {
		for j:=1; j<=n-i; j++ {
			fmt.Print("  ")
		}
		for j:=1; j<=i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
	for i :=n-1; i>=1; i-- {
		for j:=1; j<=n-i; j++ {
			fmt.Print("  ")
		}
		for j:=1; j<=i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}