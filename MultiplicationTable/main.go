package main

import "fmt"

/*

I chose to do multiplications for the numbers below the letters represent the number i have chosen
H = 8
J= 10
import "fmt"
*/

func main() {

	for H := 1; H <= 10; H++ {
		fmt.Printf("Multiplication table for %d\n", H)
		for j := 1; j <= 10; j++ {
			fmt.Printf("%d x %d = %d\n", H, j, H*j)
		
		}
	}
}
