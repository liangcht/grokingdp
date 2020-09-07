package solution

import (
	"fmt"
	"grokingdp/mathutil"
)

func editDistance(s1, s2 string) int {
	dp := make([][]int, len(s1)+1)
	for i := range dp {
		dp[i] = make([]int, len(s2)+1)
	}
	dp[0][0] = 0
	for i := 1; i < len(s1)+1; i++ {
		dp[i][0] = 1
	}
	for i := 1; i < len(s2)+1; i++ {
		dp[0][i] = 1
	}

	for i := 1; i < len(s1)+1; i++ {
		for j := 1; j < len(s2)+1; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = mathutil.MinInt(dp[i-1][j], dp[i][j-1])
				dp[i][j] = mathutil.MinInt(dp[i][j], dp[i-1][j-1])
				dp[i][j]++
			}
		}
	}
	return dp[len(s1)][len(s2)]
}

// TestEditDistance is used for testing
func TestEditDistance() {
	fmt.Println(editDistance("bat", "but"))
	fmt.Println(editDistance("abdca", "cbda"))
	fmt.Println(editDistance("passpot", "ppsspqrt"))
}
