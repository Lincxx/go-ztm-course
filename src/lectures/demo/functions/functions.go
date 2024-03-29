package main

import (
	"fmt"
)

func double(x int) int {
	return x + x
}

func add(lhs, rhs int) int {
	return lhs + rhs
}

func greet() {
	fmt.Println("Hello from greet function")
}

func main() {
	greet()

	dozen := double(6)
	fmt.Println(dozen)

	bakerDozen := add(dozen, 1)
	fmt.Println(bakerDozen)

	anotherBakersDozen := add(double(6), 1)
	fmt.Println("another baker's dozen is", anotherBakersDozen)
}
