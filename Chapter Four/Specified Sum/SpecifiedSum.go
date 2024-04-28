package main
import "fmt"


func main()  {
	var input int
	var total int
	fmt.Println("Enter limit to which numbers should be added to: ")
	_,err := fmt.Scan(&input)
	for {
		if err != nil {
			fmt.Println("Invalid input: enter a number")
			_,err = fmt.Scan(&input)
		}else{
			break
		}
	}

	for i := 0; i <= input; i++ {
		total += i
	}
	
	fmt.Printf("The total sum of numbers from 1 to %d is %d",input,total)
}