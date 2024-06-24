package main

/*


Todays lesson to see the similarities between Arrays and Strings

Each character associated with the string so we can access them.

Name [] 0  to access the string

Arrays work the same way


		BEST TIME TO USE AN ARRAY

		"ASK THE USER FOR  number of "games won"
		ASK THE USER WHAT THE AGE OF CLASS MATES

in arrays you have [5] you can have numbers inside them.

i want an array of names todays so my array will be called "name" ,
i wanted to have space for 5 names so i used strings because names are words.
If i wanted an array of number of goals i would used INT becuase
they are interger  numbers

===== 3 WAYS TO CREATE AN ARRAY ======

>> Var name string = "Aesha"
>> Names := [5]string{"Aesha", "Keshif", "Mohamed", "Rodrigue", "Mahad"}
>> Var Names = [5]string{"Aesha", "Keshif", "Mohamed", "Rodrigue", "Mahad"}

Arrays have INDEX'S they make then much more simplier

===== ALWAYS ALWAYS ALWAYS ALWAYS Always use "NAME" TO IDENTIFY VARIABLES(var) ======



How to positon Arrays in order rather then a line

"i" := 0 is used as the "Position the number 0 on the arrays we want to stop at 4 "

	for i := 0; i < len(Names); i++{
		fmt.Println(Names[i])
	}

    len is not part of the "string" variable


i++ "it means to go up in +1"
i-- "to go down"
in programming they used "i" or "j" alot stick to using "i"

===
[] = to run your program down into each row

== example of Array below ===
fmt.Println(Names[i])
====

hhow to ask the user his music option:

	}
var favouriteMusics = [5]string {} " if left empty its for user to provide information "

for j :=0; j < len(favouriteMusics); j++{
	len <- stands for length of the array "postion 0 - position 10"

	fmt.Println("Please enter the name of a music:")
	fmt.Scan(&favouriteMusics[j])
}
	for j := 0; j < len(Names); j++{
		fmt.Println(Names[j])

		=== how to read sentences ====

		reader(can give it any name):= bufio.NewReader(os.Stdin) allows us to read data(sentence)
		STDIN = reads data in byes format for the program all at once.

		scan cannot read all the variable will only read one by one.

		bufio reader allows us to read rthe phrase at once

		"Read string " allows us to read by line

		=== example below ====

		}

	var favouriteMusics = [5]string{}

	for j := 0; j < len(favouriteMusics); j++ {

		fmt.Println("Please enter the name of a music:")

		favouriteMusics[j] ,_ = reader.ReadString('\n')

		=========         ==============

		control + C to stop program

		You are able to change an Array at any moment 

fmt.Println("number in position ", j,':', numbers[j])

		APPEND - 
*/

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin) //"Read string " allows us to read by line

	Names := [7]string{"Aesha", "Keshif", "Mohamed", "Rodrigue", "Mahad", "jojo", "Matt"}

	var numbers = [10]int{}

	for j := 0; j < len(numbers); j++ {

		fmt.Println("Please enter a number:")
		fmt.Scan(&numbers[j])

	}
	for j := 0; j < len(Names); j++ {
		fmt.Println(Names[j])

		for i := 0; i < len(Names); i++ {
			fmt.Println(Names[i])
		}
		var favouriteMusics = [5]string{}

		for j := 0; j < len(favouriteMusics); j++ {

			fmt.Println("Please enter the name of a music:")

			favouriteMusics[j], _ = reader.ReadString('\n')

		}
		for j := 0; j < len(Names); j++ {
			fmt.Println(Names[j])

		}
		favouriteMusics[0]= "undefinded"

		fmt.Println(favouriteMusics[0])
	

var slice1 = numbers[0:4]

for i :=0; i < 4; i++{
slice1[i] = 20
}

for j := 0; j < len(numbers); j++ {

	fmt.Printf("number in position %d: %d\n", j,numbers[j])
	
}

var s = append(slice1, 20) // Append is to add another 20 to the slices you can use any numbers  
fmt.Println(s)
fmt.Println(numbers)

}

for index, a := range numbers{
fmt.Println(index,":", a)
}

}