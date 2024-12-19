package solutions

import (
	"errors"
	"fmt"
)

type Stack struct {
	values []int
	minVal []int
}

func NewStack() Stack {
	return Stack{}
}

func (s *Stack) Push(element int) {
	// Add value to the main stack
	s.values = append(s.values, element)
	fmt.Printf("Element %d added\n", element)

	// Update minVal stack
	if len(s.minVal) == 0 || element <= s.minVal[len(s.minVal)-1] {
		s.minVal = append(s.minVal, element)
	}
}

func (s *Stack) Pop() error {
	if len(s.values) == 0 {
		return errors.New("stack is empty")
	}

	// Pop value from main stack
	popedVal := s.values[len(s.values)-1]
	s.values = s.values[:len(s.values)-1]

	// Update minVal stack if necessary
	if len(s.minVal) > 0 && popedVal == s.minVal[len(s.minVal)-1] {
		s.minVal = s.minVal[:len(s.minVal)-1]
	}

	fmt.Printf("Element %d removed\n", popedVal)
	return nil
}

func (s *Stack) GetMinVal() int {
	if len(s.minVal) == 0 {
		fmt.Println("Stack is empty")
		return -1 // Return an invalid value to indicate error
	}

	min := s.minVal[len(s.minVal)-1]
	fmt.Println("Minimum element:", min)
	return min
}

func StackQuestions() {
	stack1 := NewStack()
	fmt.Printf("Stack1: %v\n", stack1.values)

	stack1.Push(9)
	stack1.Push(7)
	stack1.Push(5)
	stack1.Push(3)
	stack1.Push(1)

	fmt.Printf("After pushing all elements: %v\n", stack1.values)
	stack1.GetMinVal()

	stack1.Pop()
	stack1.Pop()
	fmt.Printf("After popping elements: %v\n", stack1.values)
	stack1.GetMinVal()
}
