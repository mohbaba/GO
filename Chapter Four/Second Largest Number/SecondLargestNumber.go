package main

import "fmt"

func main()  {
	var number int
	var highest int = 0
	var secondHighest int  = 0

	for i := 0; i < 10; i++ {
		fmt.Printf("Enter number %d: ",i + 1)
		fmt.Scan(&number)
		if number > highest {
			secondHighest = highest
			highest = number
		}
		if secondHighest > number {
			secondHighest = number
		}
	}
	fmt.Println("The highest number is ", highest)
	fmt.Println("The second highest is ", secondHighest)
}