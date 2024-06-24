package main

/* Create a Go program that prompts the user to enter a month (integer from 1 to 12).
Using a switch statement, identify the season associated with that month (spring, summer, autumn, or winter).
 Additionally, use logical operators to handle cases where input is outside the valid range.*/

import "fmt"

func main() {

	var month int

	fmt.Println("Enter a month in numbers between 1 - 12")// user input
	fmt.Scan(&month) // user output

	switch month {

	case 12, 1, 2:
		fmt.Println("This is the winter season")
	case 3, 4, 5:
		fmt.Println("This is Spring Season")
	case 6, 7, 8:
		fmt.Println("This is the Summer season")
	case 9, 10, 11:
		fmt.Println("This is the Autumn Season")

	default:
		fmt.Println("Entry Invalid")

	}

}
