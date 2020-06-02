package solution

import (
	"fmt"
)

func coinChange(denominations []int, amount int) int {
	n := len(denominations)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, amount+1)
	}

	for i := 0; i < n; i++ {
		dp[i][0] = 1
	}

	for i := 1; i < amount+1; i++ {
		if i%denominations[0] == 0 {
			dp[0][i] = 1
		} else {
			dp[0][i] = 0
		}
	}

	for i := 1; i < n; i++ {
		for j := 1; j < amount+1; j++ {
			dp[i][j] = dp[i-1][j]
			if j >= denominations[i] {
				dp[i][j] += dp[i][j-denominations[i]]
			}
		}
	}
	return dp[n-1][amount]
}

// TestCoinChange is used for testing
func TestCoinChange() {
	fmt.Println(coinChange([]int{1, 2, 3}, 5))
}
