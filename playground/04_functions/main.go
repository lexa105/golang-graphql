package main

import "fmt"

func greetings(name string) string {
	return "Hello" + name
}

func sum(x, y, z int) int {
	return x + y + z
}

func main() {
	fmt.Println(greetings("World"))
	fmt.Println(sum(5, 5, 3))
}
