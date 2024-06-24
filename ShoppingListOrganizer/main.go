package main

import "fmt"

/*
==============================Shopping List Organizer:==============================
Write a Go program to help users organize their shopping list.
Define a slice to store the items on the shopping list.
Prompt the user to add items to the shopping list.
Use a for loop to iterate through the shopping list and print each item.
Additionally, allow the user to remove items from the shopping list
by specifying the item's index.

TIP: Guys in here just use slice, without array.
To declare a slice that is independet from na array is like this:

shoppinglist := []string{ }

In this case the slice will be useful because as you know slices
can grow and shrink without having a limit
*/
func main() {

	var UserInput int
	shoppinglist := []string{}

	for UserInput != 4 {
		fmt.Println("Press 1 to View cart: ")
		fmt.Println("Press 2 to add item to cart: ")
		fmt.Println("Press 3 to remove item from shopping list: ")
		fmt.Println("Press 4 to exit")
		fmt.Scan(&UserInput)

		if UserInput == 1 {
			fmt.Println("Shoppinglist")
			for i := 0; i < len(shoppinglist); i++ {
				fmt.Printf("%d %s\n", i+1, shoppinglist[i])
			}
			if len(shoppinglist) == 0 {
				fmt.Println("Shopping list empty")
			}
		} else if UserInput == 2 {
			var items string
			fmt.Println("Please add item to shopping list")
			fmt.Scan(&items)
		shoppinglist = append(shoppinglist,items)	
		}else if UserInput == 3{
			var ItemNumber int
			fmt.Println("Please enter an Item number to remove from shopping list")
			fmt.Scan(ItemNumber)
			shoppinglist = append(shoppinglist[ItemNumber-1:],shoppinglist[ItemNumber])
		}else if UserInput == 4{
			var exit string
			fmt.Println("Exit shopping list ")
			fmt.Scan(exit)
		}

	}

}
