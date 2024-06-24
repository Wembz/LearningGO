package main

import "fmt"

func main() {

	var dayoftheweek int
	fmt.Println("Please enter a number between (1 - 7) representing the day of the week")
	fmt.Scan(&dayoftheweek)

	if dayoftheweek < 1 || dayoftheweek > 7 {
		fmt.Println("entry invalid please try again")

	} else if dayoftheweek == 1 {
		fmt.Println("Monday")
	}else if dayoftheweek == 2 {
		fmt.Println("Tuesday")
	}else if dayoftheweek == 3 {
		fmt.Println("Wednesday")
	}else if dayoftheweek == 4 {
		fmt.Println("Thursday")
	}else if dayoftheweek == 5 {
		fmt.Println("Friday")
	}else if dayoftheweek == 6 {
		fmt.Println("Saturday")
	}else if dayoftheweek == 7 {
		fmt.Println("Sunday")
	}

}
