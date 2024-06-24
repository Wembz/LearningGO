package main

import "fmt"

/*

Login prpoject

ask the user for the username and password until
the user gives the correct one

TIP
 FOR " to create a loop mechanism "

var. Username string -> rodrigue123
var. password string -> badbunny

when asking for details like user name or password
you must scan example  "fmt.scan(&username)""

you must ask the user interact with them

"Please enter password"
"Please enter username"
!= different


if(if they  are correct ) username == "rodrigue123" && password == "badbunny" {
		fmt.Println("Login Succesful")
	} else(if not then stop simulation) {

*/

func main() {

	var username string
	var password string

	for username != "rodrigue123" || password != "badbunny" {
		fmt.Println("Please enter Username: ")
		fmt.Scan(&username)

		fmt.Println("Please enter Password: ")
		fmt.Scan(&password)
	}

	if username != "rodrigue123" || password != "badbunny" {
		fmt.Println("Try again!!")
	}

	fmt.Println("Login Successful!!!")

}
