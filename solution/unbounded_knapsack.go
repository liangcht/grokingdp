package solution

import (
	"fmt"

	"grokingdp/mathutil"
)

func unboundedKnapsack(profits []int, weights []int, capacity int) int {
	n := len(profits)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, capacity+1)
	}

	for i := 0; i < n; i++ {
		dp[i][0] = 0
	}

	for i := 0; i < capacity+1; i++ {
		dp[0][i] = profits[0]
	}

	for i := 1; i < n; i++ {
		for j := 1; j < capacity+1; j++ {
			dp[i][j] = dp[i-1][j]
			if j >= weights[i] {
				dp[i][j] = mathutil.MaxInt(dp[i][j], profits[i]+dp[i][j-weights[i]])
			}
		}
	}
	return dp[n-1][capacity]
}

// TestUnboundedKnapsack is a test
func TestUnboundedKnapsack() {
	fmt.Println(unboundedKnapsack([]int{15, 50, 60, 90}, []int{1, 3, 4, 5}, 8))
	fmt.Println(unboundedKnapsack([]int{15, 50, 60, 90}, []int{1, 3, 4, 5}, 6))
}
