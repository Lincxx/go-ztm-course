package main

import "fmt"

func main() {
	var name = "Jeremy"
	fmt.Println("my name is", name)

	//type annotation
	var herName string = "Kathy"
	fmt.Println("my name is", herName)

	//create and assign
	userName := "Admin"
	fmt.Println("user name is", userName)

	//uninitilzed
	var sum int
	fmt.Println("the is", sum)

	//compound
	part1, other := 1, 5
	fmt.Println("part 1 is", part1, "other is", other)

	part2, other := 2, 0
	fmt.Println("part 1 is", part2, "other is", other)

	//reassign
	sum = part1 + part2
	fmt.Println("sum =", sum)

	var (
		lessonName = "Variables"
		lessonType = "Demo"
	)

	fmt.Println("lessonName=", lessonName)
	fmt.Println("lessonType=", lessonType)

	word1, word2, _ := "hello", "world", "!"
	fmt.Println(word1, word2)
}
