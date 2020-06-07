package solution

import (
	"fmt"
)

func fibonacciNumber(n int) int {
	if n < 2 {
		return n
	}
	dp := make([]int, n)
	dp[0] = 1
	dp[1] = 1

	for i := 2; i < n; i++ {
		dp[i] = dp[i-2] + dp[i-1]
	}
	return dp[n-1]
}

// TestFibinacciNumber is used for testing
func TestFibinacciNumber() {
	fmt.Printf("5th Fibonacci is ---> %d\n", fibonacciNumber(5))
	fmt.Printf("6th Fibonacci is ---> %d\n", fibonacciNumber(6))
	fmt.Printf("7th Fibonacci is ---> %d\n", fibonacciNumber(7))
}
