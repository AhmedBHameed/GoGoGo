package main

import "fmt"

func main() {
	str := "String variable with short declaration"
	str2 := `A string can create string with back tick also`
	fmt.Println(str)
	fmt.Println(str2)

	strBytes := []byte(str2) // Ascii of sliced string in str2 variable
	fmt.Println(strBytes)
	fmt.Printf("\n Type of strBytes is: %T\n", strBytes)

	for i := 0; i < len(str2); i++ {
		fmt.Printf("Hexa code for character %c is %#U\n", str2[i], str2[i])
	}
}
