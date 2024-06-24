package main

/* var floatValue float64 // Declaration of a float variable

// Assigning a value to the float variable
floatValue = 3.14

// Printing the value of the float variable
fmt.Println("Float Value:", floatValue)

// Performing operations with floats
anotherFloat := 2.5
result := floatValue * anotherFloat 

* = MULTIPLY 

fmt.Println("Result of multiplication:", result)

// Updating the float variable with a new value
floatValue = 5.678

// Printing the updated value
fmt.Println("Updated Float Value:", floatValue)
*/

import (
	"fmt"
)

func main() {
	var floatValue float64 

	floatValue = 3.14
	fmt.Println("Float Value:", floatValue)
	
	// Multiplication 
	anotherFloat := 2.5
	result := floatValue * anotherFloat
	fmt.Println("Updated Float Value:", floatValue)
	floatValue = 5.678
    
	fmt.Println("Result of multiplication:", result)
}
