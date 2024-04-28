package main

import "fmt"

func main()  {
	var number int
	var highest int = 0
	for i := 0; i < 10; i++ {
		fmt.Printf("Enter number %d: ",i + 1)
		fmt.Scan(&number)
		if number > highest {
			highest = number
		}
	}
	fmt.Print("The highest number is ", highest)
}