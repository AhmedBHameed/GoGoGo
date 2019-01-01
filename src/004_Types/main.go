package main

import (
	"fmt"
)

/* var (
	x   = 22
	str = "Hey man wassup!!"
) */

// OR you can provide tyoe of variable when declaring
var (
	x   int    = 22
	str string = "Hey man wassup!!"
)

// Raw string literal
var raw string = `A dream
is to fly
over the rainbow
"so high"...
`

func main() {
	fmt.Printf("x is: %d with type of: %T\n", x, x) // To show TYPE of VALUE
	fmt.Printf("str is: %s with type of: %T\n", str, str)
	fmt.Println(raw)
}
