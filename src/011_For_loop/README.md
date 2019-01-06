# 011 For loop

A [for](https://golang.org/ref/spec#For_statements) statement specifies repeated execution of a block. 
`for`, `break`, `continue` is a keyword

> ForStmt = "for" [ Condition | ForClause | RangeClause ] Block .

```go
package main

import "fmt"

func main() {
	
	// "for"  [ InitStmt ] ";" [ Condition ] ";" [ PostStmt ]
	for i := 0; i < 1000; i++ {
		fmt.Println(i)
	}

	fmt.Println("######## Nested loop #########")

	// "for"  [ InitStmt ] ";" [ Condition ] ";" [ PostStmt ]
	for x := 0; x < 5; x++ {
		fmt.Printf("X = %d\n", x)
		// "for"  [ InitStmt ] ";" [ Condition ] ";" [ PostStmt ]
		for y := 0; y < 2; y++ {
			fmt.Printf("\t Y = %d\n", y)
		}
	}

	
	fmt.Println("######## For loop with break #########")

	initState := 0

	for {
		if initState == 10 {
			break
		}
		fmt.Printf("Iteration number: %d\n", initState)
		initState++
	}

	fmt.Println("######## Like while loop #########")

	initial := 0
	for initial < 10 {
		fmt.Println("Condition is false")
		initial += 4
	}
	fmt.Println("Condition is true for will break")

}
```

> In Go there is no while loop. You can do while loop with `for` statement in the same way above.
- [Checkout](https://golang.org/ref/spec#For_statements) how you can deal with `For` in many cases

### Program to print even number only from 0 - 100 (100 is included)

```go
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
```