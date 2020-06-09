package solution

import (
	"fmt"
	"grokingdp/mathutil"
)

func longestCommonSubsequenceONSpace(s1 string, s2 string) int {
	dp := make([][]int, 2)
	for i := range dp {
		dp[i] = make([]int, len(s2)+1)
	}

	maxCount := 0
	for i := len(s1) - 1; i >= 0; i-- {
		for j := len(s2) - 1; j >= 0; j-- {
			iPlus := (i + 1) % 2
			count := 0
			if s1[i] == s2[j] {
				count = 1 + dp[iPlus][j+1]
			} else {
				count = mathutil.MaxInt(dp[iPlus][j], dp[i%2][j+1])
			}
			dp[i%2][j] = count
			maxCount = mathutil.MaxInt(maxCount, count)
		}
	}
	return maxCount
}

func longestCommonSubsequence(s1 string, s2 string) int {
	// return find_lcsequence_recur(s1, s2, 0, 0, 0)
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
			} else {
				count = mathutil.MaxInt(dp[i+1][j], dp[i][j+1])
			}
			dp[i][j] = count
			maxCount = mathutil.MaxInt(maxCount, count)
		}
	}
	return maxCount
}

func find_lcsequence_recur(s1, s2 string, start1, start2 int, count int) int {
	if start1 == len(s1) || start2 == len(s2) {
		return count
	}

	c1 := find_lcsequence_recur(s1, s2, start1+1, start2, count)
	c2 := find_lcsequence_recur(s1, s2, start1, start2+1, count)

	if s1[start1] == s2[start2] {
		count = find_lcsequence_recur(s1, s2, start1+1, start2+1, count+1)
	}
	return mathutil.MaxInt(count, mathutil.MaxInt(c1, c2))
}

// TestLongestCommonSubsequence is used for testing
func TestLongestCommonSubsequence() {
	fmt.Println(longestCommonSubsequenceONSpace("abdca", "cbda"))
	fmt.Println(longestCommonSubsequenceONSpace("passport", "ppsspt"))
}
