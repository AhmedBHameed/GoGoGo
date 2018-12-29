# 001- Hello World

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Go, Go, Go!!!")
}
```

## Some terms should we now about them


> Type interface{}
- It is like the any type in Typescript or no type in Go which can be anything you want. It is called empty interface.

> variadic parameters
- when you see type with a prefix of `...` then think about variadic parameters which take variables and convert them into array in one variable. For example `...interface{}` means give me any variabled of type any.

> Ignore variable or throw away it.
- use underscore `_` when you don't need to use a return variable. Because of Go nature it prevent us from using unused variable and the only way to ignore it is by using `_`.

> We use notation in Go
- `package.identifier` is the notation for using packages in Go. `fmt.Println()` is a notation. So package is `fmt` and we call identifier, constant, function from this package, so we called `Println` function from it.

> Packages
- are codes are written so that we can use them by importing them.
