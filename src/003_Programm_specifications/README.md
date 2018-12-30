# 001- Program specifications

```go
package main

import (
	"fmt"
)

var scopeVariable = "You can access me anywhere inside you main package" // Declare and assign initialization value.
var scopeVarWithIdentifier int	// Declare scope variable with identifier that has ZERO VALUE.

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
```

## Check those terms out

> Identifiers
- [Identifiers](https://golang.org/ref/spec#Identifiers) name program entities such as variables and types. An identifier is a sequence of one or more letters and digits. The first character in an identifier must be a letter.

> Predeclared identifiers
- Are set of keywords that already defined by the language see [predeclared identifiers](https://golang.org/ref/spec#Predeclared_identifiers).

> Keywords
- [Keywords](https://golang.org/ref/spec#Keywords) are special words that are reserved by the language which the user can't use them with name of a variable.

> Short declaration operator
- `:=` is used to declare a variable with a value. The variable will auto detect type of the variable upon the value that are assigned to it. When we need to reassign a new value we use `=` only to set new value.

> Expressions
- `+`, `-`, etc are expression or called operators

> Operand
- `1` + `2` one and two are operand of `+` operator. It could be more than one operand.

> Statement
- `y := 1 + 2` is a statement code line

> Parens
- `()` called parens

> Curly praces
- `{}` and the code inside `{}` of the function `()` is called a `block` of cods.

> Package scope
- It is the domain of code inside the the file.go code which can be specified by a package name.

> Zero value
- [zero value](https://golang.org/ref/spec#The_zero_value) false for booleans, 0 for numeric types, "" for strings, and nil for pointers, functions, interfaces, slices, channels, and maps.

### So what is the difference between (Short declaration) and (Var declaration)?

Short declaration can be only used inside a block of codes where can't be used in package level.
Var declaration is used to declare a variable as a package variable scope.

When use variable as a scope variable keep it `“narrow”` as possible.