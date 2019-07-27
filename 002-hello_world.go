package main

import "fmt"

func main() {
	fmt.Println("hello, world\n")

	foo()

	fmt.Println("something more")

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
// sequence
// loop, iterative
// conditional
