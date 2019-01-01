# 005- Create your own type in Go

```go
package main

import "fmt"

var a int

type kaki int // Create our own type

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
```

> Go is typed language
- Every variable in Go is based on type

> Not possible
- to assign variable of custom type to another with another type like the example above. This will generate an error.
Error says
`src/005_Create_your_own_type/main.go:18:4: cannot use b (type kaki) as type int in assignment`

> Conversion is not casting in Go
- Other language use term `Casting` to convert from type to another but go uses `conversion`