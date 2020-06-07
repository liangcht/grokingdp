package solution

import (
	"fmt"
)

func longestPalindromixSubstring(input string) int {
	n := len(input)
	if n < 1 {
		return 0
	}
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}

	bestStart := 0
	bestEnd := 0
	for i := range dp {
		dp[i][i] = true
	}

	for l := 2; l <= n; l++ {
		for start := 0; start <= n-l; start++ {
			end := start + l - 1
			if input[start] == input[end] {
				if start == end-1 || dp[start+1][end-1] {
					if end-start+1 > bestEnd-bestStart+1 {
						bestEnd = end
						bestStart = start
					}
					dp[start][end] = true
				}
			}
		}
	}
	return bestEnd - bestStart + 1
}

// TestLongestPalindromixSubstring is used for testing
func TestLongestPalindromixSubstring() {
	fmt.Println(longestPalindromixSubstring("abdbca"))
	fmt.Println(longestPalindromixSubstring("cddpd"))
	fmt.Println(longestPalindromixSubstring("pqr"))
}
