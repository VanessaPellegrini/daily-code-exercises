package main

import "fmt"

func main() {

	// #4 Create a new type, print it and reasign it.
	type dog int
	var x dog
	fmt.Print(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Print(x)

	//#5
	y := int(x)
	fmt.Print(y)
	fmt.Printf("%T\n", y)

}
