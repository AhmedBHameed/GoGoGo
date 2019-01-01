# 005- Read documentation

```go
package main

import "fmt"

var num = 100

var myname string

func main(){
	fmt.Println("Welcome 2019! Hope that you cary good things for us")
	fmt.Printf("%T\n", num)  // Print me type of num VALUE
	fmt.Printf("%b\n", num)  // Print me binary of num VALUE
	fmt.Printf("%x\n", num)  //Print me hexadicimal of num VALUE
	fmt.Printf("%#x\n", num) // Print me hexa with prefix of "0x" for the num VALUE

	str := fmt.Sprintf("Value of %v is equal to binary of %b and dicimal of %d", num, num, num)
	fmt.Println(str)

	fmt.Println("\nType your name:")
	fmt.Scanln(&myname)
	fmt.Println("Your name is: ", myname)
}
```

Go to the `godoc.org/fmt` then you can see instructions of how to use `fmt` package with all it's accessable methods.


> \n
- Means new line like you pressed enter and start new line of typing

> \t
- to give spaces like when you press `Tab` on your keyboard

> \
- used to escape character

### Some important methods that you have to know in `fmt` package are:

> Print to the standard out (Or to the terminal in words)
	- fmt.Print(a ...interface{}) (n int, err error)
	- fmt.Printf(s string, a ...interface{}) (n int, err error)
	- fmt.Println(a ...interface{}) (n int, err error)

> Print to string which you can store the return value in a variable
	- fmt.Sprint(a ...interface{}) string
	- fmt.Sprintf(s string, a ...interface{}) string
	- fmt.Sprintln(a ...interface{}) string

> Print to a `File` or web server `Response`
	- fmt.FPrint(r io.Writer, a ...interface{}) (n int, err error)
	- fmt.Fprintf(r io.Writer, format string, a ...interface{}) (n int, err error)
	- fmt.Fprintln(r io.Writer, a ...interface{}) (n int, err error)

> Scan from terminal (Or take command by user entry in terminal)
	- func Scan(a ...interface{}) (n int, err error)
	- func Scanf(format string, a ...interface{}) (n int, err error)
	- func Scanln(a ...interface{}) (n int, err error)
	


