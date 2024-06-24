package main

import "fmt"

func main() {
 
	var character string
	var house string

	fmt.Println("Please Enter the name of a Character from Games of Thrones")
	fmt.Scan(&character)
	fmt.Println("Please Enter the name of the house allegiance")
	fmt.Scan(&house)
	switch house {
	case "House of stark":
		fmt.Println(&character, "is loyal to the North")

	case "house lannister":
		fmt.Println(&character, "has a cunning nature")

	case "House Targaryen":
		fmt.Println(&character, "has a connection to dragons")

	default:
		fmt.Println(&character, "knowledging their presence but not providing specific details")
	}

	if house == "house of stark" {
		fmt.Println(character, "is loyal to the North")
	} else if house == "house lannister" {
		fmt.Println(&character, "has a cunning nature")
	} else if house == "House Targaryen" {
		fmt.Println(&character, "has a connection to dragons")
	} else {
		fmt.Println(&character, "knowledging their presence but not providing specific details")
	}
}
