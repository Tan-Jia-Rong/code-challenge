package main

import (
	"fmt"
)

// Solution 1: Iterative Approach | Time: O(n) | Space: O(1)
func sum_to_n_a(n int) int {
	var sum int = 0

	// For loop | Alternatively can use a while loop
	for i := 1; i <= n; i++ {
		sum += i
	}

	return sum
}

// Solution 2: Recursive Approach | Time: O(n) | Space: O(n) [stack space]
func sum_to_n_b(n int) int {
	if n == 1 {
		return 1
	}

	return n + sum_to_n_b(n-1)
}

// Solution 3: Math Approach | Time: O(1) | Space: O(1)
// Notice that: 1 + 2 + ... + n = ((n + 1)(n)) / 2
func sum_to_n_c(n int) int {
	return ((n + 1) * n) / 2
}

func main() {
	var n int
	fmt.Print("Enter any integer: ")
	_, err := fmt.Scanln(&n)
	if err != nil {
		fmt.Println("Error: Please enter a valid integer.")
		return
	}

	fmt.Println("Using sum_to_n_a: ", sum_to_n_a(n))
	fmt.Println("Using sum_to_n_b: ", sum_to_n_b(n))
	fmt.Println("Using sum_to_n_c: ", sum_to_n_c(n))
}
