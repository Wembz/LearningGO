package main

import "fmt"

/* creating a variable that will Print the first 50 even numbers */

   /* 
		//
		1st Initialize the variable 
		2nd Condition -> condition that is varified every iteration
		(every time the loop jump)
		3rd  Incrementing -> what makes the variable change 

		example below is for Running normal code 
		number := 2

	for number <= 100 {
	   
		    fmt.Println(number)
		number += 2

		formular below is for numeric numbers 


var number int32 
var factorial int32 = 1

for number :=5; number >= 2; number -= 1 {
	factorial = factorial · number
}

fmt.Println(factorial)
	}
	
	
	
	
	for number :=2; number <=100; number += 2 {
	 if number == 24 {
		continue //continue is a code instruction that will send the execution to 
		         // the line 45 because it ends the iteration
	 }
   fmt.Println(number)
   
   */


	// Print the factorial of a given number

// 5 x 4 * 3 * 2 * 1 = 120

// "factorial" is to print the facts of what is being printed 

func main (){



   for number :=2; number <=100; number *= 2 {
	if number == 24 {
	   break 

	}
  fmt.Println(number)

	}


	
	for number :=100; number >=2; number -= 2 {
	 if number == 24 {
		continue //continue is a code instruction that will send the execution to 
		         // the line 45 because it ends the iteration
	 }
   fmt.Println(number)

}

}