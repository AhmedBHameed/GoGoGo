package main

import (
	"fmt"
)

var scopeVariable = "You can access me anywhere inside you main package"
var scopeVarWithIdentifier int = 100

func main() {
	// DECLARING and assigning VALUE (of a certin type)
	s := "This is string"              //Short declaration operator	// Statement
	fmt.Println(s)                     // Statement
	s = "Overwrite my old string here" // Statement
	fmt.Println(s)                     // Statement

	y := 1 + 2     // Expression in a statement that are using short declaration operator.
	fmt.Println(y) // Statement

	// Another way of declaration
	var val = "Hey wassup!!"
	fmt.Println(val)

	printScopeVariable() // Call to a function (More details later)
}

func printScopeVariable() {
	fmt.Println(scopeVariable)
	fmt.Println(scopeVarWithIdentifier)
}
