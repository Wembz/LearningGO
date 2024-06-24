package main

import "fmt"

func main() {

	var username string = ("rodrigue")
	var age int
	var predifined string

	fmt.Println("Please enter username")
	fmt.Scan(&username)
	fmt.Println("Please enter age")
	fmt.Scan(&age)

	if age >= 18 && username == predifined {
		fmt.Println("authencation successful")
	} else if age < 18 {
		fmt.Println("Sorry you're uneligible you have to be 18+")
	} else {
		fmt.Println("incorrect username")
	}
}


