package main

import "fmt"

func main() {
	a := 5
	b := &a //Set a pointer to A !!adress

	fmt.Println(a, b)     // it will print 5 + memory address of a
	fmt.Printf("%T\n", a) // a is int
	fmt.Printf("%T\n", b) //b is *int * = pointer

	// * to read val from address
	fmt.Println(*b) //it will read the value of the pointer
	fmt.Println(b)  //print adress that points to.

	fmt.Println()
}
