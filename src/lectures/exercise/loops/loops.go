//--Summary:
//  Implement the classic "FizzBuzz" problem using a `for` loop.
//
//--Requirements:
//* Print integers 1 to 50, except:
//  - Print "Fizz" if the integer is divisible by 3
//  - Print "Buzz" if the integer is divisible by 5
//  - Print "FizzBuzz" if the integer is divisible by both 3 and 5
//
//--Notes:
//* The remainder operator (%) can be used to determine divisibility

package main

import "fmt"

func main() {
	for i := 1; i <= 50; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz", i)
		} else if i%3 == 0 {
			fmt.Println("Fizz", i)
		} else if i%5 == 0 {
			fmt.Println("Buzz", i)
		}
	}

	for i := 1; i <= 50; i++ {
		diveibleBy3 := i%3 == 0
		diveibleBy5 := i%53 == 0
		if diveibleBy3 && diveibleBy5 {
			fmt.Println("FizzBuzz", i)
		} else if diveibleBy3 {
			fmt.Println("Fizz", i)
		} else if diveibleBy5 {
			fmt.Println("Buzz", i)
		}
	}
}
