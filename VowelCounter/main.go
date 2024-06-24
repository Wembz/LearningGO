package main


import (
	"fmt"
	"strings"
)

     


	


/* Develop a Go program that prompts the user to enter a sentence. Using a for loop, iterate through each character
of the sentence.  If the character is a vowel (a, e, i, o, u), count it and print the total number of vowels in the sentence.
Use continue to skip counting if the character is not a vowel.


	var VowelCounter int
	fmt.Println("Please write a setence to count vowel characters: ")
	int := bufio.NewReader(os.Stdin)

	sentence, _ := in.ReadString('\n')
	sentence = strings.ToLower(sentence)
	for _, char = range sentence {
		switch char {
		case "a", "e", "i", "o", "u":
			VowelCounter++
		default:
			continue

		}
	}
	fmt.Println("Total vowels presented in the sentence: %d\n ", VowelCounter)

*/
	
	func main() {
		var sentence string
		// Prompt user to enter a sentence using "fmt.scan"
		fmt.Print("Enter a sentence: ")
		fmt.Scanln(&sentence)
	
		sentence = strings.ToLower(sentence)
	
		vowelCount := 0
	
		// Iterate through each character of the sentence
		for _, char := range sentence {
			// Check if the character is a vowel

		// purpose of the "CHAR" a single display unit of information equivalent to one alphabetic
		// symbol, digit, or letter. 
			switch char {
			case 'a', 'e', 'i', 'o', 'u':
				vowelCount++
			default:
				// Skip counting if the character is not a vowel
				continue
			}
		}
	
		// Print the total number of vowels in the sentence
		fmt.Printf("Total number of vowels in the sentence: %d\n", vowelCount)
	}
	

