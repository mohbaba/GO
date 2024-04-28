package main
import "fmt"


var commission float32 = 200
var totalSales float32

func main()  {
	var item int
	for{
		fmt.Println("Enter item value(enter 0 to end program): ")
		fmt.Scan(&item)
		if item == 0 {
			break
		}
		totalSales += getPrice(checkItem(item))
	}
	commission = commission + (totalSales * 0.09)
	fmt.Println("Your earnings for this week is ", commission)
}

func getPrice(item int)  float32{
	switch item {
	case 1:
		return 239.99
	case 2:
		return 129.75
	case 3:
		return 99.95
	case 4:
		return 350.89
	default:
		return 0
	}
}

func checkItem(item int) int{
	for item < 1 || item > 4 {
		fmt.Println("Enter item number between 1 and 4: ")
		fmt.Scan(&item)
	}
	return item
}


