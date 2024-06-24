package main

import "fmt"

/* ============================== Temperature Analysis ===========================================

Develop a Go program that analyzes daily temperatures recorded over a week.

Define an array or slice to store the temperatures for each day.
Prompt the user to enter the temperatures for each day of the week.
Use a for loop to iterate through the array/slice and calculate the average temperature,
maximum temperature, and minimum temperature for the week.*/

func main() {

	const numDays = 7 // "const is a datat type that defines something that can never change such as numbers "

	//Define an array or slice to store the temperatures for each day.
	temperatures := make([]float64, numDays)

	//Prompt the user to enter the temperatures for each day of the week.
	for i := 0; i < numDays; i++ {
		fmt.Printf("Please enter temperature for the day %d: ", i+1)
		fmt.Scan(&temperatures[i])

		//calculate the average temperature, maximum temperature, and minimum temperature for the week.
		averageTemperature := calculateAvgTemperature(temperatures)

		MaxTemperature := calculateMaximumTemperature(temperatures)

		MinTemperature := calculateMinimumTemperature(temperatures)

		fmt.Printf("\n Analysis Results: \n")
		fmt.Printf("\n Average Temperature: %2f\n", averageTemperature)
		fmt.Printf("\n Maximum Temperature: %2f\n", MaxTemperature)
		fmt.Printf("\n Minimum Temperature: %2f\n", MinTemperature)
	}
}
		func calculateAvgTemperature(temperatures []float64)float64 {
			sum := 0.0
			for _, temp := range temperatures{
				sum += temp
			}
			return sum / float64(len(temperatures))
		}

		func calculateMaximumTemperature(temperatures []float64) float64 {
			max := temperatures[0]
			for _, temp := range temperatures{
				if temp > max {
					max = temp

					return max
				}
			}
			return max
		}	
		func calculateMinimumTemperature(temperatures [] float64)float64{	
			min := temperatures[0]
			for _, temp := range temperatures{
				if temp < min {
					min = temp
				}
			}
			return min
		}
		
		
		