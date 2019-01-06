package main

import "fmt"

func main() {

	n := -1

	for {
		if n >= 100 {
			break
		}

		n++
		if n%2 != 0 {
			continue
		}
		fmt.Println(n)

	}
}
