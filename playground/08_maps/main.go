package main

import "fmt"

func main() {
	// emails := make(map[string]string)

	// //Assign key values

	// emails["Bob"] = "bob@gmail.com"
	// emails["Anna"] = "ANNNA@gmail.com"
	// emails["Mike"] = "Mike@gmail.com"

	// Declare map and add

	emails := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharon@gmail.com"}
	fmt.Println(emails)
	fmt.Println(emails["Bob"])
	fmt.Println(len(emails))

	emails["Mike"] = "mike@gmail.com"

	//Delete one from map
	delete(emails, "Mike")
	fmt.Println(emails)
	fmt.Println(len(emails))
}
