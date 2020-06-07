package solution

import (
	"fmt"
	"grokingdp/mathutil"
)

func longestPalindromixSubsequence(input string) int {
	// return lps([]byte(input), 0, len(input)-1)
	n := len(input)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := range dp {
		dp[i][i] = 1
	}
	for l := 2; l <= n; l++ {
		for start := 0; start <= n-l; start++ {
			end := start + l - 1
			if input[start] == input[end] {
				dp[start][end] = 2
				if start < end && end > 0 {
					dp[start][end] += dp[start+1][end-1]
					continue
				}
			}
			dp[start][end] = dp[start][end-1]
			if start < end {
				dp[start][end] = mathutil.MaxInt(dp[start][end], dp[start+1][end])
			}
		}
	}
	return dp[0][n-1]
}

// Recursive function
func lps(input []byte, start int, end int) int {
	if start >= end {
		return 1
	}
	if start > end {
		return 0
	}
	if input[start] == input[end] {
		return lps(input, start+1, end-1) + 2
	}
	return mathutil.MaxInt(lps(input, start, end-1), lps(input, start+1, end))
}

// TestLongestPalindomicSubsequence is used for testing
func TestLongestPalindomicSubsequence() {
	fmt.Println(longestPalindromixSubsequence("abdbca"))
	fmt.Println(longestPalindromixSubsequence("cddpd"))
	fmt.Println(longestPalindromixSubsequence("pqr"))
}
