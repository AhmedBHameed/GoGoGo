package main

import "fmt"

var num = 100

var myname string

func main() {
	fmt.Println("Welcome 2019! Hope that you cary good things for us")
	fmt.Printf("%T\n", num)    // Print me type of num VALUE
	fmt.Printf("%b\n", num)    // Print me binary of num VALUE
	fmt.Printf("%x\n", num)    //Print me hexadicimal of num VALUE
	fmt.Printf("\a%#x\n", num) // Print me hexa with prefix of "0x" for the num VALUE

	str := fmt.Sprintf("Value of %v is equal to binary of %b and dicimal of %d", num, num, num)
	fmt.Println(str)

	fmt.Println("\nType your name:")
	fmt.Scanln(&myname)
	fmt.Println("Your name is: ", myname)
}
