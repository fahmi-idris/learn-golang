package main

import "fmt"

func main() {
	/*
		Golang supports the declaration of many variables simultaneously,
		the way is to write the variables with the comma (,) delimiter.
		To fill in the value is also allowed at the same time.
	*/

	var first, second, third string = "1st", "2nd", "3rd"
	fmt.Printf(first, second, third)

	var fourth, fifth, sixth string
	fourth, fifth, sixth = "4th", "5th", "6th"
	fmt.Printf(fourth, fifth, sixth)

	seventh, eight, ninth := "7th", "8th", "9th"
	fmt.Printf(seventh, eight, ninth)

	one, isFriday, twoPointTwo, say := 1, true, 2.2, "Hello"
}
