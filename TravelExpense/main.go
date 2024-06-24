package main

import "fmt"

/*
Design a Go program to help travelers calculate their total expenses for a trip.
Define an array to store the daily expenses incurred during the journey.
Prompt the user to enter the expenses for each day of the trip.
Use a for loop to iterate through the array and calculate the total expenses.
Additionally, determine the average daily expense.
Finally, provide a summary of the trip's expenses, including the total cost and average daily expenditure.

===== NOTES =====

When calculating numbers always start from "0" when introducing numbers so for example below

for i := 0; i < NumDays; i++ {
		fmt.Printf("Enter expenses for the day %d: ", i+1)
		fmt.Scanln(&expenses[i])

*/


/* Name [] 0  to access the string / Index notation is used to specify the elements of an array*/

func main() {
	
	// create a Var/Array for the "Number of Days"
	var NumDays int
	fmt.Println("Enter number of Days for trip: ")
	fmt.Scan(&NumDays)

	expenses := make([]float64, NumDays)

	// Promot user to enter expenses
	for i := 0; i < NumDays; i++ {
		fmt.Printf("Enter expenses for the day %d: ", i+1)
		fmt.Scanln(&expenses[i])

		// calculate expenses
		totalExpenses := 0.0
		for _, expense := range expenses {
			totalExpenses += expense / float64(NumDays)

			// Print Summary
			averageExpenses := totalExpenses / float64(NumDays)

			fmt.Printf("\n Trip summary: \n")
			fmt.Printf("Total Expenses: $%.2f\n", totalExpenses)
			fmt.Printf("Average Daily Expenses: $%2f\n", averageExpenses)
		}

	}
}
