package solution

import (
	"fmt"
	"grokingdp/mathutil"
)

func longestCommonSubstring(s1 string, s2 string) int {
	// return find_lcs_recur(s1, s2, 0, 0, 0)
	dp := make([][]int, len(s1)+1)
	for i := range dp {
		dp[i] = make([]int, len(s2)+1)
	}

	maxCount := 0
	for i := len(s1) - 1; i >= 0; i-- {
		for j := len(s2) - 1; j >= 0; j-- {
			count := 0
			if s1[i] == s2[j] {
				count = 1 + dp[i+1][j+1]
			}
			dp[i][j] = count
			maxCount = mathutil.MaxInt(maxCount, count)
		}
	}
	return maxCount
}

func find_lcs_recur(s1, s2 string, start1, start2 int, count int) int {
	if start1 == len(s1) || start2 == len(s2) {
		return count
	}

	if s1[start1] == s2[start2] {
		count = find_lcs_recur(s1, s2, start1+1, start2+1, count+1)
	}

	c1 := find_lcs_recur(s1, s2, start1+1, start2, 0)
	c2 := find_lcs_recur(s1, s2, start1, start2+1, 0)
	return mathutil.MaxInt(count, mathutil.MaxInt(c1, c2))
}

// TestLongestCommonSubstring is used for testing
func TestLongestCommonSubstring() {
	fmt.Println(longestCommonSubstring("abdca", "cbda"))
	fmt.Println(longestCommonSubstring("passport", "ppsspt"))
}
