# 007- Boolean & Numeric types

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

Check out go `primitive` [types](https://golang.org/ref/spec#Numeric_types)

> Boolean types
- Which takes either `true` or `false`

> Integer type where the number that has no decimal digits or decimal point. Float type has decimal point or digits.

> There is no need to specify exact width of type. But only when you want to do it.

> Better to use `int` as integer or `float64` for float. We knew that from the language itself when you use shot declaration for integer and float numbers.

> Numeric types
- u in uint[8, 16, 32, 64] is unsiged number which start from 0 to 2 of n, which n is one of [8, 16, 32, 64] as specified in you program. [Numeric types](https://golang.org/ref/spec#Numeric_types)

> Types without size
- You can set types as absolute types without word-size of container. When compile to a specific system the compiler will set it to the specific type of the system.
System with 64 bit then `int` will take `int64` and so on.

> Alias types
- All numeric types are distinct except:
	- `byte` which is an alias for `uint8`
	- `rune` which is an alias for `int32` 

