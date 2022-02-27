package main

import (
	"fmt"
)

func main() {
	ids := []int{33, 53, 123, 52, 68, 5}

	//loop through ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d \n", i, id)
	}

	//Not using index
	for _, id := range ids {
		fmt.Printf("ID: %d \n", id)
	}

	//Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum", sum)

	//Range with map

	// emails := make(map[string]string)
	// emails["bob"] = "bob@gmail.com"
	// emails["Aneta"] = "anet@gmail.com"
	// emails["Ondrej"] = "hungryondrs@seznam.cz"
	emails := map[string]string{"Bob": "email@gmail.com"}

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}
}
