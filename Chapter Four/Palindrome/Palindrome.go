package main

import "fmt"

func main() {
	var input int
	fmt.Print("Enter a five-digit integer: ")
	_, err := fmt.Scan(&input)

	for err != nil || input < 10000 || input > 99999 {
		fmt.Println("Error: Input must be a five-digit integer.")
		_, err = fmt.Scan(&input)
		if err == nil && (input > 10000 && input < 99999) {
			break
		}
	}

	if isPalindrome(input) {
		fmt.Println("Yes, the number", input, "is a palindrome.")
	} else {
		fmt.Println("No, the number", input, "is not a palindrome.")
	}
}

func isPalindrome(n int) bool {
	reversed := 0
	original := n

	for n > 0 {
		remainder := n % 10
		reversed = reversed*10 + remainder
		n /= 10
	}

	return original == reversed
}
