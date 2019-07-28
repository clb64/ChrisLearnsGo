package main

import (
	"fmt"
)

var y = 43

func main() {

	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Printf("%b\n", y)  // binary
	fmt.Printf("%x\n", y)  // base 16 - lower case
	fmt.Printf("%X\n", y)  // base 16 - upper case
	fmt.Printf("%#x\n", y) // hex - with leading zeros
	fmt.Printf("%c\n", y)  // Unicode
	fmt.Printf("%d\n", y)  // base 10
	fmt.Printf("%o\n", y)  // base 8 - octal
	fmt.Printf("%U\n", y)  // Unicode?

}
