package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

/*

	Create a project that asks the user for 5 books, each one related to 1 of this billionaires:

	- Bill Gates
	- Elon Musk
	- Jeff Bezos
	- Warren Buffet
	- Donald Trump

	Write those 5 books to a file (you can call it books.txt)

	Then read only the 3 first books from the file

	And with those 3 first books, sort them by ascending order (alphabetically)



	TIP

	- Use an array
	- Create a file in your computer in the same folder as the project's folder
	- Use the Artists Project to your advantage (have both projects side by side)
	- Ask Google how to sort a slice/array of strings by ascending order


	75 POUNDSSSSSSSS

*/

func main() {

	var books = []string{}

	var book string

	for true {
		
		// Ask User to enter book name  
		fmt.Println("Please enter name of book")

		fmt.Scan(&book)

		if book == "exit" {
			break
		}

		books = append(books, book)
	}

	sort.Strings(books)
	// Print the sorted slice of strings.
	fmt.Println("Sorted books:", books)

	// Creates Book text file
	file, err := os.Create("books.txt") // "os" operation system || "err" Error

	//IF STATEMENT JUST IN CASE ITS WRONG
	if err != nil { // if the err variable is not empty
		fmt.Println("There was an error creating the file") // check for errors
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file) //writer is a variable that will allow us to write to the file

	for _, book := range books {

		writer.WriteString((book + "\n"))
		writer.Flush() //flush means flushing the toilet, if there's data remaining in the writer, that data will be flushed
		//if there's data that wasn't written yet, now it will
	}

	// Reopen the file
	file, ferr := os.Open("books.txt")
	if ferr != nil {
		panic(ferr)
	}
	defer file.Close() // "Defer" Used to close the file

	scanner := bufio.NewScanner(file)

	// Read first 3 books
	fmt.Println("First 3 books")
	for i := 0; i < 3; i++ { //run a for loop  3 books only
		if scanner.Scan() {
			fmt.Println(scanner.Text())
		} else {
			break // Break the loop if there nomore lines to read
		}
	}

	//for scanner.Scan() {
	//	line := scanner.Text()
	//	fmt.Println(line)
	//}

	defer file.Close() //After using the file, we need to close the file using "defer"

}
