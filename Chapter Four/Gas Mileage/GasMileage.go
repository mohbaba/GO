package main

import "fmt"

var miles int
var gallons int
var milesPerGallon  float32
var totalMilesDriven int
var totalGallonsUsed int
var totalMilesPerGallon float32

func main()  {
	
	saveMilesDriven()
	for miles != 0 || gallons != 0 {
		
	
		for miles < 0 {
			saveMilesDriven()
			
		}
		totalMilesDriven += miles

		saveGallonsUsed()
		for gallons < 0 {
			saveGallonsUsed()
		}
		totalGallonsUsed += gallons
		milesPerGallon = float32(miles)/float32(gallons)
		fmt.Print("The miles per gallon for this trip is: ",milesPerGallon)
	}
	totalMilesPerGallon = float32(totalMilesDriven)/float32(totalGallonsUsed)
	fmt.Print("The total miles per gallons used is: ",totalMilesPerGallon)
}

func print(message string)  {
	fmt.Println(message)
}

func saveMilesDriven()  {
	print("Enter miles driven(Enter 0 to stop program): ")
	fmt.Scan(&miles)
	
	
}

func saveGallonsUsed()  {
	print("Enter Gallons used(Enter 0 to stop program): ")
	fmt.Scan(&gallons)
	
}