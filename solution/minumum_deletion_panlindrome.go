package solution

import (
	"fmt"
	"grokingdp/mathutil"
)

func minimumDeletionForPanlindrome(input string) int {
	// return recursive(input, 0, len(input)-1)
	n := len(input)
	if n == 0 {
		return -1
	}
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := range dp {
		if i < n-1 && input[i] != input[i+1] {
			dp[i][i+1] = 1
		}
	}

	for l := 3; l <= n; l++ {
		for i := 0; i <= n-l; i++ {
			end := i + l - 1
			if input[i] == input[end] {
				dp[i][end] = dp[i+1][end-1]
				continue
			}
			dp[i][end] = 2 + dp[i+1][end-1]
			dp[i][end] = mathutil.MinInt(dp[i][end], 1+dp[i+1][end])
			dp[i][end] = mathutil.MinInt(dp[i][end], 1+dp[i][end-1])
		}
	}
	return dp[0][n-1]
}

func recursive(input string, start int, end int) int {
	if start == end {
		return 0
	}
	if input[start] == input[end] {
		if start == end-1 {
			return 0
		}
		return recursive(input, start+1, end-1)
	}
	if start == end-1 {
		return 1
	}
	minDeletion := 2 + recursive(input, start+1, end-1)
	minDeletion = mathutil.MinInt(minDeletion, 1+recursive(input, start+1, end))
	minDeletion = mathutil.MinInt(minDeletion, 1+recursive(input, start, end-1))
	return minDeletion
}

// TestMinimumDeletionForPanlindrome is used for testing
func TestMinimumDeletionForPanlindrome() {
	fmt.Println(minimumDeletionForPanlindrome("abdbca"))
	fmt.Println(minimumDeletionForPanlindrome("cddpd"))
	fmt.Println(minimumDeletionForPanlindrome("pqr"))
}
