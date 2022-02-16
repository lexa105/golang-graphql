package main

import "fmt"

func main() {
	//Arrays

	var fruitArray1 []int

	//Assign values

	fruitArray1[0] = 2
	fruitArray1[1] = 1

	// Declare and assign #"I assume using :="
	// Is the same as the above just shortened :)
	fruitArray := [2]string{"Apple", "Banana"}

	fmt.Println(fruitArray)
	fmt.Println(fruitArray1)

	//Slices

	fruitSlices := []string{"Apple", "Orranges", "Grape"}

	fmt.Println(fruitSlices)

	fmt.Println(fruitSlices[1:2])

}
