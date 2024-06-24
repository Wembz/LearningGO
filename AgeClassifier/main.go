package main

import "fmt"

/* Create a Go program that prompts the user to enter their age (an integer).
 Print out a message categorizing the age group as "Child" (0-12 years old),
 "Teenager" (13-19 years old),
 "Adult" (20-64 years old), or "Senior" (65 years old and above).

TIP: This one requires the using of LOGICAL OPERATORS, the ones we've seen at the end of the class,

&& - and operator
|| - or operator */

func main() {

	var age int //(Must always activate var using For loop or another data type)
	fmt.Println("Please enter age:")
	fmt.Scan(&age)

	var AgeGroup string //(Must always activate var using For loop or another data type)

	if age >= 0 && age <= 12{
		AgeGroup = "child"
	} else if age >= 13 && age <= 19{
		AgeGroup = "Teenager"
	}else if age >= 20 && age<=64{
		AgeGroup = "Adult"
	} else if age >= 65 {
		AgeGroup = "Senior"
	}else{
		AgeGroup = "Invalid entry"
	}
 fmt.Printf("you belong to the group: %s\n", AgeGroup)
}
