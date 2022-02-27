package main

import "fmt"

func main() {
	a := 5
	b := &a

	fmt.Println(a, b)     // it will print 5 + memory address of a
	fmt.Printf("%T\n", a) // a is int
	fmt.Printf("%T\n", b) //b is *int * = pointer

	// * to read val from address
	fmt.Println(*b) //it will POINT cuz it's pointer to the value
	fmt.Println(b) //print adress

	fmt.Println()
}
