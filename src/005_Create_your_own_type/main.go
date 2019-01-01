package main

import "fmt"

var a int

type kaki int

var b kaki

func main() {
	a = 100
	b = 200

	fmt.Printf("a is of type %T\n", a)
	fmt.Printf("b is of type %T\n", b)

	// a = b // This will give us error try it yourself
	a = int(b) // Type conversion.
	fmt.Printf("a is of type %T with value of %v\n", a, a)
}
