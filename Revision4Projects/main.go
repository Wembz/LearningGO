package main

import (
	"fmt"
)

func main() {
	// Initialize almanac
	almanac := make(map[string][]string)

	// Menu loop
	for {
		// Display menu options
		fmt.Println("\nMenu:")
		fmt.Println("1. Add a flower to the almanac")
		fmt.Println("2. Print the almanac")
		fmt.Println("3. Exit")

		// Prompt user for choice
		fmt.Print("Enter your option: ")
		var option int
		fmt.Scanln(&option)

		// Process user choice
		if option == 1 {
			addFlower(almanac)
		} else if option == 2 {
			printAlmanac(almanac)
		} else if option == 3 {
			fmt.Println("Exiting...")
			return
		} else {
			fmt.Println("Invalid choice. Please enter a valid option.")
		}
	}
}

// Function to add a flower to the almanac
func addFlower(almanac map[string][]string) {
	var scientificName string

	fmt.Println("Please enter the scientific name of the flower:")
	fmt.Scanln(&scientificName)

	var attributes []string
	fmt.Println("Please enter attributes for the flower (type 'done' to finish):")
	for {
		var attribute string
		fmt.Scanln(&attribute)
		if attribute == "done" {
			break
		}
		attributes = append(attributes, attribute)
	}

	almanac[scientificName] = attributes

	fmt.Println("Flower added to the almanac successfully.")
}

// Function to print the almanac
func printAlmanac(almanac map[string][]string) {
	fmt.Println("\nAlmanac:")
	for scientificName, attributes := range almanac {
		fmt.Printf("%s -> %v\n", scientificName, attributes)
	}
}


/*

Create a flower almanac to have different plants being registered in our program

This program will have 4 functionalities:

    - Add flower
	- List all flowers in the almanac
	- Show Flower Information (the user has to specify by the name the plant they want to see)
	- Save Almanac (save the data on a txt file)
	- Exit

	(after having Save Almanac, we will learn how to import data from a file)

	    Key             Value
	Flower Name -> Scientific name, Brief Description, Main color, Flower family



	//YOU DON'T HAVE TO USE FUNCTIONS FOR THIS ONE




*/

/*

func main() {

	var almanac = map[string][]string{} //our map will be a map of string slices
	//    KEY                             VALUE
	// daffodil -> [asdasdasd, it's a beautiful flower, white, asfdasdads]
	//                 0                  1               2         3
	// narcissus -> [asdasdasd, it's a beautiful flower, white, asfdasdads]
	//                 0                  1               2         3

	//REMEMBER TO DO THE MENU

	//THIS CODE WILL BE FOR OPTION 1
	var scientificName string

	fmt.Println("Please give me the scientific name of the flower:")
	fmt.Scan(&scientificName, almanac)

	for scientificName != "Name of flower" {
		fmt.Scan(&scientificName)
	}

	almanac ["Forget-me-not"] = "Myosotis"

	almanac ["daffodil"] = "Narcissus"

	almanac ["Columbine"] = "Aquilegia"

	almanac ["Bloom"] = "Chrysanthemum"

	almanac ["Broom"] = "Genista"

	var exit string
			fmt.Println("Exit shopping list ")
			fmt.Scan(exit)

} */

/*EXAMPLE ON HOW TO CREATE A NEW PAIR



                   |
				   |
				   |
				   |
				   +


var data = []string{scientificName, "Daffodil plants have a single flower on a long green stalk, with green leaves growing from the base of the stem.", "Yellow", "Amaryllidaceae"}

almanac["Daffodil"] = data

*/
