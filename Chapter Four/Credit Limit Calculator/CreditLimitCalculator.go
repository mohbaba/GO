package main

import (
	"fmt"
	"strings"
)

var accountNumber int
var beginningBalance int
var creditLimit int
var totalCharges int
var totalCredits int


func main()  {
	var newBalance int
	for{
		accountNumber = collectCustomerAccountNumber()
		beginningBalance = collectBeginningBalance()
		totalCharges = collectTotalCharges()
		totalCredits = collectTotalCredits()
		creditLimit = collectCreditLimit()
		newBalance = beginningBalance + totalCharges - totalCredits
		fmt.Printf("The new balance of customer: %v is %v\n",accountNumber, newBalance)
		if newBalance > creditLimit {
			fmt.Println("Credit limit exceeded")
		}else{
			fmt.Println("Credit limit not exceeded")

		}
		var sentinel string
		var sentinelVariable string
		sentinel = breakFromLoop("Would you like to continue operation(Yes or No)",sentinelVariable)
		if strings.EqualFold(sentinel,"no") {
			break
		}


	}
}

func breakFromLoop(prompt string, sentinel string)  string{

	for {
		println(prompt)
		_, err := fmt.Scan(&sentinel)
		if err == nil {
			break
		}	
		for err != nil ||!strings.EqualFold(sentinel,"yes") || !strings.EqualFold(sentinel,"no") {
			
			fmt.Println("Please Enter a valid Input: (Yes or No) ")
			_, err = fmt.Scan(&sentinel)
			if err == nil {
				break
			}	
			
			
		}
		
		
	}
	
	
	
	return sentinel
}
func collectCustomerAccountNumber()int  {
	for{
		fmt.Println("Please enter your account number: ")
		_, err := fmt.Scan(&accountNumber)
		if err != nil {
			fmt.Print("Invalid input\nEnter Account Number: ")
			fmt.Scan(&accountNumber)
		} else {
			break
		}

	}
	return accountNumber
}


func collectBeginningBalance()  int{
	for{
		fmt.Println("Please enter your balance at the beginning of the month: ")
		_, err := fmt.Scan(&beginningBalance)
		if err != nil {
			fmt.Print("Invalid input\nyour balance at the beginning of the month: ")
			_, err = fmt.Scan(&beginningBalance)
		} else {
			break
		}

	}
	return beginningBalance
}

func collectCreditLimit()  int{
	for {
		fmt.Println("Please enter your credit limit: ")
		_, err := fmt.Scan(&creditLimit)
		if err != nil {
			fmt.Print("Invalid input\nEnter Credit Limit: ")
			_, err = fmt.Scan(&creditLimit)
		} else {
			break
		}
	}
	return creditLimit
}

func collectTotalCharges()  int{
	for {
		fmt.Println("Please enter total charges of this month: ")
		_, err := fmt.Scan(&totalCharges)
		if err != nil {
			fmt.Print("Invalid input\nEnter total charges: ")
			_, err = fmt.Scan(&totalCharges)
		} else {
			break
		}
	}
	return totalCharges
}

func collectTotalCredits()  int{
	for {
		fmt.Println("Please enter totaldits applied this month: ")
		_, err := fmt.Scan(&totalCredits)
		if err != nil {
			fmt.Print("Invalid input\nEnter Total credits: ")
			_, err = fmt.Scan(&totalCredits)
		} else {
			break
		}
	}
	return totalCredits
}

