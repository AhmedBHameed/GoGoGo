package main

import "fmt"

const (
	_  = iota
	kb = 1 << (10 * iota)
	mb = 1 << (10 * iota)
	gb = 1 << (10 * iota)
	tb = 1 << (10 * iota)
)

func main() {
	fmt.Printf("%d of KB is equal to \t\t\t%b in binary\n", kb, kb)
	fmt.Printf("%d of MB is equal to \t\t%b in binary\n", mb, mb)
	fmt.Printf("%d of GB is equal to \t\t%b in binary\n", gb, gb)
	fmt.Printf("%d of TB is equal to \t%b in binary\n", tb, tb)
}
