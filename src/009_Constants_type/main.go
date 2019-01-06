package main

import "fmt"

// Declaration witout type which give us some flexability
// const (
// 	n   = 10
// 	pi  = 3.14
// 	str = "I dare you to change me!!!"
// )

// Declaration witout type which give us some flexability
const (
	n   int     = 10
	pi  float64 = 3.14
	str string  = "I dare you to change me!!!"
)

func main() {
	fmt.Println(n)
	fmt.Println(pi)
	fmt.Println(str)
	fmt.Printf("%T\n", n)
	fmt.Printf("%T\n", pi)
	fmt.Printf("%T\n", str)
}
