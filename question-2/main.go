package main

import "fmt"

// recursiveFunction is a recursive function that takes an integer 'n' as input
// and prints the value of 'n' at each step until 'n' becomes 1.
func recursiveFunction(n int) {
	// Base case: If 'n' is 1, terminate the recursion.
	if n == 1 {
		return
	} else {
		// Recursive call with 'n' divided by 2.
		recursiveFunction(n / 2)
		// Print the current value of 'n'.
		fmt.Println(n)
	}
}

func main() {
	// Call the recursiveFunction with an initial value of 9.
	recursiveFunction(9)
}
