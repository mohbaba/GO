package main

import "fmt"

func main()  {
	var length int
	for {
		fmt.Println("Enter the length of the base of triangle")
		_, err := fmt.Scan(&length)
		if err != nil || length < 0 {
			fmt.Println("Invalid input\nEnter the length of the base of triangle")
			fmt.Scan(&length)
		}else{
			break
		}
	}
	printTriangle(length)
	
}

func printTriangle(length int)  {
	
	for i := 1; i < length+1; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}