package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Oh look who it is:")
	fmt.Println("Please enter your name")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	name := scanner.Text()
	fmt.Println(greetingMessage(name))

}
func celsiusToFehrenheit(temperatureCelsius float64) float64 {

	return (temperatureCelsius * float64(9) / float64(5)) + 32

}

func calculateAverage(subject1, subject2, subject3 float64) float64 {

	return ((subject1 + subject2 + subject3) / 3)

}

func greetingMessage(name string) string {

	return fmt.Sprintf("Welcome to the matrix %s please enter..this is... Utophia", name)

}
