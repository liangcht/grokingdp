package solution

import (
	"fmt"
	"grokingdp/mathutil"
)

func longestRepeatingSubsequence(str string) int {
	n := len(str)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		dp[0][i] = 0
		dp[i][0] = 0
	}

	for i := 1; i < n; i++ {
		for j := i; j < n; j++ {
			if str[i-1] == str[j] {
				dp[i][j] = dp[i-1][j-1] + 1
			}
			dp[i][j] = mathutil.MaxInt(dp[i][j], dp[i-1][j])
			dp[i][j] = mathutil.MaxInt(dp[i][j], dp[i][j-1])
		}
	}
	return dp[n-1][n-1]
}

// TestLongestRepeatingSubsequence is used for testing
func TestLongestRepeatingSubsequence() {
	fmt.Println(longestRepeatingSubsequence("tomorrow"))
	fmt.Println(longestRepeatingSubsequence("aabdbcec"))
	fmt.Println(longestRepeatingSubsequence("fmff"))
	fmt.Println(longestRepeatingSubsequence("atactcgga"))
}
