package solution

import (
	"fmt"
)

func countPalindromicSubstring(input string) int {
	n := len(input)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	count := 0
	for i := range dp {
		dp[i][i] = true
		count++
	}
	for l := 2; l <= n; l++ {
		for start := 0; start <= n-l; start++ {
			end := start + l - 1
			if input[start] == input[end] {
				if start == end-1 || input[start+1] == input[end-1] {
					count++
					dp[start][end] = true
				}
			}
		}
	}
	return count
}

// TestCountPalindromicSubstring is used for testing
func TestCountPalindromicSubstring() {
	fmt.Println(countPalindromicSubstring("abdbca"))
	fmt.Println(countPalindromicSubstring("cddpd"))
	fmt.Println(countPalindromicSubstring("pqr"))
	fmt.Println(countPalindromicSubstring("qqq"))
}
