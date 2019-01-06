# 007- Iota and bit shifting in Go

`Iota` is a predefined identifier.

```go

package main

import "fmt"

const (
	a int = iota
	b
	c
	d
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
```

We can use it with more logical reasons like the example bellow

```go
package main

import "fmt"

const (
	_  = iota
	kb = 1 << (10 * iota)
	mb = 1 << (kb * iota)
	gb = 1 << (mb * iota)
	tb = 1 << (gb * iota)
)

func main() {
	fmt.Println(kb)
	fmt.Println(mb)
	fmt.Println(gb)
	fmt.Println(tb)
}
```