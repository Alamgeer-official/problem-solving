package main

import "fmt"

func main() {
	fmt.Printf("Patter problem solving....\n")
	// 1 )
	// *****
	// *****
	// *****
	// *****
	pattern1(4)

	//  2)
	// *
	// **
	// ***
	// ****
	pattern2(4)

	//  3)
	// 1
	// 22
	// 333
	// 4444
	pattern3(4)

	//  4)
	// 1
	// 12
	// 123
	// 1234
	pattern4(4)

	//  5)
	// ####
	// ###
	// ##
	// #
	pattern5(4)

	//  6)
	// 1234
	// 123
	// 12
	// 1
	pattern6(4)

	//  7)
	// 	  *
	// 	 ***
	//  *****
	// *******
	pattern7(4)

}

func pattern1(n int) {
	fmt.Print("\n1st problem output\n")
	for row := 0; row < n; row++ {
		for column := 0; column < n; column++ {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
}

func pattern2(n int) {
	fmt.Print("\n2nd problem output\n")
	for row := 0; row < n; row++ {
		for column := 0; column <= row; column++ {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
}

func pattern3(n int) {
	fmt.Print("\n3rd problem output\n")
	for row := 1; row <= n; row++ {
		for column := 1; column <= row; column++ {
			fmt.Printf("%d", row)
		}
		fmt.Printf("\n")
	}
}

func pattern4(n int) {
	fmt.Print("\n4th problem output\n")

	for row := 1; row <= n; row++ {
		for col := 1; col <= row; col++ {
			fmt.Printf("%d", col)
		}
		fmt.Printf("\n")
	}
}

func pattern5(n int) {
	fmt.Print("\n5th problem output\n")
	for row := 0; row < n; row++ {
		for col := 0; col < n-row; col++ {
			fmt.Printf("#")
		}
		fmt.Printf("\n")
	}
}

func pattern6(n int) {
	fmt.Print("\n6th problem output\n")
	for row := 1; row <= n; row++ {
		for col := 1; col <= n-row+1; col++ {
			fmt.Printf("%d", col)
		}
		fmt.Printf("\n")
	}

}

func pattern7(n int) {

	fmt.Print("\n7th problem output\n")
	for row := 0; row < n; row++ {
		//for spaces
		for col := 0; col < n-row-1; col++ {
			fmt.Printf(" ")
		}
		//for *
		for col := 0; col < 2*row+1; col++ {
			fmt.Printf("*")
		}
		//for spaces
		for col := 0; col < n-row-1; col++ {
			fmt.Printf(" ")
		}

		fmt.Println()
	}
}
