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
