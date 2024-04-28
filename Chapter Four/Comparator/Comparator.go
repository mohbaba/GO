package main

import "fmt"


func main()  {
	var input1 int
	var input2 int

	for{
		fmt.Println("Enter first number,Press Ctrl + C to end program: ")
		_, err := fmt.Scan(&input1)
		for{
			if err != nil {
				fmt.Println("Invalid input, Enter first number,Press Ctrl + C to end program: ")
				_, err = fmt.Scan(&input1)	
			}else{
				break
			}
		}
		fmt.Println("Enter second number,Press Ctrl + C to end program: ")
		_, exception := fmt.Scan(&input2)
		for{
			if exception != nil {
				fmt.Println("Invalid input, Enter second number,Press Ctrl + C to end program")
				_, exception = fmt.Scan(&input2)

			}else{
				break
			}
		}

		

		if input2 > input1 {
			println("-1")
		}else if input2 < input1 {
			println("1")
		}else if input1 == input2 {
			println("0")
		}
	}

}