package main

import "fmt"

func main() {

	/*var Grades = [10]int{}

	for i := 0; i < len(Grades); i++ {
		fmt.Println("Please enter a student Grade: ")
		fmt.Scan(&Grades[i])

	}

	//How can i add one more grade to the array? EXAMPLE BELOW
	
	- WHEN THERE NO LIMIT TO THE DATA( LIKE THIS TYPE "1,0000 OR 1.12") 
	
	- YOU SHOULD USE SLICE WHEN THERE IS A LIMIT USE AN ARRAY

	SCHOOL DATA COMPANY - ARRAY 
	JUSTEAT APP - ARRAY
	HOW MANY PRODUCTS - SLICES
	HOW MANY FOOTBAL PLAYERS - SLICES
	/*var Grades2 = [11]int{} - SLICES

	for i := 0; i < len(Grades); i++ {
		Grades2[i] = Grades[i]

		//How can i remove a grade from the array

		Grades[9] = -1
		for i := 0; i < len(Grades); i++ {
		if Grades[i] <0{

		}

		for i := 0; i < len(&Grades); i++ {
		fmt.Println("Please enter a Grade: ")
		fmt.Scan(&Grades[i])

			for i := 0; i < len(Grades); i++ {
		fmt.Printf("Grades %d : %d\n", i, Grades[i])

		fmt.Println("Grades" i, ":",Grades[i]) another way to execute program
	}
	Grades[4] = -1 // you can choose any number in your order

	// Asking user for Grades
	for i := 0; i < len(Grades); i++ {
		if Grades[i] < 0 {
			fmt.Println("Please enter a  student grade:")
			fmt.Scan(&Grades[i])
		}
	} // print out grade results
	for i := 0; i < len(Grades); i++ {
		fmt.Printf("Grades %d : %d\n", i, Grades[i])

	}

	var g = []int{} // slices have no fixed size

	// g = append(g, 1) // adding the number 1 to my slice, my slice "g" will now have size 1
	
	var num int

	for i := 0; i < 10; i++ {

		fmt.Println("Please enter a  student grade:")
		fmt.Scan(&num)
		g = append(g, num)
	}
*/
	run := true
	num := 0  
	var g = []{}//open array for user to enter

	for run == true{

		fmt.Println("Please enter a  student grade:")
		fmt.Scan(&num)
		if num < 0 {
			num = false
		}else {
		g = append(g, num)

	}
	fmt.Println("Do you want to exit: ") // for user purpose and canbe 
	fmt.Scan(&response) // activate new word "response"
	if response == "yes"|| response =="YES" { // word activated 
		run = false
	}
}
// we want to remove first grade

g[0] = g[len(g) - 1] // copying the grade of last position to the first position

g = g[0:len(g)-2] // how to remove position 

for i :=0; i < len(g); i++{


}
}
