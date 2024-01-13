//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import "fmt"

//   - Write a function that accepts a person's name as a function
//     parameter and displays a greeting to that person.
func greet(name string) {
	fmt.Println("Hello", name, "it's nice to meet you.")
}

//   - Write a function that returns any message, and call it from within
//     fmt.Println()
func anyMsg() string {
	return "This is a simple message"
}

//   - Write a function to add 3 numbers together, supplied as arguments, and
//     return the answer
func addThree(num1, num2, num3 int) int {
	return num1 + num2 + num3
}

// * Write a function that returns any number
func anyNumber(num int) int {
	return num
}

// * Write a function that returns any two numbers
func twoNumbers(num1, num2 int) (int, int) {
	return num1, num2
}

// * Add three numbers together using any combination of the existing functions.
//   - Print the result
func addThreeCombo() {
	result := addThree(2, 3, 4) + anyNumber(6) + anyNumber(4)
	fmt.Println(result)
}

//* Call every function at least once

func main() {
	greet("Jeremy")
	fmt.Println(anyMsg())
	fmt.Println(addThree(1, 2, 3))
	fmt.Println(anyNumber(33))
	fmt.Println(twoNumbers(3, 3))
	addThreeCombo()

}
