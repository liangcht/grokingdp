package solution

import (
	"fmt"
	"grokingdp/mathutil"
)

func maximumRibbonCut(ribbonLengths []int, length int) int {
	n := len(ribbonLengths)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, length+1)
	}

	for i := 0; i < n; i++ {
		dp[i][0] = 0
	}

	for i := 1; i < length+1; i++ {
		if i%ribbonLengths[0] == 0 {
			dp[0][i] = i / ribbonLengths[0]
		} else {
			dp[0][i] = -1
		}
	}

	for i := 1; i < n; i++ {
		for j := 1; j < length+1; j++ {
			dp[i][j] = dp[i-1][j]
			if j >= ribbonLengths[i] && dp[i][j-ribbonLengths[i]] >= 0 {
				dp[i][j] = mathutil.MaxInt(dp[i][j], 1+dp[i][j-ribbonLengths[i]])
			}
		}
	}
	return dp[n-1][length]
}

// TestMaximumRibbonCut is used for testing
func TestMaximumRibbonCut() {
	fmt.Println(maximumRibbonCut([]int{2, 3, 5}, 5))
	fmt.Println(maximumRibbonCut([]int{2, 3}, 7))
	fmt.Println(maximumRibbonCut([]int{3, 5, 7}, 13))
	fmt.Println(maximumRibbonCut([]int{3, 5}, 7))
}
