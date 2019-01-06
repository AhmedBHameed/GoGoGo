# 007- Strings

```go
package main

import "fmt"

var isValid bool

func main() {
	n := 20
	y := 123.232123

	fmt.Printf("Integer number %d of type %T", n, n)
	fmt.Printf("Float number %d of type %T", y, y)
	fmt.Println("Form validation ", isValid)
}
```
# 008- String type

```go
package main

import "fmt"

func main() {
	str := "String variable with short declaration"
	str2 := `A string can create string with back tick also`
	fmt.Println(str)
	fmt.Println(str2)

	strBytes := []byte(str2) // Ascii of sliced string in str2 variable
	fmt.Println(strBytes)
	fmt.Printf("\n Type of strBytes is: %T\n", strBytes)

	for i := 0; i < len(str2); i++ {
		fmt.Printf("Hexa code for character %c is %#U which is the same as 'rune' type\n", str2[i], str2[i])
	}
}
```

> A string type represents the set of string values. A string value is a (possibly empty) sequence of bytes. Strings are immutable: once created, it is impossible to change the contents of a string. The predeclared string type is string

[Checkout ASCII](https://en.wikipedia.org/wiki/ASCII)

Also checkout [Effective go](https://golang.org/doc/effective_go.html) and [Rune literals](https://golang.org/ref/spec#Rune_literals) for more character formating.