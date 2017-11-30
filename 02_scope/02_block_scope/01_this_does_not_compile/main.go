package main

import "fmt"

func main()  {
	x := 30
	fmt.Print(x)
	foo()
}

func foo() {
	fmt.Print(x)
}