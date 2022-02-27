package main

import "fmt"

//Define person struct

type Person struct {
	firstName string
	lastName  string
	age       int
	gender    string
	city      string
}

func main() {
	//init Person using struct
	person1 := Person{firstName: "Josh", lastName: "brickwall", age: 12, gender: "MALE"}

	fmt.Println(person1)
	fmt.Println(person1.firstName)
	person1.age = 5
	fmt.Println(person1.age)

}
