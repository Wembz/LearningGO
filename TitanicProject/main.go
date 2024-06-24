package main

import "fmt"

func main(){

	/* Titanic Passenger Survival Checker:

Write a Go program that prompts the user to enter the age and gender of a passenger who was on the Titanic. 
Based on historical data, determine whether the passenger was likely to survive or not. */


		var age int
		var gender string
	{
		fmt.Print("Enter your age: ")
	    fmt.Scanln(&age)

		fmt.Println("Please Enter Gender")
		fmt.Scan(&gender)

		if gender != "female" && age <= 18 {
			fmt.Println("Consider Priority")


		} else if gender != "male" && age <= 18{
			fmt.Println("status child passenger but not further details about survival")
		}
	}
	

}