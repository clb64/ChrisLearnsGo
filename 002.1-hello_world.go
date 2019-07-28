package main

import "fmt"

func main() {
	fmt.Println("hello, world\n")

	foo()

	n, err := fmt.Println("something more", 42, true)
	fmt.Println(n)
	fmt.Println(err)

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	bar()
}

func foo() {
	fmt.Println("I'm in Foo")
}

func bar() {
	fmt.Println("something bar")
}

// control flow;    See Control Flow Wiki
//     sequence
//     loop, iterative
//     conditional
