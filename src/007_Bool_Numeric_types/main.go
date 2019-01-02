package main

import "fmt"

var (
	n int
	y float64
)

func main() {
	// n := 20
	// y := 123.232123

	n = 20
	y = 123.232123

	fmt.Printf("Integer number %d of type %T", n, n)
	fmt.Printf("\nFloat number %f of type %T\n", y, y)
}
