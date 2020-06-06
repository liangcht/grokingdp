package solution

import (
	"fmt"
	"grokingdp/mathutil"
)

func minimumJumpsWithFee(fee []int) int {
	n := len(fee)
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = 0
	}
	dp[0] = 0
	dp[1] = fee[0]
	dp[2] = fee[0]
	dp[3] = fee[0]
	for i := 4; i < n+1; i++ {
		dp[i] = dp[i-1] + fee[i-1]
		dp[i] = mathutil.MinInt(dp[i], dp[i-2]+fee[i-2])
		dp[i] = mathutil.MinInt(dp[i], dp[i-3]+fee[i-3])
	}
	return dp[n]
}

// TestMinimumJumpsWithFee is used for testing
func TestMinimumJumpsWithFee() {
	fmt.Println(minimumJumpsWithFee(([]int{1, 2, 5, 2, 1, 2})))
	fmt.Println(minimumJumpsWithFee(([]int{2, 3, 4, 5})))
}
