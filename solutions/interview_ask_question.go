package solutions

import (
	"fmt"
	"time"
)

// FindFirstNameDistinctCountEmployee finds the count of distinct first names in the given list of employees.
func FindFirstNameDistinctCountEmployee() {
	// Problem Statement:
	// Given a list of employees, find the count of distinct first names.

	// Employee struct represents the structure of an employee.
	type Employee struct {
		ID          int64
		FirstName   string
		LastName    string
		Designation string
	}

	// Initialize the slice with values
	emp := []Employee{
		{
			ID:          time.Now().UnixNano(),
			FirstName:   "alamgeer",
			LastName:    "siddiqui",
			Designation: "SE",
		},
		{
			ID:          time.Now().UnixNano(),
			FirstName:   "alamgeer",
			LastName:    "shezad",
			Designation: "SE",
		},
		{
			ID:          time.Now().UnixNano(),
			FirstName:   "alamgeer",
			LastName:    "razi",
			Designation: "SE",
		},
		{
			ID:          time.Now().UnixNano(),
			FirstName:   "alamgeer",
			LastName:    "khan",
			Designation: "SE",
		},
		{
			ID:          time.Now().UnixNano(),
			FirstName:   "sarthak",
			LastName:    "sharma",
			Designation: "SE",
		},
		{
			ID:          time.Now().UnixNano(),
			FirstName:   "sarthak",
			LastName:    "ahuja",
			Designation: "SE",
		},
		{
			ID:          time.Now().UnixNano(),
			FirstName:   "sarthak",
			LastName:    "verma",
			Designation: "SE",
		},
	}

	// Logic of the Solution:
	// We use a hashtable to keep track of the count of each distinct first name.
	// Traverse the list of employees, and for each employee, check if their first name is in the hashtable.
	// If yes, increment the count; if no, add the first name to the hashtable with a count of 1.

	hashtable := map[string]int{}
	for _, employee := range emp {
		if count, exists := hashtable[employee.FirstName]; exists {
			hashtable[employee.FirstName] = count + 1
		} else {
			hashtable[employee.FirstName] = 1
		}
	}

	// Print the original list of employees and the hashtable with distinct first name counts.
	fmt.Printf("List of Employees: %v\n", emp)
	fmt.Printf("Hashtable with Distinct First Name Counts: %v\n", hashtable)

	// Time Complexity: O(n), where n is the number of employees in the list.
	// Space Complexity: O(d), where d is the number of distinct first names.
}

// ReverseGivenIntegerNumber reverses the digits of a given integer.
// It takes an integer 'number' as input and prints the reversed number.
func ReverseGivenIntegerNumber(number int) {
	// Problem Statement:
	// Given an integer, reverse its digits.

	// Logic:
	// The function uses a loop to extract the last digit of the input number,
	// adds it to the reversed number, and multiplies the reversed number by 10.
	// This process is repeated until all digits are processed.

	// Time Complexity:
	// The time complexity is O(log10(number)), where 'number' is the input integer.
	// The loop iterates once for each digit in the input number.

	// Space Complexity:
	// The space complexity is O(1) as the function uses a constant amount of space
	// regardless of the input size.

	// Initialize a variable to store the reversed number
	revNumber := 0

	// Iterate through each digit of the original number
	for number != 0 {
		// Extract the last digit of the number
		remainder := number % 10

		// Build the reversed number by adding the remainder
		// and multiplying the existing reversed number by 10
		revNumber = revNumber*10 + remainder

		// Remove the last digit from the original number
		number = number / 10
	}

	// Print the reversed number
	fmt.Printf("Reversed Number: %v\n", revNumber)
}

func FindCommonString() {

	fmt.Println("start.........")
	str1 := "abcd"
	str2 := "cdxuy"

	commonChan := make(chan rune)

	go func() {

		hash := make(map[rune]bool)
		for _, char := range str1 {
			hash[char] = true
		}

		for _, char := range str2 {
			if hash[char] {
				commonChan <- char
			}
		}
		close(commonChan)

	}()

	for c := range commonChan {
		fmt.Printf("%c\n", c)
	}

}

func ReverseStringAndPrint() {

	input := "TEST"
	// Convert string to a slice of runes to handle Unicode characters
	runes := []rune(input)
	n := len(runes) - 1 // Use len(runes) - 1 for clarity

	// Reverse the slice of runes in place
	for i := 0; i < len(runes)/2; i++ {
		runes[i], runes[n-i] = runes[n-i], runes[i]
	}

	// Convert the slice of runes back to a string
	reversed := string(runes)

	// Print the original and reversed strings
	fmt.Println("Original:", input)
	fmt.Println("Reversed:", reversed)
}
