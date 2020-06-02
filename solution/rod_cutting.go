package solution

import (
	"fmt"
	"grokingdp/mathutil"
)

func rodCutting(lengths []int, prices []int, rodLength int) int {
	n := len(lengths)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, rodLength+1)
	}

	for i := 0; i < n; i++ {
		dp[i][0] = 0
	}

	for i := 0; i < rodLength+1; i++ {
		if lengths[0] <= i {
			dp[0][i] = prices[0]
		} else {
			dp[0][i] = 0
		}
	}

	for i := 1; i < n; i++ {
		for j := 1; j < rodLength+1; j++ {
			dp[i][j] = dp[i-1][j]
			if j >= lengths[i] {
				dp[i][j] = mathutil.MaxInt(dp[i][j], prices[i]+dp[i][j-lengths[i]])
			}
		}
	}
	return dp[n-1][rodLength]
}

// TestRodCutting is used for testing
func TestRodCutting() {
	fmt.Println(rodCutting([]int{1, 2, 3, 4, 5}, []int{2, 6, 7, 10, 13}, 5))
}
