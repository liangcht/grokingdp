package solution

import (
	"fmt"
	"grokingdp/mathutil"
)

func minimumCoinChange(denominations []int, amount int) int {
	n := len(denominations)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, amount+1)
	}

	for i := 0; i < n; i++ {
		dp[i][0] = 0
	}

	for i := 1; i < amount+1; i++ {
		if i%denominations[0] == 0 {
			dp[0][i] = i / denominations[0]
		} else {
			dp[0][i] = -1
		}
	}

	for i := 1; i < n; i++ {
		for j := 1; j < amount+1; j++ {
			dp[i][j] = dp[i-1][j]
			if j >= denominations[i] && dp[i][j-denominations[i]] >= 0 {
				dp[i][j] = mathutil.MinInt(dp[i][j], 1+dp[i][j-denominations[i]])
			}
		}
	}
	return dp[n-1][amount]
}

// TestMinimumCoinChange is used for testing
func TestMinimumCoinChange() {
	fmt.Println(minimumCoinChange([]int{1, 2, 3}, 5))
	fmt.Println(minimumCoinChange([]int{1, 2, 3}, 11))
	fmt.Println(minimumCoinChange([]int{1, 2, 3}, 7))
	fmt.Println(minimumCoinChange([]int{3, 5}, 7))
}
